# Kata 023 — HTTP Status Classifier

**Focus:** Switch ranges

## Your task
Implement:

```go
func StatusClass(code int) string
```

## Rules / Expectations
- 1xx informational
- 2xx success
- 3xx redirect
- 4xx client
- 5xx server
- otherwise unknown

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
