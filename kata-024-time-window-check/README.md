# Kata 024 — Time Window Check

**Focus:** time package

## Your task
Implement:

```go
func WithinWindow(t, start, end time.Time) bool
```

### Learning goal
- Expected work: Implement `Time Window Check` with precise clock/window state updates so timing behavior remains deterministic in tests.
- Why: `Time Window Check` teaches temporal correctness, a common source of flaky and hard-to-debug behavior.

## Rules / Expectations
- inclusive start, exclusive end
- if start after end => false

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains time package by implementing `Time Window Check` under explicit constraints.
- It is important because `Time Window Check` practices time-based logic that frequently causes subtle production bugs.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
