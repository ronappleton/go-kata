# Kata 043 — Retry with Backoff

**Focus:** time

## Your task
Implement:

```go
func Retry(attempts int, baseDelay time.Duration, fn func() error) error
```

## Rules / Expectations
- attempts>=1
- exponential backoff
- stop on success

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
