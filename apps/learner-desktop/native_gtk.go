//go:build gtk4

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/diagnostics"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/evaluator"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/workspace"
	"github.com/ronappleton/golang-katas-1-100/internal/rendering"
)

// LearningMode controls how the sidebar and flashcards/quiz behave.
type LearningMode int

const (
	ModeLinear LearningMode = iota
	ModeADHD
	ModeReview
)

type nativeApp struct {
	config     desktopConfig
	track      catalog.Track
	allTracks  []catalog.TrackInfo
	trackPaths map[string]string
	paths      workspace.Paths
	workspace  *workspace.Manager
	progress   *progress.Store
	runner     *evaluator.Runner

	window       *gtk.ApplicationWindow
	headerBox    *gtk.Box
	kataList     *gtk.Box
	title        *gtk.Label
	subtitle     *gtk.Label
	status       *gtk.Label
	docs         *gtk.TextBuffer
	code         *gtk.TextBuffer
	learnerTests *gtk.TextBuffer
	output       *gtk.TextBuffer
	reflection   *gtk.TextBuffer
	stack        *gtk.Stack
	runButton    *gtk.Button
	saveButton   *gtk.Button
	flashText    *gtk.Label
	flashSide    *gtk.Label
	flashIndex   int
	quizText     *gtk.Label
	quizFeedback *gtk.Label
	quizIndex    int
	bugText      *gtk.Label
	outputSpinner *gtk.Button
	startCodingBtn *gtk.Button
	selected     catalog.Kata
	running      bool
	kataButtons  map[string]*gtk.Button
	// Learning mode
	mode        LearningMode
	modeButtons []*gtk.ToggleButton
	// Progress bars
	stageProgressBars map[string]*gtk.ProgressBar
	catProgressBars   map[string]*gtk.ProgressBar
	// ADHD mode
	breakReminder *gtk.Label
	kataCount     int
	// Cross-kata flashcard deck
	flashDeck    []flashcardEntry
	flashDeckIdx int
	// Cross-kata quiz deck
	quizDeck     []quizEntry
	quizDeckIdx  int
	quizAnswered bool
}

type flashcardEntry struct {
	kataID string
	front  string
	back   string
}

type quizEntry struct {
	kataID   string
	question string
	options  []string
	answer   string
}

func runNative(config desktopConfig) int {
	app := gtk.NewApplication("org.gokatas.LearnerStudio", gio.ApplicationFlagsNone)
	model := &nativeApp{config: config, mode: ModeLinear}
	app.ConnectActivate(func() {
		if model.window == nil {
			installTheme()
			model.build(app)
		}
		model.window.Present()
	})
	return app.Run([]string{os.Args[0]})
}

func (n *nativeApp) build(app *gtk.Application) {
	window := gtk.NewApplicationWindow(app)
	n.window = window
	window.SetTitle("GoKatas")
	window.SetDefaultSize(1440, 920)
	window.AddCSSClass("gokatas")

	track, err := catalog.LoadTrack(filepath.Join(n.config.ContentRoot, "tracks", "go-core-100", "track.json"))
	if err != nil {
		n.showFatal(window, fmt.Sprintf("Curriculum could not be loaded:\n\n%v", err))
		return
	}
	n.track = track

	paths, err := workspace.ResolvePaths("gokatas")
	if err != nil {
		n.showFatal(window, fmt.Sprintf("Application directories could not be resolved:\n\n%v", err))
		return
	}
	n.paths = paths
	n.workspace = workspace.NewManager(paths)
	if err := n.workspace.Ensure(); err != nil {
		n.showFatal(window, fmt.Sprintf("Application directories could not be created:\n\n%v", err))
		return
	}
	n.progress = progress.NewStore(filepath.Join(paths.State, "progress.json"))
	if n.config.Image != "" {
		n.runner, _ = evaluator.NewRunner(n.config.Image)
	}

	n.docs = gtk.NewTextBuffer(nil)
	n.code = gtk.NewTextBuffer(nil)
	n.learnerTests = gtk.NewTextBuffer(nil)
	n.output = gtk.NewTextBuffer(nil)
	n.stageProgressBars = make(map[string]*gtk.ProgressBar)
	n.catProgressBars = make(map[string]*gtk.ProgressBar)

	root := gtk.NewBox(gtk.OrientationVertical, 0)
	n.headerBox = n.buildHeader()
	root.Append(n.headerBox)
	body := gtk.NewPaned(gtk.OrientationHorizontal)
	body.SetPosition(380)
	body.SetStartChild(n.buildSidebar())
	body.SetEndChild(n.buildContent())
	root.Append(body)
	window.SetChild(root)

	if len(n.track.AllKatas()) > 0 {
		n.selectKata(n.track.AllKatas()[0])
	}
	n.startDiagnostics()
	n.initUpdateCheck()
}

