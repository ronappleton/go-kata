# Kata 065 — Concurrent File Downloader

**Focus:** Concurrency, net/http

## Your task
Implement:

```go
func DownloadAll(ctx context.Context, urls []string, workers int) (map[string][]byte, error)
```

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
- This kata is about implementing Concurrent File Downloader with constraints that make you practice concurrency, net/http.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
