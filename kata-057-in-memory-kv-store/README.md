# Kata 057 — In-Memory KV Store

**Focus:** Maps, RWMutex

## Your task
Implement:

```go
type KV struct
```

## Rules / Expectations
- Set/Get/Delete/Keys
- thread-safe

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing In-Memory KV Store with constraints that make you practice maps, rwmutex.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
