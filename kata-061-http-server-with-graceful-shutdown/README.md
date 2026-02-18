# Kata 061 — HTTP Server with Graceful Shutdown

**Focus:** net/http, context

## Your task
Implement:

```go
func ServeGraceful(addr string, h http.Handler) (stop func(context.Context) error, err error)
```

## Rules / Expectations
- start server
- stop shuts down with context
- no leaked goroutines

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
