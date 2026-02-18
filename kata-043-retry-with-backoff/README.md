# Kata 043 — Retry with Backoff

**Focus:** time

## Your task
Implement:

```go
func Retry(attempts int, baseDelay time.Duration, fn func() error) error
```

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
- This kata is about implementing Retry with Backoff with constraints that make you practice time.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