func (n *nativeApp) startDiagnostics() {
	image := n.config.Image
	go func() {
		report := diagnostics.Check(context.Background(), "podman", image)
		glib.MainContextDefault().InvokeFull(0, func() bool {
			n.setStatus(report.Message)
			return false
		})
	}()
}

func (n *nativeApp) buildHeader() *gtk.Box {
	header := gtk.NewBox(gtk.OrientationHorizontal, 12)
	header.AddCSSClass("header")

	brand := gtk.NewBox(gtk.OrientationVertical, 0)
	title := gtk.NewLabel("")
	title.SetMarkup(`<span color="#14b8a6" weight="800">Go</span><span color="#e7ecf3" weight="800">Katas</span>`)
	title.AddCSSClass("brand")
	title.SetHAlign(gtk.AlignStart)
	subtitle := gtk.NewLabel("Junior to Lead — Go Mastery")
	subtitle.AddCSSClass("brand-sub")
	subtitle.SetHAlign(gtk.AlignStart)
	brand.Append(title)
	brand.Append(subtitle)
	brand.SetHExpand(true)
	header.Append(brand)

	// Mode selector (toggle buttons)
	modeLabel := gtk.NewLabel("Mode:")
	modeLabel.AddCSSClass("body")
	header.Append(modeLabel)
	modeBox := gtk.NewBox(gtk.OrientationHorizontal, 2)
	n.modeButtons = make([]*gtk.ToggleButton, 3)
	modes := []string{"Linear", "ADHD", "Review"}
	for i, name := range modes {
		btn := gtk.NewToggleButtonWithLabel(name)
		btn.AddCSSClass("mode-btn")
		if i == 0 {
			btn.SetActive(true)
			btn.AddCSSClass("active")
		}
		idx := i
		btn.ConnectToggled(func() {
			if btn.Active() {
				n.mode = LearningMode(idx)
				n.rebuildSidebar()
				n.buildFlashDeck()
				n.buildQuizDeck()
				n.kataCount = 0
				// Update button styles
				for j, b := range n.modeButtons {
					if j == idx {
						b.AddCSSClass("active")
					} else {
						b.RemoveCSSClass("active")
					}
				}
			}
		})
		n.modeButtons[i] = btn
		modeBox.Append(btn)
	}
	header.Append(modeBox)

	// Track selector
	n.trackPaths = make(map[string]string)
	n.allTracks, _ = catalog.DiscoverTracks(filepath.Join(n.config.ContentRoot, "tracks"))
	if len(n.allTracks) > 1 {
		trackNames := make([]string, len(n.allTracks))
		for i, t := range n.allTracks {
			trackNames[i] = t.Title
			n.trackPaths[t.Title] = filepath.Join(n.config.ContentRoot, "tracks", t.ID, "track.json")
		}
		trackLabel := gtk.NewLabel("Track:")
		trackLabel.AddCSSClass("body")
		header.Append(trackLabel)
		trackCombo := gtk.NewDropDownFromStrings(trackNames)
		trackCombo.Connect("notify::selected", func() {
			idx := trackCombo.Selected()
			if int(idx) < len(n.allTracks) {
				title := n.allTracks[idx].Title
				if p, ok := n.trackPaths[title]; ok {
					if t, err := catalog.LoadTrack(p); err == nil {
						n.track = t
						n.rebuildSidebar()
					}
				}
			}
		})
		header.Append(trackCombo)
	}

	n.status = gtk.NewLabel("Starting…")
	n.status.AddCSSClass("status-pill")
	n.status.SetHAlign(gtk.AlignStart)
	header.Append(n.status)

	n.saveButton = gtk.NewButtonWithLabel("Save")
	n.saveButton.SetSensitive(false)
	n.saveButton.ConnectClicked(func() { n.saveCurrent() })
	header.Append(n.saveButton)

	n.runButton = gtk.NewButtonWithLabel("Run in sandbox")
	n.runButton.AddCSSClass("suggested-action")
	n.runButton.SetSensitive(false)
	n.runButton.ConnectClicked(func() { n.runCurrent() })
	header.Append(n.runButton)
	return header
}

