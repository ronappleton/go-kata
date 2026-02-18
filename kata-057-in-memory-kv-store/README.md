# Kata 057 — In-Memory KV Store

**Focus:** Maps, RWMutex

## Your task
Implement:

```go
type KV struct
```

### Learning goal
- Expected work: Implement `In-Memory KV Store` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `In-Memory KV Store` teaches core implementation habits that compound across all later katas.

## Rules / Expectations
- Set/Get/Delete/Keys
- thread-safe

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains maps, rwmutex by implementing `In-Memory KV Store` under explicit constraints.
- It is important because `In-Memory KV Store` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
