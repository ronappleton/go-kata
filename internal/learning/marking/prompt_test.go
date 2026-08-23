package marking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestBuildPromptContainsContext(t *testing.T) {
	in := PromptInput{
		TrackTitle:    "Go Core 100",
		CategoryTitle: "Basics",
		KataID:        "001",
		KataTitle:     "Build Greeting",
		Focus:         "Functions, variables",
		Signature:     "func BuildGreeting(name string) string",
		Rules:         []string{"empty name => stranger", "non-empty => Hello"},
		KataCode:      "package kata\nfunc BuildGreeting(name string) string { return \"\" }",
		TestCode:      "package kata\nfunc TestBuildGreeting(t *testing.T) {}",
	}

	prompt := BuildPrompt(in)

	// Should contain all context fields
	for _, want := range []string{
		"Go Core 100", "Basics", "001", "Build Greeting",
		"Functions, variables", "func BuildGreeting(name string) string",
		"empty name => stranger", "non-empty => Hello",
		"package kata", "Verdict: Pass / Needs Work",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("prompt missing %q", want)
		}
	}
}

func TestBuildPromptWithRunResult(t *testing.T) {
	in := PromptInput{
		TrackTitle:    "Go Core 100",
		CategoryTitle: "Basics",
		KataID:        "001",
		KataTitle:     "Build Greeting",
		Focus:         "Functions",
		LastRunResult: "fail",
		LastFailedTests: []string{"TestEmptyName", "TestTrim"},
		LastOutputTail:  "--- FAIL: TestEmptyName\nexpected \"Hello, stranger!\" got \"\"",
		KataCode:       "package kata\nfunc BuildGreeting(name string) string { return \"\" }",
		TestCode:       "package kata\nfunc TestEmptyName(t *testing.T) {}",
	}

	prompt := BuildPrompt(in)

	for _, want := range []string{
		"Result: FAIL",
		"TestEmptyName", "TestTrim",
		"--- FAIL:", "FAIL:",
	} {
		if !strings.Contains(prompt, want) {
			t.Errorf("prompt missing %q", want)
		}
	}
}

func TestBuildPromptEmptyRules(t *testing.T) {
	in := PromptInput{
		TrackTitle:    "Go Core 100",
		CategoryTitle: "Basics",
		KataID:        "001",
		KataTitle:     "Build Greeting",
		Focus:         "Functions",
		KataCode:      "package kata",
		TestCode:      "package kata",
	}

	prompt := BuildPrompt(in)

	if !strings.Contains(prompt, "Use the README contract") {
		t.Error("prompt should mention README contract when rules are empty")
	}
}

func TestBuildPromptNoSignature(t *testing.T) {
	in := PromptInput{
		TrackTitle:    "Go Core 100",
		CategoryTitle: "Basics",
		KataID:        "001",
		KataTitle:     "Build Greeting",
		Focus:         "Functions",
		KataCode:      "package kata",
		TestCode:      "package kata",
	}

	prompt := BuildPrompt(in)

	if strings.Contains(prompt, "Function contract:") {
		t.Error("prompt should not contain Function contract when signature is empty")
	}
}

func TestBuildPromptNoRunResult(t *testing.T) {
	in := PromptInput{
		TrackTitle:    "Go Core 100",
		CategoryTitle: "Basics",
		KataID:        "001",
		KataTitle:     "Build Greeting",
		Focus:         "Functions",
		KataCode:      "package kata",
		TestCode:      "package kata",
	}

	prompt := BuildPrompt(in)

	if !strings.Contains(prompt, "No run recorded yet") {
		t.Error("prompt should say 'No run recorded yet' when LastRunResult is empty")
	}
}

func TestPromptFileName(t *testing.T) {
	now := time.Date(2025, 3, 15, 14, 30, 45, 0, time.UTC)
	name := PromptFileName("001", now)
	if name != "kata-001-20250315-143045.md" {
		t.Errorf("PromptFileName = %q, want kata-001-20250315-143045.md", name)
	}
}

func TestDefaultPromptPath(t *testing.T) {
	now := time.Date(2025, 3, 15, 14, 30, 45, 0, time.UTC)
	path := DefaultPromptPath("/repo", "001", now)
	if !strings.HasSuffix(path, ".learning/marking/kata-001-20250315-143045.md") {
		t.Errorf("unexpected path: %s", path)
	}
}

func TestDefaultDataPromptPath(t *testing.T) {
	now := time.Date(2025, 3, 15, 14, 30, 45, 0, time.UTC)
	path := DefaultDataPromptPath("/data", "001", now)
	if !strings.HasSuffix(path, "marking/kata-001-20250315-143045.md") {
		t.Errorf("unexpected path: %s", path)
	}
}

func TestWritePromptCreatesFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test-prompt.md")
	content := "# Test prompt\nHello world"

	if err := WritePrompt(path, content); err != nil {
		t.Fatalf("WritePrompt: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(data) != content {
		t.Errorf("file content = %q, want %q", string(data), content)
	}
}

func TestWritePromptCreatesDirectories(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "deep", "prompt.md")

	if err := WritePrompt(path, "content"); err != nil {
		t.Fatalf("WritePrompt: %v", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("file was not created")
	}
}
