# Kata 099 — Reflection: Struct Tag Parser

**Focus:** reflect

## Your task
Implement:

```go
func ParseTags(v any, tagKey string) (map[string]string, error)
```

### Learning goal
- Expected work: Implement `Reflection: Struct Tag Parser` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `Reflection: Struct Tag Parser` teaches strict boundary handling so malformed input cannot silently corrupt results.
- When correct: `Reflection: Struct Tag Parser` should satisfy the required behavior, including: `struct only`; `map fieldName->tagValue`; and `skip empty tags`.

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
- This kata trains reflect by implementing `Reflection: Struct Tag Parser` under explicit constraints.
- It is important because `Reflection: Struct Tag Parser` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
