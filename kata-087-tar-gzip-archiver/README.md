# Kata 087 — Tar/Gzip Archiver

**Focus:** archive/tar, compress/gzip

## Your task
Implement:

```go
func ArchiveTGZ(srcDir string, w io.Writer) error
```

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
- This kata is about implementing Tar/Gzip Archiver with constraints that make you practice archive/tar, compress/gzip.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
