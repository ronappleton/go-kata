# Kata 077 — Event Sourcing Mini

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata77() error
```

### Learning goal
- What you are building: Build `func Kata77() error` as a reliable contract. Focus: real-world Go integration and testing.
- Why this matters in real projects: Good structure compounds. Clear contracts reduce future rework.
- How this grows your Go skills: You practice interface-first design and explicit dependencies.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: match the full README contract; include tests for happy path and edge cases; and keep API and error style idiomatic Go.

### Tips
- Write contract tests before implementation details.
- Keep dependencies explicit, not implicit.
- Prefer small interfaces with one reason to change.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Event Sourcing pattern](https://martinfowler.com/eaaDev/EventSourcing.html)
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
