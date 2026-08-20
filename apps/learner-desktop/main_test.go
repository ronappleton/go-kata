package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveContentRoot(t *testing.T) {
	root := t.TempDir()
	trackPath := filepath.Join(root, "tracks", "go-core-100", "track.json")
	if err := os.MkdirAll(filepath.Dir(trackPath), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(trackPath, []byte(`{"id":"test","categories":[]}`), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := resolveContentRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if got != root {
		t.Fatalf("resolveContentRoot() = %q, want %q", got, root)
	}
}

func TestResolveContentRootRejectsMissingContent(t *testing.T) {
	if _, err := resolveContentRoot(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("expected missing content to be rejected")
	}
}
