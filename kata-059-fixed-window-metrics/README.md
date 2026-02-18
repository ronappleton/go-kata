# Kata 059 — Fixed-Window Metrics

**Focus:** Time buckets

## Your task
Implement:

```go
type CounterWindow struct
```

### Learning goal
- Expected work: Implement `Fixed-Window Metrics` with precise clock/window state updates so timing behavior remains deterministic in tests.
- Why: `Fixed-Window Metrics` teaches temporal correctness, a common source of flaky and hard-to-debug behavior.
- When correct: `Fixed-Window Metrics` should satisfy the required behavior, including: `Add(value)`; `Sum(last duration)`; and `thread-safe`.

## Rules / Expectations
- Add(value)
- Sum(last duration)
- thread-safe

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains time buckets by implementing `Fixed-Window Metrics` under explicit constraints.
- It is important because `Fixed-Window Metrics` practices time-based logic that frequently causes subtle production bugs.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
