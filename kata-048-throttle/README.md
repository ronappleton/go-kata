# Kata 048 — Throttle

**Focus:** Time, channels

## Your task
Implement:

```go
func Throttle(d time.Duration, in <-chan struct{}) <-chan struct{}
```

### Learning goal
- Expected work: Implement `Throttle` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Throttle` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- max 1 event per period
- drop extras

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata trains time, channels by implementing `Throttle` under explicit constraints.
- It is important because `Throttle` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
