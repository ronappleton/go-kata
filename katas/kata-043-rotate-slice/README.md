# Kata 043 — Rotate Slice

**Focus:** Generics, indexing

## Your task
Implement:

```go
func RotateLeft[T any](items []T, k int) []T
```

### Learning goal
- What you are building: Build `func RotateLeft[T any](items []T, k int) []T` as a reliable contract. Focus: Generics, indexing.
- Why this matters in real projects: Good structure compounds. Clear contracts reduce future rework.
- How this grows your Go skills: You practice interface-first design and explicit dependencies.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: k can be > len; k can be negative (rotate right); and for `nil`, return `empty slice`.

### Tips
- Write contract tests before implementation details.
- Keep dependencies explicit, not implicit.
- Prefer small interfaces with one reason to change.

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
- Core lesson: design seams for testability and change.
- After this kata, you should be able to describe component contracts and test at those seams.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
