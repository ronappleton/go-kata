# Kata 135 — Clamp Percentage Bug

**Focus:** boundary logic

## Your task
Fix the existing implementation in `kata.go` so all tests pass.

Implement/fix:

```go
func ClampPercentage(v int) int
```

### Learning goal
- What you are building: bug-fix process for boundary validation logic.
- Why this matters in real projects: off-by-boundary defects propagate invalid state quickly.
- How this grows your Go skills: you practice making constraints explicit and testable.
- Definition of done (plain English): all tests pass and behavior matches the Rules / Expectations contract.

## Bug Hunt
- Reproduce the failing test first.
- Fix one defect at a time.
- Rerun tests after each small patch.
- Keep function signature and tests unchanged.

## Rules / Expectations
- v < 0 => 0
- v > 100 => 100
- otherwise => v

## Prior reading
- [Go: if statements](https://go.dev/tour/flowcontrol/5)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: boundary handling must be intentionally enforced.
- After this kata, you should be able to identify missing boundary guards quickly.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
