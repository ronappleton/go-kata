# Kata 055 — Mini HTTP Router

**Focus:** net/http

## Your task
Implement:

```go
func NewRouter() *Router
```

### Learning goal
- Expected work: Implement `Mini HTTP Router` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `Mini HTTP Router` teaches service-edge correctness, where request/response semantics directly drive reliability.

## Rules / Expectations
- GET/POST register
- path params /users/{id}
- 404 not found

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata trains net/http by implementing `Mini HTTP Router` under explicit constraints.
- It is important because `Mini HTTP Router` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
