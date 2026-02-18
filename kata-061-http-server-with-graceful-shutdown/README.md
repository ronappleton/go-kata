# Kata 061 — HTTP Server with Graceful Shutdown

**Focus:** net/http, context

## Your task
Implement:

```go
func ServeGraceful(addr string, h http.Handler) (stop func(context.Context) error, err error)
```

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
- This kata is about implementing HTTP Server with Graceful Shutdown with constraints that make you practice net/http, context.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
