# Kata 071 — Token Bucket (thread-safe)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata71() error
```

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata is about implementing Token Bucket (thread-safe) with constraints that make you practice system design, integration boundaries, and robust testing.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
