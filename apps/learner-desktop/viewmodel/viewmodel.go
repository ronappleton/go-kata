// Package viewmodel separates business logic from the GTK view layer.
//
// AppViewModel owns kata selection, progress tracking, sandbox run lifecycle,
// flashcard/quiz deck construction, and learning-mode state. The GTK view
// subscribes to callbacks to update widgets; the ViewModel never references
// any GTK type, so it can be tested with standard go test (no display needed).
package viewmodel

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/languages"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/diagnostics"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/evaluator"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/workspace"
)

// catalogLanguage resolves a kata's language identifier to a Language
// definition, defaulting to Go when unknown.
func catalogLanguage(id string) *languages.Language {
	reg := languages.NewRegistry()
	if lang := reg.Lookup(id); lang != nil {
		return lang
	}
	return reg.Default()
}

// ── Learning mode ──────────────────────────────────────────────────────────

type LearningMode int

const (
	ModeLinear LearningMode = iota
	ModeADHD
	ModeReview
)

func (m LearningMode) String() string {
	switch m {
	case ModeLinear:
		return "Linear"
	case ModeADHD:
		return "ADHD"
	case ModeReview:
		return "Review"
	default:
		return "Unknown"
	}
}

// ── Flashcard / Quiz types ────────────────────────────────────────────────

type FlashcardEntry struct {
	KataID string
	Front  string
	Back   string
}

type QuizEntry struct {
	KataID   string
	Question string
	Options  []string
	Answer   string
}

// ── Callbacks the view layer registers ────────────────────────────────────

type Callbacks struct {
	OnKataChanged        func(kata catalog.Kata)
	OnModeChanged        func(mode LearningMode)
	OnRunStarted         func()
	OnRunCompleted       func(result evaluator.Result, passed bool, failedTests []string)
	OnProgressUpdated    func(stageID, categoryID string, fraction float64)
	OnStatusChanged      func(message string)
	OnRunnerStateChanged func(available bool)
}

// ── AppViewModel ──────────────────────────────────────────────────────────

type AppViewModel struct {
	mu sync.Mutex

	// config
	Config desktopConfig

	// state
	Track      catalog.Track
	AllTracks  []catalog.TrackInfo
	TrackPaths map[string]string
	Mode       LearningMode
	Selected   catalog.Kata
	Running    bool

	// sub-systems
	Paths     workspace.Paths
	Workspace *workspace.Manager
	Progress  *progress.Store
	Runner    *evaluator.Runner

	// callback registration
	CB Callbacks

	// workspace-derived text (cached per-select)
	Readme       string
	Solution     string
	LearnerTests string
	Output       string

	// flashcard / quiz decks (current selection)
	FlashDeck    []FlashcardEntry
	FlashDeckIdx int
	FlashSide    string // "Front" | "Back"
	QuizDeck     []QuizEntry
	QuizDeckIdx  int
	QuizAnswered bool

	// ADHD counters
	KataCount int

	// cross-kata decks (Review mode)
	ReviewFlashDeck []FlashcardEntry
	ReviewQuizDeck  []QuizEntry
}

type desktopConfig struct {
	ContentRoot string
	Image       string
	DevMode     bool
}

func New(cfg desktopConfig) *AppViewModel {
	return &AppViewModel{
		Config:     cfg,
		Mode:       ModeLinear,
		TrackPaths: make(map[string]string),
		FlashSide:  "Front",
	}
}

// ── Lifecycle ─────────────────────────────────────────────────────────────

