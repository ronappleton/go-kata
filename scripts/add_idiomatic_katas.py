#!/usr/bin/env python3
"""Add katas for idiomatic Go patterns and conventions."""
import json, os

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
KATAS_DIR = os.path.join(ROOT, "katas")
TRACK_FILE = os.path.join(ROOT, "tracks", "go-core-100", "track.json")

NEW_KATAS = [
    {
        "id": "170",
        "slug": "error-handling-idioms",
        "title": "Error Handling Idioms",
        "focus": "if err != nil, sentinel errors, error wrapping, custom types, Is/As",
        "stage": "foundation",
        "category": "language-basics",
        "level": "junior",
        "tags": ["errors", "idioms", "wrapping", "sentinel", "errors.Is", "errors.As"],
        "prerequisites": ["003", "030"],
        "estimated_minutes": 35,
        "flashcards": [
            {"front": "What is the Go error handling pattern?", "back": "Check errors immediately after the call: result, err := fn(); if err != nil { return err }. Never ignore errors."},
            {"front": "What is a sentinel error?", "back": "A package-level variable of type error that callers can check with errors.Is(). Example: var ErrNotFound = errors.New('not found')."},
            {"front": "How do you wrap errors in Go?", "back": "Use fmt.Errorf with %w verb: fmt.Errorf('context: %w', err). This preserves the error chain for errors.Is/errors.As."},
            {"front": "When should you create a custom error type vs use sentinel errors?", "back": "Sentinel for simple 'this happened' checks. Custom types when you need to carry data (field values, context) that callers inspect with errors.As."}
        ],
        "quiz_questions": [
            {"question": "What verb preserves the error chain in fmt.Errorf?", "options": ["%v", "%s", "%w", "%d"], "answer": "%w"},
            {"question": "How do you check if an error matches a sentinel?", "options": ["err == ErrNotFound", "errors.Is(err, ErrNotFound)", "err.Error() == 'not found'", "reflect.DeepEqual(err, ErrNotFound)"], "answer": "errors.Is(err, ErrNotFound)"},
            {"question": "Should you ever ignore an error?", "options": ["Yes, if it's minor", "Yes, with _ = err", "Only in tests", "Never without explicit comment why"], "answer": "Never without explicit comment why"}
        ],
        "kata_go": '''package kata

import (
\t"errors"
\t"fmt"
\t"strings"
)

// ErrInvalidEmail is a sentinel error for invalid emails.
var ErrInvalidEmail = errors.New("invalid email format")

// ValidationError carries context about what failed.
type ValidationError struct {
\tField   string
\tMessage string
}

func (e *ValidationError) Error() string {
\treturn fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateEmail checks email format.
// Rules:
// - Must contain exactly one @
// - Must have non-empty local and domain parts
// - Domain must contain at least one dot
// - Return ErrInvalidEmail if invalid
// - Return nil if valid
func ValidateEmail(email string) error {
\t// TODO: Implement email validation
\treturn nil
}

// ValidateAge checks age is in valid range.
// Rules:
// - age < 0 or age > 150 => return &ValidationError{Field: "age", Message: "must be 0-150"}
// - otherwise => return nil
func ValidateAge(age int) error {
\t// TODO: Implement age validation
\treturn nil
}

// ClassifyError inspects an error and returns a category.
// Rules:
// - nil => "none"
// - errors.Is(err, ErrInvalidEmail) => "email"
// - *ValidationError => "validation"
// - other => "unknown"
func ClassifyError(err error) string {
\t// TODO: Classify the error type
\treturn "unknown"
}

// WrapWithContext wraps an error with operation context.
// Rules:
// - err is nil => return nil
// - err is not nil => return fmt.Errorf("%s: %w", operation, err)
func WrapWithContext(operation string, err error) error {
\t// TODO: Wrap the error with context
\treturn nil
}
''',
        "kata_test_go": '''package kata

import (
\t"errors"
\t"testing"
)

func TestValidateEmail(t *testing.T) {
\ttests := []struct {
\t\temail string
\t\twant  error
\t}{
\t{"user@example.com", nil},
\t{"test@test.co.uk", nil},
\t{"no-at-sign", ErrInvalidEmail},
\t{"@missing.com", ErrInvalidEmail},
\t{"missing@", ErrInvalidEmail},
\t{"missing.domain", ErrInvalidEmail},
\t{"", ErrInvalidEmail},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.email, func(t *testing.T) {
\t\t\tgot := ValidateEmail(tt.email)
\t\t\tif !errors.Is(got, tt.want) {
\t\t\t\tt.Errorf("ValidateEmail(%q) = %v, want %v", tt.email, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestValidateAge(t *testing.T) {
\tif err := ValidateAge(25); err != nil {
\t\tt.Errorf("ValidateAge(25) = %v, want nil", err)
\t}
\tvar ve *ValidationError
\tif err := ValidateAge(-1); !errors.As(err, &ve) {
\t\tt.Errorf("ValidateAge(-1) = %v, want *ValidationError", err)
\t}
\tif err := ValidateAge(200); !errors.As(err, &ve) {
\t\tt.Errorf("ValidateAge(200) = %v, want *ValidationError", err)
\t}
\tif ve != nil && ve.Field != "age" {
\t\tt.Errorf("Field = %q, want %q", ve.Field, "age")
\t}
}

func TestClassifyError(t *testing.T) {
\ttests := []struct {
\t\terr  error
\t\twant string
\t}{
\t{nil, "none"},
\t{ErrInvalidEmail, "email"},
\t{&ValidationError{Field: "x", Message: "bad"}, "validation"},
\t{errors.New("other"), "unknown"},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.want, func(t *testing.T) {
\t\t\tgot := ClassifyError(tt.err)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("ClassifyError(%v) = %q, want %q", tt.err, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestWrapWithContext(t *testing.T) {
\tif got := WrapWithContext("op", nil); got != nil {
\t\tt.Errorf("WrapWithContext(op, nil) = %v, want nil", got)
\t}
\terr := WrapWithContext("parse", errors.New("bad input"))
\tif err == nil {
\t\tt.Fatal("WrapWithContext(parse, err) = nil, want error")
\t}
\tif !strings.Contains(err.Error(), "parse:") {
\t\tt.Errorf("error = %q, want prefix parse:", err.Error())
\t}
\tif !errors.Is(err, errors.New("bad input")) {
\t\tt.Error("wrapped error should be unwrappable")
\t}
}
''',
        "readme": '''# Kata 170 — Error Handling Idioms

**Focus:** if err != nil, sentinel errors, error wrapping, custom types, Is/As

## Your task

Implement idiomatic error handling patterns.

### Learning goal
- What you are practicing: the Go error handling contract — always check, wrap with context, use sentinels and types appropriately.
- Why this matters: error handling is THE defining Go pattern. Bad error handling is the #1 source of bugs.
- How this grows your Go skills: you'll write robust code that fails gracefully and is debuggable.

## Rules / Expectations
- ValidateEmail returns ErrInvalidEmail sentinel
- ValidateAge returns *ValidationError with context
- ClassifyError uses errors.Is and errors.As
- WrapWithContext wraps with fmt.Errorf %w

## What this kata is about (and why it matters)
- Core lesson: Go errors are values. Handle them explicitly, wrap them with context, and use types/sentinels for callers.
- After this kata, you should handle errors idiomatically in any Go program.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "171",
        "slug": "naming-conventions",
        "title": "Naming & Stuttering",
        "focus": "No stuttering, clear names, receiver consistency, package naming",
        "stage": "foundation",
        "category": "language-basics",
        "level": "junior",
        "tags": ["naming", "stuttering", "conventions", "readability", "gofmt"],
        "prerequisites": ["167"],
        "estimated_minutes": 25,
        "flashcards": [
            {"front": "What is 'stuttering' in Go naming?", "back": "When a type name repeats the package name: http.HTTPClient is stuttering. The correct name is http.Client."},
            {"front": "What is the receiver naming convention?", "back": "Use short, consistent names. The first letter of the type name. Don't use 'this' or 'self'. Use the same name across all methods."},
            {"front": "How should package names relate to directory names?", "back": "They should match. A package in directory 'httputil' should be named 'httputil', not 'utils' or 'helpers'."},
            {"front": "What makes a good Go function name?", "back": "Verb or verb phrase. Short but clear. Avoid 'Get' prefix if the function doesn't fetch. Example: os.Open not os.GetOpen."}
        ],
        "quiz_questions": [
            {"question": "Which is idiomatic Go naming?", "options": ["client.HTTPClient", "client.Client", "HTTPClient", "GetClient"], "answer": "client.Client"},
            {"question": "What receiver name should String() use on type User?", "options": ["this", "self", "u", "user"], "answer": "u"},
            {"question": "Which package name violates conventions?", "options": ["httputil", "helpers", "http", "json"], "answer": "helpers"}
        ],
        "kata_go": '''package kata

import (
\t"strings"
\t"unicode"
)

// FixStuttering removes stuttering from Go type names.
// Rules:
// - "HTTPClient" => "Client" (remove package prefix if it matches)
// - "UserRecord" => "Record" if package is "user"
// - "ServerConfig" => "Config" if package is "server"
// - Package name is passed as context
// - If name starts with package name (case-insensitive), remove it
// - Preserve the rest of the name exactly
// - If name equals package name, return as-is
func FixStuttering(name, packageName string) string {
\t// TODO: Remove stuttering
\treturn name
}

// ValidateReceiverNames checks that all methods on a type use the same receiver.
// Rules:
// - methods is a list of "Receiver.Method" strings
// - All must use the same single-letter receiver
// - Receiver must be 1-2 characters
// - Return list of inconsistencies (empty = all good)
func ValidateReceiverNames(methods []string) []string {
\t// TODO: Check receiver consistency
\treturn nil
}

// SuggestMethodName suggests a better name following Go conventions.
// Rules:
// - If starts with "Get" and returns a value, suggest removing "Get" prefix
// - If starts with "Is" and returns bool, keep it
// - If starts with "Set" and takes a value, keep it
// - If name is all lowercase with underscores, suggest camelCase
// - Return the suggested name (may be same as input)
func SuggestMethodName(name string) string {
\t// TODO: Suggest idiomatic name
\treturn name
}

// CleanPackageName removes non-idiomatic package names.
// Rules:
// - "my-package" => "mypackage" (remove hyphens)
// - "my_package" => "mypackage" (remove underscores)
// - "MyPackage" => "mypackage" (lowercase)
// - "utils" => "util" (prefer singular, but accept both)
// - "helpers" => suggest removal (not a meaningful name)
func CleanPackageName(name string) string {
\t// TODO: Clean package name
\treturn name
}
''',
        "kata_test_go": '''package kata

import (
\t"testing"
)

func TestFixStuttering(t *testing.T) {
\ttests := []struct {
\t\tname, pkg, want string
\t}{
\t{"HTTPClient", "http", "Client"},
\t{"UserRecord", "user", "Record"},
\t{"ServerConfig", "server", "Config"},
\t{"Config", "server", "Config"},    // no stuttering
\t{"Client", "http", "Client"},      // no stuttering
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.name, func(t *testing.T) {
\t\t\tgot := FixStuttering(tt.name, tt.pkg)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("FixStuttering(%q, %q) = %q, want %q", tt.name, tt.pkg, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestValidateReceiverNames(t *testing.T) {
\tgood := []string{"u.Name", "u.Email", "u.Save"}
\tif bad := ValidateReceiverNames(good); len(bad) != 0 {
\t\tt.Errorf("ValidateReceiverNames(%v) = %v, want []", good, bad)
\t}

\tbad := []string{"u.Name", "user.Email", "u.Save"}
\tif bad2 := ValidateReceiverNames(bad); len(bad2) == 0 {
\t\tt.Errorf("ValidateReceiverNames(%v) = [], want issues", bad)
\t}
}

func TestSuggestMethodName(t *testing.T) {
\ttests := []struct {
\t\tname, want string
\t}{
\t{"GetName", "Name"},
\t{"IsAdmin", "IsAdmin"},
\t{"SetValue", "SetValue"},
\t{"get_name", "getName"},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.name, func(t *testing.T) {
\t\t\tgot := SuggestMethodName(tt.name)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("SuggestMethodName(%q) = %q, want %q", tt.name, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestCleanPackageName(t *testing.T) {
\ttests := []struct {
\t\tname, want string
\t}{
\t{"my-package", "mypackage"},
\t{"my_package", "mypackage"},
\t{"MyPackage", "mypackage"},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.name, func(t *testing.T) {
\t\t\tgot := CleanPackageName(tt.name)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("CleanPackageName(%q) = %q, want %q", tt.name, got, tt.want)
\t\t\t}
\t\t})
\t}
}
''',
        "readme": '''# Kata 171 — Naming & Stuttering

**Focus:** No stuttering, clear names, receiver consistency, package naming

## Your task

Fix non-idiomatic Go names and validate naming conventions.

### Learning goal
- What you are practicing: Go's naming conventions that make code readable and self-documenting.
- Why this matters: good naming is the most important refactoring. Stuttering names confuse readers.
- How this grows your Go skills: you'll write code that reads like prose, not code.

## Rules / Expectations
- FixStuttering removes package prefix from type names
- ValidateReceiverNames checks method receiver consistency
- SuggestMethodName follows Go naming conventions
- CleanPackageName removes non-idiomatic patterns

## What this kata is about (and why it matters)
- Core lesson: Go values clarity over cleverness. Names should be short, consistent, and non-stuttering.
- After this kata, you should name anything in Go idiomatically.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "172",
        "slug": "composition-patterns",
        "title": "Composition Over Inheritance",
        "focus": "Struct embedding, small interfaces, accept interfaces return structs",
        "stage": "foundation",
        "category": "language-basics",
        "level": "junior",
        "tags": ["composition", "embedding", "interfaces", "structural-typing"],
        "prerequisites": ["040", "167"],
        "estimated_minutes": 30,
        "flashcards": [
            {"front": "What is struct embedding in Go?", "back": "Including a struct field without a name. The embedded type's methods are promoted to the outer type, giving the appearance of inheritance."},
            {"front": "How does Go achieve polymorphism?", "back": "Through interfaces. Any type that satisfies an interface's method set implicitly implements it — no 'implements' keyword needed."},
            {"front": "What does 'accept interfaces, return structs' mean?", "back": "Function parameters should be interfaces for flexibility. Return concrete types for clarity and options. This is Go's core design principle."},
            {"front": "What is a small interface?", "back": "An interface with 1-3 methods. Go's best interfaces (io.Reader, fmt.Stringer) are tiny. Large interfaces are a design smell."}
        ],
        "quiz_questions": [
            {"question": "How does Go implement polymorphism?", "options": ["Class inheritance", "Struct embedding", "Interfaces with implicit satisfaction", "Abstract base classes"], "answer": "Interfaces with implicit satisfaction"},
            {"question": "What does 'accept interfaces, return structs' give you?", "options": ["Better performance", "Flexibility in inputs, clarity in outputs", "Stronger type safety", "Easier testing only"], "answer": "Flexibility in inputs, clarity in outputs"},
            {"question": "How many methods should a good Go interface have?", "options": ["As many as needed", "1-3 ideally", "At least 5", "Exactly 1"], "answer": "1-3 ideally"}
        ],
        "kata_go": '''package kata

import (
\t"fmt"
\t"strings"
)

// --- Small Interfaces (Go style) ---

// Stringer converts a value to a display string.
type Stringer interface {
\tString() string
}

// Counter tracks a count.
type Counter interface {
\tIncrement()
\tCount() int
}

// --- Struct Embedding ---

// BaseItem provides common fields.
type BaseItem struct {
\tID   string
\tName string
\tTags []string
}

// String implements Stringer.
func (b BaseItem) String() string {
\treturn fmt.Sprintf("[%s] %s", b.ID, b.Name)
}

// HasTag checks if the item has a specific tag.
func (b BaseItem) HasTag(tag string) bool {
\tfor _, t := range b.Tags {
\t\tif t == tag {
\t\t\treturn true
\t\t}
\t}
\treturn false
}

// Task embeds BaseItem and adds task-specific fields.
type Task struct {
\tBaseItem
\tDone     bool
\tPriority int
}

// TodoList embeds BaseItem and manages tasks.
type TodoList struct {
\tBaseItem
\tTasks []Task
}

// --- Accept Interfaces, Return Structs ---

// DescribeItem accepts any Stringer — maximum flexibility.
func DescribeItem(s Stringer) string {
\treturn "Item: " + s.String()
}

// FilterByTag accepts any item that has a HasTag method.
// In real Go this would use an interface, but here we demonstrate the pattern.
type TagChecker interface {
\tHasTag(tag string) bool
}

// FilterItems returns items matching a tag.
func FilterItems(items []TagChecker, tag string) []TagChecker {
\tvar result []TagChecker
\tfor _, item := range items {
\t\tif item.HasTag(tag) {
\t\t\tresult = append(result, item)
\t\t}
\t}
\treturn result
}

// --- Composition helpers ---

// NewCounter creates a functional counter.
func NewCounter() Counter {
\treturn &counterImpl{count: 0}
}

type counterImpl struct {
\tcount int
}

func (c *counterImpl) Increment() { c.count++ }
func (c *counterImpl) Count() int { return c.count }
''',
        "kata_test_go": '''package kata

import (
\t"strings"
\t"testing"
)

func TestBaseItemString(t *testing.T) {
\titem := BaseItem{ID: "1", Name: "Test"}
\tgot := item.String()
\tif !strings.Contains(got, "1") || !strings.Contains(got, "Test") {
\t\tt.Errorf("String() = %q, want '1' and 'Test'", got)
\t}
}

func TestBaseItemHasTag(t *testing.T) {
\titem := BaseItem{ID: "1", Name: "Test", Tags: []string{"urgent", "backend"}}
\tif !item.HasTag("urgent") {
\t\tt.Error("HasTag('urgent') = false, want true")
\t}
\tif item.HasTag("frontend") {
\t\tt.Error("HasTag('frontend') = true, want false")
\t}
}

func TestTaskEmbedding(t *testing.T) {
\ttask := Task{
\t\tBaseItem: BaseItem{ID: "t1", Name: "Do thing"},
\t\tDone:     false,
\t\tPriority: 1,
\t}
\t// Promoted method
\tif task.HasTag("x") {
\t\tt.Error("empty tags should not match")
\t}
\ttask.Tags = []string{"go"}
\tif !task.HasTag("go") {
\t\tt.Error("HasTag should work through embedding")
\t}
}

func TestDescribeItem(t *testing.T) {
\titem := BaseItem{ID: "42", Name: "Widget"}
\tdesc := DescribeItem(item) // item satisfies Stringer via promoted String()
\tif !strings.Contains(desc, "42") {
\t\tt.Errorf("DescribeItem() = %q, want contain '42'", desc)
\t}
}

func TestFilterItems(t *testing.T) {
\titems := []TagChecker{
\t\tBaseItem{ID: "1", Tags: []string{"go"}},
\t\tBaseItem{ID: "2", Tags: []string{"rust"}},
\t\tBaseItem{ID: "3", Tags: []string{"go", "web"}},
\t}
\tfiltered := FilterItems(items, "go")
\tif len(filtered) != 2 {
\t\tt.Errorf("FilterItems(go) returned %d items, want 2", len(filtered))
\t}
}

func TestCounter(t *testing.T) {
\tc := NewCounter()
\tif c.Count() != 0 {
\t\tt.Errorf("new counter count = %d, want 0", c.Count())
\t}
\tc.Increment()
\tc.Increment()
\tif c.Count() != 2 {
\t\tt.Errorf("count = %d, want 2", c.Count())
\t}
}
''',
        "readme": '''# Kata 172 — Composition Over Inheritance

**Focus:** Struct embedding, small interfaces, accept interfaces return structs

## Your task

Build composable types using embedding and small interfaces.

### Learning goal
- What you are practicing: Go's composition model — struct embedding for reuse, small interfaces for flexibility.
- Why this matters: Go has no class inheritance. Composition is the Go way to build complex types from simple ones.
- How this grows your Go skills: you'll design flexible, testable systems using Go's type system.

## Rules / Expectations
- BaseItem provides common String() and HasTag()
- Task and TodoList embed BaseItem
- DescribeItem accepts any Stringer (interface flexibility)
- FilterItems works with any TagChecker

## What this kata is about (and why it matters)
- Core lesson: Go uses composition, not inheritance. Small interfaces + struct embedding = clean, testable design.
- After this kata, you should design composable Go types naturally.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "173",
        "slug": "defer-patterns",
        "title": "Defer Patterns & Cleanup",
        "focus": "defer, resource cleanup, panic recovery, named returns",
        "stage": "foundation",
        "category": "language-basics",
        "level": "junior",
        "tags": ["defer", "cleanup", "panic", "recovery", "resource-management"],
        "prerequisites": ["166"],
        "estimated_minutes": 25,
        "flashcards": [
            {"front": "When does a deferred function execute?", "back": "After the surrounding function returns (or panics). Deferred calls execute in LIFO order (last deferred, first executed)."},
            {"front": "What is the most common use of defer?", "back": "Resource cleanup: closing files, releasing locks, closing database connections. defer ensures cleanup happens even if the function panics."},
            {"front": "Can you recover from a panic in a deferred function?", "back": "Yes. Call recover() inside a deferred function. If a panic is active, recover() returns the panic value. Otherwise it returns nil."},
            {"front": "What is a gotcha with defer and loops?", "back": "defer executes when the FUNCTION returns, not the loop iteration. In a loop, defer will accumulate until the function exits. Extract the loop body into a helper function."}
        ],
        "quiz_questions": [
            {"question": "What is the execution order of deferred calls?", "options": ["FIFO", "LIFO", "Random", "Parallel"], "answer": "LIFO"},
            {"question": "When does a deferred function run?", "back": "After the surrounding function returns", "options": ["Immediately", "After the surrounding function returns", "At program exit", "In a goroutine"], "answer": "After the surrounding function returns"},
            {"question": "How do you stop a panic from crashing the program?", "options": ["try/catch", "recover() in defer", "ignore()", "panic(false)"], "answer": "recover() in defer"}
        ],
        "kata_go": '''package kata

import (
\t"fmt"
\t"strings"
)

// CaptureOutput simulates defer-based resource cleanup.
// Rules:
// - Record starts empty
// - "open" appends "opened"
// - "close" appends "closed"
// - Must always pair open/close regardless of errors
// - Returns the recorded actions
func CaptureOutput(actions []string) []string {
\t// TODO: Implement using defer for close
\treturn nil
}

// SafeDivide divides a/b with deferred error recovery.
// Rules:
// - If b == 0, panic with "divide by zero"
// - Use defer/recover to catch the panic
// - Return 0 and error message on panic
// - Return quotient and nil error on success
func SafeDivide(a, b int) (int, string) {
\t// TODO: Implement with defer/recover
\treturn 0, ""
}

// Pipeline chains operations with cleanup.
// Rules:
// - stages is a list of operation names
// - For each stage, record "started: <name>" and "completed: <name>"
// - Use defer to ensure "completed" is always recorded
// - If a stage contains "fail", record "started: <name>" then panic
// - recover from panic, record "failed: <name>"
// - Return all recorded messages
func Pipeline(stages []string) []string {
\tvar log []string
\t// TODO: Implement pipeline with defer-based cleanup
\treturn log
}

// DeferredMultiply demonstrates argument evaluation.
// Rules:
// - x and y are multiplied immediately when defer is called
// - The result is computed and stored
// - If called with (3, 4), the deferred result is 12
// - Return the result via a channel-like pattern (here, just return the value)
func DeferredMultiply(x, y int) int {
\tresult := 0
\tdefer func() {
\t\t// This runs after the function returns
\t\t_ = result // result was set before defer
\t}()
\t// TODO: Set result = x * y, then use defer to confirm it
\tresult = x * y
\treturn result
}
''',
        "kata_test_go": '''package kata

import (
\t"strings"
\t"testing"
)

func TestCaptureOutput(t *testing.T) {
\tgot := CaptureOutput([]string{"open", "close"})
\twant := []string{"opened", "closed"}
\tif len(got) != len(want) {
\t\tt.Fatalf("CaptureOutput() returned %d items, want %d", len(got), len(want))
\t}
\tfor i, w := range want {
\t\tif got[i] != w {
\t\t\tt.Errorf("CaptureOutput()[%d] = %q, want %q", i, got[i], w)
\t\t}
\t}
}

func TestSafeDivide(t *testing.T) {
\tquotient, err := SafeDivide(10, 2)
\tif quotient != 5 || err != "" {
\t\tt.Errorf("SafeDivide(10, 2) = (%d, %q), want (5, '')", quotient, err)
\t}

\t_, err = SafeDivide(10, 0)
\tif err == "" {
\t\tt.Error("SafeDivide(10, 0) should return error")
\t}
\tif !strings.Contains(err, "divide by zero") {
\t\tt.Errorf("error = %q, want contain 'divide by zero'", err)
\t}
}

func TestPipeline(t *testing.T) {
\tgot := Pipeline([]string{"build", "test", "deploy"})
\tif len(got) != 6 {
\t\tt.Fatalf("Pipeline() returned %d messages, want 6: %v", len(got), got)
\t}
\t// Each stage should have started and completed
\tfor i := 0; i < len(got); i += 2 {
\t\tif !strings.HasPrefix(got[i], "started:") {
\t\t\tt.Errorf("got[%d] = %q, want 'started:' prefix", i, got[i])
\t\t}
\t\tif !strings.HasPrefix(got[i+1], "completed:") {
\t\t\tt.Errorf("got[%d] = %q, want 'completed:' prefix", i+1, got[i+1])
\t\t}
\t}
}

func TestPipelineWithFailure(t *testing.T) {
\tgot := Pipeline([]string{"build", "fail-test", "deploy"})
\t// build starts/completes, fail-test starts then fails
\tif len(got) < 3 {
\t\tt.Fatalf("Pipeline() with fail returned %d messages, want >= 3: %v", len(got), got)
\t}
\tfound := false
\tfor _, msg := range got {
\t\tif strings.Contains(msg, "failed:") {
\t\t\tfound = true
\t\t\tbreak
\t\t}
\t}
\tif !found {
\t\tt.Errorf("Pipeline() should record 'failed:' message, got: %v", got)
\t}
}

func TestDeferredMultiply(t *testing.T) {
\tif got := DeferredMultiply(3, 4); got != 12 {
\t\tt.Errorf("DeferredMultiply(3, 4) = %d, want 12", got)
\t}
}
''',
        "readme": '''# Kata 173 — Defer Patterns & Cleanup

**Focus:** defer, resource cleanup, panic recovery, named returns

## Your task

Use defer for cleanup and error recovery.

### Learning goal
- What you are practicing: defer for guaranteed cleanup, recover() for panic safety, and understanding defer timing.
- Why this matters: resource leaks are silent killers. defer ensures cleanup happens. recover prevents cascading panics.
- How this grows your Go skills: you'll write robust code that cleans up after itself.

## Rules / Expectations
- CaptureOutput pairs open/close with defer
- SafeDivide uses recover() to catch divide-by-zero panics
- Pipeline ensures completed/failed is always recorded
- DeferredMultiply demonstrates argument evaluation timing

## What this kata is about (and why it matters)
- Core lesson: defer is Go's RAII equivalent. Use it for file/lock/connection cleanup. Use recover for defensive boundaries.
- After this kata, you should use defer confidently and idiomatically.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "174",
        "slug": "nil-empty-slices",
        "title": "Nil vs Empty Slices & Maps",
        "focus": "nil vs empty, initialization patterns, when to use which",
        "stage": "foundation",
        "category": "language-basics",
        "level": "junior",
        "tags": ["nil", "empty", "slices", "maps", "initialization", "zero-values"],
        "prerequisites": ["163"],
        "estimated_minutes": 20,
        "flashcards": [
            {"front": "What is the difference between nil and empty slices?", "back": "nil slice: var s []int (no backing array). Empty slice: s := []int{} or s := make([]int, 0) (has backing array). Both have len 0, but JSON marshals differently."},
            {"front": "When should you return nil vs empty slices?", "back": "Return nil for 'no data exists'. Return empty for 'data exists but is empty'. JSON consumers prefer [] over null, so prefer empty slices for API responses."},
            {"front": "How do you safely append to a nil slice?", "back": "Just use append. Go handles nil slices gracefully: append(nil, 1) works and returns [1]. No special nil check needed."},
            {"front": "What is the zero value of a map?", "back": "nil. A nil map behaves like an empty map for reads (returns zero value), but writing to a nil map panics. Always initialize maps before writing."}
        ],
        "quiz_questions": [
            {"question": "What does len(nil slice) return?", "options": ["panic", "-1", "0", "undefined"], "answer": "0"},
            {"question": "What happens when you write to a nil map?", "options": ["Creates the map", "Returns zero value", "Panics", "Silently ignores"], "answer": "Panics"},
            {"question": "How do JSON encoders handle nil vs empty slices?", "options": ["Same behavior", "nil => null, empty => []", "nil => [], empty => null", "Both become null"], "answer": "nil => null, empty => []"}
        ],
        "kata_go": '''package kata

// ClassifySlice checks if a slice is nil, empty, or has values.
// Rules:
// - nil => "nil"
// - len 0 (non-nil) => "empty"
// - len > 0 => "populated"
func ClassifySlice(s []int) string {
\t// TODO: Classify the slice
\treturn ""
}

// SafeMapGet retrieves a value from a map safely.
// Rules:
// - If map is nil, return ("", false)
// - If key not found, return ("", false)
// - If found, return (value, true)
func SafeMapGet(m map[string]string, key string) (string, bool) {
\t// TODO: Safely get value
\treturn "", false
}

// EnsureEmpty returns an empty (non-nil) slice.
// Rules:
// - If input is nil, return []string{}
// - If input has values, return a copy
// - Always return non-nil slice
func EnsureEmpty(s []string) []string {
\t// TODO: Ensure non-nil return
\treturn nil
}

// NormalizeMap ensures all values are non-empty strings.
// Rules:
// - Return a new map
// - Keys preserved
// - Empty string values replaced with "default"
// - Nil input => empty map (not nil)
func NormalizeMap(m map[string]string) map[string]string {
\t// TODO: Normalize values
\treturn nil
}
''',
        "kata_test_go": '''package kata

import (
\t"testing"
)

func TestClassifySlice(t *testing.T) {
\tvar nilSlice []int
\tif got := ClassifySlice(nilSlice); got != "nil" {
\t\tt.Errorf("ClassifySlice(nil) = %q, want %q", got, "nil")
\t}

\tempty := []int{}
\tif got := ClassifySlice(empty); got != "empty" {
\t\tt.Errorf("ClassifySlice([]) = %q, want %q", got, "empty")
\t}

\tpopulated := []int{1, 2, 3}
\tif got := ClassifySlice(populated); got != "populated" {
\t\tt.Errorf("ClassifySlice([1,2,3]) = %q, want %q", got, "populated")
\t}
}

func TestSafeMapGet(t *testing.T) {
\tvar nilMap map[string]string
\tif _, ok := SafeMapGet(nilMap, "x"); ok {
\t\tt.Error("SafeMapGet(nil, x) should return false")
\t}

\tm := map[string]string{"a": "1", "b": "2"}
\tval, ok := SafeMapGet(m, "a")
\tif !ok || val != "1" {
\t\tt.Errorf("SafeMapGet(m, a) = (%q, %v), want ('1', true)", val, ok)
\t}

\t_, ok = SafeMapGet(m, "z")
\tif ok {
\t\tt.Error("SafeMapGet(m, z) should return false")
\t}
}

func TestEnsureEmpty(t *testing.T) {
\tvar nilSlice []string
\tgot := EnsureEmpty(nilSlice)
\tif got == nil {
\t\tt.Error("EnsureEmpty(nil) returned nil, want non-nil empty slice")
\t}
\tif len(got) != 0 {
\t\tt.Errorf("EnsureEmpty(nil) len = %d, want 0", len(got))
\t}

\tnonEmpty := []string{"a", "b"}
\tgot = EnsureEmpty(nonEmpty)
\tif len(got) != 2 {
\t\tt.Errorf("EnsureEmpty([a,b]) len = %d, want 2", len(got))
\t}
\t// Should be a copy
\tgot[0] = "x"
\tif nonEmpty[0] == "x" {
\t\tt.Error("EnsureEmpty should return a copy, not mutate input")
\t}
}

func TestNormalizeMap(t *testing.T) {
\tvar nilMap map[string]string
\tgot := NormalizeMap(nilMap)
\tif got == nil {
\t\tt.Error("NormalizeMap(nil) returned nil, want non-nil empty map")
\t}

\tm := map[string]string{"a": "val", "b": "", "c": "x"}
\tgot = NormalizeMap(m)
\tif got["a"] != "val" {
\t\tt.Errorf("NormalizeMap()[a] = %q, want 'val'", got["a"])
\t}
\tif got["b"] != "default" {
\t\tt.Errorf("NormalizeMap()[b] = %q, want 'default'", got["b"])
\t}
\tif got["c"] != "x" {
\t\tt.Errorf("NormalizeMap()[c] = %q, want 'x'", got["c"])
\t}
\t// Should not mutate original
\tif m["b"] != "" {
\t\tt.Error("NormalizeMap should not mutate input")
\t}
}
''',
        "readme": '''# Kata 174 — Nil vs Empty Slices & Maps

**Focus:** nil vs empty, initialization patterns, when to use which

## Your task

Handle nil and empty collections correctly.

### Learning goal
- What you are practicing: understanding nil vs empty slices/maps, safe operations, and API-friendly patterns.
- Why this matters: nil vs empty bugs cause JSON null issues, panics on nil maps, and confusing behavior.
- How this grows your Go skills: you'll write predictable code that handles zero values gracefully.

## Rules / Expectations
- ClassifySlice distinguishes nil, empty, populated
- SafeMapGet handles nil maps
- EnsureEmpty returns non-nil empty slices
- NormalizeMap produces non-nil maps

## What this kata is about (and why it matters)
- Core lesson: nil slices are valid but behave differently from empty slices. Maps must be initialized before writes.
- After this kata, you should handle nil/empty collections without fear.

## What you must submit for marking
- `kata.go`
'''
    }
]

# Add to language-basics category in foundation
for kata in NEW_KATAS:
    kid = kata["id"]
    slug = kata["slug"]
    dirname = f"kata-{kid.zfill(3)}-{slug}"
    dirpath = os.path.join(KATAS_DIR, dirname)
    os.makedirs(dirpath, exist_ok=True)
    
    meta = {
        "id": kid,
        "title": kata["title"],
        "slug": slug,
        "focus": kata["focus"],
        "stage": kata["stage"],
        "category": kata["category"],
        "level": kata["level"],
        "tags": kata["tags"],
        "prerequisites": kata["prerequisites"],
        "estimated_minutes": kata["estimated_minutes"],
        "evaluator_status": "ready",
        "flashcards": kata["flashcards"],
        "quiz_questions": kata["quiz_questions"]
    }
    with open(os.path.join(dirpath, "kata.json"), "w") as f:
        json.dump(meta, f, indent=2)
    with open(os.path.join(dirpath, "kata.go.txt"), "w") as f:
        f.write(kata["kata_go"])
    with open(os.path.join(dirpath, "kata_test.go.txt"), "w") as f:
        f.write(kata["kata_test_go"])
    with open(os.path.join(dirpath, "README.md"), "w") as f:
        f.write(kata["readme"])
    print(f"Created: {dirname}")

# Update track.json
with open(TRACK_FILE) as f:
    track = json.load(f)

new_ids = [k["id"] for k in NEW_KATAS]
for stage in track["stages"]:
    if stage["id"] == "foundation":
        for cat in stage["categories"]:
            if cat["id"] == "language-basics":
                # Add after existing katas
                cat["kata_ids"].extend(new_ids)
                print(f"Updated language-basics: {len(cat['kata_ids'])} katas")
                break

with open(TRACK_FILE, "w") as f:
    json.dump(track, f, indent=2)
print("Done!")
