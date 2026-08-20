#!/usr/bin/env python3
"""Create the 22 new kata directories with stub content."""
import json
import os
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
KATAS_DIR = ROOT / "katas"

NEW_KATAS = [
    {
        "id": "000", "slug": "go-setup-first-program",
        "title": "Go Setup & First Program",
        "focus": "go mod init, main package, go run, go build",
        "stage": "foundation", "category": "setup-and-toolchain", "level": "junior",
        "tags": ["setup", "toolchain"],
        "estimated_minutes": 10,
        "kata_go": '''package kata

// InitModule creates a new Go module by running "go mod init".
// It should return the module name created.
//
// Your task: implement the function below.
func InitModule(name string) string {
\treturn ""
}''',
        "kata_test": '''package kata

import (
\t"testing"
)

func TestInitModule(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\tinput string
\t\twant string
\t}{
\t	{"simple module", "example.com/myapp", "example.com/myapp"},
\t\t{"with path", "github.com/user/repo", "github.com/user/repo"},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot := InitModule(tc.input)
\t\t\tif got != tc.want {
\t\t\t\tt.Errorf("InitModule(%q) = %q, want %q", tc.input, got, tc.want)
\t\t\t}
\t\t})
\t}
}''',
        "readme": '''# Kata 000 — Go Setup & First Program

**Focus:** go mod init, main package, go run, go build

## Your task
Implement:

```go
func InitModule(name string) string
```

### Learning goal
- What you are practicing: understanding the Go project layout, module system, and basic toolchain commands.
- Why this matters: every Go project starts with a module. Understanding `go mod init`, `go run`, and `go build` is essential before writing any code.
- How this grows your Go skills: you learn the development workflow that every Go developer uses daily.

### Tips
- Run `go help mod` to see available commands.
- Try creating a real module with `go mod init example.com/myapp`.

## Rules / Expectations
- simple module => "example.com/myapp"
- with path => "github.com/user/repo"

## What this kata is about (and why it matters)
- Core lesson: every Go project needs a module. Understanding the toolchain is step one.
- After this kata, you should be comfortable setting up new Go projects.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What command initializes a new Go module?", "back": "go mod init <module-path>"},
            {"front": "What is the main package?", "back": "The entry point for a Go executable. It must have a main() function."},
            {"front": "What is the difference between go run and go build?", "back": "go run compiles and runs. go build compiles to a binary."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which command initializes a Go module?", "options": ["go init", "go mod init", "go new", "go create"], "answer": "go mod init"},
        ],
    },
    {
        "id": "005", "slug": "go-vet-and-lint",
        "title": "Go Vet & Lint",
        "focus": "go vet, golangci-lint, static analysis, code quality",
        "stage": "foundation", "category": "setup-and-toolchain", "level": "junior",
        "tags": ["toolchain", "testing"],
        "estimated_minutes": 10,
        "kata_go": '''package kata

// CheckCode runs static analysis on the given code string.
// It returns a list of issues found (empty if clean).
//
// Your task: implement the function below.
func CheckCode(code string) []string {
\treturn nil
}''',
        "kata_test": '''package kata

import "testing"

func TestCheckCode(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\tcode string
\t\twantLen int
\t}{
\t	{"clean code", "package main", 0},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot := CheckCode(tc.code)
\t\t\tif len(got) != tc.wantLen {
\t\t\t\tt.Errorf("CheckCode() returned %d issues, want %d", len(got), tc.wantLen)
\t\t\t}
\t\t})
\t}
}''',
        "readme": '''# Kata 005 — Go Vet & Lint

**Focus:** go vet, golangci-lint, static analysis, code quality

## Your task
Implement:

```go
func CheckCode(code string) []string
```

### Learning goal
- What you are practicing: using static analysis tools to catch bugs before they reach production.
- Why this matters: `go vet` and linters catch common mistakes that the compiler misses. Every Go team uses them.
- How this grows your Go skills: you develop the habit of checking code quality automatically.

### Tips
- Run `go vet ./...` on any Go project to see it in action.
- Try installing golangci-lint for more thorough checks.

## Rules / Expectations
- clean code => empty slice

