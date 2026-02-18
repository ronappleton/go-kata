# Kata 085 — Password Hashing (bcrypt/argon2)

**Focus:** crypto, security

## Your task
Implement:

```go
func HashPassword(pw string) (string, error)
```

### Learning goal
- Expected work: Implement `Password Hashing (bcrypt/argon2)` with strict validation and byte-level correctness for all cryptographic operations.
- Why: `Password Hashing (bcrypt/argon2)` teaches security precision, where small mistakes can become real vulnerabilities.
- When correct: `Password Hashing (bcrypt/argon2)` should satisfy the required behavior, including: `use bcrypt or argon2`; `also implement VerifyPassword(pw, hash) bool`; and `constant-time compare`.

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
- This kata trains crypto, security by implementing `Password Hashing (bcrypt/argon2)` under explicit constraints.
- It is important because `Password Hashing (bcrypt/argon2)` reflects production security code, where correctness mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
