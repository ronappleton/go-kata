# Kata 059 — Fixed-Window Metrics

**Focus:** Time buckets

## Your task
Implement:

```go
type CounterWindow struct
```

### Learning goal
- What you are practicing: Build `Fixed-Window Metrics` with precise time/window logic so behavior is stable and testable.
- Why it matters: You will use this in rate limits, retries, and scheduling logic that often causes subtle bugs.
- How this grows your Go skills: This builds precision around clocks, windows, and repeatable time-based tests.
- When correct: When your solution is correct, it should satisfy: `Add(value)`; `Sum(last duration)`; and `thread-safe`.

## Rules / Expectations
- Add(value)
- Sum(last duration)
- thread-safe

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Time buckets through `Fixed-Window Metrics`.
- You will use this in rate limits, retries, and scheduling logic that often causes subtle bugs.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
