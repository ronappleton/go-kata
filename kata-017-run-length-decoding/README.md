# Kata 017 — Run-Length Decoding

**Focus:** Parsing, errors

## Your task
Implement:

```go
func RLDecode(s string) (string, error)
```

## Rules / Expectations
- invalid format => error
- count can be multiple digits
- roundtrip with encode for valid inputs

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Run-Length Decoding with constraints that make you practice parsing, errors.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
