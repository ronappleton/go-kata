# Kata 046 — Safe Counter

**Focus:** sync/atomic or mutex

## Your task
Implement:

```go
type Counter struct
```

## Rules / Expectations
- Inc/Add/Value
- race-free

## Prior reading
- [Go sync/atomic package](https://pkg.go.dev/sync/atomic)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Safe Counter with constraints that make you practice sync/atomic or mutex.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
