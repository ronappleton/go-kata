# Kata 075 — Worker Pool

**Focus:** Concurrency patterns

## Your task
Implement:

```go
func RunPool(ctx context.Context, workers int, jobs <-chan Job) <-chan Result
```

### Learning goal
- What you are building: Build `func RunPool(ctx context.Context, workers int, jobs <-chan Job) <-chan Result` as a reliable contract. Focus: Concurrency patterns.
- Why this matters in real projects: Concurrency bugs are expensive. You are learning to prevent them by design.
- How this grows your Go skills: You practice ownership, cancellation, synchronization, and leak-free shutdown.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: stop promptly when context is canceled; close output after all work is complete; and leave no goroutines, timers, or resources behind.

### Tips
- Decide ownership first: who starts, stops, and closes.
- Test cancellation and shutdown before throughput.
- Run `go test -race ./...` regularly.

## Rules / Expectations
- cancels on ctx
- close output when done
- no leaks

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
