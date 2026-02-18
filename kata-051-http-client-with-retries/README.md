# Kata 051 — HTTP Client with Retries

**Focus:** net/http

## Your task
Implement:

```go
func DoWithRetries(client *http.Client, req *http.Request, attempts int) (*http.Response, error)
```

### Learning goal
- What you are practicing: Build `HTTP Client with Retries` as production-style HTTP boundary code with clear request handling and response behavior.
- Why it matters: You will use this in APIs and services where request/response correctness directly affects reliability.
- How this grows your Go skills: This builds confidence at service edges: inputs, status codes, retries, and shutdown behavior.
- When correct: When your solution is correct, it should satisfy: `retry on 5xx + network errors`; `no retry on 4xx`; and `safe body handling`.

## Rules / Expectations
- retry on 5xx + network errors
- no retry on 4xx
- safe body handling

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)

## What this kata is about (and why it matters)
- This kata is focused practice in net/http through `HTTP Client with Retries`.
- You will use this in APIs and services where request/response correctness directly affects reliability.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
