# Kata 138 — Read CSV Records

**Focus:** encoding/csv reader loops

## Your task
Implement:

```go
func ReadCSVRecords(r *csv.Reader) ([][]string, error)
```

### Learning goal
- What you are building: CSV ingestion logic with explicit EOF and malformed-input handling.
- Why this matters in real projects: CSV import pipelines are common and failure modes must be predictable.
- How this grows your Go skills: you practice structured parsing with looped read boundaries.
- Definition of done (plain English): behavior matches all expectations below.

## Rules / Expectations
- nil reader => non-nil error
- read all records until EOF
- preserve field ordering
- malformed CSV => non-nil error

## Prior reading
- [encoding/csv package](https://pkg.go.dev/encoding/csv)
- [Go package documentation index](https://pkg.go.dev/std)

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
