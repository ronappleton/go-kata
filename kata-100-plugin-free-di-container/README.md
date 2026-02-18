# Kata 100 — Plugin-free DI Container

**Focus:** Design

## Your task
Implement:

```go
func NewContainer() *Container
```

### Learning goal
- What you are practicing: Build `Plugin-free DI Container` as a reusable, testable abstraction with explicit contracts.
- Why it matters: You will use this when designing packages that need to stay maintainable as features grow.
- How this grows your Go skills: This builds API design judgment and composable package structure in Go.
- When correct: When your solution is correct, it should satisfy: `Register(name, ctor)`; `Resolve(name)`; and `singleton scope`.

## Rules / Expectations
- Register(name, ctor)
- Resolve(name)
- singleton scope

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Design through `Plugin-free DI Container`.
- You will use this when designing packages that need to stay maintainable as features grow.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
