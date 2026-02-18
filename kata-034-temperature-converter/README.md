# Kata 034 — Temperature Converter

**Focus:** Parsing floats

## Your task
Implement:

```go
func ConvertTemp(input string) (string, error)
```

## Rules / Expectations
- supports C/F/K suffix
- output all 3 units with 2 decimals
- invalid => error

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Temperature Converter with constraints that make you practice parsing floats.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
