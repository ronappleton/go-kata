# Kata 007 — Fibonacci

**Focus:** Iteration, slices

## Your task
Implement:

```go
func Fibonacci(n int) ([]int, error)
```

### Learning goal
- Expected work: Implement `Fibonacci` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Fibonacci` teaches core implementation habits that compound across all later katas.
- When correct: `Fibonacci` should satisfy the required behavior, including: `n<0 => error`; `n=0 => empty slice`; `n=1 => [0]`; and `sequence starts 0,1,...`.

## Rules / Expectations
- n<0 => error
- n=0 => empty slice
- n=1 => [0]
- sequence starts 0,1,...

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains iteration, slices by implementing `Fibonacci` under explicit constraints.
- It is important because `Fibonacci` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