// Init loads the default track, resolves workspace, and starts diagnostics.
// Call after constructing with New(). Returns an error for fatal startup failures.
func (vm *AppViewModel) Init() error {
	track, err := catalog.LoadTrack(filepath.Join(vm.Config.ContentRoot, "tracks", "go-core-100", "track.json"))
	if err != nil {
		return fmt.Errorf("curriculum could not be loaded: %w", err)
	}
	vm.Track = track

	paths, err := workspace.ResolvePaths("gokatas")
	if err != nil {
		return fmt.Errorf("application directories could not be resolved: %w", err)
	}
	vm.Paths = paths
	vm.Workspace = workspace.NewManager(paths)
	if err := vm.Workspace.Ensure(); err != nil {
		return fmt.Errorf("application directories could not be created: %w", err)
	}
	vm.Progress = progress.NewStore(filepath.Join(paths.State, "progress.json"))
	if vm.Config.Image != "" {
		vm.Runner, _ = evaluator.NewRunner(vm.Config.Image)
	}

	// Discover other tracks
	vm.AllTracks, _ = catalog.DiscoverTracks(filepath.Join(vm.Config.ContentRoot, "tracks"))
	for _, t := range vm.AllTracks {
		vm.TrackPaths[t.Title] = filepath.Join(vm.Config.ContentRoot, "tracks", t.ID, "track.json")
	}

	return nil
}

// StartDiagnostics checks podman availability in background.
func (vm *AppViewModel) StartDiagnostics() {
	image := vm.Config.Image
	go func() {
		report := diagnostics.Check(context.Background(), "podman", image)
		vm.setStatus(report.Message)
		if vm.CB.OnRunnerStateChanged != nil {
			vm.CB.OnRunnerStateChanged(report.PodmanAvailable)
		}
	}()
}

// ── Kata selection ────────────────────────────────────────────────────────

// SelectKata loads workspace content for a kata and notifies the view.
func (vm *AppViewModel) SelectKata(kata catalog.Kata) {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vm.Selected = kata

	lang := catalogLanguage(kata.Language)
	solution, err := vm.Workspace.ReadSolutionAs(kata.ID, lang.SourceFilename)
	if err != nil {
		vm.setStatus(fmt.Sprintf("Unable to read workspace: %v", err))
		return
	}
	if solution == "" {
		solution = string(kata.Content.KataGo)
	}
	vm.Solution = solution

	learnerTests, err := vm.Workspace.ReadLearnerTestsAs(kata.ID, lang.TestsFilename)
	if err != nil {
		vm.setStatus(fmt.Sprintf("Unable to read workspace: %v", err))
		return
	}
	vm.LearnerTests = learnerTests
	vm.Readme = kata.Content.Readme
	vm.Output = ""
	vm.FlashDeckIdx = 0
	vm.QuizDeckIdx = 0
	vm.QuizAnswered = false
	vm.FlashSide = "Front"

	// Build flashcard / quiz decks for current kata
	vm.populateCurrentFlashDeck()
	vm.populateCurrentQuizDeck()

	// Update status
	if vm.Runner == nil {
		vm.setStatus("Podman runner is not configured. Set a digest-pinned runner image in setup.")
	} else if kata.EvaluatorStatus != "ready" {
		vm.setStatus("This kata does not yet have a complete trusted evaluator.")
	} else {
		vm.setStatus("Ready. Run executes in a disposable rootless Podman container.")
	}

	if vm.CB.OnKataChanged != nil {
		vm.CB.OnKataChanged(kata)
	}
}

// ── Save ──────────────────────────────────────────────────────────────────

// SaveCurrent writes the current solution and learner tests to the workspace.
func (vm *AppViewModel) SaveCurrent() {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	if vm.Selected.ID == "" || vm.Workspace == nil {
		return
	}
	lang := catalogLanguage(vm.Selected.Language)
	if err := vm.Workspace.SaveSolutionAs(vm.Selected.ID, vm.Solution, evaluator.DefaultCodeLimitBytes, lang.SourceFilename); err != nil {
		vm.setStatus(fmt.Sprintf("Save failed: %v", err))
		return
	}
	if err := vm.Workspace.SaveLearnerTestsAs(vm.Selected.ID, vm.LearnerTests, evaluator.DefaultTestsLimitBytes, lang.TestsFilename); err != nil {
		vm.setStatus(fmt.Sprintf("Save failed: %v", err))
		return
	}
	vm.setStatus("Saved to the user workspace.")
}

// ── Run ───────────────────────────────────────────────────────────────────

