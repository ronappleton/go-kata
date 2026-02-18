# Kata 019 — ISBN-10 Validator

**Focus:** Checksums

## Your task
Implement:

```go
func IsValidISBN10(s string) bool
```

### Learning goal
- Expected work: Implement `ISBN-10 Validator` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `ISBN-10 Validator` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- allow hyphens/spaces
- X allowed as last digit meaning 10
- must be 10 digits after stripping

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains checksums by implementing `ISBN-10 Validator` under explicit constraints.
- It is important because `ISBN-10 Validator` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
