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

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [OWASP Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [golang.org/x/crypto/argon2](https://pkg.go.dev/golang.org/x/crypto/argon2)

## What this kata is about (and why it matters)
- This kata is about implementing Password Hashing (bcrypt/argon2) with constraints that make you practice crypto, security.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
