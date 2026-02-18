# Kata 129 — Reflection: Struct Tag Parser

**Focus:** reflect

## Your task
Implement:

```go
func ParseTags(v any, tagKey string) (map[string]string, error)
```

### Learning goal
- What you are building: Build `func ParseTags(v any, tagKey string) (map[string]string, error)` as a reliable contract. Focus: reflect.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: accept only `struct`; map fieldname->tagvalue; and skip empty tags.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

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
- Core lesson: parse strictly and fail loudly on bad input.
- After this kata, you should be able to state input contracts and return precise parse errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
