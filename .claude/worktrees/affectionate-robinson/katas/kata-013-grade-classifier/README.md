# Kata 013 — Grade Classifier

**Focus:** Condition ranges, input validation

## Your task
Implement:

```go
func GradeLabel(score int) string
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
- score<0 or score>100 => "invalid"
- 90-100 => "A"
- 80-89 => "B"
- 70-79 => "C"
- 60-69 => "D"
- 0-59 => "F"

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
