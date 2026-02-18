# Kata 069 — Batcher

**Focus:** Time, channels

## Your task
Implement:

```go
func Batch[T any](ctx context.Context, in <-chan T, max int, every time.Duration) <-chan []T
```

### Learning goal
- Expected work: Implement `Batcher` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Batcher` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- emit when max reached or timer fires
- flush on close
- no empty batches

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata trains time, channels by implementing `Batcher` under explicit constraints.
- It is important because `Batcher` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
