# Kata 044 — Concurrent MapReduce

**Focus:** Goroutines, channels

## Your task
Implement:

```go
func MapReduce[T any, R any](items []T, workers int, mapFn func(T) R, reduceFn func([]R) R) (R, error)
```

## Rules / Expectations
- workers>=1
- order not required
- handle empty items

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Concurrent MapReduce with constraints that make you practice goroutines, channels.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
