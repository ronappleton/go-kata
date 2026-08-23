//go:build gtk4

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangocairo"
	"github.com/ronappleton/golang-katas-1-100/internal/languages"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/content"
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

	// Language system: registry plus per-editor highlight/check state.
	langRegistry   *languages.Registry
	editorContexts []*editorCtx

	window         *gtk.ApplicationWindow
	headerBox      *gtk.Box
	kataList       *gtk.Box
	title          *gtk.Label
	subtitle       *gtk.Label
	status         *gtk.Label
	downloadBar    *gtk.ProgressBar
	docs           *gtk.TextBuffer
	code           *gtk.TextBuffer
	learnerTests   *gtk.TextBuffer
	output         *gtk.TextBuffer
	reflection     *gtk.TextBuffer
	stack          *gtk.Stack
	runButton      *gtk.Button
	saveButton     *gtk.Button
	flashText      *gtk.Label
	flashSide      *gtk.Label
	flashIndex     int
	quizText       *gtk.Label
	quizFeedback   *gtk.Label
	quizIndex      int
	bugText        *gtk.Label
	outputSpinner  *gtk.Button
	startCodingBtn *gtk.Button
	selected       catalog.Kata
	running        bool
	kataButtons    map[string]*gtk.Button
	// Learning mode
	mode            LearningMode
	modeSelector    *gtk.DropDown
	contentProvider content.ContentManager
	contentMu       sync.Mutex
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

	// Language system: registry drives highlighting, auto-pair, and checking.
	n.langRegistry = languages.NewRegistry()

	if n.config.ContentRoot != "" {
		track, err := catalog.LoadTrack(filepath.Join(n.config.ContentRoot, "tracks", "go-core-100", "track.json"))
		if err != nil {
			n.showFatal(window, fmt.Sprintf("Curriculum could not be loaded:\n\n%v", err))
			return
		}
		n.track = track
	}

	if n.config.ContentRoot == "" {
		contentDir, err := contentCacheDir()
		if err != nil {
			n.showFatal(window, fmt.Sprintf("Curriculum cache could not be resolved:\n\n%v", err))
			return
		}
		provider, err := content.NewProvider(content.ProviderConfig{ContentDir: contentDir, RemoteURL: n.config.ContentURL})
		if err != nil {
			n.showFatal(window, fmt.Sprintf("Curriculum provider could not be created:\n\n%v", err))
			return
		}
		n.contentProvider = provider
	}

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
	} else if n.contentProvider != nil {
		n.setStatus("Downloading curriculum…")
	}
	if n.contentProvider != nil {
		// Report per-kata download progress from the sync worker goroutines.
		if p, ok := n.contentProvider.(interface{ SetProgress(func(int, int)) }); ok {
			p.SetProgress(func(completed, total int) {
				n.postToMain(func() {
					n.updateDownloadBar(fmt.Sprintf("Extracting… %d/%d", completed, total), float64(completed)/float64(total))
				})
			})
		}
		// Report byte-level progress during the archive download.
		if p, ok := n.contentProvider.(interface{ SetDownloadProgress(func(int64, int64)) }); ok {
			p.SetDownloadProgress(func(bytesRead, totalBytes int64) {
				n.postToMain(func() {
					if totalBytes > 0 {
						pct := float64(bytesRead) / float64(totalBytes)
						n.updateDownloadBar(fmt.Sprintf("Downloading… %s / %s", humanBytes(bytesRead), humanBytes(totalBytes)), pct)
					}
				})
			})
		}
		go n.bootstrapRemoteContent(window)
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
			// If podman is not installed, offer to install it
			if !report.PodmanAvailable {
				n.showInstallDialog()
			}
			return false
		})
	}()
}

