# Kata 020 — Roman Numerals

**Focus:** Greedy mapping

## Your task
Implement:

```go
func ToRoman(n int) (string, error)
```

## Rules / Expectations
- 1..3999 only
- invalid => error
- use standard subtractive notation

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Roman Numerals with constraints that make you practice greedy mapping.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
