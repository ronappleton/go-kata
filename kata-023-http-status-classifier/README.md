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

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing HTTP Status Classifier with constraints that make you practice switch ranges.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