// showInstallDialog presents a dialog offering to install podman.
func (n *nativeApp) showInstallDialog() {
	plan := diagnostics.DetectInstallPlan()
	if !plan.CanAutoInstall() {
		// Can't auto-install — just show the manual instructions
		dialog := gtk.NewMessageDialog(
			&n.window.Window,
			gtk.DialogModal,
			gtk.MessageInfo,
			gtk.ButtonsClose,
		)
		dialog.SetTitle("Podman Required")
		dialog.SetMarkup("<b>Podman is required</b>\n\n" +
			"Podman is needed to run kata tests in a sandbox.\n\n" +
			"<i>" + strings.ReplaceAll(plan.Notes, "\n", "\n") + "</i>")
				dialog.ConnectResponse(func(response int) {
			dialog.Destroy()
		})
		dialog.Show()
		return
	}

	// Can auto-install — show Install / Cancel dialog
	dialog := gtk.NewMessageDialog(
		&n.window.Window,
		gtk.DialogModal,
		gtk.MessageQuestion,
		gtk.ButtonsNone,
	)
	dialog.SetTitle("Install Podman?")
	dialog.SetMarkup("<b>Podman is required</b>\n\n" +
		"GoKatas needs Podman to run kata tests in a sandbox.\n\n" +
		"The following commands will be run:\n" + formatCommands(plan))
	dialog.AddButton("Install", int(gtk.ResponseAccept))
	dialog.AddButton("Cancel", int(gtk.ResponseCancel))
	dialog.ConnectResponse(func(response int) {
		dialog.Destroy()
		if gtk.ResponseType(response) == gtk.ResponseAccept {
			n.runInstallPlan(plan)
		}
	})
	dialog.Show()
}

// runInstallPlan executes the install commands and shows progress.
func (n *nativeApp) runInstallPlan(plan diagnostics.InstallPlan) {
	n.setStatus("Installing Podman…")
	go func() {
		var allErr []string
		for _, cmd := range plan.Commands {
			n.glibInvoke(func() { n.setStatus("Running: " + cmd) })
			nested := plan.NeedsSudo
			var args []string
			if nested {
				args = []string{"-y", "--"}
			}
			args = append(args, "sh", "-c", cmd)
			var execCmd *exec.Cmd
			if nested {
				execCmd = exec.Command("sudo", args...)
			} else {
				execCmd = exec.Command("sh", "-c", cmd)
			}
			if out, err := execCmd.CombinedOutput(); err != nil {
				allErr = append(allErr, fmt.Sprintf("%s: %v\n%s", cmd, err, string(out)))
			}
		}
		if len(allErr) > 0 {
			errMsg := strings.Join(allErr, "\n")
			n.glibInvoke(func() {
				dialog := gtk.NewMessageDialog(
					&n.window.Window,
					gtk.DialogModal,
					gtk.MessageError,
					gtk.ButtonsClose,
				)
					dialog.SetTitle("Installation Failed")
				dialog.SetMarkup("<b>Podman installation failed</b>\n\n" +
					"Please install Podman manually:\n\n<i>" +
					strings.ReplaceAll(plan.Notes, "\n", "\n") + "</i>\n\n<b>Error:</b>\n" +
					errMsg)
					dialog.ConnectResponse(func(response int) {
						dialog.Destroy()
				})
					dialog.Show()
				n.setStatus("Podman installation failed")
			})
		} else {
			n.glibInvoke(func() {
				dialog := gtk.NewMessageDialog(
					&n.window.Window,
					gtk.DialogModal,
					gtk.MessageInfo,
					gtk.ButtonsClose,
				)
					dialog.SetTitle("Podman Installed")
				dialog.SetMarkup("<b>Podman installed successfully</b>\n\n" +
					"You may need to restart GoKatas for the runner to detect Podman.")
					dialog.ConnectResponse(func(response int) {
						dialog.Destroy()
					// Re-run diagnostics to confirm
						n.startDiagnostics()
					})
					dialog.Show()
				n.setStatus("Podman installed — restarting diagnostics…")
			})
		}
	}()
}

// glibInvoke is a helper to run a function on the main GTK thread.
func (n *nativeApp) glibInvoke(fn func()) {
	glib.MainContextDefault().InvokeFull(0, func() bool {
		fn()
		return false
	})
}

