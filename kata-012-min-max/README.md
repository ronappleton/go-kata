# Kata 012 — Min/Max

**Focus:** Errors, edge cases

## Your task
Implement:

```go
func MinMax(nums []int) (min int, max int, err error)
```

## Rules / Expectations
- empty or nil => error
- single element => min=max
- handles negatives

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Min/Max with constraints that make you practice errors, edge cases.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
