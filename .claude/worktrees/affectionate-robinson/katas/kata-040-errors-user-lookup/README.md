# Kata 030 — Find User Name

**Focus:** errors.Is, sentinel errors, strings.TrimSpace

## Your task
Implement:

```go
func FindUserName(users map[string]string, id string) (string, error)
```

### Learning goal
- What you are building: a lookup function with typed error signaling for invalid and missing IDs.
- Why this matters in real projects: downstream behavior depends on accurate error classification.
- How this grows your Go skills: you practice sentinel errors and wrapping for `errors.Is` checks.
- Definition of done (plain English): callers can reliably distinguish invalid ID from missing user.

### Tips
- Define explicit sentinel errors and wrap them with context.
- Trim input IDs before checks.
- Test with `errors.Is`, not string comparison.

## Rules / Expectations
- blank id after trim => return error wrapping `ErrInvalidUserID`
- unknown id => return error wrapping `ErrUserNotFound`
- found name should be returned trimmed
- successful lookup => name and nil error

## Prior reading
- [errors package](https://pkg.go.dev/errors)
- [fmt package](https://pkg.go.dev/fmt)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: error values are part of your API contract.
- After this kata, you should be able to design result/error behavior that callers can branch on safely.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
