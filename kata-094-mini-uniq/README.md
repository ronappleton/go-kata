# Kata 094 — Mini uniq

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- Expected work: Implement `Mini uniq` as a small UNIX-style command with predictable argument handling, streaming I/O, and explicit error returns.
- Why: `Mini uniq` teaches command-line contract discipline, where stable output and error behavior are critical for scripts and pipelines.

## Rules / Expectations
- implement subset of unix uniq
- no os.Exit
- return errors

## Prior reading
- [POSIX uniq utility](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/uniq.html)
- [Go bufio package](https://pkg.go.dev/bufio)
- [Go strings package](https://pkg.go.dev/strings)
- [Go io package](https://pkg.go.dev/io)

## What this kata is about (and why it matters)
- This kata trains cli, io by implementing `Mini uniq` under explicit constraints.
- It is important because `Mini uniq` mirrors real CLI tooling where compatibility and determinism are part of the API.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
