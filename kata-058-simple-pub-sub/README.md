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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
