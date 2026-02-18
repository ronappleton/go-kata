# Kata 006 — Factorial

**Focus:** Recursion vs loop, overflow checks

## Your task
Implement:

```go
func Factorial(n int) (int, error)
```

### Learning goal
- Expected work: Implement `Factorial` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Factorial` teaches core implementation habits that compound across all later katas.
- When correct: `Factorial` should satisfy the required behavior, including: `n<0 => error`; `0 => 1`; and `overflow => error`.

## Rules / Expectations
- n<0 => error
- 0 => 1
- overflow => error

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains recursion vs loop, overflow checks by implementing `Factorial` under explicit constraints.
- It is important because `Factorial` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