func (n *nativeApp) buildSidebar() *gtk.Widget {
	outer := gtk.NewBox(gtk.OrientationVertical, 0)
	outer.AddCSSClass("sidebar")
	outer.SetMarginStart(12)
	outer.SetMarginEnd(12)
	outer.SetMarginTop(14)
	outer.SetMarginBottom(14)

	heading := gtk.NewLabel("Curriculum")
	heading.AddCSSClass("sidebar-category")
	heading.SetHAlign(gtk.AlignStart)
	outer.Append(heading)

	scroll := gtk.NewScrolledWindow()
	scroll.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	n.kataList = gtk.NewBox(gtk.OrientationVertical, 2)
	n.kataButtons = make(map[string]*gtk.Button)

	for _, stage := range n.track.Stages {
		// Stage header with level badge
		stageBox := gtk.NewBox(gtk.OrientationHorizontal, 8)
		stageBox.SetMarginTop(14)
		stageBox.SetMarginBottom(4)

		stageLabel := gtk.NewLabel(fmt.Sprintf("%s", strings.ToUpper(stage.Title)))
		stageLabel.AddCSSClass("stage-label")
		stageLabel.SetHAlign(gtk.AlignStart)
		stageBox.Append(stageLabel)

		badge := gtk.NewLabel(strings.ToUpper(stage.Level))
		badge.AddCSSClass("level-badge")
		badge.AddCSSClass("level-" + stage.Level)
		stageBox.Append(badge)

		// Stage progress bar
		stageTotal := 0
		for _, cat := range stage.Categories {
			stageTotal += len(cat.Katas)
		}
		stageCompleted := n.stageCompletedCount(stage)
		percent := 0.0
		if stageTotal > 0 {
			percent = float64(stageCompleted) / float64(stageTotal)
		}
		progressBar := gtk.NewProgressBar()
		progressBar.SetFraction(percent)
		progressBar.AddCSSClass("stage-progress")
		progressBar.SetHExpand(true)
		stageBox.Append(progressBar)
		n.stageProgressBars[stage.ID] = progressBar

		n.kataList.Append(stageBox)

		for _, category := range stage.Categories {
			// Category header with progress
			catBox := gtk.NewBox(gtk.OrientationHorizontal, 6)
			catBox.SetMarginTop(6)
			catBox.SetMarginStart(8)

			catLabel := gtk.NewLabel(strings.ToUpper(category.Title))
			catLabel.AddCSSClass("sidebar-category")
			catLabel.SetHAlign(gtk.AlignStart)
			catLabel.SetHExpand(true)
			catBox.Append(catLabel)

			catCompleted := n.categoryCompletedCount(category)
			catPercent := 0.0
			if len(category.Katas) > 0 {
				catPercent = float64(catCompleted) / float64(len(category.Katas))
			}
			catProgress := gtk.NewProgressBar()
			catProgress.SetFraction(catPercent)
			catProgress.AddCSSClass("cat-progress")
			catBox.Append(catProgress)
			n.catProgressBars[category.ID] = catProgress

			n.kataList.Append(catBox)

			for _, kata := range category.Katas {
				// In linear mode, lock katas that aren't available yet
				available := true
				if n.mode == ModeLinear {
					available = n.isKataAvailable(kata, stage, category)
				}

				button := n.kataButton(kata, available)
				if available {
					chosen := kata
					button.ConnectClicked(func() { n.selectKata(chosen) })
				}
				n.kataButtons[kata.ID] = button
				n.kataList.Append(button)
			}
		}
	}

	// ADHD break reminder
	if n.mode == ModeADHD {
		n.breakReminder = gtk.NewLabel("")
		n.breakReminder.AddCSSClass("break-reminder")
		n.breakReminder.SetHAlign(gtk.AlignStart)
		n.kataList.Append(n.breakReminder)
	}

	scroll.SetChild(n.kataList)
	scroll.SetVExpand(true)
	outer.Append(scroll)
	return &outer.Widget
}

