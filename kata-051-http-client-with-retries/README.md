# Kata 051 — HTTP Client with Retries

**Focus:** net/http

## Your task
Implement:

```go
func DoWithRetries(client *http.Client, req *http.Request, attempts int) (*http.Response, error)
```

### Learning goal
- Expected work: Implement `HTTP Client with Retries` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `HTTP Client with Retries` teaches service-edge correctness, where request/response semantics directly drive reliability.
- When correct: `HTTP Client with Retries` should satisfy the required behavior, including: `retry on 5xx + network errors`; `no retry on 4xx`; and `safe body handling`.

## Rules / Expectations
- retry on 5xx + network errors
- no retry on 4xx
- safe body handling

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata trains net/http by implementing `HTTP Client with Retries` under explicit constraints.
- It is important because `HTTP Client with Retries` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
