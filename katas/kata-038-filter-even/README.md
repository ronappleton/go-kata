# Kata 038 — Filter Even

**Focus:** Slices, order, immutability

## Your task
Implement:

```go
func FilterEven(nums []int) []int
```

### Learning goal
- What you are building: Build `func FilterEven(nums []int) []int` as a reliable contract. Focus: Slices, order, immutability.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: preserve order; for `nil`, return `empty slice`; and do not mutate input.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- preserve order
- nil => empty slice
- do not mutate input

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
