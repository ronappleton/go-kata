# Kata 002 — Sum of Integers

**Focus:** Basics: loops, function, edge cases

## Your task
Implement:

```go
func Sum(nums []int) int
```

## Rules / Expectations
- nil slice => 0
- empty => 0
- works with negatives

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Sum of Integers with constraints that make you practice basics: loops, function, edge cases.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