// RunCurrent saves, then runs the sandbox. Returns immediately; the view
// is notified via OnRunStarted / OnRunCompleted callbacks.
func (vm *AppViewModel) RunCurrent() {
	vm.mu.Lock()
	if vm.Selected.ID == "" || vm.Runner == nil || vm.Running || vm.Selected.EvaluatorStatus != "ready" {
		vm.mu.Unlock()
		return
	}
	vm.Running = true
	vm.setStatus("Running in sandbox…")
	if vm.CB.OnRunStarted != nil {
		vm.CB.OnRunStarted()
	}

	kata := vm.Selected
	solution := vm.Solution
	learnerTests := vm.LearnerTests
	trustedTests := []byte(kata.Content.KataTest)
	runner := vm.Runner
	vm.mu.Unlock()

	// Save first (outside lock since Save acquires it)
	vm.mu.Lock()
	vm.Workspace.SaveSolution(kata.ID, solution, evaluator.DefaultCodeLimitBytes)
	vm.Workspace.SaveLearnerTests(kata.ID, learnerTests, evaluator.DefaultTestsLimitBytes)
	vm.mu.Unlock()

	go func() {
		result := runner.Run(context.Background(), evaluator.Request{
			KataID:       kata.ID,
			Module:       "kata" + kata.ID,
			Code:         solution,
			LearnerTests: learnerTests,
			TrustedTests: string(trustedTests),
		})

		vm.mu.Lock()
		vm.Running = false
		vm.Output = formatResult(result)
		vm.mu.Unlock()

		// Record progress
		if vm.Progress != nil {
			vm.Progress.RecordAttempt(kata.ID, progress.AttemptResult{
				Passed:      result.Passed,
				Duration:    result.Duration,
				FailedTests: result.FailedTests,
				OutputTail:  result.Output,
				RanAt:       time.Now().UTC(),
			})
			// Update all progress bars
			for _, stage := range vm.Track.Stages {
				for _, cat := range stage.Categories {
					total := len(cat.Katas)
					completed := vm.categoryCompletedCount(cat)
					frac := 0.0
					if total > 0 {
						frac = float64(completed) / float64(total)
					}
					if vm.CB.OnProgressUpdated != nil {
						vm.CB.OnProgressUpdated(stage.ID, cat.ID, frac)
					}
				}
			}
		}

		// ADHD break counter
		if vm.Mode == ModeADHD {
			vm.mu.Lock()
			vm.KataCount++
			vm.mu.Unlock()
		}

		vm.setStatus(vm.resultSummary(result))
		if vm.CB.OnRunCompleted != nil {
			vm.CB.OnRunCompleted(result, result.Passed, result.FailedTests)
		}
	}()
}

func (vm *AppViewModel) resultSummary(result evaluator.Result) string {
	switch result.Status {
	case evaluator.StatusPassed:
		return fmt.Sprintf("PASSED in %s", result.Duration.Round(time.Millisecond))
	case evaluator.StatusFailed:
		return fmt.Sprintf("FAILED in %s — %d failing tests", result.Duration.Round(time.Millisecond), len(result.FailedTests))
	case evaluator.StatusTimeout:
		return fmt.Sprintf("TIMEOUT after %s", result.Duration.Round(time.Millisecond))
	default:
		return fmt.Sprintf("%s — %s", strings.ToUpper(string(result.Status)), result.Duration.Round(time.Millisecond))
	}
}

// ── Mode switching ────────────────────────────────────────────────────────

// SetMode changes the learning mode and rebuilds decks.
func (vm *AppViewModel) SetMode(mode LearningMode) {
	vm.mu.Lock()
	vm.Mode = mode
	vm.KataCount = 0
	vm.mu.Unlock()

	if vm.CB.OnModeChanged != nil {
		vm.CB.OnModeChanged(mode)
	}
}

// ── Track switching ───────────────────────────────────────────────────────

