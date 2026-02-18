# Kata 059 — Fixed-Window Metrics

**Focus:** Time buckets

## Your task
Implement:

```go
type CounterWindow struct
```

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
- This kata is about implementing Fixed-Window Metrics with constraints that make you practice time buckets.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
