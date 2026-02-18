# Kata 090 — Mini wc

**Focus:** CLI, IO

## Your task
Implement:

```go
func Run(args []string, r io.Reader, w io.Writer) error
```

## Rules / Expectations
- implement subset of unix wc
- no os.Exit
- return errors

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
