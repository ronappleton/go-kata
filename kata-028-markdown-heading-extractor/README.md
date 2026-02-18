# Kata 028 — Markdown Heading Extractor

**Focus:** Line scanning

## Your task
Implement:

```go
func ExtractHeadings(md string) []string
```

## Rules / Expectations
- lines starting with #+space
- return heading text
- preserve order

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Markdown Heading Extractor with constraints that make you practice line scanning.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
