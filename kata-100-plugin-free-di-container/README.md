# Kata 100 — Plugin-free DI Container

**Focus:** Design

## Your task
Implement:

```go
func NewContainer() *Container
```

## Rules / Expectations
- Register(name, ctor)
- Resolve(name)
- singleton scope

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Plugin-free DI Container with constraints that make you practice design.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
