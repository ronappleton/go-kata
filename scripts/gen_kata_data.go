//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

var kataDirPattern = regexp.MustCompile(`^kata-(\d{3})-(.+)$`)

// kataMeta holds the full metadata from kata.json (preserving flashcards, quiz, etc.)
type kataMeta struct {
	ID              string            `json:"id"`
	Slug            string            `json:"slug"`
	Title           string            `json:"title"`
	Focus           string            `json:"focus"`
	Signature       string            `json:"signature"`
	Rules           []string          `json:"rules"`
	EvaluatorStatus string            `json:"evaluator_status"`
	Stage           string            `json:"stage,omitempty"`
	Category        string            `json:"category,omitempty"`
	Level           string            `json:"level,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	Prerequisites   []string          `json:"prerequisites,omitempty"`
	EstimatedMin    int               `json:"estimated_minutes,omitempty"`
	Flashcards      []json.RawMessage `json:"flashcards,omitempty"`
	QuizQuestions   []json.RawMessage `json:"quiz_questions,omitempty"`
}

type kataContent struct {
	ID        string
	Slug      string
	KataGo    string
	KataTest  string
	BuggyKata string
	Readme    string
	JSON      string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: go run gen_kata_data.go <katas-root> <output-go-file>\n")
		os.Exit(1)
	}
	katasRoot := os.Args[1]
	outputFile := os.Args[2]

	entries, err := os.ReadDir(katasRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read katas root: %v\n", err)
		os.Exit(1)
	}

	var contents []kataContent

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		matches := kataDirPattern.FindStringSubmatch(entry.Name())
		if len(matches) != 3 {
			continue
		}
		id := matches[1]
		slug := matches[2]
		dir := filepath.Join(katasRoot, entry.Name())

		kataGo := readFile(filepath.Join(dir, "kata.go.txt"))
		kataTest := readFile(filepath.Join(dir, "kata_test.go.txt"))
		buggyKata := readFile(filepath.Join(dir, "buggy_kata.go.txt"))
		readme := readFile(filepath.Join(dir, "README.md"))

		// Read existing kata.json to preserve flashcards, quiz, etc.
		existingMeta := readExistingMeta(filepath.Join(dir, "kata.json"))

		// Parse README metadata
		meta := parseReadmeMetadata(readme, entry.Name())
		evalStatus := inspectEvaluator(kataTest)

		// Build the full kata.json, preserving existing fields
		kj := kataMeta{
			ID:              id,
			Slug:            slug,
			Title:           meta.title,
			Focus:           meta.focus,
			Signature:       meta.signature,
			Rules:           meta.rules,
			EvaluatorStatus: evalStatus,
		}

		// Preserve existing metadata fields
		if existingMeta != nil {
			kj.Stage = existingMeta.Stage
			kj.Category = existingMeta.Category
			kj.Level = existingMeta.Level
			kj.Tags = existingMeta.Tags
			kj.Prerequisites = existingMeta.Prerequisites
			kj.EstimatedMin = existingMeta.EstimatedMin
			kj.Flashcards = existingMeta.Flashcards
			kj.QuizQuestions = existingMeta.QuizQuestions
		}

		jsonBytes, err := json.MarshalIndent(kj, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "marshal json for %s: %v\n", id, err)
			os.Exit(1)
		}
		jsonStr := string(jsonBytes) + "\n"

		jsonPath := filepath.Join(dir, "kata.json")
		if err := os.WriteFile(jsonPath, []byte(jsonStr), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s: %v\n", jsonPath, err)
			os.Exit(1)
		}
		fmt.Printf("wrote %s\n", jsonPath)

		contents = append(contents, kataContent{
			ID:        id,
			Slug:      slug,
			KataGo:    kataGo,
			KataTest:  kataTest,
			BuggyKata: buggyKata,
			Readme:    readme,
			JSON:      jsonStr,
		})
	}

	sort.Slice(contents, func(i, j int) bool {
		return contents[i].ID < contents[j].ID
	})

	if err := generateGoFile(outputFile, contents); err != nil {
		fmt.Fprintf(os.Stderr, "generate go file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s\n", outputFile)
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}

func readExistingMeta(path string) *kataMeta {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var meta kataMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil
	}
	return &meta
}

type readmeMetadata struct {
	title     string
	focus     string
	signature string
	rules     []string
}

func parseReadmeMetadata(readme string, dirName string) readmeMetadata {
	lines := strings.Split(readme, "\n")
	meta := readmeMetadata{title: dirName}

	for _, line := range lines {
		if strings.HasPrefix(line, "# Kata ") {
			meta.title = parseTitle(line)
		}
		if strings.HasPrefix(line, "**Focus:**") {
			meta.focus = strings.TrimSpace(strings.TrimPrefix(line, "**Focus:**"))
		}
	}

	if meta.focus == "" {
		meta.focus = "General Go practice"
	}

	meta.signature = extractSignature(lines)
	meta.rules = extractRules(lines)
	return meta
}

func parseTitle(line string) string {
	parts := strings.SplitN(line, "—", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(strings.TrimPrefix(line, "#"))
}

func extractSignature(lines []string) string {
	inGoBlock := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "```go":
			inGoBlock = true
		case inGoBlock && trimmed == "```":
			return ""
		case inGoBlock && strings.HasPrefix(trimmed, "func "):
			return trimmed
		}
	}
	return ""
}

func extractRules(lines []string) []string {
	rules := []string{}
	inRules := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "## Rules / Expectations":
			inRules = true
		case inRules && strings.HasPrefix(trimmed, "## "):
			return rules
		case inRules && strings.HasPrefix(trimmed, "- "):
			rules = append(rules, strings.TrimSpace(strings.TrimPrefix(trimmed, "- ")))
		}
	}
	return rules
}

func inspectEvaluator(source string) string {
	text := strings.TrimSpace(source)
	if text == "" {
		return "missing"
	}
	if strings.Contains(text, "t.Skip(") {
		for _, marker := range []string{"t.Run(", "t.Error(", "t.Fatal(", "t.Fail(", "t.Errorf(", "t.Fatalf("} {
			if strings.Contains(text, marker) {
				return "ready"
			}
		}
		return "incomplete"
	}
	return "ready"
}

var tmpl = template.Must(template.New("katas").Parse(`// Code generated by scripts/gen_kata_data.go; DO NOT EDIT.

package katas

import "sort"

// KataContent holds the embedded source files for a single kata.
type KataContent struct {
	ID        string
	Slug      string
	KataGo    string
	KataTest  string
	BuggyKata string
	Readme    string
	JSON      string
}

// Content maps kata ID (e.g. "001") to its embedded files.
var Content = map[string]KataContent{
{{- range . }}
	"{{ .ID }}": {
		ID:        "{{ .ID }}",
		Slug:      "{{ .Slug }}",
		KataGo:    {{ printf "%q" .KataGo }},
		KataTest:  {{ printf "%q" .KataTest }},
		BuggyKata: {{ printf "%q" .BuggyKata }},
		Readme:    {{ printf "%q" .Readme }},
		JSON:      {{ printf "%q" .JSON }},
	},
{{- end }}
}

// IDs returns all kata IDs in sorted order.
func IDs() []string {
	ids := make([]string, 0, len(Content))
	for id := range Content {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
`))

func generateGoFile(path string, contents []kataContent) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, contents)
}