// formatCommands formats the install commands for display in the dialog.
func formatCommands(plan diagnostics.InstallPlan) string {
	var sb strings.Builder
	for _, cmd := range plan.Commands {
		if plan.NeedsSudo {
			sb.WriteString(fmt.Sprintf("  sudo <b>%s</b>\n", cmd))
		} else {
			sb.WriteString(fmt.Sprintf("  <b>%s</b>\n", cmd))
		}
	}
	return sb.String()
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

	modeLabel := gtk.NewLabel("Learning mode")
	modeLabel.AddCSSClass("body")
	header.Append(modeLabel)
	modeSelector := gtk.NewDropDownFromStrings([]string{"Linear · guided progression", "ADHD · flexible quick wins", "Review · recall practice"})
	n.modeSelector = modeSelector
	modeSelector.SetSelected(0)
	modeSelector.SetTooltipText("Choose how GoKatas guides your next step")
	modeSelector.Connect("notify::selected", func() {
		idx := modeSelector.Selected()
		if idx > 2 {
			return
		}
		n.mode = LearningMode(idx)
		n.rebuildSidebar()
		n.buildFlashDeck()
		n.buildQuizDeck()
		n.kataCount = 0
	})
	header.Append(modeSelector)

	// Track selector
	n.trackPaths = make(map[string]string)
	if n.config.ContentRoot != "" {
		n.allTracks, _ = catalog.DiscoverTracks(filepath.Join(n.config.ContentRoot, "tracks"))
	}
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

	n.downloadBar = gtk.NewProgressBar()
	n.downloadBar.SetShowText(true)
	n.downloadBar.SetText("")
	n.downloadBar.AddCSSClass("download-progress")
	n.downloadBar.SetSizeRequest(160, 0)
	n.downloadBar.SetVisible(false)
	header.Append(n.downloadBar)

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
	if len(n.track.Stages) == 0 {
		loading := gtk.NewLabel("Downloading curriculum…")
		loading.AddCSSClass("empty-state")
		loading.SetHAlign(gtk.AlignStart)
		n.kataList.Append(loading)
	}

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
	if len(n.track.Stages) == 0 {
		loading := gtk.NewLabel("Downloading curriculum…")
		loading.AddCSSClass("empty-state")
		loading.SetHAlign(gtk.AlignStart)
		n.kataList.Append(loading)
	}

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

func contentCacheDir() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "gokatas", "content"), nil
}

// bootstrapRemoteContent drives the remote-first curriculum load:
//
//  1. Show any fully cached curriculum immediately (offline-safe fast start).
//     Gated on HasCachedContent so a cold cache never crawls hundreds of
//     katas one request at a time before the zipball sync has run.
//  2. Sync with the remote (zipball by default, per-kata fallback, progress
//     reporting).
//  3. Reload the freshest track and reflect partial failures in the status.
//
// All GTK mutations are marshalled to the main thread; network work happens on
// this goroutine.
func (n *nativeApp) bootstrapRemoteContent(window *gtk.ApplicationWindow) {
	ctx := context.Background()

	// Recover from any panic in this goroutine so the app doesn't
	// silently die and leave the sidebar stuck on "Downloading…".
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[bootstrap] PANIC: %v", r)
			n.postToMain(func() {
				n.updateDownloadBar("", -1)
				n.setStatus(fmt.Sprintf("Curriculum load failed: %v", r))
			})
		}
	}()

	log.Println("[bootstrap] starting")
	n.postToMain(func() { n.setStatus("Checking curriculum…") })

	// 1. Cached curriculum first — usable immediately, even offline. Only
	//    when a previous sync actually populated the cache; a cold cache must
	//    go straight to the zipball sync instead.
	log.Printf("[bootstrap] HasCachedContent=%v", n.contentProvider.HasCachedContent())
	if n.contentProvider.HasCachedContent() {
		if manifest, err := n.contentProvider.GetManifest(ctx); err == nil && len(manifest.Tracks) > 0 {
			log.Printf("[bootstrap] cached manifest has %d tracks", len(manifest.Tracks))
			if cached, err := catalog.LoadTrackFromContent(ctx, n.contentProvider, manifest.Tracks[0].ID); err == nil {
				log.Printf("[bootstrap] cached track loaded: %d stages, %d katas", len(cached.Stages), len(cached.AllKatas()))
				n.postToMain(func() {
					n.applyTrack(cached, "Using cached curriculum · updating…")
				})
			} else {
				log.Printf("[bootstrap] cached track load failed: %v", err)
			}
		} else {
			log.Printf("[bootstrap] cached manifest empty or error: %v", err)
		}
	}

	// 2. Sync with the remote.
	log.Println("[bootstrap] starting sync")
	n.postToMain(func() { n.setStatus("Downloading curriculum…") })
	syncResult, err := n.contentProvider.Sync(ctx)
	if err != nil {
		log.Printf("[bootstrap] sync failed: %v", err)
		// A cached curriculum is still usable when an update check fails.
		if n.hasCurriculum() {
			n.postToMain(func() { n.setStatus("Using cached curriculum · update unavailable") })
			return
		}
		n.postToMain(func() { n.showContentUnavailable(window, err) })
		return
	}
	log.Printf("[bootstrap] sync complete: added=%d failed=%d", syncResult.Added, len(syncResult.Failed))

	// 3. Reload the freshest track.
	log.Println("[bootstrap] loading manifest")
	manifest, err := n.contentProvider.GetManifest(ctx)
	if err != nil || len(manifest.Tracks) == 0 {
		log.Printf("[bootstrap] manifest load failed: err=%v tracks=%d", err, len(manifest.Tracks))
		if err == nil {
			err = fmt.Errorf("content manifest contains no tracks")
		}
		if !n.hasCurriculum() {
			n.postToMain(func() { n.showContentUnavailable(window, err) })
		}
		return
	}
	log.Printf("[bootstrap] manifest tracks: %d", len(manifest.Tracks))

	log.Printf("[bootstrap] loading track %q", manifest.Tracks[0].ID)
	track, loadErr := catalog.LoadTrackFromContent(ctx, n.contentProvider, manifest.Tracks[0].ID)
	log.Printf("[bootstrap] track loaded: %d stages, %d katas, err=%v", len(track.Stages), len(track.AllKatas()), loadErr)
	if loadErr != nil && len(track.Stages) == 0 {
		if !n.hasCurriculum() {
			n.postToMain(func() { n.showContentUnavailable(window, loadErr) })
		}
		return
	}

	status := fmt.Sprintf("Curriculum ready · %d katas", len(track.AllKatas()))
	if loadErr != nil {
		status += " · some katas unavailable"
	} else if len(syncResult.Failed) > 0 {
		status += fmt.Sprintf(" · %d sync warnings", len(syncResult.Failed))
	}
	log.Printf("[bootstrap] posting applyTrack: %s", status)
	n.postToMain(func() {
		log.Println("[bootstrap] applyTrack executing on main thread")
		n.applyTrack(track, status)
		log.Println("[bootstrap] applyTrack done")
	})
}

