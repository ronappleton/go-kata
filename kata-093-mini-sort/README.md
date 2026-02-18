# Kata 093 — Mini sort

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- What you are practicing: Build `Mini sort` like a real command-line tool: clear arguments, stream-friendly I/O, and predictable errors.
- Why it matters: You will use this in internal tooling, CI scripts, and automation where output format and exit behavior are part of the contract.
- How this grows your Go skills: This builds your ability to design user-facing contracts and test behavior from stdin/stdout boundaries.
- When correct: When your solution is correct, it should satisfy: `implement subset of unix sort`; `no os.Exit`; and `return errors`.

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
- This kata is focused practice in CLI, IO through `Mini sort`.
- You will use this in internal tooling, CI scripts, and automation where output format and exit behavior are part of the contract.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
