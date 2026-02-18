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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
