package languages

import (
	"strings"
	"testing"
)

func TestRegistryBuiltins(t *testing.T) {
	r := NewRegistry()
	want := []string{"go", "rust", "java", "csharp", "cpp", "c", "php", "python", "javascript", "typescript"}
	for _, id := range want {
		if r.Lookup(id) == nil {
			t.Errorf("built-in language %q missing", id)
		}
	}
	if len(r.All()) != len(want) {
		t.Fatalf("expected %d languages, got %d", len(want), len(r.All()))
	}
}

func TestRegistryExtensionLookup(t *testing.T) {
	r := NewRegistry()
	cases := map[string]string{
		".go": "go", ".rs": "rust", ".java": "java", ".cs": "csharp",
		".cpp": "cpp", ".c": "c", ".php": "php", ".py": "python",
		".js": "javascript", ".ts": "typescript",
	}
	for ext, id := range cases {
		if got := r.ForExtension(ext); got == nil || got.ID != id {
			t.Errorf("ForExtension(%q) = %v, want %q", ext, got, id)
		}
	}
	if got := r.ForFilename("solution.go"); got == nil || got.ID != "go" {
		t.Errorf("ForFilename(solution.go) = %v", got)
	}
	if got := r.ForExtension(".xyz"); got != nil {
		t.Errorf("ForExtension(.xyz) should be nil, got %v", got.ID)
	}
}

func TestRegistryCustomLanguage(t *testing.T) {
	r := NewRegistry()
	err := r.Register(&Language{
		ID: "kotlin", Name: "Kotlin", Extensions: []string{".kt"},
		SourceFilename: "Solution.kt", TestsFilename: "LearnerTest.kt",
		LineComment: "//", QuoteChars: "\"",
		AutoPairs: []AutoPair{{Open: '{', Close: '}'}},
		Keywords:  []string{"fun", "val", "var", "class"},
		Indent:    "    ",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Lookup("kotlin") == nil || r.ForExtension(".kt").ID != "kotlin" {
		t.Fatal("kotlin should be registered")
	}
}

func TestRegistryConflicts(t *testing.T) {
	r := NewRegistry()
	if err := r.Register(&Language{ID: "go", Extensions: []string{".go"}}); err == nil {
		t.Fatal("duplicate ID should error")
	}
	if err := r.Register(&Language{ID: "x", Extensions: []string{".go"}}); err == nil {
		t.Fatal("extension conflict should error")
	} else if !strings.Contains(err.Error(), ".go") {
		t.Fatalf("conflict error should name the extension: %v", err)
	}
	if err := r.Register(&Language{ID: "", Extensions: []string{".q"}}); err == nil {
		t.Fatal("missing ID should error")
	}
}

func TestLanguageMetadata(t *testing.T) {
	r := NewRegistry()
	for _, l := range r.All() {
		if l.SourceFilename == "" || l.TestsFilename == "" {
			t.Errorf("%s: workspace filenames must be set", l.ID)
		}
		if l.Indent == "" {
			t.Errorf("%s: indent must be set", l.ID)
		}
		if len(l.AutoPairs) == 0 {
			t.Errorf("%s: should define auto-pairs", l.ID)
		}
	}
	if r.Default().ID != "go" {
		t.Fatalf("default language should be go, got %s", r.Default().ID)
	}
}

func TestDefaultLanguageIsGo(t *testing.T) {
	if DefaultLanguage.ID != "go" || DefaultLanguage.SourceFilename != "solution.go" {
		t.Fatalf("default language mismatch: %+v", DefaultLanguage)
	}
}
