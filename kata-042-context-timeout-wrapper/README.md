# Kata 042 — Context Timeout Wrapper

**Focus:** context

## Your task
Implement:

```go
func WithTimeoutDo(ctx context.Context, d time.Duration, fn func(context.Context) error) error
```

## Rules / Expectations
- fn called with derived ctx
- returns ctx error on timeout
- propagates fn error

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