func (n *nativeApp) rebuildSidebar() {
	parent := n.kataList.Parent()
	if parent == nil {
		return
	}
	// Remove all children from the box
	for {
		child := n.kataList.FirstChild()
		if child == nil {
			break
		}
		n.kataList.Remove(child)
	}

	n.kataButtons = make(map[string]*gtk.Button)
	n.stageProgressBars = make(map[string]*gtk.ProgressBar)
	n.catProgressBars = make(map[string]*gtk.ProgressBar)
	n.breakReminder = nil

	for _, stage := range n.track.Stages {
		stageBox := gtk.NewBox(gtk.OrientationHorizontal, 8)
		stageBox.SetMarginTop(14)
		stageBox.SetMarginBottom(4)

		stageLabel := gtk.NewLabel(fmt.Sprintf("%s", strings.ToUpper(stage.Title)))
		stageLabel.AddCSSClass("stage-label")
		stageLabel.SetHAlign(gtk.AlignStart)
		stageBox.Append(stageLabel)

		badge := gtk.NewLabel(strings.ToUpper(stage.Level))
		badge.AddCSSClass("level-badge")
		badge.AddCSSClass("level-" + stage.Level)
		stageBox.Append(badge)

		stageTotal := 0
		for _, cat := range stage.Categories {
			stageTotal += len(cat.Katas)
		}
		stageCompleted := n.stageCompletedCount(stage)
		percent := 0.0
		if stageTotal > 0 {
			percent = float64(stageCompleted) / float64(stageTotal)
		}
		progressBar := gtk.NewProgressBar()
		progressBar.SetFraction(percent)
		progressBar.AddCSSClass("stage-progress")
		progressBar.SetHExpand(true)
		stageBox.Append(progressBar)
		n.stageProgressBars[stage.ID] = progressBar

		n.kataList.Append(stageBox)

		for _, category := range stage.Categories {
			catBox := gtk.NewBox(gtk.OrientationHorizontal, 6)
			catBox.SetMarginTop(6)
			catBox.SetMarginStart(8)

			catLabel := gtk.NewLabel(strings.ToUpper(category.Title))
			catLabel.AddCSSClass("sidebar-category")
			catLabel.SetHAlign(gtk.AlignStart)
			catLabel.SetHExpand(true)
			catBox.Append(catLabel)

			catCompleted := n.categoryCompletedCount(category)
			catPercent := 0.0
			if len(category.Katas) > 0 {
				catPercent = float64(catCompleted) / float64(len(category.Katas))
			}
			catProgress := gtk.NewProgressBar()
			catProgress.SetFraction(catPercent)
			catProgress.AddCSSClass("cat-progress")
			catBox.Append(catProgress)
			n.catProgressBars[category.ID] = catProgress

			n.kataList.Append(catBox)

			for _, kata := range category.Katas {
				available := true
				if n.mode == ModeLinear {
					available = n.isKataAvailable(kata, stage, category)
				}
				button := n.kataButton(kata, available)
				if available {
					chosen := kata
					button.ConnectClicked(func() { n.selectKata(chosen) })
				}
				n.kataButtons[kata.ID] = button
				n.kataList.Append(button)
			}
		}
	}

	if n.mode == ModeADHD {
		n.breakReminder = gtk.NewLabel("")
		n.breakReminder.AddCSSClass("break-reminder")
		n.breakReminder.SetHAlign(gtk.AlignStart)
		n.kataList.Append(n.breakReminder)
	}
}

// isKataAvailable checks if a kata should be accessible in linear mode.
// All katas in the first stage are available. In later stages, at least
// one kata from each previous stage must be completed.
func (n *nativeApp) isKataAvailable(kata catalog.Kata, stage catalog.Stage, cat catalog.Category) bool {
	// First stage: all available
	if len(n.track.Stages) == 0 || stage.ID == n.track.Stages[0].ID {
		return true
	}
	// Check that at least some katas from previous stages are completed
	for _, prevStage := range n.track.Stages {
		if prevStage.ID == stage.ID {
			break
		}
		completed := n.stageCompletedCount(prevStage)
		if completed >= 3 { // Require at least 3 completed from previous stage
			return true
		}
	}
	return false
}

func (n *nativeApp) stageCompletedCount(stage catalog.Stage) int {
	count := 0
	for _, cat := range stage.Categories {
		count += n.categoryCompletedCount(cat)
	}
	return count
}

