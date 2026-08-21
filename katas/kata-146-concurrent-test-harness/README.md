# Kata 146 — Concurrent Test Harness

**Focus:** concurrency, testing

## Your task
Implement:

```go
type TestCase struct { Name string; Run func() error }; func RunTests(tests []TestCase, workers int) (passed int, failed int)
```

### Learning goal
- What you are building: type TestCase struct { Name string; Run func() error }; func RunTests(tests []TestCase, workers int) (passed int, failed int) as a reliable contract. Focus: concurrency, testing.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): runs every test, at most workers concurrently, and returns how many passed and failed; workers < 1 means unlimited (each test its own goroutine).

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- concurrency limited by workers
- passed/failed counts
- nil error => passed
- workers < 1 => unlimited

## Prior reading
- [Go sync package](https://pkg.go.dev/sync)
- [Go testing package](https://pkg.go.dev/testing)

## What this kata is about (and why it matters)
- Core lesson: a harness is a concurrency primitive: bound the workers, count outcomes, and never leak goroutines.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
