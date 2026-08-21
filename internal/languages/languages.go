// Package languages defines the editor language system: a registry of
// programming languages the app can open, each describing how to highlight,
// auto-pair, indent, and syntax-check source code. Adding a language is a
// registry registration — no editor changes required.
package languages

// Language describes everything the editor needs to treat a source file as a
// specific language.
type Language struct {
	ID         string   // stable identifier, e.g. "go"
	Name       string   // display name, e.g. "Go"
	Extensions []string // file extensions, e.g. [".go"]
	// SourceFilenames are the canonical workspace file names for solution and
	// learner tests (e.g. "solution.go", "learner_test.go").
	SourceFilename    string
	TestsFilename     string
	LineComment       string // e.g. "//" ("" if none)
	BlockCommentStart string // e.g. "/*" ("" if none)
	BlockCommentEnd   string // e.g. "*/"
	QuoteChars        string // characters that start string literals, e.g. `"'`
	AutoPairs         []AutoPair
	Keywords          []string
	Types             []string
	// Indent is inserted by the auto-indent key handler (e.g. "\t" or "  ").
	Indent string
	// Checker reports syntax diagnostics for a source string. May be nil.
	Checker Checker
}

// AutoPair is an opening/closing character pair the editor completes
// automatically (braces, brackets, quotes...).
type AutoPair struct {
	Open  rune
	Close rune
}

// Diagnostic is a single syntax or lint message for a source range.
type Diagnostic struct {
	Line    int // 0-based line
	Col     int // 0-based column
	EndLine int // 0-based, inclusive-ish end (may equal Line)
	EndCol  int // 0-based end column
	Message string
	IsError bool
}

// Checker reports syntax diagnostics for source code.
type Checker interface {
	// Check returns diagnostics for src. It should be cheap and safe to call
	// on the main thread (no network).
	Check(src string) []Diagnostic
}

// Registry holds the known languages.
type Registry struct {
	byID        map[string]*Language
	byExtension map[string]*Language
	order       []string
}

// NewRegistry builds a registry pre-populated with the built-in languages.
func NewRegistry() *Registry {
	r := &Registry{
		byID:        make(map[string]*Language),
		byExtension: make(map[string]*Language),
	}
	for _, l := range builtins() {
		r.Register(l)
	}
	return r
}

// Register adds a language. It returns an error if the ID is empty, already
// registered, or if an extension collides with a different language.
func (r *Registry) Register(l *Language) error {
	if l == nil || l.ID == "" {
		return errMissingID
	}
	if _, exists := r.byID[l.ID]; exists {
		return errDuplicateID
	}
	for _, ext := range l.Extensions {
		if existing, ok := r.byExtension[ext]; ok && existing.ID != l.ID {
			return errExtensionConflict{ext: ext, first: existing.ID, second: l.ID}
		}
	}
	r.byID[l.ID] = l
	for _, ext := range l.Extensions {
		r.byExtension[ext] = l
	}
	r.order = append(r.order, l.ID)
	return nil
}

// Lookup finds a language by ID. Returns nil if unknown.
func (r *Registry) Lookup(id string) *Language {
	return r.byID[id]
}

// ForExtension finds the language for a file extension (including the dot).
func (r *Registry) ForExtension(ext string) *Language {
	return r.byExtension[ext]
}

// ForFilename finds the language for a file name by extension.
func (r *Registry) ForFilename(name string) *Language {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return r.byExtension[name[i:]]
		}
	}
	return nil
}

// All returns the registered languages in registration order.
func (r *Registry) All() []*Language {
	out := make([]*Language, 0, len(r.order))
	for _, id := range r.order {
		out = append(out, r.byID[id])
	}
	return out
}

// Default returns the language for an empty/unknown ID (Go).
func (r *Registry) Default() *Language {
	if goLang := r.byID["go"]; goLang != nil {
		return goLang
	}
	return r.byID[r.order[0]]
}
