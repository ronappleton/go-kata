# Kata 131 — Normalize Username Bug

**Focus:** strings.TrimSpace, strings.ToLower, strings.ReplaceAll

## Your task
Fix the existing implementation in `kata.go` so all tests pass.

Implement/fix:

```go
func NormalizeUsername(name string) string
```

### Learning goal
- What you are building: debugging discipline with test-guided bug isolation.
- Why this matters in real projects: most production work is fixing existing behavior without introducing regressions.
- How this grows your Go skills: you practice reading failures, localizing root cause, and applying minimal safe patches.
- Definition of done (plain English): all tests pass and behavior matches the Rules / Expectations contract.

## Bug Hunt
- Reproduce the failing test first.
- Fix one defect at a time.
- Rerun tests after each small patch.
- Keep function signature and tests unchanged.

## Rules / Expectations
- trim leading/trailing spaces
- lowercase all letters
- replace internal spaces with underscore (`_`)
- empty/space-only input => empty string

## Prior reading
- [strings package](https://pkg.go.dev/strings)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: robust debugging is a core engineering skill, not a side activity.
- After this kata, you should be able to explain the defect cause and why your patch is minimal and safe.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
