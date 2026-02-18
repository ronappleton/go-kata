# Kata 066 — Semaphore

**Focus:** Concurrency primitives

## Your task
Implement:

```go
func NewSemaphore(n int) (*Semaphore, error)
```

### Learning goal
- What you are practicing: Build `Semaphore` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `Acquire(ctx) error`; `Release()`; and `no deadlocks`.

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
- This kata is focused practice in Concurrency primitives through `Semaphore`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
