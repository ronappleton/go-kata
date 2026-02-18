# Kata 026 — Stack

**Focus:** Data structures, generics

## Your task
Implement:

```go
type Stack[T any] struct
```

### Learning goal
- What you are practicing: Build `Stack` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `Push, Pop, Peek, Len`; and `Pop/Peek on empty => ok false`.

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
- This kata is focused practice in Data structures, generics through `Stack`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
