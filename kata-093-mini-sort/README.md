# Kata 093 — Mini sort

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- Expected work: Implement `Mini sort` as a small UNIX-style command with predictable argument handling, streaming I/O, and explicit error returns.
- Why: `Mini sort` teaches command-line contract discipline, where stable output and error behavior are critical for scripts and pipelines.
- When correct: `Mini sort` should satisfy the required behavior, including: `implement subset of unix sort`; `no os.Exit`; and `return errors`.

## Rules / Expectations
- implement subset of unix sort
- no os.Exit
- return errors

## Prior reading
- [POSIX sort utility](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sort.html)
- [Go sort package](https://pkg.go.dev/sort)
- [Go bufio package](https://pkg.go.dev/bufio)
- [Go io package](https://pkg.go.dev/io)

## What this kata is about (and why it matters)
- This kata trains cli, io by implementing `Mini sort` under explicit constraints.
- It is important because `Mini sort` mirrors real CLI tooling where compatibility and determinism are part of the API.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
