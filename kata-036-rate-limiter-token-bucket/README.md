# Kata 036 — Rate Limiter (token bucket)

**Focus:** Time, structs

## Your task
Implement:

```go
func NewTokenBucket(ratePerSec int, burst int) (*TokenBucket, error)
```

### Learning goal
- Expected work: Implement `Rate Limiter (token bucket)` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Rate Limiter (token bucket)` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- Allow() bool
- single-thread only here
- ratePerSec>=1, burst>=1

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains time, structs by implementing `Rate Limiter (token bucket)` under explicit constraints.
- It is important because `Rate Limiter (token bucket)` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
