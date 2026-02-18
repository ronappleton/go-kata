# Kata 056 — Middleware Chain

**Focus:** net/http

## Your task
Implement:

```go
func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler
```

## Rules / Expectations
- applies in order given
- nil middleware ignored

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata is about implementing Middleware Chain with constraints that make you practice net/http.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
