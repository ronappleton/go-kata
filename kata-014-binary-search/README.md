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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
