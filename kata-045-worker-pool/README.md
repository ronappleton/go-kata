# Kata 045 — Worker Pool

**Focus:** Concurrency patterns

## Your task
Implement:

```go
func RunPool(ctx context.Context, workers int, jobs <-chan Job) <-chan Result
```

## Rules / Expectations
- cancels on ctx
- close output when done
- no leaks

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
