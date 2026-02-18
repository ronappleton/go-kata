# Kata 006 — Factorial

**Focus:** Recursion vs loop, overflow checks

## Your task
Implement:

```go
func Factorial(n int) (int, error)
```

## Rules / Expectations
- n<0 => error
- 0 => 1
- overflow => error

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Factorial with constraints that make you practice recursion vs loop, overflow checks.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