// SetTrack loads a track from the given path and rebuilds the sidebar.
func (vm *AppViewModel) SetTrack(title string) {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	p, ok := vm.TrackPaths[title]
	if !ok {
		return
	}
	t, err := catalog.LoadTrack(p)
	if err == nil {
		vm.Track = t
	}
}

// ── Sidebar data ──────────────────────────────────────────────────────────

// SidebarData returns the flattened list of items for the sidebar builder.
type SidebarStage struct {
	ID    string
	Title string
	Level string
	Katas []SidebarKata
}

type SidebarKata struct {
	ID              string
	Title           string
	Available       bool
	Completed       bool
	EvaluatorStatus string
}

// SidebarItems returns the full sidebar model (stages + katas + progress).
func (vm *AppViewModel) SidebarItems() []SidebarStage {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	var stages []SidebarStage
	for _, stage := range vm.Track.Stages {
		ss := SidebarStage{
			ID:    stage.ID,
			Title: strings.ToUpper(stage.Title),
			Level: stage.Level,
		}
		for _, cat := range stage.Categories {
			for _, kata := range cat.Katas {
				available := true
				if vm.Mode == ModeLinear {
					available = vm.isKataAvailable(kata, stage, cat)
				}
				completed := false
				if vm.Progress != nil {
					state, _ := vm.Progress.Load()
					completed = state.Attempts[kata.ID].Passes > 0
				}
				ss.Katas = append(ss.Katas, SidebarKata{
					ID:              kata.ID,
					Title:           fmt.Sprintf("%s  %s", kata.ID, kata.Title),
					Available:       available,
					Completed:       completed,
					EvaluatorStatus: kata.EvaluatorStatus,
				})
			}
		}
		stages = append(stages, ss)
	}
	return stages
}

// StageProgress returns (completed, total) for a stage.
func (vm *AppViewModel) StageProgress(stage catalog.Stage) (int, int) {
	total := 0
	completed := 0
	for _, cat := range stage.Categories {
		total += len(cat.Katas)
		completed += vm.categoryCompletedCount(cat)
	}
	return completed, total
}

// CategoryProgress returns (completed, total) for a category.
func (vm *AppViewModel) CategoryProgress(cat catalog.Category) (int, int) {
	return vm.categoryCompletedCount(cat), len(cat.Katas)
}

// ── Flashcard deck ────────────────────────────────────────────────────────

// BuildReviewDecks builds cross-kata flashcard and quiz decks for Review mode.
func (vm *AppViewModel) BuildReviewDecks() {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vm.ReviewFlashDeck = nil
	vm.ReviewQuizDeck = nil

	for _, kata := range vm.Track.AllKatas() {
		for _, fc := range kata.Flashcards {
			vm.ReviewFlashDeck = append(vm.ReviewFlashDeck, FlashcardEntry{
				KataID: kata.ID,
				Front:  fc.Front,
				Back:   fc.Back,
			})
		}
		for _, q := range kata.QuizQuestions {
			vm.ReviewQuizDeck = append(vm.ReviewQuizDeck, QuizEntry{
				KataID:   kata.ID,
				Question: q.Question,
				Options:  q.Options,
				Answer:   q.Answer,
			})
		}
	}
	// Shuffle
	shuffleFlashDeck(vm.ReviewFlashDeck)
	shuffleQuizDeck(vm.ReviewQuizDeck)
	// Limit
	if len(vm.ReviewFlashDeck) > 20 {
		vm.ReviewFlashDeck = vm.ReviewFlashDeck[:20]
	}
	if len(vm.ReviewQuizDeck) > 15 {
		vm.ReviewQuizDeck = vm.ReviewQuizDeck[:15]
	}
}

func shuffleFlashDeck(d []FlashcardEntry) {
	for i := len(d) - 1; i > 0; i-- {
		j := time.Now().UnixNano() % int64(i+1)
		d[i], d[j] = d[j], d[i]
	}
}

func shuffleQuizDeck(d []QuizEntry) {
	for i := len(d) - 1; i > 0; i-- {
		j := time.Now().UnixNano() % int64(i+1)
		d[i], d[j] = d[j], d[i]
	}
}

