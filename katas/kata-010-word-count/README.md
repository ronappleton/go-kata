# Kata 010 — Word Count

**Focus:** Strings, fields, maps

## Your task
Implement:

```go
func WordCount(s string) map[string]int
```

### Learning goal
- What you are building: Build `func WordCount(s string) map[string]int` as a reliable contract. Focus: Strings, fields, maps.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: split on whitespace; lowercase words; and for `empty`, return `empty map (not nil)`.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- split on whitespace
- lowercase words
- empty => empty map (not nil)

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- Core lesson: turn plain rules into deterministic Go behavior.
- After this kata, you should be able to write rule-first tests and explain each edge case clearly.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
