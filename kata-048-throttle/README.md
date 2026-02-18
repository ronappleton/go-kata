# Kata 048 — Throttle

**Focus:** Time, channels

## Your task
Implement:

```go
func Throttle(d time.Duration, in <-chan struct{}) <-chan struct{}
```

## Rules / Expectations
- max 1 event per period
- drop extras

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata is about implementing Throttle with constraints that make you practice time, channels.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