## What this kata is about (and why it matters)
- Core lesson: automated quality checks are free bugs prevention.
- After this kata, you should always run `go vet` before committing.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What does go vet check for?", "back": "Suspicious constructs that the compiler doesn't catch: printf format mismatches, unreachable code, etc."},
            {"front": "What is golangci-lint?", "back": "A meta-linter that runs many Go linters at once and reports issues."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which command runs Go's built-in static analysis?", "options": ["go check", "go vet", "go lint", "go analyze"], "answer": "go vet"},
        ],
    },
    {
        "id": "039", "slug": "error-wrapping-patterns",
        "title": "Error Wrapping Patterns",
        "focus": "error wrapping, fmt.Errorf, %w, errors.Is, errors.As",
        "stage": "foundation", "category": "error-handling", "level": "junior",
        "tags": ["error-handling"],
        "estimated_minutes": 15,
        "kata_go": '''package kata

import "errors"

var ErrNotFound = errors.New("not found")

// FindUser looks up a user by ID.
// Return ErrNotFound if the user doesn't exist.
// Wrap any other errors with context.
//
// Your task: implement the function below.
func FindUser(id int) (string, error) {
\treturn "", nil
}''',
        "kata_test": '''package kata

import (
\t"errors"
\t"testing"
)

func TestFindUser(t *testing.T) {
\tt.Run("existing user", func(t *testing.T) {
\t\tname, err := FindUser(1)
\t\tif err != nil {
\t\t\tt.Fatalf("unexpected error: %v", err)
\t\t}
\t\tif name == "" {
\t\t\tt.Fatal("expected non-empty name")
\t\t}
\t})

\tt.Run("missing user", func(t *testing.T) {
\t\t_, err := FindUser(999)
\t\tif err == nil {
\t\t\tt.Fatal("expected error")
\t\t}
\t\tif !errors.Is(err, ErrNotFound) {
\t\t\tt.Errorf("expected ErrNotFound, got %v", err)
\t\t}
\t})
}''',
        "readme": '''# Kata 039 — Error Wrapping Patterns

**Focus:** error wrapping, fmt.Errorf, %w, errors.Is, errors.As

## Your task
Implement:

```go
func FindUser(id int) (string, error)
```

### Learning goal
- What you are practicing: wrapping errors with context while preserving the ability to check error types.
- Why this matters: in real systems, you need to add context to errors ("failed to fetch user 123") while still being able to check the root cause.
- How this grows your Go skills: you learn the `%w` verb, `errors.Is`, and `errors.As` — the modern Go error handling toolkit.

### Tips
- Use `fmt.Errorf("context: %w", err)` to wrap errors.
- Use `errors.Is(err, ErrNotFound)` to check error identity.
- Use `errors.As(err, &target)` to extract error types.

## Rules / Expectations
- existing user => returns name, nil error
- missing user => returns ErrNotFound

## What this kata is about (and why it matters)
- Core lesson: wrap errors with context, check with errors.Is/As.
- After this kata, you should never lose error context in your code.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "How do you wrap an error with context in Go?", "back": "fmt.Errorf(\"context: %w\", err)"},
            {"front": "How do you check if an error is a specific sentinel?", "back": "errors.Is(err, ErrNotFound)"},
            {"front": "When would you use errors.As instead of errors.Is?", "back": "When you need to extract the concrete error type, not just check identity."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which verb wraps an error in fmt.Errorf?", "options": ["%s", "%v", "%w", "%+v"], "answer": "%w"},
            {"type": "fill_blank", "question": "errors.Is(err, ErrNotFound) returns true when err ___ ErrNotFound or wraps it", "answer": "is"},
        ],
    },
    {
        "id": "040", "slug": "interface-basics",
        "title": "Interface Basics",
        "focus": "interfaces, implicit satisfaction, accept interfaces return structs",
        "stage": "foundation", "category": "error-handling", "level": "junior",
        "tags": ["architecture"],
        "estimated_minutes": 15,
        "kata_go": '''package kata

// Greeter is an interface for greeting people.
type Greeter interface {
\tGreet(name string) string
}

// FormalGreeter greets formally.
// Your task: implement the Greeter interface.
type FormalGreeter struct{}

func (f FormalGreeter) Greet(name string) string {
\treturn ""
}

// CasualGreeter greets casually.
// Your task: implement the Greeter interface.
type CasualGreeter struct{}

func (c CasualGreeter) Greet(name string) string {
\treturn ""
}

// SayHello uses any Greeter to greet someone.
func SayHello(g Greeter, name string) string {
\treturn g.Greet(name)
}''',
        "kata_test": '''package kata

import "testing"

func TestFormalGreeter(t *testing.T) {
\tg := FormalGreeter{}
\tgot := g.Greet("Alice")
\tif got == "" {
\t\tt.Fatal("expected non-empty greeting")
\t}
}

func TestCasualGreeter(t *testing.T) {
\tg := CasualGreeter{}
\tgot := g.Greet("Bob")
\tif got == "" {
\t\tt.Fatal("expected non-empty greeting")
\t}
}

func TestSayHello(t *testing.T) {
\tvar formal Greeter = FormalGreeter{}
\tvar casual Greeter = CasualGreeter{}

\tif formal.Greet("X") == casual.Greet("X") {
\t\tt.Error("formal and casual greetings should differ")
\t}
}''',
        "readme": '''# Kata 040 — Interface Basics

**Focus:** interfaces, implicit satisfaction, accept interfaces return structs

## Your task
Implement the Greeter interface for both FormalGreeter and CasualGreeter.

### Learning goal
- What you are practicing: defining and implementing interfaces in Go.
- Why this matters: interfaces are how Go achieves polymorphism and testability. "Accept interfaces, return structs" is a core Go idiom.
- How this grows your Go skills: you learn that Go interfaces are satisfied implicitly — no `implements` keyword needed.

### Tips
- In Go, a type satisfies an interface just by implementing its methods.
- Keep interfaces small — 1-3 methods is ideal.

## Rules / Expectations
- FormalGreeter => formal greeting
- CasualGreeter => casual greeting
- Both implement Greeter

## What this kata is about (and why it matters)
- Core lesson: interfaces enable polymorphism without inheritance.
- After this kata, you should design your code around small interfaces.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "How does a type satisfy an interface in Go?", "back": "By implementing all the methods. No explicit declaration needed."},
            {"front": "What does 'accept interfaces, return structs' mean?", "back": "Function parameters should be interfaces (flexibility), return values should be concrete types (specificity)."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "How do you declare that a type implements an interface in Go?", "options": ["type T implements I", "type T : I", "By implementing all methods (no keyword needed)", "type T as I"], "answer": "By implementing all methods (no keyword needed)"},
        ],
    },
    {
        "id": "041", "slug": "table-driven-tests",
        "title": "Table-Driven Tests",
        "focus": "table-driven tests, subtests, testing.T, test patterns",
        "stage": "foundation", "category": "testing-fundamentals", "level": "junior",
        "tags": ["testing"],
        "estimated_minutes": 15,
        "kata_go": '''package kata

// Abs returns the absolute value of n.
func Abs(n int) int {
\tif n < 0 {
\t\treturn -n
\t}
\treturn n
}

// Max returns the larger of a and b.
func Max(a, b int) int {
\tif a > b {
\t\treturn a
\t}
\treturn b
}''',
        "kata_test": '''package kata

import "testing"

// Your task: write table-driven tests for Abs and Max.
// Use subtests (t.Run) for each case.

func TestAbs(t *testing.T) {
\t// Write table-driven tests here
\t// Include: positive, negative, zero cases
\tt.Fatal("implement table-driven tests for Abs")
}

func TestMax(t *testing.T) {
\t// Write table-driven tests here
\t// Include: a>b, b>a, a==b cases
\tt.Fatal("implement table-driven tests for Max")
}''',
        "readme": '''# Kata 041 — Table-Driven Tests

**Focus:** table-driven tests, subtests, testing.T, test patterns

## Your task
Write table-driven tests for the `Abs` and `Max` functions using subtests.

### Learning goal
- What you are practicing: the standard Go testing pattern — table-driven tests with subtests.
- Why this matters: table-driven tests are THE way to write tests in Go. They're readable, maintainable, and easy to extend.
- How this grows your Go skills: you learn `t.Run()` for subtests, test organization, and the Go testing idioms used in every production codebase.

### Tips
- Define a slice of test cases with input and expected output.
- Use `t.Run("name", func(t *testing.T) {...})` for each case.
- Include edge cases: zero, negative, equal values.

## Rules / Expectations
- Abs of positive => same number
- Abs of negative => positive
- Abs of zero => zero
- Max of a>b => a
- Max of b>a => b
- Max of a==b => a

## What this kata is about (and why it matters)
- Core lesson: table-driven tests are the Go testing standard.
- After this kata, every test you write should follow this pattern.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is a table-driven test?", "back": "A test that defines a slice of test cases (input/expected) and loops through them."},
            {"front": "How do you create a subtest in Go?", "back": "t.Run(\"name\", func(t *testing.T) { ... })"},
            {"front": "Why are table-driven tests preferred in Go?", "back": "They're concise, easy to add cases, and each case gets its own name for clear failure messages."},
        ],
        "quiz_questions": [
            {"type": "fill_blank", "question": "In Go, you create a subtest with t.___(\"name\", func(t *testing.T) { ... })", "answer": "Run"},
            {"type": "multiple_choice", "question": "What is the main advantage of table-driven tests?", "options": ["They run faster", "They're easier to extend with new cases", "They use less memory", "They automatically generate coverage"], "answer": "They're easier to extend with new cases"},
        ],
    },
    {
        "id": "078", "slug": "net-http-server-basics",
        "title": "net/http Server Basics",
        "focus": "net/http, http.HandleFunc, http.ListenAndServe, request/response",
        "stage": "practitioner", "category": "web-networking", "level": "mid",
        "tags": ["http"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

import "net/http"

// CreateMux returns an http.ServeMux with routes registered.
// Register:
//   - GET /health => return 200 OK
//   - GET /hello?name=XXX => return "Hello, XXX!"
//
// Your task: implement the function below.
func CreateMux() *http.ServeMux {
\treturn http.NewServeMux()
}''',
        "kata_test": '''package kata

import (
\t"net/http"
\t"net/http/httptest"
\t"testing"
)

func TestHealthEndpoint(t *testing.T) {
\tmux := CreateMux()
\treq := httptest.NewRequest("GET", "/health", nil)
\tw := httptest.NewRecorder()
\tmux.ServeHTTP(w, req)
\tif w.Code != http.StatusOK {
\t\tt.Errorf("expected 200, got %d", w.Code)
\t}
}

func TestHelloEndpoint(t *testing.T) {
\tmux := CreateMux()
\treq := httptest.NewRequest("GET", "/hello?name=World", nil)
\tw := httptest.NewRecorder()
\tmux.ServeHTTP(w, req)
\tif w.Code != http.StatusOK {
\t\tt.Errorf("expected 200, got %d", w.Code)
\t}
\tbody := w.Body.String()
\tif body == "" {
\t\tt.Fatal("expected non-empty body")
\t}
}''',
        "readme": '''# Kata 078 — net/http Server Basics

**Focus:** net/http, http.HandleFunc, http.ListenAndServe, request/response

## Your task
Implement:

```go
func CreateMux() *http.ServeMux
```

### Learning goal
- What you are practicing: building HTTP handlers with Go's standard library.
- Why this matters: `net/http` is the foundation of every Go web service. Understanding handlers, muxes, and request/response is essential.
- How this grows your Go skills: you learn the handler interface, request parsing, and response writing.

### Tips
- Use `http.NewRequest` and `httptest.NewRecorder` for testing.
- Register handlers with `mux.HandleFunc`.

## Rules / Expectations
- /health => 200 OK
- /hello?name=X => "Hello, X!"

## What this kata is about (and why it matters)
- Core lesson: Go's net/http is powerful enough for production use.
- After this kata, you can build HTTP services without any framework.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What interface must an HTTP handler implement?", "back": "http.Handler: ServeHTTP(ResponseWriter, *Request)"},
            {"front": "How do you test HTTP handlers without starting a server?", "back": "httptest.NewRecorder() + httptest.NewRequest()"},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which function registers a handler on the default mux?", "options": ["http.Handle", "http.HandleFunc", "http.Register", "http.Route"], "answer": "http.HandleFunc"},
        ],
    },
    {
        "id": "088", "slug": "yaml-toml-parsing",
        "title": "YAML/TOML Parsing",
        "focus": "encoding/json patterns, struct tags, config file parsing",
        "stage": "practitioner", "category": "data-serialization", "level": "mid",
        "tags": ["io", "data-structures"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// Config represents a simplified config file structure.
type Config struct {
\tName    string            `json:"name" yaml:"name" toml:"name"`
\tVersion string            `json:"version" yaml:"version" toml:"version"`
\tDebug   bool              `json:"debug" yaml:"debug" toml:"debug"`
\tTags    map[string]string `json:"tags" yaml:"tags" toml:"tags"`
}

// ParseConfig parses a JSON string into a Config struct.
// Your task: implement the function below.
func ParseConfig(data string) (Config, error) {
\treturn Config{}, nil
}''',
        "kata_test": '''package kata

import "testing"

func TestParseConfig(t *testing.T) {
\tdata := `{"name":"app","version":"1.0","debug":true,"tags":{"env":"prod"}}`
\tcfg, err := ParseConfig(data)
\tif err != nil {
\t\tt.Fatalf("unexpected error: %v", err)
\t}
\tif cfg.Name != "app" {
\t\tt.Errorf("expected name 'app', got %q", cfg.Name)
\t}
\tif !cfg.Debug {
\t\tt.Error("expected debug to be true")
\t}
}''',
        "readme": '''# Kata 088 — YAML/TOML Parsing

**Focus:** encoding/json patterns, struct tags, config file parsing

## Your task
Implement:

```go
func ParseConfig(data string) (Config, error)
```

### Learning goal
- What you are practicing: parsing structured config data with struct tags.
- Why this matters: every Go application needs configuration. Understanding struct tags and config parsing patterns is essential.
- How this grows your Go skills: you learn how struct tags control serialization and how to design config types.

### Tips
- Use `encoding/json` for JSON parsing.
- Struct tags like `json:"name"` control field mapping.
- Always handle parse errors explicitly.

## Rules / Expectations
- valid JSON => Config with correct fields
- invalid JSON => error

## What this kata is about (and why it matters)
- Core lesson: struct tags are Go's way of controlling serialization.
- After this kata, you can parse any JSON/YAML/TOML config.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What are struct tags in Go?", "back": "Metadata strings after struct fields that control serialization: `json:\"name\"`"},
            {"front": "How do you parse JSON into a struct?", "back": "json.Unmarshal([]byte(data), &target)"},
        ],
        "quiz_questions": [
            {"type": "fill_blank", "question": "Struct tags are written after the field type in backticks, like `json:\"___\"`", "answer": "name"},
        ],
    },
    {
        "id": "089", "slug": "structured-logging-slog",
        "title": "Structured Logging (slog)",
        "focus": "log/slog, structured logging, key-value pairs, log levels",
        "stage": "practitioner", "category": "real-world-packages", "level": "mid",
        "tags": ["io"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

import "log/slog"

// AppLogger wraps slog for application logging.
// Your task: implement the methods below.
type AppLogger struct {
\tlogger *slog.Logger
}

// NewAppLogger creates a new AppLogger.
func NewAppLogger(logger *slog.Logger) *AppLogger {
\treturn &AppLogger{logger: logger}
}

// LogRequest logs an HTTP request with structured fields.
func (l *AppLogger) LogRequest(method, path string, status int) {
\t// Your implementation here
}

// LogError logs an error with context.
func (l *AppLogger) LogError(msg string, err error) {
\t// Your implementation here
}''',
        "kata_test": '''package kata

import (
\t"log/slog"
\t"os"
\t"testing"
)

func TestLogRequest(t *testing.T) {
\tlogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
\tapp := NewAppLogger(logger)
\t// Should not panic
\tapp.LogRequest("GET", "/health", 200)
}

func TestLogError(t *testing.T) {
\tlogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
\tapp := NewAppLogger(logger)
\t// Should not panic
\tapp.LogError("something failed", nil)
}''',
        "readme": '''# Kata 089 — Structured Logging (slog)

**Focus:** log/slog, structured logging, key-value pairs, log levels

## Your task
Implement structured logging with Go's `log/slog` package.

### Learning goal
- What you are practicing: structured logging with key-value pairs and log levels.
- Why this matters: structured logs are machine-parseable, searchable, and essential for production debugging.
- How this grows your Go skills: you learn the slog API that replaced most third-party logging in Go 1.21+.

### Tips
- Use `slog.Info("msg", "key", value)` for structured logging.
- Use `slog.NewJSONHandler` for JSON output.
- Always include context in log messages.

## Rules / Expectations
- LogRequest logs method, path, status as structured fields
- LogError logs message and error

## What this kata is about (and why it matters)
- Core lesson: structured logs > unstructured logs. Always.
- After this kata, you'll never use fmt.Println for logging again.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is structured logging?", "back": "Logging with key-value pairs instead of free-form text. Makes logs searchable and parseable."},
            {"front": "How do you create a JSON logger in slog?", "back": "slog.New(slog.NewJSONHandler(os.Stdout, nil))"},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which Go package provides structured logging as of Go 1.21?", "options": ["log", "log/slog", "logrus", "zap"], "answer": "log/slog"},
        ],
    },
    {
        "id": "090", "slug": "cobra-cli",
        "title": "Cobra CLI",
        "focus": "github.com/spf13/cobra, CLI commands, flags, help text",
        "stage": "practitioner", "category": "real-world-packages", "level": "mid",
        "tags": ["cli"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// CLIConfig holds the parsed CLI arguments.
type CLIConfig struct {
\tCommand string
\tName    string
\tVerbose bool
}

// ParseCLI parses command-line arguments into a CLIConfig.
// Commands: "greet --name=World --verbose"
// Your task: implement the function below.
func ParseCLI(args []string) (CLIConfig, error) {
\treturn CLIConfig{}, nil
}''',
        "kata_test": '''package kata

import "testing"

func TestParseCLI(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\targs []string
\t\twant CLIConfig
\t}{
\t\t{"greet", []string{"greet", "--name=World"}, CLIConfig{Command: "greet", Name: "World"}},
\t\t{"greet verbose", []string{"greet", "--name=Alice", "--verbose"}, CLIConfig{Command: "greet", Name: "Alice", Verbose: true}},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot, err := ParseCLI(tc.args)
\t\t\tif err != nil {
\t\t\t\tt.Fatalf("unexpected error: %v", err)
\t\t\t}
\t\t\tif got != tc.want {
\t\t\t\tt.Errorf("got %+v, want %+v", got, tc.want)
\t\t\t}
\t\t})
\t}
}''',
        "readme": '''# Kata 090 — Cobra CLI

**Focus:** github.com/spf13/cobra, CLI commands, flags, help text

## Your task
Implement CLI argument parsing.

### Learning goal
- What you are practicing: parsing command-line arguments into structured config.
- Why this matters: most Go tools are CLIs. Understanding flag parsing and command structure is essential.
- How this grows your Go skills: you learn the patterns used by every popular Go CLI tool.

### Tips
- Handle both `--flag=value` and `--flag value` formats.
- Always provide helpful error messages for invalid arguments.

## Rules / Expectations
- Parse command name from first argument
- Parse --name= value
- Parse --verbose boolean flag

## What this kata is about (and why it matters)
- Core lesson: good CLIs have clear argument parsing and helpful errors.
- After this kata, you can build production-quality CLI tools.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is Cobra in Go?", "back": "A popular library for building CLI applications with commands, flags, and help text."},
            {"front": "What are the two common flag formats?", "back": "--flag=value and --flag value"},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which Go library is most commonly used for building CLIs?", "options": ["flag", "cobra", "cli", "argparse"], "answer": "cobra"},
        ],
    },
    {
        "id": "091", "slug": "testify-assertions",
        "title": "testify Assertions",
        "focus": "github.com/stretchr/testify, assert/require, test helpers",
        "stage": "practitioner", "category": "real-world-packages", "level": "mid",
        "tags": ["testing"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

// Multiply returns the product of a and b.
func Multiply(a, b int) int {
\treturn a * b
}

// Split splits a string by delimiter and returns non-empty parts.
func Split(s, delim string) []string {
\tif s == "" {
\t\treturn nil
\t}
\tresult := []string{}
\tfor _, part := range splitSimple(s, delim) {
\t\tif part != "" {
\t\t\tresult = append(result, part)
\t\t}
\t}
\treturn result
}

func splitSimple(s, delim string) []string {
\tvar parts []string
\tcurrent := ""
\tfor _, ch := range s {
\t\tif string(ch) == delim {
\t\t\tparts = append(parts, current)
\t\t\tcurrent = ""
\t\t} else {
\t\t\tcurrent += string(ch)
\t\t}
\t}
\tparts = append(parts, current)
\treturn parts
}''',
        "kata_test": '''package kata

import (
\t"testing"
)

// Rewrite these tests using testify assertions (if available).
// For now, use standard library patterns that testify would improve.

func TestMultiply(t *testing.T) {
\ttests := []struct {
\t\ta, b, want int
\t}{
\t\t{2, 3, 6},
\t\t{0, 5, 0},
\t\t{-1, 4, -4},
\t}

\tfor _, tc := range tests {
\t\tt.Run("", func(t *testing.T) {
\t\t\tgot := Multiply(tc.a, tc.b)
\t\t\tif got != tc.want {
\t\t\t\tt.Errorf("Multiply(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
\t\t\t}
\t\t})
\t}
}''',
        "readme": '''# Kata 091 — testify Assertions

**Focus:** github.com/stretchr/testify, assert/require, test helpers

## Your task
Rewrite tests using testify-style assertions.

### Learning goal
- What you are practicing: using the testify library for cleaner, more expressive tests.
- Why this matters: testify is the most popular Go testing library. Its assert/require packages reduce boilerplate.
- How this grows your Go skills: you learn assert.Equal, require.NoError, and testify suites.

### Tips
- `assert.Equal(t, want, got)` replaces manual if-checks.
- `require.NoError(t, err)` stops the test immediately on failure.
- Use `assert.Contains`, `assert.Len`, `assert.True` for specific checks.

## Rules / Expectations
- Multiply(2,3) => 6
- Multiply(0,5) => 0
- Tests use assertion patterns

## What this kata is about (and why it matters)
- Core lesson: good assertions make tests readable and failures obvious.
- After this kata, your tests will be cleaner and more maintainable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What does assert.Equal do?", "back": "Compares two values and fails the test with a clear message if they differ."},
            {"front": "What is the difference between assert and require?", "back": "assert continues the test on failure. require stops immediately."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which testify function stops the test immediately on failure?", "options": ["assert.Equal", "require.NoError", "assert.True", "require.Contains"], "answer": "require.NoError"},
        ],
    },
    {
        "id": "092", "slug": "pgx-database-basics",
        "title": "pgx Database Basics",
        "focus": "database/sql, pgx, connection, query, scan",
        "stage": "practitioner", "category": "real-world-packages", "level": "mid",
        "tags": ["database"],
        "estimated_minutes": 30,
        "kata_go": '''package kata

// User represents a database record.
type User struct {
\tID    int
\tName  string
\tEmail string
}

// UserStore provides access to user data.
// Your task: implement the interface.
type UserStore interface {
\tGetByID(id int) (User, error)
\tCreate(name, email string) (int, error)
}''',
        "kata_test": '''package kata

import "testing"

// Test with a mock store (no real database needed)
type mockStore struct {
\tusers map[int]User
\tnextID int
}

func (m *mockStore) GetByID(id int) (User, error) {
\tu, ok := m.users[id]
\tif !ok {
\t\treturn User{}, fmt.Errorf("user %d not found", id)
\t}
\treturn u, nil
}

func (m *mockStore) Create(name, email string) (int, error) {
\tm.nextID++
\tm.users[m.nextID] = User{ID: m.nextID, Name: name, Email: email}
\treturn m.nextID, nil
}

func TestUserStore(t *testing.T) {
\tstore := &mockStore{users: make(map[int]User), nextID: 0}

\tid, err := store.Create("Alice", "alice@example.com")
\tif err != nil {
\t\tt.Fatalf("create: %v", err)
\t}

\tuser, err := store.GetByID(id)
\tif err != nil {
\t\tt.Fatalf("get: %v", err)
\t}
\tif user.Name != "Alice" {
\t\tt.Errorf("expected name Alice, got %q", user.Name)
\t}
}''',
        "readme": '''# Kata 092 — pgx Database Basics

**Focus:** database/sql, pgx, connection, query, scan

## Your task
Implement a UserStore interface with a mock backing store.

### Learning goal
- What you are practicing: database access patterns with Go interfaces.
- Why this matters: every production Go app talks to a database. Understanding the interface pattern enables testing without a real DB.
- How this grows your Go skills: you learn the repository pattern and how to mock database access.

### Tips
- Define an interface for your store.
- Test with a mock implementation.
- Use context.Context for cancellation support.

## Rules / Expectations
- Create returns an ID
- GetByID returns the user or error
- Tests use mock store (no real DB)

## What this kata is about (and why it matters)
- Core lesson: mock your database boundary for fast, reliable tests.
- After this kata, you can design data access layers that are testable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is the repository pattern?", "back": "An interface that abstracts data access, enabling mock implementations for testing."},
            {"front": "Why use interfaces for database access?", "back": "So you can test with mocks instead of a real database."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Why do we mock the database in tests?", "options": ["It's faster", "No real DB needed, tests are reliable", "It uses less memory", "All of the above"], "answer": "All of the above"},
        ],
    },
    {
        "id": "093", "slug": "fuzz-testing",
        "title": "Fuzz Testing",
        "focus": "testing.F, fuzzing, property-based testing, crash detection",
        "stage": "practitioner", "category": "real-world-packages", "level": "mid",
        "tags": ["testing"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

// ParseInt safely parses a string to int.
// Return 0 and an error for invalid input.
func ParseInt(s string) (int, error) {
\t// Your implementation
\treturn 0, nil
}''',
        "kata_test": '''package kata

import "testing"

func TestParseInt(t *testing.T) {
\ttests := []struct {
\t\tinput string
\t\twant int
\t\twantErr bool
\t}{
\t	{"42", 42, false},
\t	{"0", 0, false},
\t	{"abc", 0, true},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.input, func(t *testing.T) {
\t\t\tgot, err := ParseInt(tc.input)
\t\t\tif (err != nil) != tc.wantErr {
\t\t\t\tt.Errorf("ParseInt(%q) error = %v, wantErr %v", tc.input, err, tc.wantErr)
\t\t\t}
\t\t\tif got != tc.want {
\t\t\t\tt.Errorf("ParseInt(%q) = %d, want %d", tc.input, got, tc.want)
\t\t\t}
\t\t})
\t}
}

// FuzzParseInt is a fuzz test that feeds random inputs to ParseInt.
// Uncomment and run with: go test -fuzz=FuzzParseInt
/*
func FuzzParseInt(f *testing.F) {
\tf.Add("42")
\tf.Add("0")
\tf.Add("-1")
\tf.Add("abc")

\tf.Fuzz(func(t *testing.T, input string) {
\t\tresult, err := ParseInt(input)
\t\tif err != nil {
\t\t\treturn
\t\t}
\t\t// Property: if no error, result should be parseable
\t\t_ = result
\t})
}
*/''',
        "readme": '''# Kata 093 — Fuzz Testing

**Focus:** testing.F, fuzzing, property-based testing, crash detection

## Your task
Implement ParseInt and write a fuzz test for it.

### Learning goal
- What you are practicing: property-based testing with Go's built-in fuzzing.
- Why this matters: fuzz tests find edge cases that human-written tests miss. They're essential for parsers and input handlers.
- How this grows your Go skills: you learn Go's fuzz testing API and property-based testing thinking.

### Tips
- Use `f.Add()` to seed the fuzzer with known inputs.
- Use `f.Fuzz()` to define the property to check.
- Run with `go test -fuzz=FuzzXxx` for 30 seconds.

## Rules / Expectations
- ParseInt("42") => 42, nil
- ParseInt("abc") => 0, error
- Fuzz test runs without crashes

## What this kata is about (and why it matters)
- Core lesson: fuzz tests find bugs you didn't know existed.
- After this kata, you'll fuzz-test every parser you write.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
go test -fuzz=FuzzParseInt -fuzztime=30s
```
''',
        "flashcards": [
            {"front": "How do you define a fuzz test in Go?", "back": "func FuzzXxx(f *testing.F) { f.Add(seed); f.Fuzz(func(t *testing.T, args...) { ... }) }"},
            {"front": "What is property-based testing?", "back": "Testing that certain properties hold for ALL inputs, not just specific examples."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "How do you run a fuzz test in Go?", "options": ["go test -fuzz", "go fuzz", "go test -fuzz=Xxx", "go test -fuzztest"], "answer": "go test -fuzz=Xxx"},
        ],
    },
    {
        "id": "113", "slug": "race-detection",
        "title": "Race Detection",
        "focus": "go test -race, data races, sync.Mutex, race detector",
        "stage": "practitioner", "category": "concurrency", "level": "mid",
        "tags": ["concurrency", "testing"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

import "sync"

// Counter is a concurrent-safe counter.
// Your task: make it safe for concurrent use.
type Counter struct {
\tvalue int
}

// Inc increments the counter by 1.
func (c *Counter) Inc() {
\tc.value++
}

// Value returns the current counter value.
func (c *Counter) Value() int {
\treturn c.value
}''',
        "kata_test": '''package kata

import (
\t"sync"
\t"testing"
)

func TestCounterConcurrent(t *testing.T) {
\tvar c Counter
\tvar wg sync.WaitGroup

\tfor i := 0; i < 1000; i++ {
\t\twg.Add(1)
\t\tgo func() {
\t\t\tdefer wg.Done()
\t\t\tc.Inc()
\t\t}()
\t}

\twg.Wait()
\tif c.Value() != 1000 {
\t\tt.Errorf("expected 1000, got %d", c.Value())
\t}
}''',
        "readme": '''# Kata 113 — Race Detection

**Focus:** go test -race, data races, sync.Mutex, race detector

## Your task
Make the Counter type safe for concurrent use and verify with the race detector.

### Learning goal
- What you are practicing: identifying and fixing data races using Go's race detector.
- Why this matters: data races cause non-deterministic bugs that are nearly impossible to reproduce without the detector.
- How this grows your Go skills: you learn to use `go test -race` as a standard part of your workflow.

### Tips
- Use `sync.Mutex` to protect shared state.
- Run `go test -race ./...` to detect races.
- The race detector adds ~10x overhead — use it in tests, not production.

## Rules / Expectations
- Counter is safe for concurrent use
- 1000 goroutines each increment once => value is 1000
- Passes with -race flag

## What this kata is about (and why it matters)
- Core lesson: always run the race detector. Always.
- After this kata, `go test -race` will be part of your muscle memory.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test -race ./...
```
''',
        "flashcards": [
            {"front": "How do you detect data races in Go?", "back": "go test -race ./..."},
            {"front": "What does the race detector do?", "back": "Finds concurrent memory accesses where at least one is a write, with no synchronization."},
            {"front": "What is the performance overhead of the race detector?", "back": "About 2-10x slower, with 5-10x more memory usage."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which flag enables the race detector?", "options": ["-race", "-racy", "-detect", "-safe"], "answer": "-race"},
            {"type": "fill_blank", "question": "Run the race detector with go test ___ ./...", "answer": "-race"},
        ],
    },
    {
        "id": "123", "slug": "opentelemetry-basics",
        "title": "OpenTelemetry Basics",
        "focus": "tracing, spans, context propagation, observability",
        "stage": "senior", "category": "observability", "level": "senior",
        "tags": ["architecture"],
        "estimated_minutes": 35,
        "kata_go": '''package kata

import "context"

// Span represents a tracing span (simplified for this kata).
type Span struct {
\tName     string
\tParent   string
\tFinished bool
}

// Tracer creates and manages spans.
type Tracer struct {
\tspans []Span
}

// NewTracer creates a new Tracer.
func NewTracer() *Tracer {
\treturn &Tracer{}
}

// StartSpan creates a new span with the given name.
// If ctx contains a parent span, link to it.
func (t *Tracer) StartSpan(ctx context.Context, name string) context.Context {
\t// Your implementation
\treturn ctx
}

// FinishSpan marks the current span as finished.
func (t *Tracer) FinishSpan(ctx context.Context) {
\t// Your implementation
}''',
        "kata_test": '''package kata

import (
\t"context"
\t"testing"
)

func TestTracerSpans(t *testing.T) {
\ttracer := NewTracer()

\tctx := tracer.StartSpan(context.Background(), "request")
\ttracer.StartSpan(ctx, "database")
\ttracer.FinishSpan(ctx)

\tif len(tracer.spans) == 0 {
\t\tt.Fatal("expected spans")
\t}
}''',
        "readme": '''# Kata 123 — OpenTelemetry Basics

**Focus:** tracing, spans, context propagation, observability

## Your task
Implement a simplified tracing system.

### Learning goal
- What you are practicing: distributed tracing concepts with spans and context propagation.
- Why this matters: in production systems, you need to trace requests across services. OpenTelemetry is the standard.
- How this grows your Go skills: you learn how context carries trace information through call chains.

### Tips
- Spans have a name, parent, and lifecycle (start/finish).
- Context propagation links parent and child spans.
- In real code, use go.opentelemetry.io/otel.

## Rules / Expectations
- StartSpan creates a span
- FinishSpan marks it done
- Parent-child relationships are tracked

## What this kata is about (and why it matters)
- Core lesson: tracing is essential for understanding distributed systems.
- After this kata, you can add observability to any Go service.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is a span in distributed tracing?", "back": "A named, timed operation that represents a unit of work in a trace."},
            {"front": "How does context propagation work?", "back": "Trace context is stored in context.Context and passed through function calls."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What is the standard for distributed tracing in Go?", "options": ["OpenTracing", "OpenCensus", "OpenTelemetry", "Zipkin"], "answer": "OpenTelemetry"},
        ],
    },
    {
        "id": "125", "slug": "pprof-profiling",
        "title": "pprof Profiling",
        "focus": "runtime/pprof, profiling, CPU/memory analysis, benchmarks",
        "stage": "senior", "category": "performance", "level": "senior",
        "tags": ["testing"],
        "estimated_minutes": 35,
        "kata_go": '''package kata

// FindDuplicates returns duplicate integers from a slice.
// Your task: implement an efficient version.
func FindDuplicates(nums []int) []int {
\t// Your implementation
\treturn nil
}''',
        "kata_test": '''package kata

import "testing"

func TestFindDuplicates(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\tnums []int
\t\twant []int
\t}{
\t\t{"no dups", []int{1, 2, 3}, nil},
\t\t{"one dup", []int{1, 2, 2, 3}, []int{2}},
\t\t{"multiple", []int{1, 1, 2, 2}, []int{1, 2}},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot := FindDuplicates(tc.nums)
\t\t\tif len(got) != len(tc.want) {
\t\t\t\tt.Errorf("got %v, want %v", got, tc.want)
\t\t\t}
\t\t})
\t}
}

func BenchmarkFindDuplicates(b *testing.B) {
\tnums := make([]int, 1000)
\tfor i := range nums {
\t\tnums[i] = i % 100
\t}
\tb.ResetTimer()
\tfor i := 0; i < b.N; i++ {
\t\tFindDuplicates(nums)
\t}
}''',
        "readme": '''# Kata 125 — pprof Profiling

**Focus:** runtime/pprof, profiling, CPU/memory analysis, benchmarks

## Your task
Implement FindDuplicates and write a benchmark for it.

### Learning goal
- What you are practicing: writing benchmarks and using profiling to find performance bottlenecks.
- Why this matters: "make it work, then make it fast" — profiling tells you where to optimize.
- How this grows your Go skills: you learn to use `go test -bench`, `go tool pprof`, and benchmark-driven optimization.

### Tips
- Write benchmarks with `func BenchmarkXxx(b *testing.B)`.
- Use `b.ResetTimer()` after setup.
- Run `go test -bench=. -benchmem` for memory allocation info.
- Use `go tool pprof` to analyze CPU/memory profiles.

## Rules / Expectations
- FindDuplicates returns duplicate values
- Benchmark runs without errors
- Code is efficient (O(n) expected)

## What this kata is about (and why it matters)
- Core lesson: measure before optimizing. Profiling finds real bottlenecks.
- After this kata, you'll benchmark before claiming code is "fast enough."

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
go test -bench=. -benchmem ./...
```
''',
        "flashcards": [
            {"front": "How do you write a benchmark in Go?", "back": "func BenchmarkXxx(b *testing.B) { for i := 0; i < b.N; i++ { ... } }"},
            {"front": "What does -benchmem show?", "back": "Bytes allocated per operation and number of allocations."},
            {"front": "How do you profile a running Go application?", "back": "Import _ \"net/http/pprof\" and use go tool pprof to analyze the profile."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What flag shows memory allocations in benchmarks?", "options": ["-bench", "-benchmem", "-mem", "-alloc"], "answer": "-benchmem"},
        ],
    },
    {
        "id": "126", "slug": "benchmarking-deep-dive",
        "title": "Benchmarking Deep Dive",
        "focus": "benchmark patterns, sub-benchmarks, comparison, regression",
        "stage": "senior", "category": "performance", "level": "senior",
        "tags": ["testing"],
        "estimated_minutes": 35,
        "kata_go": '''package kata

// ConcatStrings concatenates strings with a loop.
func ConcatStrings(parts []string) string {
\tresult := ""
\tfor _, s := range parts {
\t\tresult += s
\t}
\treturn result
}

// ConcatStringsBuilder concatenates strings using strings.Builder.
// Your task: implement this more efficient version.
func ConcatStringsBuilder(parts []string) string {
\t// Your implementation
\treturn ""
}''',
        "kata_test": '''package kata

import "testing"

func TestConcatStringsBuilder(t *testing.T) {
\tparts := []string{"hello", " ", "world"}
\tgot := ConcatStringsBuilder(parts)
\twant := "hello world"
\tif got != want {
\t\tt.Errorf("got %q, want %q", got, want)
\t}
}

func BenchmarkConcatStrings(b *testing.B) {
\tparts := []string{"a", "b", "c", "d", "e"}
\tb.Run("loop", func(b *testing.B) {
\t\tfor i := 0; i < b.N; i++ {
\t\t\tConcatStrings(parts)
\t\t}
\t})
\tb.Run("builder", func(b *testing.B) {
\t\tfor i := 0; i < b.N; i++ {
\t\t\tConcatStringsBuilder(parts)
\t\t}
\t})
}''',
        "readme": '''# Kata 126 — Benchmarking Deep Dive

**Focus:** benchmark patterns, sub-benchmarks, comparison, regression

## Your task
Implement ConcatStringsBuilder and compare performance with the loop version.

### Learning goal
- What you are practicing: comparing algorithm implementations with benchmarks.
- Why this matters: choosing the right implementation can mean 10x performance differences.
- How this grows your Go skills: you learn sub-benchmarks, benchmark comparison, and performance regression testing.

### Tips
- Use sub-benchmarks (`b.Run`) to compare implementations.
- Run `go test -bench=. -benchmem` to see allocation differences.
- Use `benchstat` to compare benchmark results statistically.

## Rules / Expectations
- ConcatStringsBuilder produces correct output
- Benchmark compares both implementations
- Builder version should be more efficient

## What this kata is about (and why it matters)
- Core lesson: benchmarks prove which implementation is actually faster.
- After this kata, you'll benchmark before choosing algorithms.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test -bench=. -benchmem ./...
```
''',
        "flashcards": [
            {"front": "How do you compare two implementations in benchmarks?", "back": "Use sub-benchmarks: b.Run(\"name\", func(b *testing.B) { ... })"},
            {"front": "What tool statistically compares benchmark results?", "back": "benchstat (golang.org/x/perf/cmd/benchstat)"},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "How do you create sub-benchmarks for comparison?", "options": ["b.Sub()", "b.Run()", "b.Compare()", "b.Bench()"], "answer": "b.Run()"},
        ],
    },
    {
        "id": "131", "slug": "security-hardening",
        "title": "Security Hardening",
        "focus": "input validation, SQL injection, XSS, secrets management",
        "stage": "senior", "category": "security", "level": "senior",
        "tags": ["architecture"],
        "estimated_minutes": 35,
        "kata_go": '''package kata

import "strings"

// SanitizeInput removes potentially dangerous characters from user input.
// Your task: implement proper input sanitization.
func SanitizeInput(input string) string {
\t// Your implementation
\treturn input
}

// ValidateEmail checks if an email address is valid.
// Your task: implement basic email validation.
func ValidateEmail(email string) bool {
\t// Your implementation
\treturn false
}''',
        "kata_test": '''package kata

import "testing"

func TestSanitizeInput(t *testing.T) {
\ttests := []struct {
\t\tname, input, want string
\t}{
\t\t{"clean", "hello", "hello"},
\t\t{"script tag", "<script>alert(1)</script>", ""},
\t\t{"sql inject", "'; DROP TABLE users;--", ""},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot := SanitizeInput(tc.input)
\t\t\tif got != tc.want {
\t\t\t\tt.Errorf("SanitizeInput(%q) = %q, want %q", tc.input, got, tc.want)
\t\t\t}
\t\t})
\t}
}

func TestValidateEmail(t *testing.T) {
\ttests := []struct {
\t\tname, email string
\t\twant bool
\t}{
\t\t{"valid", "user@example.com", true},
\t\t{"no at", "userexample.com", false},
\t\t{"no domain", "user@", false},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot := ValidateEmail(tc.email)
\t\t\tif got != tc.want {
\t\t\t\tt.Errorf("ValidateEmail(%q) = %v, want %v", tc.email, got, tc.want)
\t\t\t}
\t\t})
\t}
}''',
        "readme": '''# Kata 131 — Security Hardening

**Focus:** input validation, SQL injection, XSS, secrets management

## Your task
Implement input sanitization and email validation.

### Learning goal
- What you are practicing: defensive coding against common security vulnerabilities.
- Why this matters: security bugs are the most expensive bugs. Input validation is your first line of defense.
- How this grows your Go skills: you learn to think like an attacker and validate everything.

### Tips
- Strip or encode HTML/script tags.
- Validate email format with regex or string checks.
- Never trust user input.

## Rules / Expectations
- Script tags => removed
- SQL injection patterns => removed
- Valid email => true
- Invalid email => false

## What this kata is about (and why it matters)
- Core lesson: validate and sanitize all user input. Always.
- After this kata, you'll think about security in every function that handles input.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is XSS?", "back": "Cross-Site Scripting: injecting malicious scripts into web pages via user input."},
            {"front": "What is SQL injection?", "back": "Injecting SQL commands through user input to manipulate database queries."},
            {"front": "What is the golden rule of input handling?", "back": "Never trust user input. Validate and sanitize everything."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What is the first defense against security vulnerabilities?", "options": ["Firewall", "Input validation", "Encryption", "Authentication"], "answer": "Input validation"},
        ],
    },
    {
        "id": "147", "slug": "code-review-kata",
        "title": "Code Review Kata",
        "focus": "code review patterns, refactoring, readability, maintainability",
        "stage": "lead", "category": "code-quality-review", "level": "lead",
        "tags": ["architecture"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// ReviewCode analyzes code quality and returns improvement suggestions.
// Your task: implement the function below.
func ReviewCode(code string) []string {
\t// Return a list of code review comments
\treturn nil
}''',
        "kata_test": '''package kata

import "testing"

func TestReviewCode(t *testing.T) {
\tcode := `func f(x int) int { return x * x }`
\tcomments := ReviewCode(code)
\tif len(comments) == 0 {
\t\tt.Error("expected at least one review comment")
\t}
}''',
        "readme": '''# Kata 147 — Code Review Kata

**Focus:** code review patterns, refactoring, readability, maintainability

## Your task
Implement code review analysis.

### Learning goal
- What you are practicing: identifying code quality issues and suggesting improvements.
- Why this matters: code review is how teams maintain quality. Knowing what to look for makes you a better reviewer and writer.
- How this grows your Go skills: you learn to see code from a reviewer's perspective.

### Tips
- Look for: naming, error handling, test coverage, documentation.
- Suggest refactoring, not rewriting.
- Be constructive, not critical.

## Rules / Expectations
- Returns at least one suggestion for unclear code
- Suggestions are actionable

## What this kata is about (and why it matters)
- Core lesson: good code review is a skill. Practice it deliberately.
- After this kata, you'll give better reviews and write more reviewable code.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What makes a good code review?", "back": "Specific, actionable feedback. Focus on correctness, readability, and maintainability."},
            {"front": "What are common code review anti-patterns?", "back": "Nitpicking style, suggesting rewrites instead of improvements, being vague."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What should code review focus on first?", "options": ["Style", "Correctness", "Performance", "Documentation"], "answer": "Correctness"},
        ],
    },
    {
        "id": "148", "slug": "docker-multi-stage-build",
        "title": "Docker Multi-Stage Build",
        "focus": "Dockerfile, multi-stage builds, Go binary, containerization",
        "stage": "lead", "category": "build-deploy", "level": "lead",
        "tags": ["cli"],
        "estimated_minutes": 30,
        "kata_go": '''package kata

// BuildDockerfile generates a multi-stage Dockerfile for a Go app.
// Your task: implement the function below.
func BuildDockerfile(appName string) string {
\t// Return a multi-stage Dockerfile as a string
\treturn ""
}''',
        "kata_test": '''package kata

import (
\t"strings"
\t"testing"
)

func TestBuildDockerfile(t *testing.T) {
\tdf := BuildDockerfile("myapp")
\tif df == "" {
\t\tt.Fatal("expected non-empty Dockerfile")
\t}
\tif !strings.Contains(df, "FROM") {
\t\tt.Error("Dockerfile should contain FROM")
\t}
\tif !strings.Contains(df, "COPY") {
\t\tt.Error("Dockerfile should contain COPY")
\t}
}''',
        "readme": '''# Kata 148 — Docker Multi-Stage Build

**Focus:** Dockerfile, multi-stage builds, Go binary, containerization

## Your task
Implement a Dockerfile generator for Go applications.

### Learning goal
- What you are practicing: creating efficient Docker images for Go applications.
- Why this matters: multi-stage builds create tiny production images (10MB vs 1GB+).
- How this grows your Go skills: you learn the deployment side of Go development.

### Tips
- Use `golang` as the build stage, `alpine` or `scratch` as runtime.
- Copy only the binary in the final stage.
- Use `CGO_ENABLED=0` for static binaries.

## Rules / Expectations
- Output contains FROM, COPY, RUN, EXPOSE
- Multi-stage (at least 2 FROM statements)
- Copies binary in final stage

## What this kata is about (and why it matters)
- Core lesson: containerization is a core DevOps skill for Go developers.
- After this kata, you can containerize any Go application efficiently.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is a multi-stage Docker build?", "back": "A Dockerfile with multiple FROM statements. Build in one stage, copy only the binary to a minimal runtime stage."},
            {"front": "Why use CGO_ENABLED=0 for Go Docker images?", "back": "Creates a statically linked binary that runs on scratch/alpine without glibc."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What is the benefit of multi-stage Docker builds?", "options": ["Faster builds", "Smaller images", "Better security", "All of the above"], "answer": "All of the above"},
        ],
    },
    {
        "id": "149", "slug": "cicd-pipeline",
        "title": "CI/CD Pipeline",
        "focus": "GitHub Actions, CI/CD, automated testing, deployment",
        "stage": "lead", "category": "build-deploy", "level": "lead",
        "tags": ["testing"],
        "estimated_minutes": 30,
        "kata_go": '''package kata

// PipelineConfig represents a CI/CD pipeline configuration.
type PipelineConfig struct {
\tName    string
\tTrigger string
\tSteps   []string
}

// GeneratePipeline creates a basic CI/CD pipeline config.
// Your task: implement the function below.
func GeneratePipeline(name string) PipelineConfig {
\treturn PipelineConfig{}
}''',
        "kata_test": '''package kata

import "testing"

func TestGeneratePipeline(t *testing.T) {
\tp := GeneratePipeline("go-ci")
\tif p.Name == "" {
\t\tt.Fatal("expected non-empty name")
\t}
\tif len(p.Steps) == 0 {
\t\tt.Fatal("expected at least one step")
\t}
}''',
        "readme": '''# Kata 149 — CI/CD Pipeline

**Focus:** GitHub Actions, CI/CD, automated testing, deployment

## Your task
Implement a CI/CD pipeline configuration generator.

### Learning goal
- What you are practicing: designing automated build/test/deploy pipelines.
- Why this matters: CI/CD is how modern teams ship code. Every Go project needs it.
- How this grows your Go skills: you learn GitHub Actions, automated testing, and deployment workflows.

### Tips
- Include: checkout, build, test, lint, deploy steps.
- Use matrix builds for multiple Go versions.
- Cache dependencies for faster builds.

## Rules / Expectations
- Pipeline has a name
- Pipeline has at least one step
- Steps include test and build

## What this kata is about (and why it matters)
- Core lesson: automate everything. Manual processes break.
- After this kata, you can set up CI/CD for any Go project.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What are the stages of a CI/CD pipeline?", "back": "Build, Test, Lint, Deploy (and optionally: Security Scan, Package, Release)"},
            {"front": "What is GitHub Actions?", "back": "GitHub's CI/CD platform. Define workflows in .github/workflows/*.yml"},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which stage should run first in a CI pipeline?", "options": ["Deploy", "Test", "Build", "Lint"], "answer": "Build"},
        ],
    },
    {
        "id": "153", "slug": "architecture-decision-records",
        "title": "Architecture Decision Records",
        "focus": "ADR, technical decision documentation, RFC process",
        "stage": "lead", "category": "leadership-communication", "level": "lead",
        "tags": ["architecture"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// ADR represents an Architecture Decision Record.
type ADR struct {
\tTitle   string
\tStatus  string // "proposed", "accepted", "deprecated", "superseded"
\tContext string
\tDecision string
\tConsequences string
}

// CreateADR generates an ADR document from the given fields.
// Your task: implement the function below.
func CreateADR(title, context, decision, consequences string) ADR {
\treturn ADR{}
}''',
        "kata_test": '''package kata

import "testing"

func TestCreateADR(t *testing.T) {
\tadr := CreateADR("Use PostgreSQL", "We need a relational DB", "Use PostgreSQL for all services", "Requires DBA expertise")
\tif adr.Title == "" {
\t\tt.Fatal("expected non-empty title")
\t}
\tif adr.Status != "proposed" {
\t\tt.Errorf("expected status 'proposed', got %q", adr.Status)
\t}
}''',
        "readme": '''# Kata 153 — Architecture Decision Records

**Focus:** ADR, technical decision documentation, RFC process

## Your task
Implement an ADR generator.

### Learning goal
- What you are practicing: documenting architectural decisions for future reference.
- Why this matters: teams forget why decisions were made. ADRs preserve institutional knowledge.
- How this grows your Go skills: you learn to communicate technical decisions clearly.

### Tips
- ADRs are short (1-2 pages), not essays.
- Focus on context, decision, and consequences.
- Use a consistent format (Michael Nygard's template is popular).

## Rules / Expectations
- ADR has a title, status, context, decision, consequences
- Status defaults to "proposed"
- Document is well-structured

## What this kata is about (and why it matters)
- Core lesson: document decisions, not just code.
- After this kata, you'll ADR every significant technical choice.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is an Architecture Decision Record?", "back": "A short document capturing a significant architectural decision: context, decision, and consequences."},
            {"front": "Why are ADRs important?", "back": "They preserve why decisions were made, helping future developers understand the reasoning."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What are the three key sections of an ADR?", "options": ["Code, Tests, Deploy", "Context, Decision, Consequences", "Problem, Solution, Result", "Plan, Execute, Review"], "answer": "Context, Decision, Consequences"},
        ],
    },
    {
        "id": "154", "slug": "mentoring-scenario",
        "title": "Mentoring Scenario",
        "focus": "mentoring, code review feedback, teaching patterns, knowledge transfer",
        "stage": "lead", "category": "leadership-communication", "level": "lead",
        "tags": ["architecture"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// Feedback represents mentoring feedback on code.
type Feedback struct {
\tCategory string // "positive", "improvement", "question"
\tMessage  string
\tPriority string // "high", "medium", "low"
}

// ReviewAndMentor analyzes code and provides mentoring feedback.
// Your task: implement the function below.
func ReviewAndMentor(code string) []Feedback {
\treturn nil
}''',
        "kata_test": '''package kata

import "testing"

func TestReviewAndMentor(t *testing.T) {
\tfeedback := ReviewAndMentor("func f() {}")
\tif len(feedback) == 0 {
\t\tt.Fatal("expected at least one feedback item")
\t}
}''',
        "readme": '''# Kata 154 — Mentoring Scenario

**Focus:** mentoring, code review feedback, teaching patterns, knowledge transfer

## Your task
Implement a mentoring feedback generator.

### Learning goal
- What you are practicing: giving constructive, actionable feedback that helps others grow.
- Why this matters: lead developers grow their teams through effective mentoring, not just code review.
- How this grows your Go skills: you learn to explain Go patterns and idioms to others.

### Tips
- Balance positive feedback with improvements.
- Ask questions to guide discovery.
- Prioritize: focus on the most impactful improvements first.

## Rules / Expectations
- Returns at least one feedback item
- Includes positive, improvement, or question categories
- Has priority levels

## What this kata is about (and why it matters)
- Core lesson: mentoring is a skill. Practice giving feedback that helps people grow.
- After this kata, you'll be a more effective tech lead.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What makes mentoring feedback effective?", "back": "Specific, actionable, balanced (positive + improvement), and prioritized."},
            {"front": "How do you balance positive vs improvement feedback?", "back": "Acknowledge what's good, then focus on the highest-impact improvements."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What is the most important quality of mentoring feedback?", "options": ["It's positive", "It's specific and actionable", "It's long", "It's technical"], "answer": "It's specific and actionable"},
        ],
    },
]


def main():
    created = 0
    for kata in NEW_KATAS:
        kata_dir = KATAS_DIR / f"kata-{kata['id']}-{kata['slug']}"
        if kata_dir.exists():
            print(f"  SKIP {kata_dir.name} (exists)")
            continue

        kata_dir.mkdir(parents=True, exist_ok=True)

        # kata.go.txt
        (kata_dir / "kata.go.txt").write_text(kata["kata_go"])

        # kata_test.go.txt
        (kata_dir / "kata_test.go.txt").write_text(kata["kata_test"])

        # README.md
        (kata_dir / "README.md").write_text(kata["readme"])

        # kata.json
        meta = {
            "id": kata["id"],
            "slug": kata["slug"],
            "title": kata["title"],
            "focus": kata["focus"],
            "signature": "",
            "rules": [],
            "evaluator_status": "incomplete",
            "stage": kata["stage"],
            "category": kata["category"],
            "level": kata["level"],
            "tags": kata["tags"],
            "prerequisites": [],
            "estimated_minutes": kata["estimated_minutes"],
            "flashcards": kata["flashcards"],
            "quiz_questions": kata["quiz_questions"],
        }
        (kata_dir / "kata.json").write_text(json.dumps(meta, indent=2) + "\n")

        print(f"  CREATED {kata_dir.name}")
        created += 1

    print(f"\nCreated {created} new katas")


if __name__ == "__main__":
    main()
