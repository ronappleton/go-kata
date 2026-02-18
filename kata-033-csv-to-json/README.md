# Kata 033 — CSV to JSON

**Focus:** I/O, encoding

## Your task
Implement:

```go
func CSVToJSON(r io.Reader) ([]byte, error)
```

### Learning goal
- What you are building: Build `func CSVToJSON(r io.Reader) ([]byte, error)` as a reliable contract. Focus: I/O, encoding.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: use the first row as field names; emit a JSON array of objects; and trim required surrounding spaces.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

## Rules / Expectations
- first row headers
- JSON array of objects
- trim spaces

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go encoding/csv package](https://pkg.go.dev/encoding/csv)
- [RFC 4180 (CSV format)](https://www.rfc-editor.org/rfc/rfc4180)

## What this kata is about (and why it matters)
- Core lesson: parse strictly and fail loudly on bad input.
- After this kata, you should be able to state input contracts and return precise parse errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
