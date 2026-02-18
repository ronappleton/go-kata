# Kata 001 — FizzBuzz

**Focus:** Basics: loops, conditionals, slices, strconv

## Your task
Implement:

```go
func FizzBuzz(n int) []string
```

### Learning goal
- What you are building: Build `func FizzBuzz(n int) []string` as a reliable contract. Focus: Basics: loops, conditionals, slices, strconv.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: if `n<=0`, return `empty slice (not nil)`; for multiples of 3, return `Fizz`; for multiples of 5, return `Buzz`; when both rules apply, return `FizzBuzz`; and otherwise return `number string`.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- n<=0 => empty slice (not nil)
- multiples of 3 => Fizz
- multiples of 5 => Buzz
- both => FizzBuzz
- otherwise number string

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
