# Kata 132 — Bloom Filter

**Focus:** data structures, hashing

## Your task
Implement:

```go
func NewBloomFilter(size int, hashCount int) (*BloomFilter, error); func (b *BloomFilter) Add(item string); func (b *BloomFilter) Contains(item string) bool
```

### Learning goal
- What you are building: func NewBloomFilter(size int, hashCount int) (*BloomFilter, error); func (b *BloomFilter) Add(item string); func (b *BloomFilter) Contains(item string) bool as a reliable contract. Focus: data structures, hashing.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): Add never causes Contains to go false (no false negatives); Contains may rarely return true for unseen items (false positive) but only within the bloom-filter tradeoff.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- no false negatives
- deterministic hash family
- error on invalid size/hashCount
- Contains true for all added

## Prior reading
- [Bloom filter (Wikipedia)](https://en.wikipedia.org/wiki/Bloom_filter)
- [Go hash/fnv package](https://pkg.go.dev/hash/fnv)

## What this kata is about (and why it matters)
- Core lesson: probabilistic structures trade exactness for space; the invariant that must never break is the false-negative rate.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
