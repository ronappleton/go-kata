# Kata 070 — LRU Cache (thread-safe)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata70() error
```

### Learning goal
- Expected work: Implement `LRU Cache (thread-safe)` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `LRU Cache (thread-safe)` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `LRU Cache (thread-safe)` under explicit constraints.
- It is important because `LRU Cache (thread-safe)` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