func (n *nativeApp) categoryCompletedCount(cat catalog.Category) int {
	if n.progress == nil {
		return 0
	}
	state, _ := n.progress.Load()
	count := 0
	for _, kata := range cat.Katas {
		if state.Attempts[kata.ID].Passes > 0 {
			count++
		}
	}
	return count
}

func (n *nativeApp) kataButton(kata catalog.Kata, available bool) *gtk.Button {
	button := gtk.NewButton()
	button.AddCSSClass("kata-row")
	button.SetHAlign(gtk.AlignFill)

	row := gtk.NewBox(gtk.OrientationHorizontal, 8)
	dot := gtk.NewLabel("●")
	dot.AddCSSClass("kata-dot")

	// Check completion
	completed := false
	if n.progress != nil {
		state, _ := n.progress.Load()
		completed = state.Attempts[kata.ID].Passes > 0
	}

	switch {
	case completed:
		dot.SetMarkup(`<span color="#34d399">●</span>`)
	case kata.EvaluatorStatus == "ready":
		dot.SetMarkup(`<span color="#fbbf24">●</span>`)
	default:
		dot.SetMarkup(`<span color="#6b7686">●</span>`)
	}

	label := gtk.NewLabel(fmt.Sprintf("%s  %s", kata.ID, kata.Title))
	label.SetHAlign(gtk.AlignStart)
	label.SetHExpand(true)
	row.Append(dot)
	row.Append(label)

	if !available {
		lock := gtk.NewLabel("🔒")
		lock.SetHAlign(gtk.AlignEnd)
		row.Append(lock)
		button.SetSensitive(false)
		button.AddCSSClass("locked")
	}

	button.SetChild(row)
	return button
}

func (n *nativeApp) buildContent() *gtk.Widget {
	outer := gtk.NewBox(gtk.OrientationVertical, 10)
	outer.SetMarginStart(16)
	outer.SetMarginEnd(16)
	outer.SetMarginTop(14)
	outer.SetMarginBottom(16)

	n.title = gtk.NewLabel("Select a kata")
	n.title.AddCSSClass("heading")
	n.title.SetHAlign(gtk.AlignStart)
	n.subtitle = gtk.NewLabel("Choose a kata from the curriculum.")
	n.subtitle.AddCSSClass("heading-sub")
	n.subtitle.SetHAlign(gtk.AlignStart)
	outer.Append(n.title)
	outer.Append(n.subtitle)

	n.stack = gtk.NewStack()
	n.stack.AddNamed(n.buildDocs(), "docs")
	n.stack.AddNamed(n.buildWorkbench(), "workbench")
	n.stack.AddNamed(n.buildFlashcards(), "flashcards")
	n.stack.AddNamed(n.buildQuiz(), "quiz")
	n.stack.AddNamed(n.buildBugHunt(), "bug")
	n.stack.AddNamed(n.buildReflection(), "reflection")
	n.stack.SetVisibleChildName("docs")
	n.stack.SetVExpand(true)

	tabs := gtk.NewStackSwitcher()
	tabs.SetStack(n.stack)
	outer.Append(tabs)
	outer.Append(n.stack)
	return &outer.Widget
}

func (n *nativeApp) buildDocs() *gtk.Widget {
	box := gtk.NewBox(gtk.OrientationVertical, 10)
	scroll := gtk.NewScrolledWindow()
	scroll.SetPolicy(gtk.PolicyAutomatic, gtk.PolicyAutomatic)
	view := gtk.NewTextViewWithBuffer(n.docs)
	view.SetEditable(false)
	view.SetWrapMode(gtk.WrapWordChar)
	view.SetVExpand(true)
	scroll.SetChild(view)
	scroll.SetVExpand(true)
	box.Append(scroll)

	// Start Coding button
	n.startCodingBtn = gtk.NewButtonWithLabel("\u25b6  Start Coding")
	n.startCodingBtn.AddCSSClass("suggested-action")
	n.startCodingBtn.SetTooltipText("Switch to the Workbench to write and test your solution")
	n.startCodingBtn.ConnectClicked(func() {
		n.stack.SetVisibleChildName("workbench")
	})
	btnBox := gtk.NewBox(gtk.OrientationHorizontal, 0)
	btnBox.SetMarginTop(8)
	btnBox.Append(n.startCodingBtn)
	box.Append(btnBox)

	return &box.Widget
}

