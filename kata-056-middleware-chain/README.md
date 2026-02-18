# Kata 056 — Middleware Chain

**Focus:** net/http

## Your task
Implement:

```go
func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler
```

### Learning goal
- Expected work: Implement `Middleware Chain` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `Middleware Chain` teaches service-edge correctness, where request/response semantics directly drive reliability.
- When correct: `Middleware Chain` should satisfy the required behavior, including: `applies in order given`; and `nil middleware ignored`.

## Rules / Expectations
- applies in order given
- nil middleware ignored

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata trains net/http by implementing `Middleware Chain` under explicit constraints.
- It is important because `Middleware Chain` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
