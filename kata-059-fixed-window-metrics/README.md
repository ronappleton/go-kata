# Kata 059 — Fixed-Window Metrics

**Focus:** Time buckets

## Your task
Implement:

```go
type CounterWindow struct
```

### Learning goal
- What you are building: Build `Fixed-Window Metrics` as a reliable contract. Focus: Time buckets.
- Why this matters in real projects: Time logic gets flaky fast. Precision here prevents hard-to-reproduce bugs.
- How this grows your Go skills: You practice boundary testing at just-before/at/just-after moments.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: add(value); sum(last duration); and thread-safe.

### Tips
- Test just-before, at, and just-after boundaries.
- Keep clock math in one place.
- Prefer deterministic tests over sleeps.

## Rules / Expectations
- Add(value)
- Sum(last duration)
- thread-safe

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- Core lesson: model time explicitly and test boundaries directly.
- After this kata, you should be able to test temporal boundaries without flaky sleeps.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
