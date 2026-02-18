# Kata 091 — Mini cut

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- What you are practicing: Build `Mini cut` like a real command-line tool: clear arguments, stream-friendly I/O, and predictable errors.
- Why it matters: You will use this in internal tooling, CI scripts, and automation where output format and exit behavior are part of the contract.
- How this grows your Go skills: This builds your ability to design user-facing contracts and test behavior from stdin/stdout boundaries.
- When correct: When your solution is correct, it should satisfy: `implement subset of unix cut`; `no os.Exit`; and `return errors`.

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
- This kata is focused practice in CLI, IO through `Mini cut`.
- You will use this in internal tooling, CI scripts, and automation where output format and exit behavior are part of the contract.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
