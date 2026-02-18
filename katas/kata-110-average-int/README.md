# Kata 110 — Average (Integer)

**Focus:** Slices, integer arithmetic, validation

## Your task
Implement:

```go
func Average(values []int) (int, bool)
```

### Learning goal
- What you are practicing: translate simple requirements into correct, testable function behavior.
- Why this matters: this is the foundation for every larger Go service, API, and data pipeline you will build later.
- How this grows your Go skills: you build confidence with function contracts, edge cases, and predictable outputs.
- When correct: your function satisfies every rule in the Rules / Expectations section below.

### Tips
- Start from one rule, write one test, then implement only enough code to pass it.
- Keep input handling explicit; do not hide behavior behind clever shortcuts.
- Treat nil/empty and boundary values as first-class cases.

## Rules / Expectations
- nil or empty slice => 0 and false
- non-empty slice => integer average using floor division and true
- all values contribute once to total

## Prior reading
- [A Tour of Go](https://go.dev/tour/list)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: clear thinking first, code second.
- You are training the habit of proving behavior with tests before moving on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
