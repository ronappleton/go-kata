# Kata 087 — Tar/Gzip Archiver

**Focus:** archive/tar, compress/gzip

## Your task
Implement:

```go
func ArchiveTGZ(srcDir string, w io.Writer) error
```

### Learning goal
- What you are practicing: Build `Tar/Gzip Archiver` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `recursive archive`; `preserve relative paths`; and `skip symlinks`.

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
- This kata is focused practice in archive/tar, compress/gzip through `Tar/Gzip Archiver`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
