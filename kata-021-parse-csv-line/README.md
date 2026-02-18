# Kata 021 — Parse CSV Line

**Focus:** Parsing quoted CSV

## Your task
Implement:

```go
func ParseCSVLine(line string) ([]string, error)
```

### Learning goal
- Expected work: Implement `Parse CSV Line` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `Parse CSV Line` teaches strict boundary handling so malformed input cannot silently corrupt results.
- When correct: `Parse CSV Line` should satisfy the required behavior, including: `commas separate fields`; `double quotes allow commas`; and `escaped quotes as "" inside quoted field`.

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
- This kata trains parsing quoted csv by implementing `Parse CSV Line` under explicit constraints.
- It is important because `Parse CSV Line` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
