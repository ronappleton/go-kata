# Kata 014 — Binary Search

**Focus:** Algorithms, generics

## Your task
Implement:

```go
func BinarySearch[T ~int | ~int64 | ~float64 | ~string](sorted []T, target T) int
```

### Learning goal
- What you are practicing: Build `Binary Search` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `return index or -1`; `empty => -1`; and `assumes sorted input`.

## Rules / Expectations
- return index or -1
- empty => -1
- assumes sorted input

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)

## What this kata is about (and why it matters)
- This kata is focused practice in Algorithms, generics through `Binary Search`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
