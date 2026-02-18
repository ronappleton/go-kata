# Kata 016 — Run-Length Encoding

**Focus:** Strings.Builder, runes

## Your task
Implement:

```go
func RLEncode(s string) string
```

## Rules / Expectations
- encode runs as <char><count> where count>1
- single chars stay as char only
- Unicode-safe

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Run-Length Encoding with constraints that make you practice strings.builder, runes.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
