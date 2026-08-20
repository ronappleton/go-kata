#!/usr/bin/env python3
"""Inject flashcard and quiz data into kata.json files.
Run this AFTER gen_kata_data.go to restore hand-written flashcards/quiz data.
"""
import json, os

FLASHCARD_DATA = {
    "000": [
        {"front": "What command initializes a new Go module?", "back": "go mod init <module-path>"},
        {"front": "What is the main package?", "back": "The entry point for a Go executable. It must have a main() function."},
        {"front": "What is the difference between go run and go build?", "back": "go run compiles and runs. go build compiles to a binary."},
    ],
    "005": [
        {"front": "What does go vet check for?", "back": "Suspicious constructs that the compiler doesn't catch."},
        {"front": "What is golangci-lint?", "back": "A meta-linter that runs many Go linters at once."},
    ],
    "039": [
        {"front": "How do you wrap an error with context in Go?", "back": 'fmt.Errorf("context: %w", err)'},
        {"front": "How do you check if an error is a specific sentinel?", "back": "errors.Is(err, ErrNotFound)"},
        {"front": "When would you use errors.As instead of errors.Is?", "back": "When you need to extract the concrete error type."},
    ],
    "040": [
        {"front": "How does a type satisfy an interface in Go?", "back": "By implementing all the methods. No explicit declaration needed."},
        {"front": "What does 'accept interfaces, return structs' mean?", "back": "Function parameters should be interfaces, return values should be concrete types."},
    ],
    "041": [
        {"front": "What is a table-driven test?", "back": "A test that defines a slice of test cases and loops through them."},
        {"front": "How do you create a subtest in Go?", "back": 't.Run("name", func(t *testing.T) { ... })'},
        {"front": "Why are table-driven tests preferred in Go?", "back": "Concise, easy to add cases, clear failure messages."},
    ],
    "078": [
        {"front": "What interface must an HTTP handler implement?", "back": "http.Handler: ServeHTTP(ResponseWriter, *Request)"},
        {"front": "How do you test HTTP handlers without starting a server?", "back": "httptest.NewRecorder() + httptest.NewRequest()"},
    ],
    "088": [
        {"front": "What are struct tags in Go?", "back": "Metadata strings after struct fields that control serialization."},
        {"front": "How do you parse JSON into a struct?", "back": "json.Unmarshal([]byte(data), &target)"},
    ],
    "089": [
        {"front": "What is structured logging?", "back": "Logging with key-value pairs instead of free-form text."},
        {"front": "How do you create a JSON logger in slog?", "back": "slog.New(slog.NewJSONHandler(os.Stdout, nil))"},
    ],
    "090": [
        {"front": "What is Cobra in Go?", "back": "A popular library for building CLI applications."},
        {"front": "What are the two common flag formats?", "back": "--flag=value and --flag value"},
    ],
    "091": [
        {"front": "What does assert.Equal do?", "back": "Compares two values and fails the test if they differ."},
        {"front": "What is the difference between assert and require?", "back": "assert continues on failure. require stops immediately."},
    ],
    "092": [
        {"front": "What is the repository pattern?", "back": "An interface that abstracts data access for mock implementations."},
        {"front": "Why use interfaces for database access?", "back": "So you can test with mocks instead of a real database."},
    ],
    "093": [
        {"front": "How do you define a fuzz test in Go?", "back": "func FuzzXxx(f *testing.F) { f.Add(seed); f.Fuzz(...) }"},
        {"front": "What is property-based testing?", "back": "Testing that properties hold for ALL inputs, not just examples."},
    ],
    "113": [
        {"front": "How do you detect data races in Go?", "back": "go test -race ./..."},
        {"front": "What does the race detector do?", "back": "Finds concurrent memory accesses with no synchronization."},
        {"front": "What is the overhead of the race detector?", "back": "About 2-10x slower, 5-10x more memory."},
    ],
    "123": [
        {"front": "What is a span in distributed tracing?", "back": "A named, timed operation representing a unit of work."},
        {"front": "How does context propagation work?", "back": "Trace context is stored in context.Context and passed through calls."},
    ],
    "125": [
        {"front": "How do you write a benchmark in Go?", "back": "func BenchmarkXxx(b *testing.B) { for i := 0; i < b.N; i++ {} }"},
        {"front": "What does -benchmem show?", "back": "Bytes allocated per operation and number of allocations."},
        {"front": "How do you profile a running Go app?", "back": 'Import _ "net/http/pprof" and use go tool pprof.'},
    ],
    "126": [
        {"front": "How do you compare implementations in benchmarks?", "back": "Use sub-benchmarks: b.Run(\"name\", func(b *testing.B) {})"},
        {"front": "What tool statistically compares benchmark results?", "back": "benchstat"},
    ],
    "131": [
        {"front": "What is XSS?", "back": "Cross-Site Scripting: injecting malicious scripts via user input."},
        {"front": "What is SQL injection?", "back": "Injecting SQL commands through user input."},
        {"front": "What is the golden rule of input handling?", "back": "Never trust user input. Validate and sanitize everything."},
    ],
    "147": [
        {"front": "What makes a good code review?", "back": "Specific, actionable feedback focused on correctness."},
        {"front": "What are common code review anti-patterns?", "back": "Nitpicking style, suggesting rewrites, being vague."},
    ],
    "148": [
        {"front": "What is a multi-stage Docker build?", "back": "Multiple FROM statements. Build in one stage, copy binary to minimal runtime."},
        {"front": "Why use CGO_ENABLED=0 for Go Docker images?", "back": "Creates a statically linked binary for scratch/alpine."},
    ],
    "149": [
        {"front": "What are the stages of a CI/CD pipeline?", "back": "Build, Test, Lint, Deploy."},
        {"front": "What is GitHub Actions?", "back": "GitHub's CI/CD platform. Define workflows in .github/workflows/."},
    ],
    "153": [
        {"front": "What is an Architecture Decision Record?", "back": "A document capturing context, decision, and consequences."},
        {"front": "Why are ADRs important?", "back": "They preserve why decisions were made."},
    ],
    "154": [
        {"front": "What makes mentoring feedback effective?", "back": "Specific, actionable, balanced, and prioritized."},
        {"front": "How do you balance positive vs improvement feedback?", "back": "Acknowledge what's good, then focus on highest-impact improvements."},
    ],
}

