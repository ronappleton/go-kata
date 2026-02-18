# Kata 029 — Longest Common Prefix

**Focus:** Strings

## Your task
Implement:

```go
func LongestCommonPrefix(items []string) string
```

### Learning goal
- Expected work: Implement `Longest Common Prefix` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Longest Common Prefix` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- empty => empty
- single => itself
- case-sensitive

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains strings by implementing `Longest Common Prefix` under explicit constraints.
- It is important because `Longest Common Prefix` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
