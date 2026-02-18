# Kata 026 — Stack

**Focus:** Data structures, generics

## Your task
Implement:

```go
type Stack[T any] struct
```

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
- This kata is about implementing Stack with constraints that make you practice data structures, generics.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