// postToMain schedules fn on the GTK main loop from any goroutine.
func (n *nativeApp) postToMain(fn func()) {
	glib.MainContextDefault().InvokeFull(0, func() bool {
		fn()
		return false
	})
}

// applyTrack installs a loaded track into the UI (main thread only).
func (n *nativeApp) applyTrack(track catalog.Track, status string) {
	n.contentMu.Lock()
	n.track = track
	n.contentMu.Unlock()
	n.rebuildSidebar()
	all := n.track.AllKatas()
	if len(all) > 0 && n.selected.ID == "" {
		n.selectKata(all[0])
	}
	n.updateDownloadBar("", -1) // hide the progress bar
	n.setStatus(status)
}

func (n *nativeApp) hasCurriculum() bool {
	n.contentMu.Lock()
	defer n.contentMu.Unlock()
	return len(n.track.AllKatas()) > 0
}

func (n *nativeApp) showContentUnavailable(window *gtk.ApplicationWindow, err error) {
	if n.status != nil {
		n.setStatus("Curriculum unavailable · check connection and retry")
	}
	if n.title != nil {
		n.title.SetText("Curriculum unavailable")
	}
	if n.subtitle != nil {
		n.subtitle.SetText(err.Error())
	}
	if n.docs != nil {
		n.docs.SetText("Curriculum is not available yet. Check your connection, then restart GoKatas to retry the download.")
	}
	if n.kataList != nil && !n.hasCurriculum() {
		n.rebuildSidebar()
	}
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
	view.AddCSSClass("docs-view")
	view.SetEditable(false)
	view.SetWrapMode(gtk.WrapWordChar)
	view.SetVExpand(true)
	scroll.SetChild(view)
	scroll.SetVExpand(true)
	box.Append(scroll)

	// Start Coding button
	n.startCodingBtn = gtk.NewButtonWithLabel("\u25b6  Start Coding")
	n.startCodingBtn.AddCSSClass("suggested-action")
	n.startCodingBtn.AddCSSClass("start-coding-btn")
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

// editorCtx is the state attached to one editable code pane: its highlight and
// error tags, the currently selected language, and a debounced checker.
type editorCtx struct {
	buffer  *gtk.TextBuffer
	view    *gtk.TextView
	tags    map[string]*gtk.TextTag
	errTag  *gtk.TextTag
	lang    *languages.Language
	checker languages.Checker
	checkID uint // debounce generation counter
}

// editorPane builds a code editor: a line-number gutter beside the text view,
// language-driven syntax highlighting, auto-pairing, and (for editable panes)
// syntax error indication. Language can be changed later via setEditorLanguage.
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
	// Padding so the cursor never sits flush against the edges.
	view.SetLeftMargin(16)
	view.SetRightMargin(16)
	view.SetTopMargin(10)
	view.SetBottomMargin(10)

	scroll := gtk.NewScrolledWindow()
	scroll.SetChild(view)
	scroll.SetVExpand(true)

	// Line-number gutter, drawn on demand and kept in sync with scrolling.
	gutter := gtk.NewDrawingArea()
	gutter.SetSizeRequest(52, -1)
	gutter.AddCSSClass("editor-gutter")
	gutter.SetDrawFunc(func(area *gtk.DrawingArea, cr *cairo.Context, width, height int) {
		n.drawLineNumbers(area, cr, width, height, view, buffer, scroll)
	})
	adj := scroll.VAdjustment()
	adj.Connect("value-changed", func() { gutter.QueueDraw() })
	buffer.ConnectChanged(func() { gutter.QueueDraw() })

	editorBox := gtk.NewBox(gtk.OrientationHorizontal, 0)
	editorBox.Append(gutter)
	editorBox.Append(scroll)
	editorBox.SetVExpand(true)
	box.Append(editorBox)

	if editable {
		ctx := &editorCtx{buffer: buffer, view: view, tags: make(map[string]*gtk.TextTag)}
		ctx.lang = n.langRegistry.Default()
		n.installEditorTags(ctx)
		n.editorContexts = append(n.editorContexts, ctx)
		// Auto-pair + smart indent + backspace handling.
		n.installEditorKeys(ctx)
	}
	return &box.Widget
}