func (n *nativeApp) buildWorkbench() *gtk.Widget {
	outer := gtk.NewBox(gtk.OrientationVertical, 12)
	paned := gtk.NewPaned(gtk.OrientationHorizontal)
	paned.SetPosition(560)
	paned.SetStartChild(n.editorPane("Solution", n.code, true))
	paned.SetEndChild(n.editorPane("Learner tests", n.learnerTests, true))
	paned.SetVExpand(true)
	outer.Append(paned)

	// Output header with spinner
	outputHeaderBox := gtk.NewBox(gtk.OrientationHorizontal, 8)
	outputHeading := gtk.NewLabel("Output")
	outputHeading.AddCSSClass("panel-title")
	outputHeading.SetHAlign(gtk.AlignStart)
	outputHeaderBox.Append(outputHeading)
	n.outputSpinner = gtk.NewButtonWithLabel("\u23f3 Running\u2026")
	n.outputSpinner.AddCSSClass("warning")
	n.outputSpinner.SetVisible(false)
	n.outputSpinner.SetTooltipText("Sandbox execution in progress\u2026 press Escape to cancel")
	outputHeaderBox.Append(n.outputSpinner)
	outer.Append(outputHeaderBox)

	outputView := gtk.NewTextViewWithBuffer(n.output)
	outputView.SetEditable(false)
	outputView.SetMonospace(true)
	outputView.AddCSSClass("console")
	outputScroll := gtk.NewScrolledWindow()
	outputScroll.SetChild(outputView)
	outputScroll.SetVExpand(true)
	outer.Append(outputScroll)
	return &outer.Widget
}

func (n *nativeApp) editorPane(label string, buffer *gtk.TextBuffer, editable bool) *gtk.Widget {
	box := gtk.NewBox(gtk.OrientationVertical, 6)
	heading := gtk.NewLabel(label)
	heading.AddCSSClass("panel-title")
	heading.SetHAlign(gtk.AlignStart)
	box.Append(heading)
	view := gtk.NewTextViewWithBuffer(buffer)
	view.SetEditable(editable)
	view.SetMonospace(true)
	view.SetWrapMode(gtk.WrapNone)
	view.SetVExpand(true)
	view.SetHExpand(true)
	view.AddCSSClass("editor")
	scroll := gtk.NewScrolledWindow()
	scroll.SetChild(view)
	scroll.SetVExpand(true)
	box.Append(scroll)
	return &box.Widget
}

func (n *nativeApp) selectKata(kata catalog.Kata) {
	n.selected = kata
	for id, button := range n.kataButtons {
		if id == kata.ID {
			button.AddCSSClass("active")
		} else {
			button.RemoveCSSClass("active")
		}
	}
	readme := []byte(kata.Content.Readme)
	starter := []byte(kata.Content.KataGo)
	solution, err := n.workspace.ReadSolution(kata.ID)
	if err != nil {
		n.setStatus(fmt.Sprintf("Unable to read workspace: %v", err))
		return
	}
	if solution == "" {
		solution = string(starter)
	}
	learnerTests, err := n.workspace.ReadLearnerTests(kata.ID)
	if err != nil {
		n.setStatus(fmt.Sprintf("Unable to read learner tests: %v", err))
		return
	}

	n.title.SetText(fmt.Sprintf("%s — %s", kata.ID, kata.Title))
	n.subtitle.SetText(fmt.Sprintf("%s · %s · %s · evaluator: %s", kata.Focus, kata.Signature, strings.ToUpper(kata.EvaluatorStatus), kata.EvaluatorStatus))
	pango := rendering.MarkdownToPango(string(readme))
	n.docs.SetText(pango)
	n.code.SetText(solution)
	n.learnerTests.SetText(learnerTests)
	n.output.SetText("")
	n.flashIndex = 0
	n.quizIndex = 0
	n.quizAnswered = false
	n.updateModes()
	n.saveButton.SetSensitive(true)
	n.runButton.SetSensitive(n.runner != nil && kata.EvaluatorStatus == "ready")
	if n.runner == nil {
		n.setStatus("Podman runner is not configured. Set a digest-pinned runner image in setup.")
	} else if kata.EvaluatorStatus != "ready" {
		n.setStatus("This kata does not yet have a complete trusted evaluator.")
	} else {
		n.setStatus("Ready. Run executes in a disposable rootless Podman container.")
	}
}

