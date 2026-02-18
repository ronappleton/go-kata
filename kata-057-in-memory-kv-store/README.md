# Kata 057 — In-Memory KV Store

**Focus:** Maps, RWMutex

## Your task
Implement:

```go
type KV struct
```

### Learning goal
- What you are practicing: Build `In-Memory KV Store` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `Set/Get/Delete/Keys`; and `thread-safe`.

## Rules / Expectations
- Set/Get/Delete/Keys
- thread-safe

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Maps, RWMutex through `In-Memory KV Store`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
