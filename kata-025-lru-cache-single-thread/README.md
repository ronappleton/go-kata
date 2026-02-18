# Kata 025 — LRU Cache (single-thread)

**Focus:** Structs, maps, list

## Your task
Implement:

```go
func NewLRU(capacity int) (*LRU, error)
```

## Rules / Expectations
- capacity>0
- Get/Put methods
- evict least recently used

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
