# Kata 058 — Simple Pub/Sub

**Focus:** Fanout

## Your task
Implement:

```go
type PubSub struct
```

## Rules / Expectations
- Subscribe(topic) chan
- Publish(topic,msg)
- Unsubscribe closes chan

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Simple Pub/Sub with constraints that make you practice fanout.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
