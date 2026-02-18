# Kata 017 — Run-Length Decoding

**Focus:** Parsing, errors

## Your task
Implement:

```go
func RLDecode(s string) (string, error)
```

## Rules / Expectations
- invalid format => error
- count can be multiple digits
- roundtrip with encode for valid inputs

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
