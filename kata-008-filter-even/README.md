# Kata 008 — Filter Even

**Focus:** Slices, order, immutability

## Your task
Implement:

```go
func FilterEven(nums []int) []int
```

## Rules / Expectations
- preserve order
- nil => empty slice
- do not mutate input

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Filter Even with constraints that make you practice slices, order, immutability.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
