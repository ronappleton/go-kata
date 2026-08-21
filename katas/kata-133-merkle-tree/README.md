# Kata 133 — Merkle Tree

**Focus:** data structures, crypto

## Your task
Implement:

```go
type MerkleNode struct { Hash string; Left, Right *MerkleNode }; func BuildMerkleTree(data [][]byte) *MerkleNode; func VerifyMerkleTree(root *MerkleNode, data [][]byte) bool
```

### Learning goal
- What you are building: type MerkleNode struct { Hash string; Left, Right *MerkleNode }; func BuildMerkleTree(data [][]byte) *MerkleNode; func VerifyMerkleTree(root *MerkleNode, data [][]byte) bool as a reliable contract. Focus: data structures, crypto.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): BuildMerkleTree hashes leaves with sha256 and combines parents as sha256(left+right); VerifyMerkleTree rebuilds from data and compares roots.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- sha256 hashing
- deterministic root for same data
- tampered data => verify false
- single-item tree has one leaf root

## Prior reading
- [Merkle tree (Wikipedia)](https://en.wikipedia.org/wiki/Merkle_tree)
- [Go crypto/sha256 package](https://pkg.go.dev/crypto/sha256)

## What this kata is about (and why it matters)
- Core lesson: hash trees give tamper evidence: any change to input changes the root.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
