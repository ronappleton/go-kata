# Kata 069 — Batcher

**Focus:** Time, channels

## Your task
Implement:

```go
func Batch[T any](ctx context.Context, in <-chan T, max int, every time.Duration) <-chan []T
```

## Rules / Expectations
- emit when max reached or timer fires
- flush on close
- no empty batches

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata is about implementing Batcher with constraints that make you practice time, channels.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
