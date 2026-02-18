# Kata 036 — Rate Limiter (token bucket)

**Focus:** Time, structs

## Your task
Implement:

```go
func NewTokenBucket(ratePerSec int, burst int) (*TokenBucket, error)
```

## Rules / Expectations
- Allow() bool
- single-thread only here
- ratePerSec>=1, burst>=1

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
