# Kata 083 — JWT Sign/Verify

**Focus:** crypto, encoding

## Your task
Implement:

```go
func SignJWT(claims map[string]any, secret []byte, ttl time.Duration) (string, error)
```

## Rules / Expectations
- HS256
- include exp
- also implement VerifyJWT(token, secret)

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
