# Kata 046 — Safe Counter

**Focus:** sync/atomic or mutex

## Your task
Implement:

```go
type Counter struct
```

### Learning goal
- Expected work: Implement `Safe Counter` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Safe Counter` teaches safe coordination patterns that prevent costly production race conditions.
- When correct: `Safe Counter` should satisfy the required behavior, including: `Inc/Add/Value`; and `race-free`.

## Rules / Expectations
- Inc/Add/Value
- race-free

## Prior reading
- [Go sync/atomic package](https://pkg.go.dev/sync/atomic)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains sync/atomic or mutex by implementing `Safe Counter` under explicit constraints.
- It is important because `Safe Counter` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
