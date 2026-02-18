# Kata 021 — Parse CSV Line

**Focus:** Parsing quoted CSV

## Your task
Implement:

```go
func ParseCSVLine(line string) ([]string, error)
```

### Learning goal
- What you are practicing: Build `Parse CSV Line` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `commas separate fields`; `double quotes allow commas`; and `escaped quotes as "" inside quoted field`.

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
- This kata is focused practice in Parsing quoted CSV through `Parse CSV Line`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
