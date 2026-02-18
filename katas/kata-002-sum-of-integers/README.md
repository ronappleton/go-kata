# Kata 002 — Sum of Integers

**Focus:** Basics: loops, function, edge cases

## Your task
Implement:

```go
func Sum(nums []int) int
```

### Learning goal
- What you are building: Build `func Sum(nums []int) int` as a reliable contract. Focus: Basics: loops, function, edge cases.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: for `nil slice`, return `0`; for `empty`, return `0`; and work correctly with negative values.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- nil slice => 0
- empty => 0
- works with negatives

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
