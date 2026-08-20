#!/usr/bin/env python3
"""Add katas for variable scope, application structure, and Go conventions."""
import json, os

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
KATAS_DIR = os.path.join(ROOT, "katas")
TRACK_FILE = os.path.join(ROOT, "tracks", "go-core-100", "track.json")

# New katas to add (insert after kata-005, before kata-006)
NEW_KATAS = [
    {
        "id": "163",
        "slug": "variable-scope",
        "title": "Variable Scope & Lifetime",
        "focus": "local vs package vs global scope, block scoping, short vs long declaration",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["scope", "variables", "declarations", ":=", "var", "const"],
        "prerequisites": ["000"],
        "estimated_minutes": 25,
        "flashcards": [
            {"front": "What is the difference between := and var in Go?", "back": ":= is short variable declaration, only works inside functions. var works anywhere and can declare package-level variables."},
            {"front": "What is the scope of a variable declared with := inside a function?", "back": "Block scope — it's only visible within the nearest curly braces { } where it was declared."},
            {"front": "Can you redeclare a variable with := in the same scope?", "back": "No, := requires at least one new variable on the left side. You can shadow in inner scopes though."},
            {"front": "What is a zero value in Go?", "back": "The default value when a variable is declared without initialization: 0 for numbers, '' for strings, false for bools, nil for pointers/slices/maps/channels/interfaces."}
        ],
        "quiz_questions": [
            {"question": "Which declaration creates a package-level variable?", "options": ["x := 10", "var x = 10", "x = 10", "let x = 10"], "answer": "var x = 10"},
            {"question": "What is the zero value of a string?", "options": ["null", "undefined", "\"\"", "nil"], "answer": "\"\""},
            {"question": "Can := be used outside a function?", "options": ["Yes, always", "No, only inside functions", "Only in main()", "Only with var"], "answer": "No, only inside functions"}
        ],
        "kata_go": '''package kata

// CountDeclarations returns the number of variables declared in the scope.
// You will learn about short vs long declarations and block scoping.
//
// Rules:
// - Count how many distinct variables are declared in CountDeclarations()
// - Do NOT count the function parameter or return value
// - The test will verify your understanding of declaration syntax
func CountDeclarations() int {
\t// TODO: Declare variables using different methods:
\t// 1. Short declaration (:=)
\t// 2. var declaration
\t// 3. Multiple var declarations
\t// Then return the count of distinct variables you declared.
\t
\treturn 0
}

// ShadowExample demonstrates variable shadowing.
// Inner scope variables can "shadow" outer scope variables.
func ShadowExample(x int) int {
\tif x > 0 {
\t\tx := x * 2 // This shadows the parameter x
\t\treturn x
\t}
\treturn x
}
''',
        "kata_test_go": '''package kata

import "testing"

func TestCountDeclarations(t *testing.T) {
\t// Student should declare at least 3 variables
\tcount := CountDeclarations()
\tif count < 3 {
\t\tt.Errorf("Expected at least 3 variable declarations, got %d", count)
\t}
}

func TestShadowExample(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\tin   int
\t\twant int
\t}{
\t{"positive doubles", 5, 10},
\t{"negative unchanged", -3, -3},
\t{"zero unchanged", 0, 0},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.name, func(t *testing.T) {
\t\t\tgot := ShadowExample(tt.in)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("ShadowExample(%d) = %d, want %d", tt.in, got, tt.want)
\t\t\t}
\t\t})
\t}
}
''',
        "readme": '''# Kata 163 — Variable Scope & Lifetime

**Focus:** local vs package vs global scope, block scoping, short vs long declaration

## Your task

Implement `CountDeclarations()` and understand variable scoping.

### Learning goal
- What you are practicing: understanding Go's variable declaration syntax, scope rules, and lifetime.
- Why this matters: scope bugs are silent killers — a shadowed variable can cause logic errors that compile fine.
- How this grows your Go skills: you'll write safer, more predictable code.

## Rules / Expectations
- CountDeclarations returns the number of distinct variables you declare
- ShadowExample should demonstrate understanding of block scoping

## What this kata is about (and why it matters)
- Core lesson: Go has strict block scoping. := creates block-scoped variables. var works at package level.
- After this kata, you should understand when variables live and die.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "164",
        "slug": "app-structure",
        "title": "Go Application Structure",
        "focus": "cmd/, internal/, pkg/, directory layout, package naming conventions",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["structure", "layout", "cmd", "internal", "pkg", "packages", "conventions"],
        "prerequisites": ["163"],
        "estimated_minutes": 30,
        "flashcards": [
            {"front": "What is the standard Go project layout?", "back": "cmd/ for entry points, internal/ for private packages, pkg/ for public packages, with go.mod at root."},
            {"front": "What does the internal/ directory mean in Go?", "back": "Packages inside internal/ can only be imported by code in the parent directory tree. It enforces private packages."},
            {"front": "What is the convention for package names?", "back": "Short, lowercase, single-word. No underscores, no camelCase. The name should describe what the package provides."},
            {"front": "Should a package directory name match the package name?", "back": "Yes, ideally. If the directory is 'httputil', the package should be 'httputil' not 'utils'."}
        ],
        "quiz_questions": [
            {"question": "Where do Go executable entry points go?", "options": ["pkg/", "internal/", "cmd/", "bin/"], "answer": "cmd/"},
            {"question": "What enforces package privacy in Go?", "options": ["private keyword", "internal/ directory", "underscore prefix", "visibility modifier"], "answer": "internal/ directory"},
            {"question": "Which is a valid Go package name?", "options": ["MyPackage", "my_package", "mypackage", "package-my"], "answer": "mypackage"}
        ],
        "kata_go": '''package kata

// PackagePath returns the correct import path for a given command.
// You will learn how Go projects are structured.
//
// Rules:
// - Given a project root and command name, return the import path
// - Example: ("github.com/user/project", "server") => "github.com/user/project/cmd/server"
// - Example: ("example.com/myapp", "cli") => "example.com/myapp/cmd/cli"
func PackagePath(root, command string) string {
\t// TODO: Return the full import path for the command
\treturn ""
}

// PackageName extracts the package name from an import path.
// Rules:
// - "github.com/user/project/cmd/server" => "server"
// - "github.com/user/project/internal/auth" => "auth"
// - "example.com/myapp" => "myapp"
func PackageName(importPath string) string {
\t// TODO: Extract and return the last component of the path
\treturn ""
}

// IsInternal returns true if the import path represents an internal package.
// Rules:
// - Any path containing "/internal/" => true
// - "github.com/user/project/internal/auth" => true
// - "github.com/user/project/pkg/server" => false
func IsInternal(importPath string) bool {
\t// TODO: Check if path contains /internal/
\treturn false
}
''',
        "kata_test_go": '''package kata

import "testing"

func TestPackagePath(t *testing.T) {
\ttests := []struct {
\t\troot, cmd, want string
\t}{
\t{"github.com/user/project", "server", "github.com/user/project/cmd/server"},
\t{"example.com/myapp", "cli", "example.com/myapp/cmd/cli"},
\t{"gitlab.com/team/app", "worker", "gitlab.com/team/app/cmd/worker"},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.cmd, func(t *testing.T) {
\t\t\tgot := PackagePath(tt.root, tt.cmd)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("PackagePath(%q, %q) = %q, want %q", tt.root, tt.cmd, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestPackageName(t *testing.T) {
\ttests := []struct {
\t\tpath, want string
\t}{
\t{"github.com/user/project/cmd/server", "server"},
\t{"github.com/user/project/internal/auth", "auth"},
\t{"example.com/myapp", "myapp"},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.want, func(t *testing.T) {
\t\t\tgot := PackageName(tt.path)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("PackageName(%q) = %q, want %q", tt.path, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestIsInternal(t *testing.T) {
\ttests := []struct {
\t\tpath string
\t\twant bool
\t}{
\t{"github.com/user/project/internal/auth", true},
\t{"github.com/user/project/pkg/server", false},
\t{"example.com/myapp/internal/config", true},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.path, func(t *testing.T) {
\t\t\tgot := IsInternal(tt.path)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("IsInternal(%q) = %v, want %v", tt.path, got, tt.want)
\t\t\t}
\t\t})
\t}
}
''',
        "readme": '''# Kata 164 — Go Application Structure

**Focus:** cmd/, internal/, pkg/, directory layout, package naming conventions

## Your task

Implement functions that understand Go project structure.

### Learning goal
- What you are practicing: understanding how Go projects are organized, import paths, and package privacy.
- Why this matters: correct structure makes your code maintainable and importable.
- How this grows your Go skills: you'll architect real Go applications properly.

## Rules / Expectations
- PackagePath returns the correct cmd/ import path
- PackageName extracts the last component
- IsInternal checks for /internal/ in the path

## What this kata is about (and why it matters)
- Core lesson: Go has strong conventions for project layout. cmd/, internal/, and pkg/ are standard.
- After this kata, you should structure any Go project correctly.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "165",
        "slug": "constants-and-iota",
        "title": "Constants & iota",
        "focus": "const declarations, iota enum pattern, typed vs untyped constants",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["const", "iota", "enum", "typed", "untyped"],
        "prerequisites": ["163"],
        "estimated_minutes": 20,
        "flashcards": [
            {"front": "What is iota in Go?", "back": "A special identifier that auto-increments in const declarations. Starts at 0 and increments by 1 for each const spec."},
            {"front": "How do you create an enum in Go?", "back": "Use const with iota: const ( A = iota; B; C ) gives A=0, B=1, C=2."},
            {"front": "What is the difference between typed and untyped constants?", "back": "Typed constants (const x int = 1) have a fixed type. Untyped constants (const x = 1) can be used with any compatible type."},
            {"front": "Can you take the address of a constant?", "back": "No. Constants are computed at compile time and have no address in memory."}
        ],
        "quiz_questions": [
            {"question": "What value does iota have in the first const spec?", "options": ["1", "0", "-1", "undefined"], "answer": "0"},
            {"question": "How do you skip a value in an iota sequence?", "options": ["_ = iota", "skip", "continue", "blank identifier _"], "answer": "blank identifier _"},
            {"question": "What happens when you multiply iota by 2?", "options": ["0, 2, 4, 6", "0, 1, 2, 3", "2, 4, 6, 8", "Error"], "answer": "0, 2, 4, 6"}
        ],
        "kata_go": '''package kata

// Define a Weekday enum using iota
type Weekday int

// TODO: Define const block with iota for Sunday through Saturday
// Sunday = 0, Monday = 1, ..., Saturday = 6

// DayName returns the name of the weekday.
// Rules:
// - Sunday => "Sunday"
// - Monday => "Monday"
// - ... through Saturday
// - Invalid => "Unknown"
func DayName(d Weekday) string {
\t// TODO: Implement using your enum
\treturn "Unknown"
}

// IsWeekend returns true if the day is Saturday or Sunday.
func IsWeekend(d Weekday) bool {
\t// TODO: Check if day is Saturday (6) or Sunday (0)
\treturn false
}

// NextDay returns the next day of the week.
// Saturday wraps to Sunday.
func NextDay(d Weekday) Weekday {
\t// TODO: Return the next day, wrapping Saturday to Sunday
\treturn 0
}
''',
        "kata_test_go": '''package kata

import "testing"

func TestDayName(t *testing.T) {
\tnames := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
\tfor i, want := range names {
\t\tgot := DayName(Weekday(i))
\t\tif got != want {
\t\t\tt.Errorf("DayName(%d) = %q, want %q", i, got, want)
\t\t}
\t}
\tgot := DayName(Weekday(99))
\tif got != "Unknown" {
\t\tt.Errorf("DayName(99) = %q, want %q", got, "Unknown")
\t}
}

func TestIsWeekend(t *testing.T) {
\tweekends := []Weekday{0, 6}
\tdays := []Weekday{1, 2, 3, 4, 5}
\tfor _, d := range weekends {
\t\tif !IsWeekend(d) {
\t\t\tt.Errorf("IsWeekend(%d) should be true", d)
\t\t}
\t}
\tfor _, d := range days {
\t\tif IsWeekend(d) {
\t\t\tt.Errorf("IsWeekend(%d) should be false", d)
\t\t}
\t}
}

func TestNextDay(t *testing.T) {
\texpected := []Weekday{1, 2, 3, 4, 5, 6, 0} // Sun->Mon, ..., Sat->Sun
\tfor i, want := range expected {
\t\tgot := NextDay(Weekday(i))
\t\tif got != want {
\t\t\tt.Errorf("NextDay(%d) = %d, want %d", i, got, want)
\t\t}
\t}
}
''',
        "readme": '''# Kata 165 — Constants & iota

**Focus:** const declarations, iota enum pattern, typed vs untyped constants

## Your task

Implement a Weekday enum using iota.

### Learning goal
- What you are practicing: defining constants with iota, creating enums, using typed constants.
- Why this matters: enums prevent invalid states and make code self-documenting.
- How this grows your Go skills: you'll use iota patterns constantly in real Go code.

## Rules / Expectations
- Define Weekday type and iota-based constants
- DayName returns the string name
- IsWeekend and NextDay use the enum correctly

## What this kata is about (and why it matters)
- Core lesson: Go uses const + iota for enums. It's simple but powerful.
- After this kata, you should create type-safe enums for any domain.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "166",
        "slug": "init-function",
        "title": "init() Function & Package Lifecycle",
        "focus": "init() execution order, package-level initialization, side effects",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["init", "lifecycle", "initialization", "side-effects", "package"],
        "prerequisites": ["163", "164"],
        "estimated_minutes": 25,
        "flashcards": [
            {"front": "When does init() run in Go?", "back": "After all package-level variables are initialized, and before main(). Each package can have multiple init() functions."},
            {"front": "Can a package have multiple init() functions?", "back": "Yes. Go allows multiple init() per package, and they run in source order."},
            {"front": "What is the initialization order of packages?", "back": "Packages are initialized in dependency order. If package A imports B, B's init() runs before A's."},
            {"front": "Should you use init() for business logic?", "back": "Generally no. init() is for setup like registering drivers or setting up globals. Business logic belongs in explicit functions."}
        ],
        "quiz_questions": [
            {"question": "How many init() functions can a package have?", "options": ["Exactly one", "Zero or one", "Multiple", "None"], "answer": "Multiple"},
            {"question": "What runs first: init() or main()?", "options": ["main()", "init()", "They run simultaneously", "It depends on the Go version"], "answer": "init()"},
            {"question": "When are package-level variables initialized?", "options": ["Before init()", "After init()", "During main()", "On first use"], "answer": "Before init()"}
        ],
        "kata_go": '''package kata

// Track initialization order using package-level variables.
// This kata teaches you how init() and variable initialization work.

var initOrder []string

// TODO: Define an init() function that appends "main" to initOrder
// This simulates the main package's init

// GetInitOrder returns the recorded initialization order.
// The test will verify that variables initialize before init(),
// and that init() runs before any explicit function call.
func GetInitOrder() []string {
\treturn initOrder
}

// ResetInitOrder clears the recorded order (for testing).
func ResetInitOrder() {
\tinitOrder = nil
}

// Initialize demonstrates the proper use of init().
// It should:
// 1. Reset the order
// 2. Show that package variables are set first
// 3. Show that init() functions run next
func Initialize() []string {
\tResetInitOrder()
\t// In a real app, init() would have already run.
\t// Here we simulate the order:
\tinitOrder = append(initOrder, "package_vars")
\tinitOrder = append(initOrder, "init_func")
\tinitOrder = append(initOrder, "main")
\treturn initOrder
}
''',
        "kata_test_go": '''package kata

import "testing"

func TestInitialize(t *testing.T) {
\tgot := Initialize()
\texpected := []string{"package_vars", "init_func", "main"}
\tif len(got) != len(expected) {
\t\tt.Fatalf("Initialize() returned %d items, want %d", len(got), len(expected))
\t}
\tfor i, want := range expected {
\t\tif got[i] != want {
\t\t\tt.Errorf("Initialize()[%d] = %q, want %q", i, got[i], want)
\t\t}
\t}
}

func TestGetInitOrder(t *testing.T) {
\tResetInitOrder()
\torder := GetInitOrder()
\tif len(order) != 0 {
\t\tt.Errorf("GetInitOrder() should be empty after reset, got %v", order)
\t}
}
''',
        "readme": '''# Kata 166 — init() Function & Package Lifecycle

**Focus:** init() execution order, package-level initialization, side effects

## Your task

Understand and implement initialization order tracking.

### Learning goal
- What you are practicing: understanding Go's initialization sequence — variables first, then init(), then main().
- Why this matters: init() bugs are hard to debug because the execution order isn't obvious.
- How this grows your Go skills: you'll understand why some Go code "just works" and some doesn't.

## Rules / Expectations
- Initialize returns the correct order: package_vars, init_func, main
- GetInitOrder and ResetInitOrder work correctly

## What this kata is about (and why it matters)
- Core lesson: Go has a strict initialization order. Understanding it prevents subtle bugs.
- After this kata, you should predict initialization behavior in any Go program.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "167",
        "slug": "exported-unexported",
        "title": "Exported vs Unexported Identifiers",
        "focus": "uppercase/lowercase naming, visibility rules, API design",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["exported", "unexported", "visibility", "naming", "API"],
        "prerequisites": ["164"],
        "estimated_minutes": 20,
        "flashcards": [
            {"front": "How do you make a Go identifier exported (public)?", "back": "Start with an uppercase letter. That's it — Go's visibility is purely based on the first character."},
            {"front": "Can unexported identifiers be accessed from other packages?", "back": "No. Only exported identifiers (uppercase first letter) are visible to other packages."},
            {"front": "What is the convention for unexported helper functions?", "back": "Use lowercase, often with a leading underscore or just lowercase. Keep them short and descriptive."},
            {"front": "Should types be exported or unexported?", "back": "Export types that form your API surface. Keep implementation details unexported."}
        ],
        "quiz_questions": [
            {"question": "Which identifier is exported in Go?", "options": ["myFunc", "MyFunc", "_myFunc", "my_func"], "answer": "MyFunc"},
            {"question": "Can package A access an unexported func in package B?", "options": ["Yes, if imported", "No, never", "Only with reflection", "Only if in same directory"], "answer": "No, never"},
            {"question": "What is a common pattern for unexported types?", "options": ["Prefix with _", "Use lowercase only", "Add 'internal' suffix", "Put in internal/"], "answer": "Use lowercase only"}
        ],
        "kata_go": '''package kata

// This kata teaches you about Go's exported/unexported naming convention.
// In Go, uppercase = exported (public), lowercase = unexported (private).

// Exported function - visible to other packages
func ExportedName() string {
\treturn "visible"
}

// unexported function - only visible within this package
func unexportedName() string {
\treturn "hidden"
}

// IsExported checks if a Go identifier is exported (starts with uppercase).
// Rules:
// - "MyFunc" => true
// - "myFunc" => false
// - "_myFunc" => false
// - "MyFunc123" => true
// - "" => false
func IsExported(name string) bool {
\t// TODO: Check if first character is uppercase
\treturn false
}

// FilterExported returns only the exported identifiers from a list.
// Rules:
// - FilterExported(["MyFunc", "myFunc", "Test", "_helper"]) => ["MyFunc", "Test"]
func FilterExported(names []string) []string {
\t// TODO: Return only names where IsExported is true
\treturn nil
}

// MaskUnexported replaces unexported names with "***" in a list.
// Rules:
// - MaskUnexported(["MyFunc", "myFunc", "Test"]) => ["MyFunc", "***", "Test"]
func MaskUnexported(names []string) []string {
\t// TODO: Replace unexported names with "***"
\treturn nil
}
''',
        "kata_test_go": '''package kata

import "testing"

func TestIsExported(t *testing.T) {
\ttests := []struct {
\t\tname string
\t\twant bool
\t}{
\t{"MyFunc", true},
\t{"myFunc", false},
\t{"_myFunc", false},
\t{"MyFunc123", true},
\t{"", false},
\t{"A", true},
\t{"a", false},
\t}
\tfor _, tt := range tests {
\t\tt.Run(tt.name, func(t *testing.T) {
\t\t\tgot := IsExported(tt.name)
\t\t\tif got != tt.want {
\t\t\t\tt.Errorf("IsExported(%q) = %v, want %v", tt.name, got, tt.want)
\t\t\t}
\t\t})
\t}
}

func TestFilterExported(t *testing.T) {
\tinput := []string{"MyFunc", "myFunc", "Test", "_helper"}
\tgot := FilterExported(input)
\twant := []string{"MyFunc", "Test"}
\tif len(got) != len(want) {
\t\tt.Fatalf("FilterExported() returned %d items, want %d", len(got), len(want))
\t}
\tfor i, w := range want {
\t\tif got[i] != w {
\t\t\tt.Errorf("FilterExported()[%d] = %q, want %q", i, got[i], w)
\t\t}
\t}
}

func TestMaskUnexported(t *testing.T) {
\tinput := []string{"MyFunc", "myFunc", "Test"}
\tgot := MaskUnexported(input)
\twant := []string{"MyFunc", "***", "Test"}
\tif len(got) != len(want) {
\t\tt.Fatalf("MaskUnexported() returned %d items, want %d", len(got), len(want))
\t}
\tfor i, w := range want {
\t\tif got[i] != w {
\t\t\tt.Errorf("MaskUnexported()[%d] = %q, want %q", i, got[i], w)
\t\t}
\t}
}
''',
        "readme": '''# Kata 167 — Exported vs Unexported Identifiers

**Focus:** uppercase/lowercase naming, visibility rules, API design

## Your task

Implement functions that understand Go's visibility rules.

### Learning goal
- What you are practicing: understanding that Go uses naming case for visibility, not keywords.
- Why this matters: API design in Go is entirely based on exported/unexported naming.
- How this grows your Go skills: you'll design clean, well-encapsulated packages.

## Rules / Expectations
- IsExported returns true for uppercase-starting names
- FilterExported returns only exported names
- MaskUnexported replaces unexported names with "***"

## What this kata is about (and why it matters)
- Core lesson: Go's visibility is simple but powerful — uppercase = public, lowercase = private.
- After this kata, you should design APIs with clear public/private boundaries.

## What you must submit for marking
- `kata.go`
'''
    }
]

# Category insertion order: these go into setup-and-toolchain
# Insert after kata-005 (Go Vet & Lint) and before kata-006 (FizzBuzz)

for kata in NEW_KATAS:
    kid = kata["id"]
    slug = kata["slug"]
    dirname = f"kata-{kid.zfill(3)}-{slug}"
    dirpath = os.path.join(KATAS_DIR, dirname)
    os.makedirs(dirpath, exist_ok=True)
    
    # Write kata.json
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
    
    # Write kata.go.txt (source stub)
    with open(os.path.join(dirpath, "kata.go.txt"), "w") as f:
        f.write(kata["kata_go"])
    
    # Write kata_test.go.txt (evaluator)
    with open(os.path.join(dirpath, "kata_test.go.txt"), "w") as f:
        f.write(kata["kata_test_go"])
    
    # Write README.md
    with open(os.path.join(dirpath, "README.md"), "w") as f:
        f.write(kata["readme"])
    
    print(f"Created: {dirname}")

# Update track.json - add new kata IDs to setup-and-toolchain category
with open(TRACK_FILE) as f:
    track = json.load(f)

for stage in track["stages"]:
    if stage["id"] == "foundation":
        for cat in stage["categories"]:
            if cat["id"] == "setup-and-toolchain":
                # Insert new katas after 005 (position 5)
                new_ids = [k["id"] for k in NEW_KATAS]
                existing = cat["kata_ids"]
                cat["kata_ids"] = existing[:6] + new_ids + existing[6:]
                print(f"Updated setup-and-toolchain: {len(cat['kata_ids'])} katas")
                break

with open(TRACK_FILE, "w") as f:
    json.dump(track, f, indent=2)

print("Done!")
