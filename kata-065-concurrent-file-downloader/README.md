# Kata 065 — Concurrent File Downloader

**Focus:** Concurrency, net/http

## Your task
Implement:

```go
func DownloadAll(ctx context.Context, urls []string, workers int) (map[string][]byte, error)
```

### Learning goal
- Expected work: Implement `Concurrent File Downloader` as boundary-focused HTTP logic with explicit parsing, status handling, and deterministic responses.
- Why: `Concurrent File Downloader` teaches service-edge correctness, where request/response semantics directly drive reliability.

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
- This kata trains concurrency, net/http by implementing `Concurrent File Downloader` under explicit constraints.
- It is important because `Concurrent File Downloader` maps directly to production HTTP boundaries and failure handling.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
