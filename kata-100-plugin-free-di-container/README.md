# Kata 100 — Plugin-free DI Container

**Focus:** Design

## Your task
Implement:

```go
func NewContainer() *Container
```

### Learning goal
- Expected work: Implement `Plugin-free DI Container` as a composable abstraction with explicit contracts and testable behavior.
- Why: `Plugin-free DI Container` teaches maintainable design through explicit contracts and low coupling.

## Rules / Expectations
- Register(name, ctor)
- Resolve(name)
- singleton scope

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains design by implementing `Plugin-free DI Container` under explicit constraints.
- It is important because `Plugin-free DI Container` develops design choices that improve extensibility and testability over time.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
