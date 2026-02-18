# Kata 100 — Plugin-free DI Container

**Focus:** Design

## Your task
Implement:

```go
func NewContainer() *Container
```

### Learning goal
- What you are building: Build `func NewContainer() *Container` as a reliable contract. Focus: Design.
- Why this matters in real projects: Good structure compounds. Clear contracts reduce future rework.
- How this grows your Go skills: You practice interface-first design and explicit dependencies.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: register(name, ctor); resolve(name); and singleton scope.

### Tips
- Write contract tests before implementation details.
- Keep dependencies explicit, not implicit.
- Prefer small interfaces with one reason to change.

## Rules / Expectations
- Register(name, ctor)
- Resolve(name)
- singleton scope

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- Core lesson: design seams for testability and change.
- After this kata, you should be able to describe component contracts and test at those seams.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
