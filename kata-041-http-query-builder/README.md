# Kata 041 — HTTP Query Builder

**Focus:** net/url

## Your task
Implement:

```go
func BuildQuery(params map[string][]string) string
```

### Learning goal
- What you are practicing: Build `HTTP Query Builder` as production-style HTTP boundary code with clear request handling and response behavior.
- Why it matters: You will use this in APIs and services where request/response correctness directly affects reliability.
- How this grows your Go skills: This builds confidence at service edges: inputs, status codes, retries, and shutdown behavior.
- When correct: When your solution is correct, it should satisfy: `keys sorted`; `values preserve given order`; and `proper escaping`.

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
- This kata is focused practice in net/url through `HTTP Query Builder`.
- You will use this in APIs and services where request/response correctness directly affects reliability.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
