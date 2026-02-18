# Kata 094 — Mini uniq

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

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
- This kata is about implementing Mini uniq with constraints that make you practice cli, io.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
