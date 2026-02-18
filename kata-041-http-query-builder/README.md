# Kata 041 — HTTP Query Builder

**Focus:** net/url

## Your task
Implement:

```go
func BuildQuery(params map[string][]string) string
```

### Learning goal
- Expected work: Implement `HTTP Query Builder` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `HTTP Query Builder` teaches service-edge correctness, where request/response semantics directly drive reliability.
- When correct: `HTTP Query Builder` should satisfy the required behavior, including: `keys sorted`; `values preserve given order`; and `proper escaping`.

## Rules / Expectations
- keys sorted
- values preserve given order
- proper escaping

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net/url package](https://pkg.go.dev/net/url)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata trains net/url by implementing `HTTP Query Builder` under explicit constraints.
- It is important because `HTTP Query Builder` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
