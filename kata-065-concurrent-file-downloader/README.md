# Kata 065 — Concurrent File Downloader

**Focus:** Concurrency, net/http

## Your task
Implement:

```go
func DownloadAll(ctx context.Context, urls []string, workers int) (map[string][]byte, error)
```

### Learning goal
- What you are practicing: Build `Concurrent File Downloader` as production-style HTTP boundary code with clear request handling and response behavior.
- Why it matters: You will use this in APIs and services where request/response correctness directly affects reliability.
- How this grows your Go skills: This builds confidence at service edges: inputs, status codes, retries, and shutdown behavior.
- When correct: When your solution is correct, it should satisfy: `limit concurrency`; `cancel on ctx`; and `return data per url`.

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
- This kata is focused practice in Concurrency, net/http through `Concurrent File Downloader`.
- You will use this in APIs and services where request/response correctness directly affects reliability.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
