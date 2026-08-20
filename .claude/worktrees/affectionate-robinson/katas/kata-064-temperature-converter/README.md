# Kata 064 — Temperature Converter

**Focus:** Parsing floats

## Your task
Implement:

```go
func ConvertTemp(input string) (string, error)
```

### Learning goal
- What you are building: Build `func ConvertTemp(input string) (string, error)` as a reliable contract. Focus: Parsing floats.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: supports c/f/k suffix; output all 3 units with 2 decimals; and return `error` for invalid input.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- supports C/F/K suffix
- output all 3 units with 2 decimals
- invalid => error

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

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
