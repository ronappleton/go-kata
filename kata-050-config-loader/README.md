# Kata 050 — Config Loader

**Focus:** env + file

## Your task
Implement:

```go
func LoadConfig(path string, envPrefix string) (map[string]string, error)
```

### Learning goal
- What you are practicing: Build `Config Loader` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `load key=value lines`; `override with env`; and `trim spaces`.

## Rules / Expectations
- load key=value lines
- override with env
- trim spaces

## Prior reading
- [Go os package](https://pkg.go.dev/os)
- [Go path/filepath package](https://pkg.go.dev/path/filepath)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in env + file through `Config Loader`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
