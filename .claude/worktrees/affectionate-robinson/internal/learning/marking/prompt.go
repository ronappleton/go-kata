package marking

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PromptInput struct {
	TrackTitle      string
	CategoryTitle   string
	KataID          string
	KataTitle       string
	Focus           string
	Signature       string
	Rules           []string
	KataCode        string
	TestCode        string
	LastRunResult   string
	LastFailedTests []string
	LastOutputTail  string
}

func BuildPrompt(in PromptInput) string {
	var b strings.Builder

	b.WriteString("# Kata Marking Request\n\n")
	b.WriteString("You are reviewing a learner's kata solution. Mark it as a coach, not just a test checker.\n\n")

	b.WriteString("## Context\n")
	b.WriteString(fmt.Sprintf("- Track: %s\n", in.TrackTitle))
	b.WriteString(fmt.Sprintf("- Category: %s\n", in.CategoryTitle))
	b.WriteString(fmt.Sprintf("- Kata: %s — %s\n", in.KataID, in.KataTitle))
	b.WriteString(fmt.Sprintf("- Focus: %s\n", in.Focus))
	if in.Signature != "" {
		b.WriteString(fmt.Sprintf("- Function contract: `%s`\n", in.Signature))
	}
	b.WriteString("\n")

	b.WriteString("## Required Behavior\n")
	if len(in.Rules) == 0 {
		b.WriteString("- Use the README contract in this kata folder.\n")
	} else {
		for _, rule := range in.Rules {
			b.WriteString(fmt.Sprintf("- %s\n", rule))
		}
	}
	b.WriteString("\n")

	b.WriteString("## Last Test Run\n")
	if in.LastRunResult == "" {
		b.WriteString("- No run recorded yet.\n")
	} else {
		b.WriteString(fmt.Sprintf("- Result: %s\n", strings.ToUpper(in.LastRunResult)))
		if len(in.LastFailedTests) > 0 {
			b.WriteString(fmt.Sprintf("- Failing tests: %s\n", strings.Join(in.LastFailedTests, ", ")))
		}
	}
	if strings.TrimSpace(in.LastOutputTail) != "" {
		b.WriteString("\n```text\n")
		b.WriteString(strings.TrimSpace(in.LastOutputTail))
		b.WriteString("\n```\n")
	}
	b.WriteString("\n")

	b.WriteString("## Learner Code (`kata.go`)\n")
	b.WriteString("```go\n")
	b.WriteString(strings.TrimSpace(in.KataCode))
	b.WriteString("\n```\n\n")

	b.WriteString("## Learner Tests (`kata_test.go`)\n")
	b.WriteString("```go\n")
	b.WriteString(strings.TrimSpace(in.TestCode))
	b.WriteString("\n```\n\n")

	b.WriteString("## Feedback Format (please follow)\n")
	b.WriteString("1. Verdict: Pass / Needs Work\n")
	b.WriteString("2. Contract coverage: which required behaviors are satisfied, missing, or ambiguous\n")
	b.WriteString("3. Go quality: correctness, edge-case handling, readability, and idiomatic style\n")
	b.WriteString("4. Test quality: what tests are strong, missing, or redundant\n")
	b.WriteString("5. Next iteration plan: 3 concrete changes, highest impact first\n")

	return b.String()
}

func DefaultPromptPath(repoRoot, kataID string, now time.Time) string {
	fileName := fmt.Sprintf("kata-%s-%s.md", kataID, now.Format("20060102-150405"))
	return filepath.Join(repoRoot, ".learning", "marking", fileName)
}

func WritePrompt(path, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}
