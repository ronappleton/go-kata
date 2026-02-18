# Kata 014 — Binary Search

**Focus:** Algorithms, generics

## Your task
Implement:

```go
func BinarySearch[T ~int | ~int64 | ~float64 | ~string](sorted []T, target T) int
```

### Learning goal
- Expected work: Implement `Binary Search` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Binary Search` teaches complexity-aware correctness that impacts throughput and latency at scale.
- When correct: `Binary Search` should satisfy the required behavior, including: `return index or -1`; `empty => -1`; and `assumes sorted input`.

## Rules / Expectations
- return index or -1
- empty => -1
- assumes sorted input

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)

## What this kata is about (and why it matters)
- This kata trains algorithms, generics by implementing `Binary Search` under explicit constraints.
- It is important because `Binary Search` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
