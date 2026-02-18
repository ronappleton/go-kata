# Kata 044 — Concurrent MapReduce

**Focus:** Goroutines, channels

## Your task
Implement:

```go
func MapReduce[T any, R any](items []T, workers int, mapFn func(T) R, reduceFn func([]R) R) (R, error)
```

### Learning goal
- Expected work: Implement `Concurrent MapReduce` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Concurrent MapReduce` teaches safe coordination patterns that prevent costly production race conditions.
- When correct: `Concurrent MapReduce` should satisfy the required behavior, including: `workers>=1`; `order not required`; and `handle empty items`.

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
- This kata trains goroutines, channels by implementing `Concurrent MapReduce` under explicit constraints.
- It is important because `Concurrent MapReduce` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
