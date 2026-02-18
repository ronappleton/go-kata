# Kata 067 — INI Parser

**Focus:** Sections, maps

## Your task
Implement:

```go
func ParseINI(s string) (map[string]map[string]string, error)
```

### Learning goal
- What you are building: Build `func ParseINI(s string) (map[string]map[string]string, error)` as a reliable contract. Focus: Sections, maps.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: sections [name]; key=value; and comments start with ; or #.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

## Rules / Expectations
- sections [name]
- key=value
- comments start with ; or #

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [INI format overview](https://en.wikipedia.org/wiki/INI_file)
- [Go bufio.Scanner docs](https://pkg.go.dev/bufio#Scanner)

## What this kata is about (and why it matters)
- Core lesson: parse strictly and fail loudly on bad input.
- After this kata, you should be able to state input contracts and return precise parse errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
