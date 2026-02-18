# Kata 052 — TCP Echo Server

**Focus:** net

## Your task
Implement:

```go
func StartEchoServer(addr string) (stop func() error, err error)
```

### Learning goal
- What you are building: Build `func StartEchoServer(addr string) (stop func() error, err error)` as a reliable contract. Focus: net.
- Why this matters in real projects: Network paths fail in partial states. Robust lifecycle handling matters.
- How this grows your Go skills: You practice stream I/O discipline and connection lifecycle control.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: multiple clients; graceful stop; and per-connection goroutine.

### Tips
- Treat EOF and short reads as normal states.
- Keep connection lifecycle explicit.
- Prove cleanup on disconnect in tests.

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
- Core lesson: handle partial I/O and disconnects as normal paths.
- After this kata, you should be able to handle disconnect cleanup without hidden leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
