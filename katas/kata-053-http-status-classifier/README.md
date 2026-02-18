# Kata 053 — HTTP Status Classifier

**Focus:** Switch ranges

## Your task
Implement:

```go
func StatusClass(code int) string
```

### Learning goal
- What you are building: Build `func StatusClass(code int) string` as a reliable contract. Focus: Switch ranges.
- Why this matters in real projects: HTTP boundaries are product behavior. Callers depend on exact semantics.
- How this grows your Go skills: You practice context-aware request handling and stable status/error contracts.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: 1xx informational; 2xx success; 3xx redirect; 4xx client; and 5xx server.

### Tips
- Pin boundary behavior with `httptest`.
- Cover both success and failure responses.
- Test retries/timeouts with targeted cases.

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
- Core lesson: treat request/response behavior as a hard contract.
- After this kata, you should be able to justify status/error choices at the service boundary.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
