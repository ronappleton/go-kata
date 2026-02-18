# Kata 037 — INI Parser

**Focus:** Sections, maps

## Your task
Implement:

```go
func ParseINI(s string) (map[string]map[string]string, error)
```

### Learning goal
- What you are practicing: Build `INI Parser` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `sections [name]`; `key=value`; and `comments start with ; or #`.

## Rules / Expectations
- sections [name]
- key=value
- comments start with ; or #

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [INI format overview](https://en.wikipedia.org/wiki/INI_file)
- [Go bufio.Scanner docs](https://pkg.go.dev/bufio#Scanner)

## What this kata is about (and why it matters)
- This kata is focused practice in Sections, maps through `INI Parser`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
