# Kata 008 — Filter Even

**Focus:** Slices, order, immutability

## Your task
Implement:

```go
func FilterEven(nums []int) []int
```

### Learning goal
- Expected work: Implement `Filter Even` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Filter Even` teaches core implementation habits that compound across all later katas.

## Rules / Expectations
- preserve order
- nil => empty slice
- do not mutate input

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains slices, order, immutability by implementing `Filter Even` under explicit constraints.
- It is important because `Filter Even` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
