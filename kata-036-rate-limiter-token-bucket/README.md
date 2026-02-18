# Kata 036 — Rate Limiter (token bucket)

**Focus:** Time, structs

## Your task
Implement:

```go
func NewTokenBucket(ratePerSec int, burst int) (*TokenBucket, error)
```

## Rules / Expectations
- Allow() bool
- single-thread only here
- ratePerSec>=1, burst>=1

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Rate Limiter (token bucket) with constraints that make you practice time, structs.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
