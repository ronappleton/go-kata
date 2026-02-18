# Kata 031 — File Extension Counter

**Focus:** os, filepath

## Your task
Implement:

```go
func CountExt(root string) (map[string]int, error)
```

## Rules / Expectations
- walk dir recursively
- count by extension (lowercased)
- skip directories

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
