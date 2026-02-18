# Kata 039 — UUID v4 Generator

**Focus:** crypto/rand

## Your task
Implement:

```go
func NewUUIDv4() (string, error)
```

## Rules / Expectations
- RFC4122 version/variant bits
- format 8-4-4-4-12 lowercase

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
