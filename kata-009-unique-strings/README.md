# Kata 009 — Unique Strings

**Focus:** Maps as sets, order preservation

## Your task
Implement:

```go
func Unique(items []string) []string
```

## Rules / Expectations
- preserve first occurrence order
- nil => empty slice
- case-sensitive

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata is about implementing Unique Strings with constraints that make you practice maps as sets, order preservation.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
