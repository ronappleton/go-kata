# Kata 058 — Simple Pub/Sub

**Focus:** Fanout

## Your task
Implement:

```go
type PubSub struct
```

### Learning goal
- Expected work: Implement `Simple Pub/Sub` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Simple Pub/Sub` teaches safe coordination patterns that prevent costly production race conditions.

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
- This kata trains fanout by implementing `Simple Pub/Sub` under explicit constraints.
- It is important because `Simple Pub/Sub` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
