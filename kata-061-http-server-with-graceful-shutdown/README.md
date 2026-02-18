# Kata 061 — HTTP Server with Graceful Shutdown

**Focus:** net/http, context

## Your task
Implement:

```go
func ServeGraceful(addr string, h http.Handler) (stop func(context.Context) error, err error)
```

### Learning goal
- What you are practicing: Build `HTTP Server with Graceful Shutdown` as production-style HTTP boundary code with clear request handling and response behavior.
- Why it matters: You will use this in APIs and services where request/response correctness directly affects reliability.
- How this grows your Go skills: This builds confidence at service edges: inputs, status codes, retries, and shutdown behavior.
- When correct: When your solution is correct, it should satisfy: `start server`; `stop shuts down with context`; and `no leaked goroutines`.

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
- This kata is focused practice in net/http, context through `HTTP Server with Graceful Shutdown`.
- You will use this in APIs and services where request/response correctness directly affects reliability.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
