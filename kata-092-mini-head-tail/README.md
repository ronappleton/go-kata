# Kata 092 — Mini head/tail

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- Expected work: Implement `Mini head/tail` as a small UNIX-style command with predictable argument handling, streaming I/O, and explicit error returns.
- Why: `Mini head/tail` teaches command-line contract discipline, where stable output and error behavior are critical for scripts and pipelines.

## Rules / Expectations
- implement subset of unix head/tail
- no os.Exit
- return errors

## Prior reading
- [POSIX head utility](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/head.html)
- [POSIX tail utility](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/tail.html)
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)

## What this kata is about (and why it matters)
- This kata trains cli, io by implementing `Mini head/tail` under explicit constraints.
- It is important because `Mini head/tail` mirrors real CLI tooling where compatibility and determinism are part of the API.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
