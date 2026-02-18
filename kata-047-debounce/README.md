# Kata 047 — Debounce

**Focus:** Time, channels

## Your task
Implement:

```go
func Debounce(d time.Duration, in <-chan struct{}) <-chan struct{}
```

## Rules / Expectations
- emit after quiet period
- close output when input closes

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
