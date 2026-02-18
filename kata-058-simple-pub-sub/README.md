# Kata 058 — Simple Pub/Sub

**Focus:** Fanout

## Your task
Implement:

```go
type PubSub struct
```

### Learning goal
- What you are practicing: Build `Simple Pub/Sub` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `Subscribe(topic) chan`; `Publish(topic,msg)`; and `Unsubscribe closes chan`.

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
- This kata is focused practice in Fanout through `Simple Pub/Sub`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
