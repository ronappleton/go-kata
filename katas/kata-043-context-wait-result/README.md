# Kata 028 — Wait For Context Or Duration

**Focus:** context cancellation, time.Timer selection

## Your task
Implement:

```go
func WaitForContextOrDuration(ctx context.Context, d time.Duration) string
```

### Learning goal
- What you are building: a cancellation-aware wait function.
- Why this matters in real projects: cancellation and timeout handling is central to reliable concurrent code.
- How this grows your Go skills: you practice `select` logic with context and timers.
- Definition of done (plain English): function reports whether cancellation or duration completion happened first.

### Tips
- Handle `d <= 0` before starting timers.
- Use timer cleanup patterns to avoid resource leaks.
- Test both cancellation-first and duration-first paths.

## Rules / Expectations
- d <= 0 => return `completed` immediately
- if context is done before duration => return `cancelled`
- if duration elapses first => return `completed`
- avoid leaking timer resources

## Prior reading
- [context package](https://pkg.go.dev/context)
- [time package](https://pkg.go.dev/time)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: cancellation-aware control flow is mandatory for safe concurrency.
- After this kata, you should be able to encode race-free timeout/cancel decisions.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
