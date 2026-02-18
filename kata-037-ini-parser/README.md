# Kata 037 — INI Parser

**Focus:** Sections, maps

## Your task
Implement:

```go
func ParseINI(s string) (map[string]map[string]string, error)
```

## Rules / Expectations
- sections [name]
- key=value
- comments start with ; or #

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [INI format overview](https://en.wikipedia.org/wiki/INI_file)
- [Go bufio.Scanner docs](https://pkg.go.dev/bufio#Scanner)

## What this kata is about (and why it matters)
- This kata is about implementing INI Parser with constraints that make you practice sections, maps.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
