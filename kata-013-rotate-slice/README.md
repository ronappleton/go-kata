# Kata 013 — Rotate Slice

**Focus:** Generics, indexing

## Your task
Implement:

```go
func RotateLeft[T any](items []T, k int) []T
```

## Rules / Expectations
- k can be > len
- k can be negative (rotate right)
- nil => empty slice

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
