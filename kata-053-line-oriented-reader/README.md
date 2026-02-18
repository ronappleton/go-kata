# Kata 053 — Line-Oriented Reader

**Focus:** bufio

## Your task
Implement:

```go
func ReadLines(r io.Reader, maxLine int) ([]string, error)
```

## Rules / Expectations
- error if any line exceeds maxLine
- keep line content only
- handles last line

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
