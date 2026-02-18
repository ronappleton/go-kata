# Kata 031 — File Extension Counter

**Focus:** os, filepath

## Your task
Implement:

```go
func CountExt(root string) (map[string]int, error)
```

### Learning goal
- Expected work: Implement `File Extension Counter` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `File Extension Counter` teaches core implementation habits that compound across all later katas.

## Rules / Expectations
- walk dir recursively
- count by extension (lowercased)
- skip directories

## Prior reading
- [Go os package](https://pkg.go.dev/os)
- [Go path/filepath package](https://pkg.go.dev/path/filepath)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains os, filepath by implementing `File Extension Counter` under explicit constraints.
- It is important because `File Extension Counter` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
