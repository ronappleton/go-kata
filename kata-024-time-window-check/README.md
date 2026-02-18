# Kata 024 — Time Window Check

**Focus:** time package

## Your task
Implement:

```go
func WithinWindow(t, start, end time.Time) bool
```

## Rules / Expectations
- inclusive start, exclusive end
- if start after end => false

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
