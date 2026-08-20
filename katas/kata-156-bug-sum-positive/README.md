# Kata 132 — Sum Positive Bug

**Focus:** loops, integer filtering

## Your task
Fix the existing implementation in `kata.go` so all tests pass.

Implement/fix:

```go
func SumPositive(nums []int) int
```

### Learning goal
- What you are building: bug-fix workflow for simple aggregate logic.
- Why this matters in real projects: incorrect aggregation logic quietly corrupts downstream behavior.
- How this grows your Go skills: you practice isolating behavior defects via focused tests.
- Definition of done (plain English): all tests pass and behavior matches the Rules / Expectations contract.

## Bug Hunt
- Reproduce the failing test first.
- Fix one defect at a time.
- Rerun tests after each small patch.
- Keep function signature and tests unchanged.

## Rules / Expectations
- include only values strictly greater than zero
- ignore zero and negative numbers
- empty input => 0

## Prior reading
- [Go specification: For statements](https://go.dev/ref/spec#For_statements)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: small logic bugs can have large business impact.
- After this kata, you should be able to describe exactly which values are included and excluded.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
