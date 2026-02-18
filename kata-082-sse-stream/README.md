# Kata 082 — SSE Stream

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata82() error
```

### Learning goal
- Expected work: Implement `SSE Stream` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `SSE Stream` teaches service-edge correctness, where request/response semantics directly drive reliability.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `SSE Stream` under explicit constraints.
- It is important because `SSE Stream` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
