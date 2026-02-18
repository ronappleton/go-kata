# Kata 064 — Circuit Breaker

**Focus:** State machines, time

## Your task
Implement:

```go
func NewBreaker(failureThreshold int, resetAfter time.Duration) (*Breaker, error)
```

### Learning goal
- Expected work: Implement `Circuit Breaker` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Circuit Breaker` teaches safe coordination patterns that prevent costly production race conditions.

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
- This kata trains state machines, time by implementing `Circuit Breaker` under explicit constraints.
- It is important because `Circuit Breaker` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
