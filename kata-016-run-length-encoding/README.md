# Kata 016 — Run-Length Encoding

**Focus:** Strings.Builder, runes

## Your task
Implement:

```go
func RLEncode(s string) string
```

### Learning goal
- Expected work: Implement `Run-Length Encoding` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Run-Length Encoding` teaches core implementation habits that compound across all later katas.
- When correct: `Run-Length Encoding` should satisfy the required behavior, including: `encode runs as <char><count> where count>1`; `single chars stay as char only`; and `Unicode-safe`.

## Rules / Expectations
- encode runs as <char><count> where count>1
- single chars stay as char only
- Unicode-safe

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains strings.builder, runes by implementing `Run-Length Encoding` under explicit constraints.
- It is important because `Run-Length Encoding` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
