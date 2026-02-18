# Kata 047 — Debounce

**Focus:** Time, channels

## Your task
Implement:

```go
func Debounce(d time.Duration, in <-chan struct{}) <-chan struct{}
```

### Learning goal
- Expected work: Implement `Debounce` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Debounce` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- emit after quiet period
- close output when input closes

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata trains time, channels by implementing `Debounce` under explicit constraints.
- It is important because `Debounce` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