func (n *nativeApp) saveCurrent() {
	if n.selected.ID == "" || n.workspace == nil {
		return
	}
	code := n.bufferText(n.code)
	tests := n.bufferText(n.learnerTests)
	if err := n.workspace.SaveSolution(n.selected.ID, code, evaluator.DefaultCodeLimitBytes); err != nil {
		n.setStatus(fmt.Sprintf("Save failed: %v", err))
		return
	}
	if err := n.workspace.SaveLearnerTests(n.selected.ID, tests, evaluator.DefaultTestsLimitBytes); err != nil {
		n.setStatus(fmt.Sprintf("Save failed: %v", err))
		return
	}
	n.setStatus("Saved to the user workspace.")
}

func (n *nativeApp) runCurrent() {
	if n.selected.ID == "" || n.runner == nil || n.running || n.selected.EvaluatorStatus != "ready" {
		return
	}
	n.saveCurrent()
	code := n.bufferText(n.code)
	learnerTests := n.bufferText(n.learnerTests)
	trustedTests := []byte(n.selected.Content.KataTest)

	n.running = true
	n.outputSpinner.SetVisible(true)
	n.runButton.SetSensitive(false)
	n.output.SetText("Running in rootless Podman…")
	kata := n.selected
	runner := n.runner
	go func() {
		result := runner.Run(context.Background(), evaluator.Request{
			KataID:       kata.ID,
			Module:       "kata" + kata.ID,
			Code:         code,
			LearnerTests: learnerTests,
			TrustedTests: string(trustedTests),
		})
		glib.MainContextDefault().InvokeFull(0, func() bool {
			n.running = false
			n.outputSpinner.SetVisible(false)
			n.runButton.SetSensitive(n.runner != nil && kata.EvaluatorStatus == "ready")
			n.output.SetText(formatResult(result))
			if n.progress != nil {
				_, _ = n.progress.RecordAttempt(kata.ID, progress.AttemptResult{
					Passed:      result.Passed,
					Duration:    result.Duration,
					FailedTests: result.FailedTests,
					OutputTail:  result.Output,
					RanAt:       time.Now().UTC(),
				})
				// Update progress bars
				n.updateProgressBars()
			}
			// ADHD break reminder
			if n.mode == ModeADHD {
				n.kataCount++
				if n.kataCount%5 == 0 && n.breakReminder != nil {
					n.breakReminder.SetText("☕ You've done 5 katas — take a break!")
				} else if n.breakReminder != nil {
					n.breakReminder.SetText("")
				}
			}
			return false
		})
	}()
}

func (n *nativeApp) updateProgressBars() {
	for _, stage := range n.track.Stages {
		if pb, ok := n.stageProgressBars[stage.ID]; ok {
			total := 0
			for _, cat := range stage.Categories {
				total += len(cat.Katas)
			}
			completed := n.stageCompletedCount(stage)
			percent := 0.0
			if total > 0 {
				percent = float64(completed) / float64(total)
			}
			pb.SetFraction(percent)
		}
		for _, cat := range stage.Categories {
			if pb, ok := n.catProgressBars[cat.ID]; ok {
				completed := n.categoryCompletedCount(cat)
				percent := 0.0
				if len(cat.Katas) > 0 {
					percent = float64(completed) / float64(len(cat.Katas))
				}
				pb.SetFraction(percent)
			}
		}
	}
}

func (n *nativeApp) bufferText(buffer *gtk.TextBuffer) string {
	return buffer.Text(buffer.StartIter(), buffer.EndIter(), true)
}

func (n *nativeApp) setStatus(status string) {
	if n.status == nil {
		return
	}
	n.status.SetText(status)
	n.status.RemoveCSSClass("ok")
	n.status.RemoveCSSClass("warn")
	lower := strings.ToLower(status)
	switch {
	case strings.Contains(lower, "ready"), strings.Contains(lower, "passed"), strings.Contains(lower, "saved"):
		n.status.AddCSSClass("ok")
	case strings.Contains(lower, "not configured"), strings.Contains(lower, "not yet"),
		strings.Contains(lower, "unavailable"), strings.Contains(lower, "failed"),
		strings.Contains(lower, "missing"), strings.Contains(lower, "error"),
		strings.Contains(lower, "incomplete"):
		n.status.AddCSSClass("warn")
	}
}

func (n *nativeApp) showFatal(window *gtk.ApplicationWindow, message string) {
	window.SetChild(gtk.NewLabel(message))
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
