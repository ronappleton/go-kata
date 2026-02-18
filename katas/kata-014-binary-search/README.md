# Kata 014 — Binary Search

**Focus:** Algorithms, generics

## Your task
Implement:

```go
func BinarySearch[T ~int | ~int64 | ~float64 | ~string](sorted []T, target T) int
```

### Learning goal
- What you are building: Build `func BinarySearch[T ~int | ~int64 | ~float64 | ~string](sorted []T, target T) int` as a reliable contract. Focus: Algorithms, generics.
- Why this matters in real projects: This is how you keep features fast when input size grows.
- How this grows your Go skills: You practice invariants and complexity reasoning, then prove both with tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: return index or -1; for `empty`, return `-1`; and assumes sorted input.

### Tips
- State the invariant first, then code.
- Test shape edges early: empty, one item, duplicates.
- Check complexity after correctness.

## Rules / Expectations
- return index or -1
- empty => -1
- assumes sorted input

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)

## What this kata is about (and why it matters)
- Core lesson: hold invariants first, then optimize safely.
- After this kata, you should be able to defend algorithm choice, complexity, and corner-case behavior.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
