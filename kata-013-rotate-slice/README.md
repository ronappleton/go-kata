# Kata 013 — Rotate Slice

**Focus:** Generics, indexing

## Your task
Implement:

```go
func RotateLeft[T any](items []T, k int) []T
```

## Rules / Expectations
- k can be > len
- k can be negative (rotate right)
- nil => empty slice

## Prior reading
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Rotate Slice with constraints that make you practice generics, indexing.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
