# Kata 016 — Run-Length Encoding

**Focus:** Strings.Builder, runes

## Your task
Implement:

```go
func RLEncode(s string) string
```

## Rules / Expectations
- encode runs as <char><count> where count>1
- single chars stay as char only
- Unicode-safe

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
