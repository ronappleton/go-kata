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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
