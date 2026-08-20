# Kata 023 — Is Within Business Hours

**Focus:** time.Time, time.Location, weekday/hour checks

## Your task
Implement:

```go
func IsWithinBusinessHours(ts time.Time, loc *time.Location) bool
```

### Learning goal
- What you are building: timezone-aware business-hours validation.
- Why this matters in real projects: scheduling, notifications, and rate windows break if time-zone handling is vague.
- How this grows your Go skills: you practice explicit time conversion and temporal boundary checks.
- Definition of done (plain English): function results are consistent regardless of caller timezone.

### Tips
- Convert once, then perform weekday/hour checks.
- Test both weekday and weekend paths.
- Include boundary-hour tests (09:00, 16:59, 17:00).

## Rules / Expectations
- if loc is nil => use ts location
- convert ts into chosen location before checks
- true only on Monday-Friday
- true only when hour is >= 9 and < 17
- false for all other times

## Prior reading
- [time package](https://pkg.go.dev/time)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: time logic should be explicit about location and boundaries.
- After this kata, you should be able to explain and test time behavior without relying on local machine assumptions.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
