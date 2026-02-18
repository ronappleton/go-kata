package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/marking"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/runner"
)

const (
	defaultTrackConfigRelative = "tracks/go-core-100/track.json"
	defaultOutputTailLines     = 80
)

type env struct {
	repoRoot string
	track    catalog.Track
	store    *progress.Store
	state    progress.State
}

func main() {
	repoRoot, args, err := parseGlobalArgs(os.Args[1:])
	if err != nil {
		fatalf("%v", err)
	}

	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	loaded, err := loadEnv(repoRoot)
	if err != nil {
		fatalf("load learner environment: %v", err)
	}

	switch args[0] {
	case "list":
		err = cmdList(loaded)
	case "katas":
		err = cmdKatas(loaded, args[1:])
	case "show":
		err = cmdShow(loaded, args[1:])
	case "run":
		err = cmdRun(loaded, args[1:])
	case "mark":
		err = cmdMark(loaded, args[1:])
	case "stats":
		err = cmdStats(loaded)
	default:
		err = fmt.Errorf("unknown command: %s", args[0])
	}

	if err != nil {
		fatalf("%v", err)
	}
}

func parseGlobalArgs(args []string) (string, []string, error) {
	repoRoot := "."
	if len(args) >= 2 && args[0] == "--repo" {
		repoRoot = args[1]
		args = args[2:]
	}
	if len(args) == 0 {
		return "", nil, errors.New("no command provided")
	}
	return repoRoot, args, nil
}

func loadEnv(repoRoot string) (*env, error) {
	absRepoRoot, err := filepath.Abs(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve repo root: %w", err)
	}

	trackPath := filepath.Join(absRepoRoot, defaultTrackConfigRelative)
	track, err := catalog.LoadTrack(trackPath, absRepoRoot)
	if err != nil {
		return nil, err
	}

	store := progress.NewStore(filepath.Join(absRepoRoot, ".learning", "progress.json"))
	state, err := store.Load()
	if err != nil {
		return nil, fmt.Errorf("load progress state: %w", err)
	}

	return &env{
		repoRoot: absRepoRoot,
		track:    track,
		store:    store,
		state:    state,
	}, nil
}

func cmdList(app *env) error {
	fmt.Printf("%s (%s)\n", app.track.Title, app.track.ID)
	fmt.Printf("%s\n\n", app.track.Description)

	totalKatas := 0
	totalCompleted := 0

	for _, category := range app.track.Categories {
		categoryIDs := make([]string, 0, len(category.Katas))
		for _, kata := range category.Katas {
			categoryIDs = append(categoryIDs, kata.ID)
		}

		completed := progress.CompletedCount(app.state, categoryIDs)
		total := len(category.Katas)
		percent := 0.0
		if total > 0 {
			percent = (float64(completed) / float64(total)) * 100
		}

		fmt.Printf("- %s [%s]\n", category.Title, category.ID)
		fmt.Printf("  Goal: %s\n", category.LearningGoal)
		fmt.Printf("  Progress: %d/%d (%.0f%%)\n\n", completed, total, percent)

		totalKatas += total
		totalCompleted += completed
	}

	overall := 0.0
	if totalKatas > 0 {
		overall = (float64(totalCompleted) / float64(totalKatas)) * 100
	}
	fmt.Printf("Overall progress: %d/%d (%.0f%%)\n", totalCompleted, totalKatas, overall)
	return nil
}

func cmdKatas(app *env, args []string) error {
	fs := newFlagSet("katas")
	categoryID := fs.String("category", "", "Category ID (optional)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if strings.TrimSpace(*categoryID) == "" {
		for _, category := range app.track.Categories {
			printCategoryKatas(app, category)
		}
		return nil
	}

	category, ok := app.track.FindCategory(*categoryID)
	if !ok {
		return fmt.Errorf("unknown category: %s", *categoryID)
	}

	printCategoryKatas(app, category)
	return nil
}

func printCategoryKatas(app *env, category catalog.Category) {
	fmt.Printf("%s [%s]\n", category.Title, category.ID)
	for _, kata := range category.Katas {
		status := " "
		if app.state.Attempts[kata.ID].Passes > 0 {
			status = "x"
		}
		fmt.Printf("  [%s] %s  %s (%s)\n", status, kata.ID, kata.Title, kata.Focus)
	}
	fmt.Println()
}

func cmdShow(app *env, args []string) error {
	fs := newFlagSet("show")
	kataID := fs.String("kata", "", "Kata ID, e.g. 001")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*kataID) == "" {
		return errors.New("show requires --kata")
	}

	kata, category, ok := app.track.FindKata(*kataID)
	if !ok {
		return fmt.Errorf("kata not found: %s", *kataID)
	}

	entry := app.state.Attempts[kata.ID]

	fmt.Printf("%s — %s\n", kata.ID, kata.Title)
	fmt.Printf("Category: %s [%s]\n", category.Title, category.ID)
	fmt.Printf("Focus: %s\n", kata.Focus)
	if kata.Signature != "" {
		fmt.Printf("Function: %s\n", kata.Signature)
	}
	fmt.Printf("Directory: %s\n", kata.Dir)
	fmt.Printf("README: %s\n", kata.ReadmePath)
	fmt.Println()

	fmt.Println("Expected behavior:")
	for _, rule := range kata.Rules {
		fmt.Printf("- %s\n", rule)
	}
	fmt.Println()

	fmt.Printf("Progress: attempts=%d, passes=%d, last=%s\n",
		entry.Attempts, entry.Passes, valueOr(entry.LastResult, "n/a"))

	return nil
}

