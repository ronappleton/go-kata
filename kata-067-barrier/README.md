# Kata 067 — Barrier

**Focus:** Concurrency coordination

## Your task
Implement:

```go
func NewBarrier(parties int) (*Barrier, error)
```

## Rules / Expectations
- Wait(ctx) error
- reusable
- no leaks

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
