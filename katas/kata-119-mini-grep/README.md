# Kata 119 — Mini grep

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

### Learning goal
- What you are building: Build `func Run(args []string, r io.Reader, w io.Writer) error` as a reliable contract. Focus: CLI, IO.
- Why this matters in real projects: CLI tools become infrastructure. Scripts depend on stable output and errors.
- How this grows your Go skills: You practice `io.Reader`/`io.Writer`-first design for automation-safe commands.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: match the required subset of Unix `grep` behavior; never call `os.Exit`; and return errors instead of swallowing failures.

### Tips
- Test exact stdout/stderr text.
- Keep flag parsing separate from core logic.
- Cover bad args and bad input paths.

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
- Core lesson: CLI output and errors are part of your API.
- After this kata, you should be able to treat stdout/stderr/error behavior as an explicit contract.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
