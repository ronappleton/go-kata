# Kata 031 — File Extension Counter

**Focus:** os, filepath

## Your task
Implement:

```go
func CountExt(root string) (map[string]int, error)
```

### Learning goal
- What you are practicing: Build `File Extension Counter` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `walk dir recursively`; `count by extension (lowercased)`; and `skip directories`.

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
- This kata is focused practice in os, filepath through `File Extension Counter`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
