package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/marking"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/runner"
)

const (
	defaultTrackConfigRelative = "tracks/go-core-100/track.json"
	defaultRunTimeoutSeconds   = 90
	defaultOutputTailLines     = 100
)

//go:embed static/*
var embeddedStatic embed.FS

type studioServer struct {
	repoRoot string
	track    catalog.Track
	store    *progress.Store
	files    fs.FS
	mu       sync.Mutex
}

type trackResponse struct {
	ID              string                     `json:"id"`
	Title           string                     `json:"title"`
	Description     string                     `json:"description"`
	OverallDone     int                        `json:"overall_done"`
	OverallTotal    int                        `json:"overall_total"`
	OverallPercent  int                        `json:"overall_percent"`
	CoachMessage    string                     `json:"coach_message"`
	NextRecommended *nextKataRecommendation    `json:"next_recommended,omitempty"`
	Categories      []trackCategorySummaryItem `json:"categories"`
}

type trackCategorySummaryItem struct {
	ID                string                 `json:"id"`
	Title             string                 `json:"title"`
	Description       string                 `json:"description"`
	LearningGoal      string                 `json:"learning_goal"`
	Done              int                    `json:"done"`
	Total             int                    `json:"total"`
	Percent           int                    `json:"percent"`
	MilestoneLabel    string                 `json:"milestone_label"`
	MilestoneMessage  string                 `json:"milestone_message"`
	NextTargetPercent int                    `json:"next_target_percent"`
	RemainingToNext   int                    `json:"remaining_to_next"`
	Katas             []trackKataSummaryItem `json:"katas"`
}

type trackKataSummaryItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Focus     string `json:"focus"`
	Completed bool   `json:"completed"`
}

type kataResponse struct {
	ID        string                `json:"id"`
	Title     string                `json:"title"`
	Focus     string                `json:"focus"`
	Signature string                `json:"signature"`
	Rules     []string              `json:"rules"`
	Category  kataCategoryReference `json:"category"`
	Readme    string                `json:"readme"`
	Code      string                `json:"code"`
	Tests     string                `json:"tests"`
	Progress  progress.KataProgress `json:"progress"`
}

type kataCategoryReference struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type saveRequest struct {
	KataID string `json:"kata_id"`
	Code   string `json:"code"`
	Tests  string `json:"tests"`
}

type runRequest struct {
	KataID         string  `json:"kata_id"`
	Code           *string `json:"code,omitempty"`
	Tests          *string `json:"tests,omitempty"`
	SaveBeforeRun  bool    `json:"save_before_run"`
	TimeoutSeconds int     `json:"timeout_seconds"`
}

type formatRequest struct {
	KataID string  `json:"kata_id"`
	Code   *string `json:"code,omitempty"`
	Tests  *string `json:"tests,omitempty"`
}

type formatResponse struct {
	Code  string `json:"code"`
	Tests string `json:"tests"`
}

type runResponse struct {
	Passed          bool                    `json:"passed"`
	DurationMS      int64                   `json:"duration_ms"`
	FailedTests     []string                `json:"failed_tests"`
	OutputTail      string                  `json:"output_tail"`
	FailureInsights []failureInsight        `json:"failure_insights"`
	CoachHint       string                  `json:"coach_hint"`
	NextRecommended *nextKataRecommendation `json:"next_recommended,omitempty"`
	Progress        progress.KataProgress   `json:"progress"`
}

type markRequest struct {
	KataID string `json:"kata_id"`
}

