# Kata 051 — HTTP Client with Retries

**Focus:** net/http

## Your task
Implement:

```go
func DoWithRetries(client *http.Client, req *http.Request, attempts int) (*http.Response, error)
```

## Rules / Expectations
- retry on 5xx + network errors
- no retry on 4xx
- safe body handling

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
