# Kata 056 — Middleware Chain

**Focus:** net/http

## Your task
Implement:

```go
func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler
```

## Rules / Expectations
- applies in order given
- nil middleware ignored

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
