# Kata 021 — Parse CSV Line

**Focus:** Parsing quoted CSV

## Your task
Implement:

```go
func ParseCSVLine(line string) ([]string, error)
```

## Rules / Expectations
- commas separate fields
- double quotes allow commas
- escaped quotes as "" inside quoted field

## Prior reading
- [Go encoding/csv package](https://pkg.go.dev/encoding/csv)
- [RFC 4180 (CSV format)](https://www.rfc-editor.org/rfc/rfc4180)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Parse CSV Line with constraints that make you practice parsing quoted csv.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
