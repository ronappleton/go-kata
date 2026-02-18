# Kata 047 — Debounce

**Focus:** Time, channels

## Your task
Implement:

```go
func Debounce(d time.Duration, in <-chan struct{}) <-chan struct{}
```

### Learning goal
- What you are practicing: Build `Debounce` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `emit after quiet period`; and `close output when input closes`.

## Rules / Expectations
- emit after quiet period
- close output when input closes

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata is focused practice in Time, channels through `Debounce`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
