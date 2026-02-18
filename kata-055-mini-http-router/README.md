# Kata 055 — Mini HTTP Router

**Focus:** net/http

## Your task
Implement:

```go
func NewRouter() *Router
```

### Learning goal
- What you are practicing: Build `Mini HTTP Router` as production-style HTTP boundary code with clear request handling and response behavior.
- Why it matters: You will use this in APIs and services where request/response correctness directly affects reliability.
- How this grows your Go skills: This builds confidence at service edges: inputs, status codes, retries, and shutdown behavior.
- When correct: When your solution is correct, it should satisfy: `GET/POST register`; `path params /users/{id}`; and `404 not found`.

## Rules / Expectations
- GET/POST register
- path params /users/{id}
- 404 not found

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata is focused practice in net/http through `Mini HTTP Router`.
- You will use this in APIs and services where request/response correctness directly affects reliability.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
