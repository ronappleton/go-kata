# Kata 052 — TCP Echo Server

**Focus:** net

## Your task
Implement:

```go
func StartEchoServer(addr string) (stop func() error, err error)
```

### Learning goal
- What you are practicing: Build `TCP Echo Server` as robust socket-level code that handles disconnects and cleanup cleanly.
- Why it matters: You will use this in daemons and services that keep long-running connections healthy.
- How this grows your Go skills: This builds operational thinking: lifecycle management, resource cleanup, and fault tolerance.
- When correct: When your solution is correct, it should satisfy: `multiple clients`; `graceful stop`; and `per-connection goroutine`.

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
- This kata is focused practice in net through `TCP Echo Server`.
- You will use this in daemons and services that keep long-running connections healthy.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
