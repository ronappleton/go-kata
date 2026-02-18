# Kata 052 — TCP Echo Server

**Focus:** net

## Your task
Implement:

```go
func StartEchoServer(addr string) (stop func() error, err error)
```

## Rules / Expectations
- multiple clients
- graceful stop
- per-connection goroutine

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
