package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
)

func TestCmdListShowsTrackAndProgress(t *testing.T) {
	app := mustTestEnv(t)

	out := captureStdout(t, func() {
		if err := cmdList(app); err != nil {
			t.Fatalf("cmdList returned error: %v", err)
		}
	})

	if !strings.Contains(out, "go-core-100") {
		t.Fatalf("expected track title in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Overall progress:") {
		t.Fatalf("expected overall progress line, got:\n%s", out)
	}
}

func TestCmdShowKataDetails(t *testing.T) {
	app := mustTestEnv(t)

	out := captureStdout(t, func() {
		if err := cmdShow(app, []string{"--kata", "001"}); err != nil {
			t.Fatalf("cmdShow returned error: %v", err)
		}
	})

	if !strings.Contains(out, "001 — Build Greeting") {
		t.Fatalf("expected kata heading in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Expected behavior:") {
		t.Fatalf("expected expected-behavior section, got:\n%s", out)
	}
}

func TestCmdKatasCategoryFilter(t *testing.T) {
	app := mustTestEnv(t)

	out := captureStdout(t, func() {
		if err := cmdKatas(app, []string{"--category", "foundations"}); err != nil {
			t.Fatalf("cmdKatas returned error: %v", err)
		}
	})

	if !strings.Contains(out, "Go Foundations [foundations]") {
		t.Fatalf("expected foundations heading, got:\n%s", out)
	}
	if !strings.Contains(out, "031  FizzBuzz") {
		t.Fatalf("expected kata listing, got:\n%s", out)
	}
}

func TestCmdMarkGeneratesPromptFile(t *testing.T) {
	app := mustTestEnv(t)
	promptPath := filepath.Join(t.TempDir(), "prompt.md")

	out := captureStdout(t, func() {
		if err := cmdMark(app, []string{"--kata", "001", "--out", promptPath}); err != nil {
			t.Fatalf("cmdMark returned error: %v", err)
		}
	})

	if !strings.Contains(out, "AI marking prompt generated:") {
		t.Fatalf("expected success output, got:\n%s", out)
	}

	prompt, err := os.ReadFile(promptPath)
	if err != nil {
		t.Fatalf("read prompt file: %v", err)
	}
	if !strings.Contains(string(prompt), "# Kata Marking Request") {
		t.Fatalf("unexpected prompt content:\n%s", string(prompt))
	}
}

func TestCmdRunInvalidKataReturnsError(t *testing.T) {
	app := mustTestEnv(t)
	err := cmdRun(app, []string{"--kata", "999"})
	if err == nil {
		t.Fatalf("expected error for missing kata")
	}
	if !strings.Contains(err.Error(), "kata not found") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func mustTestEnv(t *testing.T) *env {
	t.Helper()

	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("resolve repo root: %v", err)
	}

	track, err := catalog.LoadTrack(filepath.Join(repoRoot, defaultTrackConfigRelative), repoRoot)
	if err != nil {
		t.Fatalf("load track: %v", err)
	}

	store := progress.NewStore(filepath.Join(t.TempDir(), "progress.json"))
	state, err := store.Load()
	if err != nil {
		t.Fatalf("load temp progress: %v", err)
	}

	return &env{
		repoRoot: repoRoot,
		track:    track,
		store:    store,
		state:    state,
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("create stdout pipe: %v", err)
	}

	os.Stdout = w
	fn()

	_ = w.Close()
	os.Stdout = origStdout
	data, readErr := io.ReadAll(r)
	_ = r.Close()
	if readErr != nil {
		t.Fatalf("read captured stdout: %v", readErr)
	}
	return string(data)
}
