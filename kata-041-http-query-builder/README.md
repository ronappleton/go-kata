# Kata 041 — HTTP Query Builder

**Focus:** net/url

## Your task
Implement:

```go
func BuildQuery(params map[string][]string) string
```

### Learning goal
- What you are building: Build `func BuildQuery(params map[string][]string) string` as a reliable contract. Focus: net/url.
- Why this matters in real projects: HTTP boundaries are product behavior. Callers depend on exact semantics.
- How this grows your Go skills: You practice context-aware request handling and stable status/error contracts.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: keys sorted; values preserve given order; and proper escaping.

### Tips
- Pin boundary behavior with `httptest`.
- Cover both success and failure responses.
- Test retries/timeouts with targeted cases.

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
- Core lesson: treat request/response behavior as a hard contract.
- After this kata, you should be able to justify status/error choices at the service boundary.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
