# Kata 047 — Debounce

**Focus:** Time, channels

## Your task
Implement:

```go
func Debounce(d time.Duration, in <-chan struct{}) <-chan struct{}
```

## Rules / Expectations
- emit after quiet period
- close output when input closes

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata is about implementing Debounce with constraints that make you practice time, channels.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