type markResponse struct {
	PromptPath string `json:"prompt_path"`
	Prompt     string `json:"prompt"`
	ChatGPTURL string `json:"chatgpt_url"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type nextKataRecommendation struct {
	KataID        string `json:"kata_id"`
	KataTitle     string `json:"kata_title"`
	CategoryID    string `json:"category_id"`
	CategoryTitle string `json:"category_title"`
	Reason        string `json:"reason"`
}

type failureInsight struct {
	Kind     string `json:"kind"`
	Summary  string `json:"summary"`
	Expected string `json:"expected,omitempty"`
	Actual   string `json:"actual,omitempty"`
}

var expectedGotPattern = regexp.MustCompile(`(?i)expected[: ]+(.+?)[,; ]+got[: ]+(.+)$`)
var goTestPrefixPattern = regexp.MustCompile(`^[^:]+:\d+:\s*`)

func main() {
	repoRootFlag := flag.String("repo", ".", "Path to kata repository root")
	addrFlag := flag.String("addr", "127.0.0.1:7777", "HTTP listen address")
	flag.Parse()

	server, err := newStudioServer(*repoRootFlag)
	if err != nil {
		fatalf("init learner studio: %v", err)
	}

	fmt.Printf("Learner Studio running at http://%s\n", *addrFlag)
	if err := http.ListenAndServe(*addrFlag, server.routes()); err != nil {
		fatalf("listen: %v", err)
	}
}

func newStudioServer(repoRoot string) (*studioServer, error) {
	absRepoRoot, err := filepath.Abs(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve repo root: %w", err)
	}

	track, err := catalog.LoadTrack(filepath.Join(absRepoRoot, defaultTrackConfigRelative), absRepoRoot)
	if err != nil {
		return nil, err
	}

	store := progress.NewStore(filepath.Join(absRepoRoot, ".learning", "progress.json"))
	if _, err := store.Load(); err != nil {
		return nil, fmt.Errorf("load progress: %w", err)
	}

	staticFS, err := fs.Sub(embeddedStatic, "static")
	if err != nil {
		return nil, fmt.Errorf("load embedded static files: %w", err)
	}

	return &studioServer{
		repoRoot: absRepoRoot,
		track:    track,
		store:    store,
		files:    staticFS,
	}, nil
}

func (s *studioServer) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/app.js", s.handleStaticFile("app.js", "text/javascript; charset=utf-8"))
	mux.HandleFunc("/styles.css", s.handleStaticFile("styles.css", "text/css; charset=utf-8"))

	mux.HandleFunc("/api/track", s.handleTrack)
	mux.HandleFunc("/api/kata", s.handleKata)
	mux.HandleFunc("/api/save", s.handleSave)
	mux.HandleFunc("/api/format", s.handleFormat)
	mux.HandleFunc("/api/run", s.handleRun)
	mux.HandleFunc("/api/mark", s.handleMark)

	return mux
}

func (s *studioServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	s.serveFile(w, "index.html", "text/html; charset=utf-8")
}

func (s *studioServer) handleStaticFile(name, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.serveFile(w, name, contentType)
	}
}

func (s *studioServer) serveFile(w http.ResponseWriter, name, contentType string) {
	data, err := fs.ReadFile(s.files, name)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write(data)
}

func (s *studioServer) handleTrack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	resp := trackResponse{
		ID:          s.track.ID,
		Title:       s.track.Title,
		Description: s.track.Description,
		Categories:  make([]trackCategorySummaryItem, 0, len(s.track.Categories)),
	}

	allIDs := make([]string, 0, len(s.track.AllKatas()))
	for _, kata := range s.track.AllKatas() {
		allIDs = append(allIDs, kata.ID)
	}
	resp.OverallDone = progress.CompletedCount(state, allIDs)
	resp.OverallTotal = len(allIDs)
	if resp.OverallTotal > 0 {
		resp.OverallPercent = int((float64(resp.OverallDone) / float64(resp.OverallTotal)) * 100)
	}
	resp.CoachMessage = coachMessage(resp.OverallDone, resp.OverallTotal)
	resp.NextRecommended = s.nextRecommendation(state)

	for _, category := range s.track.Categories {
		item := trackCategorySummaryItem{
			ID:           category.ID,
			Title:        category.Title,
			Description:  category.Description,
			LearningGoal: category.LearningGoal,
			Katas:        make([]trackKataSummaryItem, 0, len(category.Katas)),
		}

		ids := make([]string, 0, len(category.Katas))
		for _, kata := range category.Katas {
			ids = append(ids, kata.ID)
			item.Katas = append(item.Katas, trackKataSummaryItem{
				ID:        kata.ID,
				Title:     kata.Title,
				Focus:     kata.Focus,
				Completed: state.Attempts[kata.ID].Passes > 0,
			})
		}

		item.Done = progress.CompletedCount(state, ids)
		item.Total = len(category.Katas)
		if item.Total > 0 {
			item.Percent = int((float64(item.Done) / float64(item.Total)) * 100)
		}
		item.MilestoneLabel, item.MilestoneMessage, item.NextTargetPercent, item.RemainingToNext = categoryMilestone(item.Done, item.Total)

		resp.Categories = append(resp.Categories, item)
	}

	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleKata(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	rawID := strings.TrimSpace(r.URL.Query().Get("id"))
	if rawID == "" {
		writeError(w, http.StatusBadRequest, "missing id query parameter")
		return
	}

	kata, category, ok := s.track.FindKata(rawID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	readme, err := os.ReadFile(kata.ReadmePath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read readme: %v", err))
		return
	}
	code, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata.go: %v", err))
		return
	}
	tests, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	resp := kataResponse{
		ID:        kata.ID,
		Title:     kata.Title,
		Focus:     kata.Focus,
		Signature: kata.Signature,
		Rules:     append([]string(nil), kata.Rules...),
		Category: kataCategoryReference{
			ID:    category.ID,
			Title: category.Title,
		},
		Readme:   string(readme),
		Code:     string(code),
		Tests:    string(tests),
		Progress: state.Attempts[kata.ID],
	}

	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req saveRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), []byte(req.Code), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
		return
	}
	if err := os.WriteFile(filepath.Join(kata.Dir, "kata_test.go"), []byte(req.Tests), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata_test.go: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "saved"})
}

func (s *studioServer) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req runRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	if req.TimeoutSeconds <= 0 {
		req.TimeoutSeconds = defaultRunTimeoutSeconds
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if req.SaveBeforeRun {
		if req.Code != nil {
			if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), []byte(*req.Code), 0o644); err != nil {
				writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
				return
			}
		}
		if req.Tests != nil {
			if err := os.WriteFile(filepath.Join(kata.Dir, "kata_test.go"), []byte(*req.Tests), 0o644); err != nil {
				writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata_test.go: %v", err))
				return
			}
		}
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(req.TimeoutSeconds)*time.Second)
	defer cancel()

	start := time.Now()
	runResult, _ := runner.RunKataTests(ctx, kata.Dir)
	duration := runResult.Elapsed
	if duration <= 0 {
		duration = time.Since(start)
	}

	attemptedAt := time.Now().UTC()
	outputTail := tailLines(runResult.RawOutput, defaultOutputTailLines)
	failureInsights := extractFailureInsights(outputTail)
	state, err := s.store.RecordAttempt(kata.ID, progress.AttemptResult{
		Passed:      runResult.Passed,
		Duration:    duration,
		FailedTests: runResult.FailedTests,
		OutputTail:  outputTail,
		RanAt:       attemptedAt,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("record progress: %v", err))
		return
	}

	resp := runResponse{
		Passed:          runResult.Passed,
		DurationMS:      duration.Milliseconds(),
		FailedTests:     append([]string(nil), runResult.FailedTests...),
		OutputTail:      outputTail,
		FailureInsights: failureInsights,
		CoachHint:       coachHint(runResult.Passed, runResult.FailedTests, failureInsights),
		NextRecommended: s.nextRecommendation(state),
		Progress:        state.Attempts[kata.ID],
	}
	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleFormat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req formatRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, _, ok := s.track.FindKata(req.KataID); !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
	defer cancel()

	code := derefOrEmpty(req.Code)
	tests := derefOrEmpty(req.Tests)

	formattedCode, codeErr := formatGoSource(ctx, code)
	formattedTests, testsErr := formatGoSource(ctx, tests)
	if codeErr != nil || testsErr != nil {
		errs := []string{}
		if codeErr != nil {
			errs = append(errs, fmt.Sprintf("code: %v", codeErr))
		}
		if testsErr != nil {
			errs = append(errs, fmt.Sprintf("tests: %v", testsErr))
		}
		writeError(w, http.StatusBadRequest, "format failed: "+strings.Join(errs, "; "))
		return
	}

	writeJSON(w, http.StatusOK, formatResponse{
		Code:  formattedCode,
		Tests: formattedTests,
	})
}

func (s *studioServer) handleMark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req markRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, category, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	kataCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata.go: %v", err))
		return
	}
	testCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	last := state.Attempts[kata.ID]
	prompt := marking.BuildPrompt(marking.PromptInput{
		TrackTitle:      s.track.Title,
		CategoryTitle:   category.Title,
		KataID:          kata.ID,
		KataTitle:       kata.Title,
		Focus:           kata.Focus,
		Signature:       kata.Signature,
		Rules:           kata.Rules,
		KataCode:        string(kataCode),
		TestCode:        string(testCode),
		LastRunResult:   last.LastResult,
		LastFailedTests: last.LastFailedTests,
		LastOutputTail:  last.LastOutputTail,
	})

	target := marking.DefaultPromptPath(s.repoRoot, kata.ID, time.Now().UTC())
	if err := marking.WritePrompt(target, prompt); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write prompt: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, markResponse{
		PromptPath: target,
		Prompt:     prompt,
		ChatGPTURL: "https://chatgpt.com/",
	})
}

func decodeJSON(r *http.Request, out any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(out); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func tailLines(input string, maxLines int) string {
	if maxLines <= 0 {
		return ""
	}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) <= maxLines {
		return strings.Join(lines, "\n")
	}
	return strings.Join(lines[len(lines)-maxLines:], "\n")
}

func coachMessage(done, total int) string {
	if total == 0 {
		return "Track ready. Start with one kata and complete the full loop: read, code, test, reflect."
	}
	if done == 0 {
		return "Start small: complete one kata end-to-end before moving on."
	}
	if done >= total {
		return "Track complete. Revisit a weak category and refactor one kata for readability and tests."
	}
	return fmt.Sprintf("You have %d/%d complete. Keep focus: one kata per session, then review what you learned.", done, total)
}

func categoryMilestone(done, total int) (label, message string, nextTargetPercent, remainingToNext int) {
	if total <= 0 {
		return "No katas", "This category has no katas configured.", 0, 0
	}

	percent := int((float64(done) / float64(total)) * 100)
	thresholds := []int{25, 50, 75, 100}
	nextTargetPercent = 100
	for _, threshold := range thresholds {
		if percent < threshold {
			nextTargetPercent = threshold
			break
		}
	}

	targetCount := int((float64(nextTargetPercent) / 100.0) * float64(total))
	if targetCount < 1 {
		targetCount = 1
	}
	if targetCount > total {
		targetCount = total
	}
	remainingToNext = targetCount - done
	if remainingToNext < 0 {
		remainingToNext = 0
	}

	switch {
	case percent == 0:
		label = "Start line"
		message = "Complete your first kata in this category to establish momentum."
	case percent < 25:
		label = "Early momentum"
		message = "Good start. Keep the same category to build pattern recognition."
	case percent < 50:
		label = "Building consistency"
		message = "You are stacking wins. Tighten tests around edge cases."
	case percent < 75:
		label = "Solid core"
		message = "Halfway and stable. Focus on failure paths and clarity."
	case percent < 100:
		label = "Finishing strong"
		message = "You are close. Polish behavior contracts and naming quality."
	default:
		label = "Category complete"
		message = "Great work. Move to the next category and keep the same discipline."
		nextTargetPercent = 100
		remainingToNext = 0
	}

	return label, message, nextTargetPercent, remainingToNext
}

func (s *studioServer) nextRecommendation(state progress.State) *nextKataRecommendation {
	for _, category := range s.track.Categories {
		for _, kata := range category.Katas {
			if state.Attempts[kata.ID].Passes > 0 {
				continue
			}
			return &nextKataRecommendation{
				KataID:        kata.ID,
				KataTitle:     kata.Title,
				CategoryID:    category.ID,
				CategoryTitle: category.Title,
				Reason:        fmt.Sprintf("Next best step: %s in %s. %s", kata.Title, category.Title, category.LearningGoal),
			}
		}
	}
	return nil
}

func coachHint(passed bool, failedTests []string, insights []failureInsight) string {
	if passed {
		return "Nice pass. Capture one sentence about what changed, then move to the next kata."
	}
	if len(failedTests) == 0 {
		return "Start with the first failure line. Make one small fix, run again, repeat."
	}
	if len(insights) > 0 {
		return "Focus on the first mismatch only. Get that passing before touching anything else."
	}
	return "Use the failing test names as your map. Fix in order, smallest behavior gap first."
}

func extractFailureInsights(output string) []failureInsight {
	lines := strings.Split(output, "\n")
	insights := make([]failureInsight, 0, 6)
	seen := map[string]bool{}

	add := func(item failureInsight) {
		key := item.Kind + "|" + item.Summary + "|" + item.Expected + "|" + item.Actual
		if seen[key] || item.Summary == "" {
			return
		}
		seen[key] = true
		insights = append(insights, item)
	}

	for i := 0; i < len(lines) && len(insights) < 6; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		normalized := strings.TrimSpace(goTestPrefixPattern.ReplaceAllString(line, ""))
		lower := strings.ToLower(normalized)

		if strings.HasPrefix(normalized, "--- FAIL:") {
			add(failureInsight{Kind: "test", Summary: clip(normalized, 140)})
			continue
		}
		if strings.Contains(lower, "panic:") {
			add(failureInsight{Kind: "panic", Summary: clip(normalized, 140)})
			continue
		}

		if strings.HasPrefix(lower, "got:") || strings.HasPrefix(lower, "actual:") {
			actual := strings.TrimSpace(strings.SplitN(normalized, ":", 2)[1])
			expected := ""
			if i+1 < len(lines) {
				next := strings.TrimSpace(lines[i+1])
				next = strings.TrimSpace(goTestPrefixPattern.ReplaceAllString(next, ""))
				nextLower := strings.ToLower(next)
				if strings.HasPrefix(nextLower, "want:") || strings.HasPrefix(nextLower, "expected:") {
					expected = strings.TrimSpace(strings.SplitN(next, ":", 2)[1])
				}
			}
			add(failureInsight{
				Kind:     "mismatch",
				Summary:  "Expected and actual values do not match.",
				Expected: clip(expected, 120),
				Actual:   clip(actual, 120),
			})
			continue
		}

		match := expectedGotPattern.FindStringSubmatch(normalized)
		if len(match) == 3 {
			add(failureInsight{
				Kind:     "mismatch",
				Summary:  "Expected and actual values do not match.",
				Expected: clip(strings.TrimSpace(match[1]), 120),
				Actual:   clip(strings.TrimSpace(match[2]), 120),
			})
		}
	}

	return insights
}

func clip(input string, max int) string {
	trimmed := strings.TrimSpace(input)
	if max <= 0 || len(trimmed) <= max {
		return trimmed
	}
	if max <= 3 {
		return trimmed[:max]
	}
	return trimmed[:max-3] + "..."
}

func formatGoSource(ctx context.Context, src string) (string, error) {
	cmd := exec.CommandContext(ctx, "gofmt")
	cmd.Stdin = strings.NewReader(src)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errText := strings.TrimSpace(stderr.String())
		if errText == "" {
			errText = err.Error()
		}
		return "", errors.New(errText)
	}
	return stdout.String(), nil
}

func derefOrEmpty(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func fatalf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
