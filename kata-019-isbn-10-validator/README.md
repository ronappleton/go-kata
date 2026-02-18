# Kata 019 — ISBN-10 Validator

**Focus:** Checksums

## Your task
Implement:

```go
func IsValidISBN10(s string) bool
```

## Rules / Expectations
- allow hyphens/spaces
- X allowed as last digit meaning 10
- must be 10 digits after stripping

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