// CurrentFlashDeck returns the active flashcard deck.
func (vm *AppViewModel) CurrentFlashDeck() []FlashcardEntry {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	if vm.Mode == ModeReview {
		return vm.ReviewFlashDeck
	}
	return vm.FlashDeck
}

// CurrentQuizDeck returns the active quiz deck.
func (vm *AppViewModel) CurrentQuizDeck() []QuizEntry {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	if vm.Mode == ModeReview {
		return vm.ReviewQuizDeck
	}
	return vm.QuizDeck
}

// FlipFlashcard toggles front/back.
func (vm *AppViewModel) FlipFlashcard() {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	if vm.FlashSide == "Front" {
		vm.FlashSide = "Back"
	} else {
		vm.FlashSide = "Front"
	}
}

// NextFlashcard advances the deck index.
func (vm *AppViewModel) NextFlashcard() {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	deck := vm.currentFlashDeckLocked()
	if vm.FlashDeckIdx < len(deck)-1 {
		vm.FlashDeckIdx++
		vm.FlashSide = "Front"
	}
}

// PrevFlashcard goes back one card.
func (vm *AppViewModel) PrevFlashcard() {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	if vm.FlashDeckIdx > 0 {
		vm.FlashDeckIdx--
		vm.FlashSide = "Front"
	}
}

// RevealQuizAnswer shows the answer for the current quiz question.
func (vm *AppViewModel) RevealQuizAnswer() string {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	deck := vm.currentQuizDeckLocked()
	if vm.QuizDeckIdx >= len(deck) {
		return ""
	}
	vm.QuizAnswered = true
	return deck[vm.QuizDeckIdx].Answer
}

// NextQuizQuestion advances the quiz deck index.
func (vm *AppViewModel) NextQuizQuestion() {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	deck := vm.currentQuizDeckLocked()
	if vm.QuizDeckIdx < len(deck)-1 {
		vm.QuizDeckIdx++
		vm.QuizAnswered = false
	}
}

// ── Reflection ────────────────────────────────────────────────────────────

// SaveReflection writes the reflection journal entry.
func (vm *AppViewModel) SaveReflection(text string) error {
	if vm.Selected.ID == "" || vm.Workspace == nil {
		return fmt.Errorf("no kata selected")
	}
	root, err := vm.Workspace.Workspace(vm.Selected.ID)
	if err != nil {
		return err
	}
	return workspace.AtomicWrite(filepath.Join(root, "reflection.md"), []byte(text), 0o600)
}

// ReadReflection reads the reflection for the current kata.
func (vm *AppViewModel) ReadReflection() string {
	if vm.Selected.ID == "" || vm.Workspace == nil {
		return ""
	}
	root, err := vm.Workspace.Workspace(vm.Selected.ID)
	if err != nil {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(root, "reflection.md"))
	if err != nil {
		return ""
	}
	return string(data)
}

// ── Status message ────────────────────────────────────────────────────────

func (vm *AppViewModel) setStatus(msg string) {
	if vm.CB.OnStatusChanged != nil {
		vm.CB.OnStatusChanged(msg)
	}
}

// ── Internal helpers ──────────────────────────────────────────────────────

func (vm *AppViewModel) isKataAvailable(kata catalog.Kata, stage catalog.Stage, cat catalog.Category) bool {
	if len(vm.Track.Stages) == 0 || stage.ID == vm.Track.Stages[0].ID {
		return true
	}
	for _, prevStage := range vm.Track.Stages {
		if prevStage.ID == stage.ID {
			break
		}
		completed := vm.stageCompletedCount(prevStage)
		if completed >= 3 {
			return true
		}
	}
	return false
}

func (vm *AppViewModel) stageCompletedCount(stage catalog.Stage) int {
	count := 0
	for _, cat := range stage.Categories {
		count += vm.categoryCompletedCount(cat)
	}
	return count
}

