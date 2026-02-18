# Kata 015 — Merge Two Sorted Lists

**Focus:** Two pointers

## Your task
Implement:

```go
func MergeSorted(a, b []int) []int
```

## Rules / Expectations
- result sorted
- handles duplicates
- nil treated as empty

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Merge Two Sorted Lists with constraints that make you practice two pointers.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
