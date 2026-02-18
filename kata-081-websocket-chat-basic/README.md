# Kata 081 — Websocket Chat (basic)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata81() error
```

### Learning goal
- Expected work: Implement `Websocket Chat (basic)` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `Websocket Chat (basic)` teaches service-edge correctness, where request/response semantics directly drive reliability.
- When correct: `Websocket Chat (basic)` should satisfy the required behavior, including: `follow README spec`; `write tests`; and `keep it idiomatic`.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `Websocket Chat (basic)` under explicit constraints.
- It is important because `Websocket Chat (basic)` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
