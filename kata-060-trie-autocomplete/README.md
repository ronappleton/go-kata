# Kata 060 — Trie Autocomplete

**Focus:** Data structures

## Your task
Implement:

```go
func NewTrie() *Trie
```

### Learning goal
- What you are practicing: Build `Trie Autocomplete` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `Insert(word)`; `Suggest(prefix, limit)`; and `lexicographic order`.

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
- This kata is focused practice in Data structures through `Trie Autocomplete`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