// setEditorLanguage swaps the highlighting/checker language for every editor.
func (n *nativeApp) setEditorLanguage(lang *languages.Language) {
	if lang == nil {
		lang = n.langRegistry.Default()
	}
	for _, ctx := range n.editorContexts {
		ctx.lang = lang
		ctx.checker = lang.Checker
		n.rehighlight(ctx)
		n.scheduleCheck(ctx)
	}
}

// installEditorTags creates the highlight and error text tags for an editor.
func (n *nativeApp) installEditorTags(ctx *editorCtx) {
	table := ctx.buffer.TagTable()
	for name, fg := range map[string]string{
		languages.TagKeyword: "#ff7b72",
		languages.TagString:  "#a5d6ff",
		languages.TagComment: "#8b949e",
		languages.TagNumber:  "#79c0ff",
		languages.TagType:    "#7ee787",
		languages.TagFunc:    "#d2a8ff",
	} {
		tag := gtk.NewTextTag(name)
		tag.SetObjectProperty("foreground", fg)
		table.Add(tag)
		ctx.tags[name] = tag
	}
	// Syntax error squiggle.
	errTag := gtk.NewTextTag("hl-error")
	errTag.SetObjectProperty("foreground", "#f87171")
	errTag.SetObjectProperty("underline", int(pango.UnderlineError))
	table.Add(errTag)
	ctx.errTag = errTag

	ctx.buffer.ConnectChanged(func() {
		n.rehighlight(ctx)
		n.scheduleCheck(ctx)
	})
	n.rehighlight(ctx)
}

// rehighlight re-lexes the editor buffer with the current language.
func (n *nativeApp) rehighlight(ctx *editorCtx) {
	start, end := ctx.buffer.Bounds()
	for _, tag := range ctx.tags {
		ctx.buffer.RemoveTag(tag, start, end)
	}
	spans := languages.Lex(ctx.buffer.Text(start, end, false), ctx.lang)
	for _, s := range spans {
		if tag, ok := ctx.tags[s.Tag]; ok {
			ts := ctx.buffer.IterAtOffset(s.Start)
			te := ctx.buffer.IterAtOffset(s.End)
			ctx.buffer.ApplyTag(tag, ts, te)
		}
	}
}

// scheduleCheck runs the language syntax checker after a short debounce.
func (n *nativeApp) scheduleCheck(ctx *editorCtx) {
	if ctx.checker == nil {
		return
	}
	ctx.checkID++
	gen := ctx.checkID
	glib.TimeoutAdd(600, func() bool {
		if gen != ctx.checkID {
			return false // superseded by a newer edit
		}
		src := ctx.buffer.Text(ctx.buffer.StartIter(), ctx.buffer.EndIter(), false)
		go func() {
			diags := ctx.checker.Check(src)
			glib.MainContextDefault().InvokeFull(0, func() bool {
				if gen != ctx.checkID {
					return false
				}
				n.applyDiagnostics(ctx, diags)
				return false
			})
		}()
		return false
	})
}

