# Kata 033 — CSV to JSON

**Focus:** I/O, encoding

## Your task
Implement:

```go
func CSVToJSON(r io.Reader) ([]byte, error)
```

## Rules / Expectations
- first row headers
- JSON array of objects
- trim spaces

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go encoding/csv package](https://pkg.go.dev/encoding/csv)
- [RFC 4180 (CSV format)](https://www.rfc-editor.org/rfc/rfc4180)

## What this kata is about (and why it matters)
- This kata is about implementing CSV to JSON with constraints that make you practice i/o, encoding.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
