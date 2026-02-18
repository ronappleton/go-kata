# Kata 044 — Concurrent MapReduce

**Focus:** Goroutines, channels

## Your task
Implement:

```go
func MapReduce[T any, R any](items []T, workers int, mapFn func(T) R, reduceFn func([]R) R) (R, error)
```

### Learning goal
- What you are practicing: Build `Concurrent MapReduce` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `workers>=1`; `order not required`; and `handle empty items`.

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
- This kata is focused practice in Goroutines, channels through `Concurrent MapReduce`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
