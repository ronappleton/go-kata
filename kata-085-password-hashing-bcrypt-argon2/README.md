# Kata 085 — Password Hashing (bcrypt/argon2)

**Focus:** crypto, security

## Your task
Implement:

```go
func HashPassword(pw string) (string, error)
```

## Rules / Expectations
- use bcrypt or argon2
- also implement VerifyPassword(pw, hash) bool
- constant-time compare

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
