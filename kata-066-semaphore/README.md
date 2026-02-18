# Kata 066 — Semaphore

**Focus:** Concurrency primitives

## Your task
Implement:

```go
func NewSemaphore(n int) (*Semaphore, error)
```

### Learning goal
- Expected work: Implement `Semaphore` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Semaphore` teaches safe coordination patterns that prevent costly production race conditions.

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
- This kata trains concurrency primitives by implementing `Semaphore` under explicit constraints.
- It is important because `Semaphore` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
