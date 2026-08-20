package catalog

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/katas"
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

func TestLoadTrackFromEmbeddedContent(t *testing.T) {
	// Verify that embedded content has the expected katas
	if len(katas.Content) == 0 {
		t.Fatal("expected embedded katas to be non-empty")
	}

	// Verify a known kata exists with correct metadata
	content, ok := katas.Content["001"]
	if !ok {
		t.Fatal("expected kata 001 in embedded content")
	}
	if content.Slug != "build-greeting" {
		t.Fatalf("expected slug 'build-greeting', got %q", content.Slug)
	}
	if content.KataGo == "" {
		t.Fatal("expected non-empty KataGo")
	}
	if content.KataTest == "" {
		t.Fatal("expected non-empty KataTest")
	}
	if content.Readme == "" {
		t.Fatal("expected non-empty Readme")
	}

	// Parse the embedded JSON metadata
	var meta katas.KataMeta
	if err := json.Unmarshal([]byte(content.JSON), &meta); err != nil {
		t.Fatalf("failed to parse embedded JSON: %v", err)
	}
	if meta.Title != "Build Greeting" {
		t.Fatalf("expected title 'Build Greeting', got %q", meta.Title)
	}
	if meta.EvaluatorStatus != "incomplete" {
		t.Fatalf("expected evaluator status 'incomplete', got %q", meta.EvaluatorStatus)
	}
}

func TestLoadTrackIntegration(t *testing.T) {
	// This test requires the track.json to exist relative to the repo root.
	// It uses the embedded kata content, so no filesystem kata dirs needed.
	track, err := LoadTrack("../../tracks/go-core-100/track.json")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}
	first, _, ok := track.FindKata("001")
	if !ok || first.EvaluatorStatus != "incomplete" {
		t.Fatalf("expected kata 001 to be incomplete, got %+v", first)
	}
	bug, _, ok := track.FindKata("131")
	if !ok || bug.EvaluatorStatus != "ready" {
		t.Fatalf("expected kata 131 to be ready, got %+v", bug)
	}
}

func TestAllKatasReturnsAllEmbedded(t *testing.T) {
	track, err := LoadTrack("../../tracks/go-core-100/track.json")
	if err != nil {
		t.Skipf("skipping: %v", err)
	}
	all := track.AllKatas()
	if len(all) != len(katas.Content) {
		t.Fatalf("expected %d katas, got %d", len(katas.Content), len(all))
	}
}
