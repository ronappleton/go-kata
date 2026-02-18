# Kata 054 — JSON Lines Filter

**Focus:** Stream processing

## Your task
Implement:

```go
func FilterJSONL(r io.Reader, w io.Writer, predicate func(map[string]any) bool) error
```

## Rules / Expectations
- each line JSON object
- write only matches
- skip blank lines

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
