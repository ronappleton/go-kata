# Kata 018 — Caesar Cipher

**Focus:** ASCII letters shifting

## Your task
Implement:

```go
func Caesar(s string, shift int) string
```

## Rules / Expectations
- only shift A-Z and a-z
- preserve case
- non-letters unchanged

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Caesar Cipher with constraints that make you practice ascii letters shifting.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
