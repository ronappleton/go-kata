# Kata 020 — Roman Numerals

**Focus:** Greedy mapping

## Your task
Implement:

```go
func ToRoman(n int) (string, error)
```

### Learning goal
- Expected work: Implement `Roman Numerals` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Roman Numerals` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- 1..3999 only
- invalid => error
- use standard subtractive notation

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains greedy mapping by implementing `Roman Numerals` under explicit constraints.
- It is important because `Roman Numerals` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
