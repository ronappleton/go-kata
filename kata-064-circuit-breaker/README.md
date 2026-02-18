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

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Circuit Breaker with constraints that make you practice state machines, time.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
