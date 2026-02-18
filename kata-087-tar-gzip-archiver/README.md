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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
