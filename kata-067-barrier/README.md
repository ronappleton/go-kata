# Kata 067 — Barrier

**Focus:** Concurrency coordination

## Your task
Implement:

```go
func NewBarrier(parties int) (*Barrier, error)
```

### Learning goal
- Expected work: Implement `Barrier` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Barrier` teaches safe coordination patterns that prevent costly production race conditions.
- When correct: `Barrier` should satisfy the required behavior, including: `Wait(ctx) error`; `reusable`; and `no leaks`.

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
- This kata trains concurrency coordination by implementing `Barrier` under explicit constraints.
- It is important because `Barrier` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
