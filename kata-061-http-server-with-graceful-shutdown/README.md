# Kata 061 — HTTP Server with Graceful Shutdown

**Focus:** net/http, context

## Your task
Implement:

```go
func ServeGraceful(addr string, h http.Handler) (stop func(context.Context) error, err error)
```

### Learning goal
- Expected work: Implement `HTTP Server with Graceful Shutdown` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `HTTP Server with Graceful Shutdown` teaches service-edge correctness, where request/response semantics directly drive reliability.

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
- This kata trains net/http, context by implementing `HTTP Server with Graceful Shutdown` under explicit constraints.
- It is important because `HTTP Server with Graceful Shutdown` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
