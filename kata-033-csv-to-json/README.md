# Kata 033 — CSV to JSON

**Focus:** I/O, encoding

## Your task
Implement:

```go
func CSVToJSON(r io.Reader) ([]byte, error)
```

### Learning goal
- What you are practicing: Build `CSV to JSON` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `first row headers`; `JSON array of objects`; and `trim spaces`.

## Rules / Expectations
- first row headers
- JSON array of objects
- trim spaces

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go encoding/csv package](https://pkg.go.dev/encoding/csv)
- [RFC 4180 (CSV format)](https://www.rfc-editor.org/rfc/rfc4180)

## What this kata is about (and why it matters)
- This kata is focused practice in I/O, encoding through `CSV to JSON`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
