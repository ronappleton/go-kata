# Kata 065 — Concurrent File Downloader

**Focus:** Concurrency, net/http

## Your task
Implement:

```go
func DownloadAll(ctx context.Context, urls []string, workers int) (map[string][]byte, error)
```

### Learning goal
- What you are building: Build `func DownloadAll(ctx context.Context, urls []string, workers int) (map[string][]byte, error)` as a reliable contract. Focus: Concurrency, net/http.
- Why this matters in real projects: HTTP boundaries are product behavior. Callers depend on exact semantics.
- How this grows your Go skills: You practice context-aware request handling and stable status/error contracts.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: limit concurrency; cancel on ctx; and return data per url.

### Tips
- Pin boundary behavior with `httptest`.
- Cover both success and failure responses.
- Test retries/timeouts with targeted cases.

## Rules / Expectations
- limit concurrency
- cancel on ctx
- return data per url

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [RFC 9110 (HTTP semantics)](https://www.rfc-editor.org/rfc/rfc9110)
- [Go net package](https://pkg.go.dev/net)
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)

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
