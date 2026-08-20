# Kata 086 — Middleware Chain

**Focus:** net/http

## Your task
Implement:

```go
func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler
```

### Learning goal
- What you are building: Build `func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler` as a reliable contract. Focus: net/http.
- Why this matters in real projects: HTTP boundaries are product behavior. Callers depend on exact semantics.
- How this grows your Go skills: You practice context-aware request handling and stable status/error contracts.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: applies in order given; and nil middleware ignored.

### Tips
- Pin boundary behavior with `httptest`.
- Cover both success and failure responses.
- Test retries/timeouts with targeted cases.

## Rules / Expectations
- applies in order given
- nil middleware ignored

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

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
