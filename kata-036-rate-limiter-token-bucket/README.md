# Kata 036 — Rate Limiter (token bucket)

**Focus:** Time, structs

## Your task
Implement:

```go
func NewTokenBucket(ratePerSec int, burst int) (*TokenBucket, error)
```

### Learning goal
- What you are practicing: Build `Rate Limiter (token bucket)` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `Allow() bool`; `single-thread only here`; and `ratePerSec>=1, burst>=1`.

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
- This kata is focused practice in Time, structs through `Rate Limiter (token bucket)`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
