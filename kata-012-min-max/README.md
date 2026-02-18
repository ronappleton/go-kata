# Kata 012 — Min/Max

**Focus:** Errors, edge cases

## Your task
Implement:

```go
func MinMax(nums []int) (min int, max int, err error)
```

## Rules / Expectations
- empty or nil => error
- single element => min=max
- handles negatives

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