// applyDiagnostics paints error squiggles for the given diagnostics.
func (n *nativeApp) applyDiagnostics(ctx *editorCtx, diags []languages.Diagnostic) {
	start, end := ctx.buffer.Bounds()
	ctx.buffer.RemoveTag(ctx.errTag, start, end)
	for _, d := range diags {
		ts, ok := ctx.buffer.IterAtLineOffset(d.Line, d.Col)
		if !ok {
			ts = ctx.buffer.EndIter()
		}
		var te *gtk.TextIter
		if d.EndLine >= 0 {
			var okEnd bool
			te, okEnd = ctx.buffer.IterAtLineOffset(d.EndLine, d.EndCol)
			if !okEnd {
				te = ctx.buffer.EndIter()
			}
		} else {
			te = ts.Copy()
			te.ForwardChar()
		}
		ctx.buffer.ApplyTag(ctx.errTag, ts, te)
	}
}

// installEditorKeys wires auto-pairing and smart indentation for an editable
// editor: typing an opening char inserts its closer, Enter inside an empty pair
// puts the closer on the next line, Backspace deletes an empty pair, and typing
// a closer that already follows the cursor skips over it.
func (n *nativeApp) installEditorKeys(ctx *editorCtx) {
	keyCtrl := gtk.NewEventControllerKey()
	keyCtrl.ConnectKeyPressed(func(keyval, keycode uint, state gdk.ModifierType) bool {
		// Let shortcuts (Ctrl/Cmd) through untouched.
		if state&(gdk.ControlMask|gdk.AltMask|gdk.MetaMask) != 0 {
			return false
		}
		buffer := ctx.buffer
		cursor := buffer.IterAtMark(buffer.GetInsert())

		if keyval == gdk.KEY_Return {
			return n.handleEnter(ctx, cursor)
		}
		if keyval == gdk.KEY_BackSpace {
			return n.handleBackspace(ctx, cursor)
		}

		// Only single-char pairs participate (printable ASCII keyvals).
		if keyval > 0x7F {
			return false
		}
		ch := rune(keyval)

		// Typing a closer that already follows the cursor: skip over it.
		if closer, ok := n.closeForOpen(ch); ok {
			// Skip only when the next char is exactly the closer.
			next := charAt(buffer, cursor)
			if next == closer {
				cursor.ForwardChar()
				buffer.PlaceCursor(cursor)
				return true
			}
			return false
		}

		// Typing an opener: insert open+close, cursor between.
		if closer, ok := n.openForClose(ch); ok {
			buffer.Insert(cursor, string(ch))
			cursor = buffer.IterAtMark(buffer.GetInsert())
			buffer.Insert(cursor, string(closer))
			cursor.BackwardChar()
			buffer.PlaceCursor(cursor)
			return true
		}
		return false
	})
	ctx.view.AddController(keyCtrl)
}

// handleEnter implements smart indentation. Inside an empty pair it moves the
// closer to a new line (the "closing brace on the next line" behaviour);
// otherwise it indents to match the previous line.
func (n *nativeApp) handleEnter(ctx *editorCtx, cursor *gtk.TextIter) bool {
	buffer := ctx.buffer
	prev := cursor.Copy()
	prev.BackwardChar()

	if prev.Char() != 0 && prev.Char() != '\n' {
		closer, isOpen := n.openForClose(rune(prev.Char()))
		next := charAt(buffer, cursor)
		if isOpen && next == closer {
			indent := currentLineIndent(buffer, cursor)
			inner := indent + ctx.lang.Indent
			buffer.Insert(cursor, "\n"+inner)
			cursor = buffer.IterAtMark(buffer.GetInsert())
			buffer.Insert(cursor, "\n"+indent+string(closer))
			cursor = buffer.IterAtMark(buffer.GetInsert())
			cursor.BackwardChars(1 + len(indent))
			buffer.PlaceCursor(cursor)
			return true
		}
	}

	// Plain enter: indent the new line to match the previous line.
	indent := currentLineIndent(buffer, cursor)
	buffer.Insert(cursor, "\n"+indent)
	cursor = buffer.IterAtMark(buffer.GetInsert())
	buffer.PlaceCursor(cursor)
	return true
}

