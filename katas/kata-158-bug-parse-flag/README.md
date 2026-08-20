# Kata 134 — Parse Flag Bug

**Focus:** strings.ToLower, explicit parsing, error handling

## Your task
Fix the existing implementation in `kata.go` so all tests pass.

Implement/fix:

```go
func ParseFlag(input string) (bool, error)
```

### Learning goal
- What you are building: strict bool parsing with clear invalid-input handling.
- Why this matters in real projects: permissive parsing can cause silent feature misconfiguration.
- How this grows your Go skills: you practice writing explicit parse rules with predictable errors.
- Definition of done (plain English): all tests pass and behavior matches the Rules / Expectations contract.

## Bug Hunt
- Reproduce the failing test first.
- Fix one defect at a time.
- Rerun tests after each small patch.
- Keep function signature and tests unchanged.

## Rules / Expectations
- accept true values: `true`, `1`, `yes` (case-insensitive, trim input)
- accept false values: `false`, `0`, `no` (case-insensitive, trim input)
- all other values => non-nil error

## Prior reading
- [strings package](https://pkg.go.dev/strings)
- [errors package](https://pkg.go.dev/errors)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: parse boundaries should reject ambiguous input.
- After this kata, you should be able to design input contracts that are easy to reason about.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
