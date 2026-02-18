# Kata 087 — Tar/Gzip Archiver

**Focus:** archive/tar, compress/gzip

## Your task
Implement:

```go
func ArchiveTGZ(srcDir string, w io.Writer) error
```

### Learning goal
- Expected work: Implement `Tar/Gzip Archiver` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Tar/Gzip Archiver` teaches core implementation habits that compound across all later katas.

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
- This kata trains archive/tar, compress/gzip by implementing `Tar/Gzip Archiver` under explicit constraints.
- It is important because `Tar/Gzip Archiver` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
