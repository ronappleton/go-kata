# Kata 013 — Rotate Slice

**Focus:** Generics, indexing

## Your task
Implement:

```go
func RotateLeft[T any](items []T, k int) []T
```

### Learning goal
- What you are practicing: Build `Rotate Slice` as a reusable, testable abstraction with explicit contracts.
- Why it matters: You will use this when designing packages that need to stay maintainable as features grow.
- How this grows your Go skills: This builds API design judgment and composable package structure in Go.
- When correct: When your solution is correct, it should satisfy: `k can be > len`; `k can be negative (rotate right)`; and `nil => empty slice`.

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
- This kata is focused practice in Generics, indexing through `Rotate Slice`.
- You will use this when designing packages that need to stay maintainable as features grow.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
