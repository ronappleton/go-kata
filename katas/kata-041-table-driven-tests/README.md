# Kata 041 — Table-Driven Tests

**Focus:** table-driven tests, subtests, testing.T, test patterns

## Your task
Write table-driven tests for the `Abs` and `Max` functions using subtests.

### Learning goal
- What you are practicing: the standard Go testing pattern — table-driven tests with subtests.
- Why this matters: table-driven tests are THE way to write tests in Go. They're readable, maintainable, and easy to extend.
- How this grows your Go skills: you learn `t.Run()` for subtests, test organization, and the Go testing idioms used in every production codebase.

### Tips
- Define a slice of test cases with input and expected output.
- Use `t.Run("name", func(t *testing.T) {...})` for each case.
- Include edge cases: zero, negative, equal values.

## Rules / Expectations
- Abs of positive => same number
- Abs of negative => positive
- Abs of zero => zero
- Max of a>b => a
- Max of b>a => b
- Max of a==b => a

## What this kata is about (and why it matters)
- Core lesson: table-driven tests are the Go testing standard.
- After this kata, every test you write should follow this pattern.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
