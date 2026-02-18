# Kata 037 — INI Parser

**Focus:** Sections, maps

## Your task
Implement:

```go
func ParseINI(s string) (map[string]map[string]string, error)
```

### Learning goal
- Expected work: Implement `INI Parser` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `INI Parser` teaches strict boundary handling so malformed input cannot silently corrupt results.

## Rules / Expectations
- sections [name]
- key=value
- comments start with ; or #

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [INI format overview](https://en.wikipedia.org/wiki/INI_file)
- [Go bufio.Scanner docs](https://pkg.go.dev/bufio#Scanner)

## What this kata is about (and why it matters)
- This kata trains sections, maps by implementing `INI Parser` under explicit constraints.
- It is important because `INI Parser` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
