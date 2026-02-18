# Kata 091 — Mini cut

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- Expected work: Implement `Mini cut` as a small UNIX-style command with predictable argument handling, streaming I/O, and explicit error returns.
- Why: `Mini cut` teaches command-line contract discipline, where stable output and error behavior are critical for scripts and pipelines.

## Rules / Expectations
- implement subset of unix cut
- no os.Exit
- return errors

## Prior reading
- [POSIX cut utility](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/cut.html)
- [Go strings package](https://pkg.go.dev/strings)
- [Go bufio package](https://pkg.go.dev/bufio)
- [Go io package](https://pkg.go.dev/io)

## What this kata is about (and why it matters)
- This kata trains cli, io by implementing `Mini cut` under explicit constraints.
- It is important because `Mini cut` mirrors real CLI tooling where compatibility and determinism are part of the API.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
