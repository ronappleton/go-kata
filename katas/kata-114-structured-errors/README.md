# Kata 114 — Structured Errors

**Focus:** Error handling, idioms

## Your task
Implement:

```go
var ErrNotFound, ErrUnauthorized error; type AppError; func IsNotFound(err error) bool; func IsUnauthorized(err error) bool
```

### Learning goal
- What you are building: var ErrNotFound, ErrUnauthorized error; type AppError; func IsNotFound(err error) bool; func IsUnauthorized(err error) bool as a reliable contract. Focus: Error handling, idioms.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): AppError carries Code/Message/cause; errors.Is/As work; sentinels ErrNotFound/ErrUnauthorized; Is* helpers unwrap chains.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- AppError type with Code int, Message string, cause
- sentinels ErrNotFound and ErrUnauthorized
- IsNotFound/IsUnauthorized helpers
- errors.Is and errors.As interoperate

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go errors package](https://pkg.go.dev/errors)
- [Go error handling overview](https://go.dev/blog/go1.13-errors)

## What this kata is about (and why it matters)
- Core lesson: treat errors as data: structured, comparable via errors.Is, and never string-matched.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
