# Kata 064 — Circuit Breaker

**Focus:** State machines, time

## Your task
Implement:

```go
func NewBreaker(failureThreshold int, resetAfter time.Duration) (*Breaker, error)
```

### Learning goal
- What you are practicing: Build `Circuit Breaker` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `Closed/Open/HalfOpen states`; `trip after threshold`; and `reset after duration`.

## Rules / Expectations
- Closed/Open/HalfOpen states
- trip after threshold
- reset after duration

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in State machines, time through `Circuit Breaker`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
