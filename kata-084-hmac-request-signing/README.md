# Kata 084 — HMAC Request Signing

**Focus:** crypto/hmac

## Your task
Implement:

```go
func SignRequest(method, path string, body []byte, secret []byte) string
```

## Rules / Expectations
- canonicalize input
- hex signature
- write VerifyRequest helper

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
