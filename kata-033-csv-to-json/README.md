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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
