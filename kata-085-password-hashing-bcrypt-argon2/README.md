# Kata 085 — Password Hashing (bcrypt/argon2)

**Focus:** crypto, security

## Your task
Implement:

```go
func HashPassword(pw string) (string, error)
```

### Learning goal
- What you are building: Build `func HashPassword(pw string) (string, error)` as a reliable contract. Focus: crypto, security.
- Why this matters in real projects: Security paths have no soft failures. Exact behavior is non-negotiable.
- How this grows your Go skills: You practice fail-closed validation and byte-accurate handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: use bcrypt or argon2; also provide `VerifyPassword(pw, hash) bool`; and constant-time compare.

### Tips
- Fail closed on malformed data.
- Never ignore crypto or encoding errors.
- Add tamper and expiry negative tests.

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
- Core lesson: validate early, fail closed, and be exact.
- After this kata, you should be able to justify validation rules and add negative tests for tampering.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
