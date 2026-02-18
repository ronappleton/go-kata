# Kata 066 — Semaphore

**Focus:** Concurrency primitives

## Your task
Implement:

```go
func NewSemaphore(n int) (*Semaphore, error)
```

## Rules / Expectations
- Acquire(ctx) error
- Release()
- no deadlocks

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