func (vm *AppViewModel) categoryCompletedCount(cat catalog.Category) int {
	if vm.Progress == nil {
		return 0
	}
	state, _ := vm.Progress.Load()
	count := 0
	for _, kata := range cat.Katas {
		if state.Attempts[kata.ID].Passes > 0 {
			count++
		}
	}
	return count
}

func (vm *AppViewModel) populateCurrentFlashDeck() {
	vm.FlashDeck = nil
	vm.FlashDeckIdx = 0
	if vm.Selected.ID == "" {
		return
	}
	if len(vm.Selected.Flashcards) > 0 {
		for _, fc := range vm.Selected.Flashcards {
			vm.FlashDeck = append(vm.FlashDeck, FlashcardEntry{
				KataID: vm.Selected.ID,
				Front:  fc.Front,
				Back:   fc.Back,
			})
		}
		return
	}
	// Fallback: generate from rules
	if vm.Selected.Signature != "" {
		vm.FlashDeck = append(vm.FlashDeck, FlashcardEntry{
			KataID: vm.Selected.ID,
			Front:  "What function contract are you implementing?",
			Back:   vm.Selected.Signature,
		})
	}
	if vm.Selected.Focus != "" {
		vm.FlashDeck = append(vm.FlashDeck, FlashcardEntry{
			KataID: vm.Selected.ID,
			Front:  "What skill is this kata training?",
			Back:   vm.Selected.Focus,
		})
	}
	for _, rule := range vm.Selected.Rules {
		parts := strings.SplitN(rule, "=>", 2)
		front := "What behavior is required?"
		back := rule
		if len(parts) == 2 {
			front = "What should happen when " + strings.TrimSpace(parts[0]) + "?"
			back = strings.TrimSpace(parts[1])
		}
		vm.FlashDeck = append(vm.FlashDeck, FlashcardEntry{
			KataID: vm.Selected.ID,
			Front:  front,
			Back:   back,
		})
	}
}

func (vm *AppViewModel) populateCurrentQuizDeck() {
	vm.QuizDeck = nil
	vm.QuizDeckIdx = 0
	vm.QuizAnswered = false
	if vm.Selected.ID == "" {
		return
	}
	if len(vm.Selected.QuizQuestions) > 0 {
		for _, q := range vm.Selected.QuizQuestions {
			vm.QuizDeck = append(vm.QuizDeck, QuizEntry{
				KataID:   vm.Selected.ID,
				Question: q.Question,
				Options:  q.Options,
				Answer:   q.Answer,
			})
		}
		return
	}
	// Fallback: generate from rules
	for _, rule := range vm.Selected.Rules {
		parts := strings.SplitN(rule, "=>", 2)
		question := "Which behavior is required by the kata?"
		answer := rule
		if len(parts) == 2 {
			question = "What is the correct outcome when " + strings.TrimSpace(parts[0]) + "?"
			answer = strings.TrimSpace(parts[1])
		}
		vm.QuizDeck = append(vm.QuizDeck, QuizEntry{
			KataID:   vm.Selected.ID,
			Question: question,
			Answer:   answer,
		})
	}
}

func (vm *AppViewModel) currentFlashDeckLocked() []FlashcardEntry {
	if vm.Mode == ModeReview {
		return vm.ReviewFlashDeck
	}
	return vm.FlashDeck
}

func (vm *AppViewModel) currentQuizDeckLocked() []QuizEntry {
	if vm.Mode == ModeReview {
		return vm.ReviewQuizDeck
	}
	return vm.QuizDeck
}

func formatResult(result evaluator.Result) string {
	lines := []string{
		strings.ToUpper(string(result.Status)),
		fmt.Sprintf("Duration: %s", result.Duration.Round(time.Millisecond)),
	}
	if len(result.FailedTests) > 0 {
		lines = append(lines, "Failing tests: "+strings.Join(result.FailedTests, ", "))
	}
	if result.EvaluatorError != "" {
		lines = append(lines, "Runner detail: "+result.EvaluatorError)
	}
	if strings.TrimSpace(result.Output) != "" {
		lines = append(lines, "", result.Output)
	}
	return strings.Join(lines, "\n")
}
