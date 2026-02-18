# Kata 060 — Trie Autocomplete

**Focus:** Data structures

## Your task
Implement:

```go
func NewTrie() *Trie
```

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
- This kata is about implementing Trie Autocomplete with constraints that make you practice data structures.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
