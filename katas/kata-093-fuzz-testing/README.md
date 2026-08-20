# Kata 093 — Fuzz Testing

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
