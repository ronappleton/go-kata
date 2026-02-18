# Kata 066 — Semaphore

**Focus:** Concurrency primitives

## Your task
Implement:

```go
func NewSemaphore(n int) (*Semaphore, error)
```

## Rules / Expectations
- Acquire(ctx) error
- Release()
- no deadlocks

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Semaphore with constraints that make you practice concurrency primitives.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
