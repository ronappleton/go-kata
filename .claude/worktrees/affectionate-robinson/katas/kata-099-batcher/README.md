# Kata 099 — Batcher

**Focus:** Time, channels

## Your task
Implement:

```go
func Batch[T any](ctx context.Context, in <-chan T, max int, every time.Duration) <-chan []T
```

### Learning goal
- What you are building: Build `func Batch[T any](ctx context.Context, in <-chan T, max int, every time.Duration) <-chan []T` as a reliable contract. Focus: Time, channels.
- Why this matters in real projects: Concurrency bugs are expensive. You are learning to prevent them by design.
- How this grows your Go skills: You practice ownership, cancellation, synchronization, and leak-free shutdown.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: emit when max reached or timer fires; flush on close; and no empty batches.

### Tips
- Decide ownership first: who starts, stops, and closes.
- Test cancellation and shutdown before throughput.
- Run `go test -race ./...` regularly.

## Rules / Expectations
- emit when max reached or timer fires
- flush on close
- no empty batches

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

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