// handleBackspace deletes an empty auto-paired pair in one keystroke.
func (n *nativeApp) handleBackspace(ctx *editorCtx, cursor *gtk.TextIter) bool {
	buffer := ctx.buffer
	if cursor.Offset() == 0 {
		return false
	}
	prev := cursor.Copy()
	prev.BackwardChar()
	opener := prev.Char()
	if opener == 0 {
		return false
	}
	closer, isOpen := n.openForClose(rune(opener))
	if isOpen && charAt(buffer, cursor) == closer {
		// Delete the closer first, then the opener.
		end := cursor.Copy()
		end.ForwardChar()
		buffer.Delete(cursor, end)
		start := buffer.IterAtOffset(cursor.Offset() - 1)
		end2 := buffer.IterAtOffset(cursor.Offset())
		buffer.Delete(start, end2)
		buffer.PlaceCursor(buffer.IterAtOffset(cursor.Offset() - 1))
		return true
	}
	return false
}

// openForClose returns the opening char that pairs with ch, if ch is an
// opener in the current language's auto-pair table.
func (n *nativeApp) openForClose(ch rune) (rune, bool) {
	if ctx := n.activeEditorCtx(); ctx != nil && ctx.lang != nil {
		for _, p := range ctx.lang.AutoPairs {
			if p.Open == ch {
				return p.Close, true
			}
		}
	}
	return 0, false
}

// closeForOpen reports whether ch is a closer in the current language's
// auto-pair table.
func (n *nativeApp) closeForOpen(ch rune) (rune, bool) {
	if ctx := n.activeEditorCtx(); ctx != nil && ctx.lang != nil {
		for _, p := range ctx.lang.AutoPairs {
			if p.Close == ch {
				return p.Close, true
			}
		}
	}
	return 0, false
}

// activeEditorCtx returns the most recently created editor context (the
// solution pane) so pair lookup has a language to consult.
func (n *nativeApp) activeEditorCtx() *editorCtx {
	if len(n.editorContexts) == 0 {
		return nil
	}
	return n.editorContexts[0]
}

// charAt returns the character at iter, or 0 at end of buffer.
func charAt(buffer *gtk.TextBuffer, iter *gtk.TextIter) rune {
	if iter == nil || iter.Offset() >= buffer.CharCount() {
		return 0
	}
	return rune(iter.Char())
}

// currentLineIndent returns the leading whitespace of the line containing iter.
func currentLineIndent(buffer *gtk.TextBuffer, iter *gtk.TextIter) string {
	lineStart := iter.Copy()
	lineStart.SetLineOffset(0)
	lineEnd := iter.Copy()
	lineEnd.ForwardToLineEnd()
	line := buffer.Slice(lineStart, lineEnd, false)
	i := 0
	for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
		i++
	}
	return line[:i]
}

// drawLineNumbers paints the current line numbers into the gutter, aligned with
// the text view's visible lines and scrolled offset.
func (n *nativeApp) drawLineNumbers(area *gtk.DrawingArea, cr *cairo.Context, width, height int, view *gtk.TextView, buffer *gtk.TextBuffer, scroll *gtk.ScrolledWindow) {
	// Gutter background.
	cr.SetSourceRGB(0.051, 0.067, 0.09)
	cr.Rectangle(0, 0, float64(width), float64(height))
	cr.Fill()

	if view == nil || buffer == nil || scroll == nil {
		return
	}

	layout := pangocairo.CreateLayout(cr)
	layout.SetFontDescription(n.editorFont(view))
	layout.SetText("0")
	_, lineHeight := layout.PixelSize()
	if lineHeight <= 0 {
		lineHeight = 18
	}
	topMargin := 10
	first := int(scroll.VAdjustment().Value()) / lineHeight
	total := buffer.LineCount()
	visible := height/lineHeight + 2

	cr.SetSourceRGB(0.42, 0.47, 0.55)
	for i := first; i < first+visible && i < total; i++ {
		layout.SetText(strconv.Itoa(i + 1))
		w, _ := layout.PixelSize()
		y := float64(i*lineHeight+topMargin) - scroll.VAdjustment().Value()
		cr.MoveTo(float64(width)-float64(w)-12, y)
		pangocairo.ShowLayout(cr, layout)
	}
}

