# Kata 032 — Top N Words

**Focus:** Sorting, ties

## Your task
Implement:

```go
func TopNWords(s string, n int) []string
```

## Rules / Expectations
- lowercase, split whitespace
- ties sorted alphabetically
- n<=0 => empty slice

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Top N Words with constraints that make you practice sorting, ties.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
