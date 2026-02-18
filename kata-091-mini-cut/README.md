# Kata 091 — Mini cut

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

## Rules / Expectations
- implement subset of unix cut
- no os.Exit
- return errors

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
