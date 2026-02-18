# Kata 007 — Fibonacci

**Focus:** Iteration, slices

## Your task
Implement:

```go
func Fibonacci(n int) ([]int, error)
```

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
- This kata is about implementing Fibonacci with constraints that make you practice iteration, slices.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
