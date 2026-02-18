# Kata 089 — Mini grep

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- Expected work: Implement `Mini grep` as a small UNIX-style command with predictable argument handling, streaming I/O, and explicit error returns.
- Why: `Mini grep` teaches command-line contract discipline, where stable output and error behavior are critical for scripts and pipelines.
- When correct: `Mini grep` should satisfy the required behavior, including: `implement subset of unix grep`; `no os.Exit`; and `return errors`.

## Rules / Expectations
- implement subset of unix grep
- no os.Exit
- return errors

## Prior reading
- [POSIX grep utility](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/grep.html)
- [Go regexp package](https://pkg.go.dev/regexp)
- [Go bufio package](https://pkg.go.dev/bufio)
- [Go io package](https://pkg.go.dev/io)

## What this kata is about (and why it matters)
- This kata trains cli, io by implementing `Mini grep` under explicit constraints.
- It is important because `Mini grep` mirrors real CLI tooling where compatibility and determinism are part of the API.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
