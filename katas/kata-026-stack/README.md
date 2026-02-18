# Kata 026 — Stack

**Focus:** Data structures, generics

## Your task
Implement:

```go
type Stack[T any] struct
```

### Learning goal
- What you are building: Build `Stack` as a reliable contract. Focus: Data structures, generics.
- Why this matters in real projects: This is how you keep features fast when input size grows.
- How this grows your Go skills: You practice invariants and complexity reasoning, then prove both with tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: push, pop, peek, len; and when `Pop/Peek on empty`, return `ok false`.

### Tips
- State the invariant first, then code.
- Test shape edges early: empty, one item, duplicates.
- Check complexity after correctness.

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
- Core lesson: hold invariants first, then optimize safely.
- After this kata, you should be able to defend algorithm choice, complexity, and corner-case behavior.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
