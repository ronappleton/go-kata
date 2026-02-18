# Kata 096 — Benchmarking Kata

**Focus:** testing/benchmark

## Your task
Implement:

```go
func Concat(parts []string) string
```

## Rules / Expectations
- write 2 benchmarks: naive vs builder
- compare allocations
- document results in README

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
