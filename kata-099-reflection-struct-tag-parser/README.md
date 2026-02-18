# Kata 099 — Reflection: Struct Tag Parser

**Focus:** reflect

## Your task
Implement:

```go
func ParseTags(v any, tagKey string) (map[string]string, error)
```

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
- This kata is about implementing Reflection: Struct Tag Parser with constraints that make you practice reflect.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
