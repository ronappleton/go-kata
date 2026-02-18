# Kata 035 — Log Line Parser

**Focus:** Key=value parsing

## Your task
Implement:

```go
func ParseLogLine(line string) (map[string]string, error)
```

## Rules / Expectations
- key=value pairs separated by spaces
- values may be quoted
- duplicate keys last wins

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
