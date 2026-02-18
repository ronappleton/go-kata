# Kata 026 — Stack

**Focus:** Data structures, generics

## Your task
Implement:

```go
type Stack[T any] struct
```

### Learning goal
- Expected work: Implement `Stack` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Stack` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- Push, Pop, Peek, Len
- Pop/Peek on empty => ok false

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains data structures, generics by implementing `Stack` under explicit constraints.
- It is important because `Stack` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
