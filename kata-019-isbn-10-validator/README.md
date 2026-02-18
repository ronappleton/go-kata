# Kata 019 — ISBN-10 Validator

**Focus:** Checksums

## Your task
Implement:

```go
func IsValidISBN10(s string) bool
```

## Rules / Expectations
- allow hyphens/spaces
- X allowed as last digit meaning 10
- must be 10 digits after stripping

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing ISBN-10 Validator with constraints that make you practice checksums.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
