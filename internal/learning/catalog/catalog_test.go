package catalog

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNormalizeKataID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "already padded", input: "001", want: "001"},
		{name: "unpadded", input: "7", want: "007"},
		{name: "with spaces", input: " 45 ", want: "045"},
		{name: "invalid text", input: "abc", wantErr: true},
		{name: "out of range", input: "1000", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NormalizeKataID(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestExpandKataIDs(t *testing.T) {
	cfg := categoryConfig{
		KataRanges: []rangeConfig{
			{Start: 1, End: 3},
		},
		KataIDs: []string{"010", "2"},
	}

	got, err := expandKataIDs(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"001", "002", "003", "010"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestParseReadmeMetadata(t *testing.T) {
	tempDir := t.TempDir()
	readmePath := filepath.Join(tempDir, "README.md")

	content := "# Kata 001 — FizzBuzz\n\n" +
		"**Focus:** Basics: loops, conditionals, slices, strconv\n\n" +
		"## Your task\n" +
		"Implement:\n\n" +
		"```go\n" +
		"func FizzBuzz(n int) []string\n" +
		"```\n\n" +
		"## Rules / Expectations\n" +
		"- n<=0 => empty slice (not nil)\n" +
		"- multiples of 3 => Fizz\n"

	if err := os.WriteFile(readmePath, []byte(content), 0o644); err != nil {
		t.Fatalf("write readme: %v", err)
	}

	meta, err := parseReadmeMetadata(readmePath)
	if err != nil {
		t.Fatalf("parseReadmeMetadata returned error: %v", err)
	}

	if meta.title != "FizzBuzz" {
		t.Fatalf("title mismatch: got %q", meta.title)
	}
	if meta.focus != "Basics: loops, conditionals, slices, strconv" {
		t.Fatalf("focus mismatch: got %q", meta.focus)
	}
	if meta.signature != "func FizzBuzz(n int) []string" {
		t.Fatalf("signature mismatch: got %q", meta.signature)
	}

	wantRules := []string{
		"n<=0 => empty slice (not nil)",
		"multiples of 3 => Fizz",
	}
	if !reflect.DeepEqual(meta.rules, wantRules) {
		t.Fatalf("rules mismatch: got %v, want %v", meta.rules, wantRules)
	}
}
