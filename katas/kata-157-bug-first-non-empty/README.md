# Kata 133 — First Non-Empty Bug

**Focus:** strings.TrimSpace, loop selection

## Your task
Fix the existing implementation in `kata.go` so all tests pass.

Implement/fix:

```go
func FirstNonEmpty(values []string, fallback string) string
```

### Learning goal
- What you are building: robust selection logic for text data.
- Why this matters in real projects: fallback behavior is often relied on by API and config code.
- How this grows your Go skills: you practice bug diagnosis around edge-case inputs.
- Definition of done (plain English): all tests pass and behavior matches the Rules / Expectations contract.

## Bug Hunt
- Reproduce the failing test first.
- Fix one defect at a time.
- Rerun tests after each small patch.
- Keep function signature and tests unchanged.

## Rules / Expectations
- return first value that is non-empty after trim
- if none found => fallback
- preserve selected non-empty value

## Prior reading
- [strings package](https://pkg.go.dev/strings)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: handling blank/empty input correctly avoids surprising defaults.
- After this kata, you should be able to reason about fallback behavior with confidence.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
