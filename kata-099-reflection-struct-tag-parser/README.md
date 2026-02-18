# Kata 099 — Reflection: Struct Tag Parser

**Focus:** reflect

## Your task
Implement:

```go
func ParseTags(v any, tagKey string) (map[string]string, error)
```

### Learning goal
- What you are practicing: Build `Reflection: Struct Tag Parser` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `struct only`; `map fieldName->tagValue`; and `skip empty tags`.

## Rules / Expectations
- struct only
- map fieldName->tagValue
- skip empty tags

## Prior reading
- [Go reflect package](https://pkg.go.dev/reflect)
- [Go struct tags in language spec](https://go.dev/ref/spec#Struct_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in reflect through `Reflection: Struct Tag Parser`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
