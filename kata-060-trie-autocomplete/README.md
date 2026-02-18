# Kata 060 — Trie Autocomplete

**Focus:** Data structures

## Your task
Implement:

```go
func NewTrie() *Trie
```

### Learning goal
- Expected work: Implement `Trie Autocomplete` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Trie Autocomplete` teaches complexity-aware correctness that impacts throughput and latency at scale.
- When correct: `Trie Autocomplete` should satisfy the required behavior, including: `Insert(word)`; `Suggest(prefix, limit)`; and `lexicographic order`.

## Rules / Expectations
- Insert(word)
- Suggest(prefix, limit)
- lexicographic order

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains data structures by implementing `Trie Autocomplete` under explicit constraints.
- It is important because `Trie Autocomplete` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
