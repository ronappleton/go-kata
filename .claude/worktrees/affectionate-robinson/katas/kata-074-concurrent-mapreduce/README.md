# Kata 074 — Concurrent MapReduce

**Focus:** Goroutines, channels

## Your task
Implement:

```go
func MapReduce[T any, R any](items []T, workers int, mapFn func(T) R, reduceFn func([]R) R) (R, error)
```

### Learning goal
- What you are building: Build `func MapReduce[T any, R any](items []T, workers int, mapFn func(T) R, reduceFn func([]R) R) (R, error)` as a reliable contract. Focus: Goroutines, channels.
- Why this matters in real projects: Concurrency bugs are expensive. You are learning to prevent them by design.
- How this grows your Go skills: You practice ownership, cancellation, synchronization, and leak-free shutdown.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: workers>=1; order not required; and handle empty items.

### Tips
- Decide ownership first: who starts, stops, and closes.
- Test cancellation and shutdown before throughput.
- Run `go test -race ./...` regularly.

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
- Core lesson: own lifecycle and shutdown before chasing throughput.
- After this kata, you should be able to explain who starts, who stops, and who closes every path.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
