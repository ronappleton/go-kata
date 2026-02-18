# Kata 067 — Barrier

**Focus:** Concurrency coordination

## Your task
Implement:

```go
func NewBarrier(parties int) (*Barrier, error)
```

## Rules / Expectations
- Wait(ctx) error
- reusable
- no leaks

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Barrier with constraints that make you practice concurrency coordination.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
