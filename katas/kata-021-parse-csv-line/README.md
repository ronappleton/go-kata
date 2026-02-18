# Kata 021 — Parse CSV Line

**Focus:** Parsing quoted CSV

## Your task
Implement:

```go
func ParseCSVLine(line string) ([]string, error)
```

### Learning goal
- What you are building: Build `func ParseCSVLine(line string) ([]string, error)` as a reliable contract. Focus: Parsing quoted CSV.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: commas separate fields; double quotes allow commas; and escaped quotes as "" inside quoted field.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

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
- Core lesson: parse strictly and fail loudly on bad input.
- After this kata, you should be able to state input contracts and return precise parse errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
