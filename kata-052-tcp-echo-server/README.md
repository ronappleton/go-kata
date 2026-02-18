# Kata 052 — TCP Echo Server

**Focus:** net

## Your task
Implement:

```go
func StartEchoServer(addr string) (stop func() error, err error)
```

### Learning goal
- Expected work: Implement `TCP Echo Server` as resilient network code that handles partial I/O, disconnects, and cleanup safely.
- Why: `TCP Echo Server` teaches socket-level robustness needed for long-running network services.
- When correct: `TCP Echo Server` should satisfy the required behavior, including: `multiple clients`; `graceful stop`; and `per-connection goroutine`.

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
- This kata trains net by implementing `TCP Echo Server` under explicit constraints.
- It is important because `TCP Echo Server` practices the connection-management behavior required in reliable services.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
