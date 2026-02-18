# Kata 022 — JSON Pretty Print

**Focus:** encoding/json

## Your task
Implement:

```go
func PrettyJSON(input []byte) ([]byte, error)
```

## Rules / Expectations
- invalid json => error
- indent 2 spaces
- stable formatting

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
