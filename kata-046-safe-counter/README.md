# Kata 046 — Safe Counter

**Focus:** sync/atomic or mutex

## Your task
Implement:

```go
type Counter struct
```

### Learning goal
- What you are practicing: Build `Safe Counter` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `Inc/Add/Value`; and `race-free`.

## Rules / Expectations
- Inc/Add/Value
- race-free

## Prior reading
- [Go sync/atomic package](https://pkg.go.dev/sync/atomic)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in sync/atomic or mutex through `Safe Counter`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