// editorFont returns the monospace font description used by the editor views,
// falling back to a default if the widget is not realized yet.
func (n *nativeApp) editorFont(view *gtk.TextView) *pango.FontDescription {
	if ctx := view.PangoContext(); ctx != nil {
		if fd := ctx.FontDescription(); fd != nil {
			return fd
		}
	}
	fd := pango.NewFontDescription()
	fd.SetFamily("monospace")
	fd.SetSize(13 * 1024)
	return fd
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

	// Resolve the kata's language (defaults to Go) and its workspace filenames.
	lang := n.langRegistry.Lookup(kata.Language)
	if lang == nil {
		lang = n.langRegistry.Default()
	}
	n.setEditorLanguage(lang)

	solution, err := n.workspace.ReadSolutionAs(kata.ID, lang.SourceFilename)
	if err != nil {
		n.setStatus(fmt.Sprintf("Unable to read workspace: %v", err))
		return
	}
	if solution == "" {
		solution = string(starter)
	}
	learnerTests, err := n.workspace.ReadLearnerTestsAs(kata.ID, lang.TestsFilename)
	if err != nil {
		n.setStatus(fmt.Sprintf("Unable to read learner tests: %v", err))
		return
	}

	n.title.SetText(fmt.Sprintf("%s — %s", kata.ID, kata.Title))
	langLabel := lang.Name
	if lang.ID != "go" {
		langLabel += " · sandbox runner currently Go-only"
	}
	n.subtitle.SetText(fmt.Sprintf("%s · %s · %s · %s · evaluator: %s", kata.Focus, langLabel, kata.Signature, strings.ToUpper(kata.EvaluatorStatus), kata.EvaluatorStatus))
	pango := rendering.MarkdownToPango(string(readme))
	n.docs.SetText("")
	n.docs.InsertMarkup(n.docs.EndIter(), pango)
	n.code.SetText(solution)
	n.learnerTests.SetText(learnerTests)
	n.output.SetText("")
	n.flashIndex = 0
	n.quizIndex = 0
	n.quizAnswered = false
	n.updateModes()
	n.saveButton.SetSensitive(true)
	canRun := n.runner != nil && kata.EvaluatorStatus == "ready" && (lang == nil || lang.ID == "go")
	n.runButton.SetSensitive(canRun)
	switch {
	case n.runner == nil:
		n.setStatus("Podman runner is not configured. Set a digest-pinned runner image in setup.")
	case kata.EvaluatorStatus != "ready":
		n.setStatus("This kata does not yet have a complete trusted evaluator.")
	case lang != nil && lang.ID != "go":
		n.setStatus(fmt.Sprintf("%s katas are editable and highlighted now; the sandbox runner currently supports Go only.", lang.Name))
	default:
		n.setStatus("Ready. Run executes in a disposable rootless Podman container.")
	}
}

func (n *nativeApp) saveCurrent() {
	if n.selected.ID == "" || n.workspace == nil {
		return
	}
	code := n.bufferText(n.code)
	tests := n.bufferText(n.learnerTests)
	lang := n.langRegistry.Lookup(n.selected.Language)
	if lang == nil {
		lang = n.langRegistry.Default()
	}
	if err := n.workspace.SaveSolutionAs(n.selected.ID, code, evaluator.DefaultCodeLimitBytes, lang.SourceFilename); err != nil {
		n.setStatus(fmt.Sprintf("Save failed: %v", err))
		return
	}
	if err := n.workspace.SaveLearnerTestsAs(n.selected.ID, tests, evaluator.DefaultTestsLimitBytes, lang.TestsFilename); err != nil {
		n.setStatus(fmt.Sprintf("Save failed: %v", err))
		return
	}
	n.setStatus("Saved to the user workspace.")
}

func (n *nativeApp) runCurrent() {
	lang := n.langRegistry.Lookup(n.selected.Language)
	if lang != nil && lang.ID != "go" {
		n.setStatus(fmt.Sprintf("%s execution isn't wired to the sandbox runner yet.", lang.Name))
		return
	}
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

// updateDownloadBar shows the download progress bar with the given text and
// fraction, or hides it when fraction < 0. It must be called on the GTK main
// thread.
func (n *nativeApp) updateDownloadBar(text string, fraction float64) {
	if n.downloadBar == nil {
		return
	}
	if fraction < 0 {
		n.downloadBar.SetVisible(false)
		return
	}
	n.downloadBar.SetVisible(true)
	n.downloadBar.SetText(text)
	n.downloadBar.SetFraction(fraction)
}

// humanBytes formats a byte count as a human-readable string.
func humanBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
