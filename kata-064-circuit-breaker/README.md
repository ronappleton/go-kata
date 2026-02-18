# Kata 064 — Circuit Breaker

**Focus:** State machines, time

## Your task
Implement:

```go
func NewBreaker(failureThreshold int, resetAfter time.Duration) (*Breaker, error)
```

## Rules / Expectations
- Closed/Open/HalfOpen states
- trip after threshold
- reset after duration

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
