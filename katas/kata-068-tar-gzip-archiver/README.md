# Kata 117 — Tar/Gzip Archiver

**Focus:** archive/tar, compress/gzip

## Your task
Implement:

```go
func ArchiveTGZ(srcDir string, w io.Writer) error
```

### Learning goal
- What you are building: Build `func ArchiveTGZ(srcDir string, w io.Writer) error` as a reliable contract. Focus: archive/tar, compress/gzip.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: recursive archive; preserve relative paths; and skip symlinks.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- recursive archive
- preserve relative paths
- skip symlinks

## Prior reading
- [Go os package](https://pkg.go.dev/os)
- [Go path/filepath package](https://pkg.go.dev/path/filepath)
- [Go archive/tar package](https://pkg.go.dev/archive/tar)
- [Go compress/gzip package](https://pkg.go.dev/compress/gzip)

## What this kata is about (and why it matters)
- Core lesson: turn plain rules into deterministic Go behavior.
- After this kata, you should be able to write rule-first tests and explain each edge case clearly.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
