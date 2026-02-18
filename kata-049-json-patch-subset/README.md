# Kata 049 — JSON Patch (subset)

**Focus:** encoding/json

## Your task
Implement:

```go
func ApplyPatch(doc []byte, patch []byte) ([]byte, error)
```

## Rules / Expectations
- support add/replace/remove at top-level keys only
- invalid => error

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
