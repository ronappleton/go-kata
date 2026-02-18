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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
