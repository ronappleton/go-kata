# Kata 075 — Dijkstra Shortest Path

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata75() error
```

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Dijkstra Shortest Path with constraints that make you practice system design, integration boundaries, and robust testing.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
