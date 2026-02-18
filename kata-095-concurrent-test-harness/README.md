# Kata 095 — Concurrent Test Harness

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata95() error
```

### Learning goal
- Expected work: Implement `Concurrent Test Harness` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Concurrent Test Harness` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `Concurrent Test Harness` under explicit constraints.
- It is important because `Concurrent Test Harness` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
