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

## Prior reading
- [Go net package](https://pkg.go.dev/net)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing TCP Echo Server with constraints that make you practice net.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
