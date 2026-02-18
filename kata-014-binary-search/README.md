# Kata 014 — Binary Search

**Focus:** Algorithms, generics

## Your task
Implement:

```go
func BinarySearch[T ~int | ~int64 | ~float64 | ~string](sorted []T, target T) int
```

## Rules / Expectations
- return index or -1
- empty => -1
- assumes sorted input

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)

## What this kata is about (and why it matters)
- This kata is about implementing Binary Search with constraints that make you practice algorithms, generics.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
