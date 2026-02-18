# Kata 043 — Retry with Backoff

**Focus:** time

## Your task
Implement:

```go
func Retry(attempts int, baseDelay time.Duration, fn func() error) error
```

### Learning goal
- Expected work: Implement `Retry with Backoff` with precise clock/window state updates so timing behavior remains deterministic in tests.
- Why: `Retry with Backoff` teaches temporal correctness, a common source of flaky and hard-to-debug behavior.

## Rules / Expectations
- attempts>=1
- exponential backoff
- stop on success

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains time by implementing `Retry with Backoff` under explicit constraints.
- It is important because `Retry with Backoff` practices time-based logic that frequently causes subtle production bugs.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
