# Kata 040 — Password Strength Scorer

**Focus:** Rules

## Your task
Implement:

```go
func PasswordScore(pw string) int
```

### Learning goal
- Expected work: Implement `Password Strength Scorer` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Password Strength Scorer` teaches core implementation habits that compound across all later katas.

## Rules / Expectations
- score 0-100
- length + variety
- simple penalties

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains rules by implementing `Password Strength Scorer` under explicit constraints.
- It is important because `Password Strength Scorer` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
