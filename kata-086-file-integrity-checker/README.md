# Kata 086 — File Integrity Checker

**Focus:** io, crypto/sha256

## Your task
Implement:

```go
func SHA256File(path string) (string, error)
```

## Rules / Expectations
- stream file
- hex digest
- error on read

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
