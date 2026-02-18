# Kata 048 — Throttle

**Focus:** Time, channels

## Your task
Implement:

```go
func Throttle(d time.Duration, in <-chan struct{}) <-chan struct{}
```

## Rules / Expectations
- max 1 event per period
- drop extras

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
