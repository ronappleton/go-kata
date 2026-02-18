# Kata 030 — Kebab/Snake → Camel

**Focus:** String transforms

## Your task
Implement:

```go
func ToCamel(s string) string
```

## Rules / Expectations
- input kebab-case or snake_case
- output lowerCamelCase
- collapse multiple separators

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Kebab/Snake → Camel with constraints that make you practice string transforms.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
