# Kata 085 — Password Hashing (bcrypt/argon2)

**Focus:** crypto, security

## Your task
Implement:

```go
func HashPassword(pw string) (string, error)
```

### Learning goal
- What you are practicing: Build `Password Hashing (bcrypt/argon2)` with careful validation and exact byte-level behavior for security-sensitive inputs/outputs.
- Why it matters: You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.
- How this grows your Go skills: This builds your discipline for correctness-first coding in high-risk code paths.
- When correct: When your solution is correct, it should satisfy: `use bcrypt or argon2`; `also implement VerifyPassword(pw, hash) bool`; and `constant-time compare`.

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
- This kata is focused practice in crypto, security through `Password Hashing (bcrypt/argon2)`.
- You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
