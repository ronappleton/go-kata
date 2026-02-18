# Kata 025 — LRU Cache (single-thread)

**Focus:** Structs, maps, list

## Your task
Implement:

```go
func NewLRU(capacity int) (*LRU, error)
```

## Rules / Expectations
- capacity>0
- Get/Put methods
- evict least recently used

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing LRU Cache (single-thread) with constraints that make you practice structs, maps, list.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
