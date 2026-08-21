package content

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestArchiveURLFromRaw(t *testing.T) {
	cases := map[string]string{
		"https://raw.githubusercontent.com/ronappleton/gokatas-content/main":  "https://codeload.github.com/ronappleton/gokatas-content/zip/refs/heads/main",
		"https://raw.githubusercontent.com/a/b/feature%2Fbranch":              "https://codeload.github.com/a/b/zip/refs/heads/feature%2Fbranch",
		"https://raw.githubusercontent.com/ronappleton/gokatas-content/main/": "https://codeload.github.com/ronappleton/gokatas-content/zip/refs/heads/main/",
		"https://example.com/not/raw":                                         "",
		"https://raw.githubusercontent.com/just-owner":                        "",
	}
	for in, want := range cases {
		if got := archiveURLFromRaw(in); got != want {
			t.Errorf("archiveURLFromRaw(%q) = %q, want %q", in, got, want)
		}
	}
}

func makeZip(t *testing.T, entries map[string]string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for name, content := range entries {
		w, err := zw.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestExtractArchiveStripsRoot(t *testing.T) {
	dir := t.TempDir()
	data := makeZip(t, map[string]string{
		"gokatas-content-main/manifest.json":                     `{"version":"1"}`,
		"gokatas-content-main/tracks/go-core-100/track.json":     `{"id":"go-core-100"}`,
		"gokatas-content-main/tracks/go-core-100/katas/000.json": `{"id":"000"}`,
		"gokatas-content-main/tracks/go-core-100/katas/001.json": `{"id":"001"}`,
		"gokatas-content-main/README.md":                         "readme",
		"gokatas-content-main/tracks/other/katas/002.json":       `{"id":"002"}`,
	})

	var progress []int
	added, err := extractArchive(data, dir, func(done, total int) { progress = append(progress, done) })
	if err != nil {
		t.Fatal(err)
	}
	if added != 3 {
		t.Fatalf("expected 3 katas extracted, got %d", added)
	}
	for _, want := range []string{
		"manifest.json",
		"tracks/go-core-100/track.json",
		"tracks/go-core-100/katas/000.json",
		"tracks/go-core-100/katas/001.json",
	} {
		if _, err := os.Stat(filepath.Join(dir, want)); err != nil {
			t.Errorf("expected %s to exist: %v", want, err)
		}
	}
	if len(progress) != 3 || progress[2] != 3 {
		t.Fatalf("unexpected progress: %v", progress)
	}
	// README.md must not be extracted (irrelevant content).
	if _, err := os.Stat(filepath.Join(dir, "README.md")); err == nil {
		t.Fatal("README.md should not be extracted")
	}
}

func TestExtractArchiveRejectsPathTraversal(t *testing.T) {
	dir := t.TempDir()
	data := makeZip(t, map[string]string{
		"gokatas-content-main/../../evil.json": `{"id":"evil"}`,
	})
	if _, err := extractArchive(data, dir, nil); err == nil {
		t.Fatal("expected zip-slip path to be rejected")
	}
	if _, err := os.Stat(filepath.Join(filepath.Dir(dir), "evil.json")); err == nil {
		t.Fatal("evil.json must not be written outside the content dir")
	}
}

func TestIsKataFile(t *testing.T) {
	cases := map[string]bool{
		"tracks/go-core-100/katas/000.json":          true,
		"gokatas-content-main/tracks/x/katas/1.json": true,
		"tracks/go-core-100/track.json":              false,
		"manifest.json":                              false,
		"tracks/go-core-100/katas/notes.md":          false,
	}
	for name, want := range cases {
		if got := isKataFile(name); got != want {
			t.Errorf("isKataFile(%q) = %v, want %v", name, got, want)
		}
	}
}

func TestStripArchiveRoot(t *testing.T) {
	cases := map[string]string{
		"gokatas-content-main/manifest.json":       "manifest.json",
		"gokatas-content-main/tracks/x/track.json": "tracks/x/track.json",
		"manifest.json":       "manifest.json",
		"tracks/x/track.json": "tracks/x/track.json",
	}
	for in, want := range cases {
		if got := stripArchiveRoot(in); got != want {
			t.Errorf("stripArchiveRoot(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSafeJoin(t *testing.T) {
	dir := "/tmp/content"
	if _, err := safeJoin(dir, "tracks/x/katas/1.json"); err != nil {
		t.Fatalf("valid path rejected: %v", err)
	}
	if _, err := safeJoin(dir, "../escape.json"); err == nil {
		t.Fatal("escaping path accepted")
	}
	if _, err := safeJoin(dir, "a/../../escape.json"); err == nil {
		t.Fatal("nested escaping path accepted")
	}
	if _, err := safeJoin(dir, strings.Repeat("a", 5000)+".json"); err == nil {
		t.Fatal("over-long path accepted")
	}
}
