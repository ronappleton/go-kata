# Kata 071 — Token Bucket (thread-safe)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata71() error
```

### Learning goal
- Expected work: Implement `Token Bucket (thread-safe)` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Token Bucket (thread-safe)` teaches safe coordination patterns that prevent costly production race conditions.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `Token Bucket (thread-safe)` under explicit constraints.
- It is important because `Token Bucket (thread-safe)` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
