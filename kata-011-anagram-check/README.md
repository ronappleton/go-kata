# Kata 011 — Anagram Check

**Focus:** Rune counts, normalization

## Your task
Implement:

```go
func IsAnagram(a, b string) bool
```

## Rules / Expectations
- case-insensitive
- ignore spaces/punctuation
- Unicode-safe via rune counts

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Anagram Check with constraints that make you practice rune counts, normalization.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
