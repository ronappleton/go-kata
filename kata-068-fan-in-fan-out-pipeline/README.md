# Kata 068 — Fan-in/Fan-out Pipeline

**Focus:** Channels, composition

## Your task
Implement:

```go
func Pipeline[T any](ctx context.Context, in <-chan T, stages ...func(context.Context, <-chan T) <-chan T) <-chan T
```

## Rules / Expectations
- compose stages in order
- cancel cleanly
- close outputs

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