func cmdRun(app *env, args []string) error {
	fs := newFlagSet("run")
	kataID := fs.String("kata", "", "Kata ID, e.g. 001")
	timeoutSec := fs.Int("timeout", 90, "Timeout in seconds")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*kataID) == "" {
		return errors.New("run requires --kata")
	}
	if *timeoutSec <= 0 {
		return errors.New("timeout must be > 0")
	}

	kata, _, ok := app.track.FindKata(*kataID)
	if !ok {
		return fmt.Errorf("kata not found: %s", *kataID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeoutSec)*time.Second)
	defer cancel()

	start := time.Now()
	result, runErr := runner.RunKataTests(ctx, kata.Dir)
	duration := result.Elapsed
	if duration <= 0 {
		duration = time.Since(start)
	}

	outputTail := tailLines(result.RawOutput, defaultOutputTailLines)
	state, err := app.store.RecordAttempt(kata.ID, progress.AttemptResult{
		Passed:      result.Passed,
		Duration:    duration,
		FailedTests: result.FailedTests,
		OutputTail:  outputTail,
		RanAt:       time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("record progress: %w", err)
	}
	app.state = state

	if result.Passed {
		fmt.Printf("PASS %s — %s\n", kata.ID, kata.Title)
		fmt.Printf("Duration: %s\n", duration.Round(time.Millisecond))
		fmt.Println("Well done. Next: run `learner-cli mark --kata " + kata.ID + "` for AI coaching feedback.")
		return nil
	}

	fmt.Printf("FAIL %s — %s\n", kata.ID, kata.Title)
	fmt.Printf("Duration: %s\n", duration.Round(time.Millisecond))
	if len(result.FailedTests) > 0 {
		fmt.Printf("Failing tests: %s\n", strings.Join(result.FailedTests, ", "))
	}
	if runErr != nil {
		fmt.Printf("Runner error: %v\n", runErr)
	}
	if strings.TrimSpace(outputTail) != "" {
		fmt.Println("\nRecent test output:")
		fmt.Println(outputTail)
	}

	return nil
}

func cmdMark(app *env, args []string) error {
	fs := newFlagSet("mark")
	kataID := fs.String("kata", "", "Kata ID, e.g. 001")
	outPath := fs.String("out", "", "Optional output path for prompt markdown")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*kataID) == "" {
		return errors.New("mark requires --kata")
	}

	kata, category, ok := app.track.FindKata(*kataID)
	if !ok {
		return fmt.Errorf("kata not found: %s", *kataID)
	}

	kataCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		return fmt.Errorf("read kata.go: %w", err)
	}
	testCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		return fmt.Errorf("read kata_test.go: %w", err)
	}

	lastAttempt := app.state.Attempts[kata.ID]

	prompt := marking.BuildPrompt(marking.PromptInput{
		TrackTitle:      app.track.Title,
		CategoryTitle:   category.Title,
		KataID:          kata.ID,
		KataTitle:       kata.Title,
		Focus:           kata.Focus,
		Signature:       kata.Signature,
		Rules:           kata.Rules,
		KataCode:        string(kataCode),
		TestCode:        string(testCode),
		LastRunResult:   lastAttempt.LastResult,
		LastFailedTests: lastAttempt.LastFailedTests,
		LastOutputTail:  lastAttempt.LastOutputTail,
	})

	target := strings.TrimSpace(*outPath)
	if target == "" {
		target = marking.DefaultPromptPath(app.repoRoot, kata.ID, time.Now().UTC())
	}

	if err := marking.WritePrompt(target, prompt); err != nil {
		return fmt.Errorf("write prompt: %w", err)
	}

	fmt.Printf("AI marking prompt generated: %s\n", target)
	fmt.Println("Next step: paste this file content into ChatGPT for structured review.")
	return nil
}

func cmdStats(app *env) error {
	totalAttempts := 0
	totalPasses := 0
	for _, entry := range app.state.Attempts {
		totalAttempts += entry.Attempts
		totalPasses += entry.Passes
	}

	allKataIDs := make([]string, 0, len(app.track.AllKatas()))
	for _, kata := range app.track.AllKatas() {
		allKataIDs = append(allKataIDs, kata.ID)
	}
	totalCompleted := progress.CompletedCount(app.state, allKataIDs)

	fmt.Printf("Track: %s\n", app.track.Title)
	fmt.Printf("Katas completed: %d/%d\n", totalCompleted, len(allKataIDs))
	fmt.Printf("Attempts logged: %d\n", totalAttempts)
	fmt.Printf("Passing runs logged: %d\n\n", totalPasses)

	fmt.Println("Per-category progress:")
	for _, category := range app.track.Categories {
		ids := make([]string, 0, len(category.Katas))
		for _, kata := range category.Katas {
			ids = append(ids, kata.ID)
		}
		completed := progress.CompletedCount(app.state, ids)
		fmt.Printf("- %s: %d/%d\n", category.Title, completed, len(ids))
	}

	return nil
}

func newFlagSet(name string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs
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

func valueOr(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func printHelp() {
	fmt.Println("Learner CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  learner-cli [--repo <path>] <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  list                     Show categories and completion")
	fmt.Println("  katas [--category <id>]  Show katas (all or category)")
	fmt.Println("  show --kata <id>         Show kata contract details")
	fmt.Println("  run --kata <id>          Run kata tests and record attempt")
	fmt.Println("  mark --kata <id>         Generate AI marking prompt markdown")
	fmt.Println("  stats                    Show global progress metrics")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  learner-cli list")
	fmt.Println("  learner-cli katas --category foundations")
	fmt.Println("  learner-cli run --kata 045")
	fmt.Println("  learner-cli mark --kata 045")
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
