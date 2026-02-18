# Kata 033 — CSV to JSON

**Focus:** I/O, encoding

## Your task
Implement:

```go
func CSVToJSON(r io.Reader) ([]byte, error)
```

### Learning goal
- Expected work: Implement `CSV to JSON` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `CSV to JSON` teaches strict boundary handling so malformed input cannot silently corrupt results.
- When correct: `CSV to JSON` should satisfy the required behavior, including: `first row headers`; `JSON array of objects`; and `trim spaces`.

## Rules / Expectations
- first row headers
- JSON array of objects
- trim spaces

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go encoding/csv package](https://pkg.go.dev/encoding/csv)
- [RFC 4180 (CSV format)](https://www.rfc-editor.org/rfc/rfc4180)

## What this kata is about (and why it matters)
- This kata trains i/o, encoding by implementing `CSV to JSON` under explicit constraints.
- It is important because `CSV to JSON` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