QUIZ_DATA = {
    "000": [{"type": "multiple_choice", "question": "Which command initializes a Go module?", "options": ["go init", "go mod init", "go new", "go create"], "answer": "go mod init"}],
    "005": [{"type": "multiple_choice", "question": "Which command runs Go's built-in static analysis?", "options": ["go check", "go vet", "go lint", "go analyze"], "answer": "go vet"}],
    "039": [{"type": "multiple_choice", "question": "Which verb wraps an error in fmt.Errorf?", "options": ["%s", "%v", "%w", "%+v"], "answer": "%w"}],
    "040": [{"type": "multiple_choice", "question": "How does a type implement an interface in Go?", "options": ["type T implements I", "By implementing all methods", "type T : I", "type T as I"], "answer": "By implementing all methods"}],
    "041": [{"type": "fill_blank", "question": "In Go, you create a subtest with t.Run(...)", "answer": "Run"}],
    "078": [{"type": "multiple_choice", "question": "Which function registers a handler on the mux?", "options": ["http.Handle", "http.HandleFunc", "http.Register", "http.Route"], "answer": "http.HandleFunc"}],
    "089": [{"type": "multiple_choice", "question": "Which package provides structured logging in Go 1.21+?", "options": ["log", "log/slog", "logrus", "zap"], "answer": "log/slog"}],
    "090": [{"type": "multiple_choice", "question": "Which library is most used for Go CLIs?", "options": ["flag", "cobra", "cli", "argparse"], "answer": "cobra"}],
    "091": [{"type": "multiple_choice", "question": "Which testify function stops immediately on failure?", "options": ["assert.Equal", "require.NoError", "assert.True", "require.Contains"], "answer": "require.NoError"}],
    "092": [{"type": "multiple_choice", "question": "Why mock the database in tests?", "options": ["Faster", "No real DB needed", "Less memory", "All of the above"], "answer": "All of the above"}],
    "093": [{"type": "multiple_choice", "question": "How do you run a fuzz test?", "options": ["go test -fuzz", "go fuzz", "go test -fuzz=Xxx", "go test -fuzztest"], "answer": "go test -fuzz=Xxx"}],
    "113": [{"type": "multiple_choice", "question": "Which flag enables the race detector?", "options": ["-race", "-racy", "-detect", "-safe"], "answer": "-race"}],
    "123": [{"type": "multiple_choice", "question": "What is the standard for distributed tracing in Go?", "options": ["OpenTracing", "OpenCensus", "OpenTelemetry", "Zipkin"], "answer": "OpenTelemetry"}],
    "125": [{"type": "multiple_choice", "question": "What flag shows memory allocations in benchmarks?", "options": ["-bench", "-benchmem", "-mem", "-alloc"], "answer": "-benchmem"}],
    "126": [{"type": "multiple_choice", "question": "How do you create sub-benchmarks?", "options": ["b.Sub()", "b.Run()", "b.Compare()", "b.Bench()"], "answer": "b.Run()"}],
    "131": [{"type": "multiple_choice", "question": "What is the first defense against security vulnerabilities?", "options": ["Firewall", "Input validation", "Encryption", "Authentication"], "answer": "Input validation"}],
    "147": [{"type": "multiple_choice", "question": "What should code review focus on first?", "options": ["Style", "Correctness", "Performance", "Documentation"], "answer": "Correctness"}],
    "148": [{"type": "multiple_choice", "question": "Benefit of multi-stage Docker builds?", "options": ["Faster builds", "Smaller images", "Better security", "All of the above"], "answer": "All of the above"}],
    "149": [{"type": "multiple_choice", "question": "Which CI stage runs first?", "options": ["Deploy", "Test", "Build", "Lint"], "answer": "Build"}],
    "153": [{"type": "multiple_choice", "question": "Three key sections of an ADR?", "options": ["Code, Tests, Deploy", "Context, Decision, Consequences", "Problem, Solution, Result", "Plan, Execute, Review"], "answer": "Context, Decision, Consequences"}],
    "154": [{"type": "multiple_choice", "question": "Most important quality of mentoring feedback?", "options": ["Positive", "Specific and actionable", "Long", "Technical"], "answer": "Specific and actionable"}],
}

updated = 0
for kid, flashcards in FLASHCARD_DATA.items():
    for d in os.listdir('katas'):
        if d.startswith(f'kata-{kid}-'):
            path = os.path.join('katas', d, 'kata.json')
            if os.path.exists(path):
                with open(path) as f:
                    meta = json.load(f)
                meta['flashcards'] = flashcards
                meta['quiz_questions'] = QUIZ_DATA.get(kid, [])
                with open(path, 'w') as f:
                    json.dump(meta, f, indent=2)
                    f.write('\n')
                updated += 1
            break

print(f'Injected flashcards/quiz into {updated} kata.json files')
