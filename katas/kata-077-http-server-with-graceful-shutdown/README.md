# Kata 091 — HTTP Server with Graceful Shutdown

**Focus:** net/http, context

## Your task
Implement:

```go
func ServeGraceful(addr string, h http.Handler) (stop func(context.Context) error, err error)
```

### Learning goal
- What you are building: Build `func ServeGraceful(addr string, h http.Handler) (stop func(context.Context) error, err error)` as a reliable contract. Focus: net/http, context.
- Why this matters in real projects: HTTP boundaries are product behavior. Callers depend on exact semantics.
- How this grows your Go skills: You practice context-aware request handling and stable status/error contracts.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: start server; stop shuts down with context; and no leaked goroutines.

### Tips
- Pin boundary behavior with `httptest`.
- Cover both success and failure responses.
- Test retries/timeouts with targeted cases.

## Rules / Expectations
- start server
- stop shuts down with context
- no leaked goroutines

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)
- [Go context package](https://pkg.go.dev/context)

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
