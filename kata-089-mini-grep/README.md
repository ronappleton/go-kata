# Kata 089 — Mini grep

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

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
- This kata is about implementing Mini grep with constraints that make you practice cli, io.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
