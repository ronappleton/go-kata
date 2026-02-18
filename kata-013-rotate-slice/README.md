# Kata 013 — Rotate Slice

**Focus:** Generics, indexing

## Your task
Implement:

```go
func RotateLeft[T any](items []T, k int) []T
```

### Learning goal
- Expected work: Implement `Rotate Slice` as a composable abstraction with explicit contracts and testable behavior.
- Why: `Rotate Slice` teaches maintainable design through explicit contracts and low coupling.
- When correct: `Rotate Slice` should satisfy the required behavior, including: `k can be > len`; `k can be negative (rotate right)`; and `nil => empty slice`.

## Rules / Expectations
- k can be > len
- k can be negative (rotate right)
- nil => empty slice

## Prior reading
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains generics, indexing by implementing `Rotate Slice` under explicit constraints.
- It is important because `Rotate Slice` develops design choices that improve extensibility and testability over time.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
