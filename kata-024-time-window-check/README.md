# Kata 024 — Time Window Check

**Focus:** time package

## Your task
Implement:

```go
func WithinWindow(t, start, end time.Time) bool
```

### Learning goal
- What you are practicing: Build `Time Window Check` with precise time/window logic so behavior is stable and testable.
- Why it matters: You will use this in rate limits, retries, and scheduling logic that often causes subtle bugs.
- How this grows your Go skills: This builds precision around clocks, windows, and repeatable time-based tests.
- When correct: When your solution is correct, it should satisfy: `inclusive start, exclusive end`; and `if start after end => false`.

## Rules / Expectations
- inclusive start, exclusive end
- if start after end => false

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in time package through `Time Window Check`.
- You will use this in rate limits, retries, and scheduling logic that often causes subtle bugs.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
