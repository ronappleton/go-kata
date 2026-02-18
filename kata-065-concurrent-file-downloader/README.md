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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
