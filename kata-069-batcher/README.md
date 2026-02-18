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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
