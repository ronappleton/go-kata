# Kata 023 — HTTP Status Classifier

**Focus:** Switch ranges

## Your task
Implement:

```go
func StatusClass(code int) string
```

### Learning goal
- Expected work: Implement `HTTP Status Classifier` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `HTTP Status Classifier` teaches service-edge correctness, where request/response semantics directly drive reliability.

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
- This kata trains switch ranges by implementing `HTTP Status Classifier` under explicit constraints.
- It is important because `HTTP Status Classifier` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
