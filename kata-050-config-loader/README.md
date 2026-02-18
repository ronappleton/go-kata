# Kata 050 — Config Loader

**Focus:** env + file

## Your task
Implement:

```go
func LoadConfig(path string, envPrefix string) (map[string]string, error)
```

### Learning goal
- Expected work: Implement `Config Loader` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Config Loader` teaches core implementation habits that compound across all later katas.

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
- This kata trains env + file by implementing `Config Loader` under explicit constraints.
- It is important because `Config Loader` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
