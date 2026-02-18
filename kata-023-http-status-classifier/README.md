# Kata 023 — HTTP Status Classifier

**Focus:** Switch ranges

## Your task
Implement:

```go
func StatusClass(code int) string
```

### Learning goal
- What you are practicing: Build `HTTP Status Classifier` as production-style HTTP boundary code with clear request handling and response behavior.
- Why it matters: You will use this in APIs and services where request/response correctness directly affects reliability.
- How this grows your Go skills: This builds confidence at service edges: inputs, status codes, retries, and shutdown behavior.
- When correct: When your solution is correct, it should satisfy: `1xx informational`; `2xx success`; `3xx redirect`; `4xx client`; and `5xx server`.

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
- This kata is focused practice in Switch ranges through `HTTP Status Classifier`.
- You will use this in APIs and services where request/response correctness directly affects reliability.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
