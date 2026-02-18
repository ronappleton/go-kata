# Kata 043 — Retry with Backoff

**Focus:** time

## Your task
Implement:

```go
func Retry(attempts int, baseDelay time.Duration, fn func() error) error
```

### Learning goal
- What you are practicing: Build `Retry with Backoff` with precise time/window logic so behavior is stable and testable.
- Why it matters: You will use this in rate limits, retries, and scheduling logic that often causes subtle bugs.
- How this grows your Go skills: This builds precision around clocks, windows, and repeatable time-based tests.
- When correct: When your solution is correct, it should satisfy: `attempts>=1`; `exponential backoff`; and `stop on success`.

## Rules / Expectations
- attempts>=1
- exponential backoff
- stop on success

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in time through `Retry with Backoff`.
- You will use this in rate limits, retries, and scheduling logic that often causes subtle bugs.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
