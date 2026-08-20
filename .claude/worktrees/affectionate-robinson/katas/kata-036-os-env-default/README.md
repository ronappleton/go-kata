# Kata 026 — Read Env Or Default

**Focus:** os.LookupEnv, strings.TrimSpace

## Your task
Implement:

```go
func ReadEnvOrDefault(key, fallback string) string
```

### Learning goal
- What you are building: predictable configuration lookup with explicit fallback behavior.
- Why this matters in real projects: env config is a common startup boundary; defaults must be intentional.
- How this grows your Go skills: you learn to differentiate unset, empty, and valid env values.
- Definition of done (plain English): callers always receive a usable value according to contract rules.

### Tips
- Use `LookupEnv` to distinguish unset from empty.
- Test unset and empty env states separately.
- Keep key normalization explicit.

## Rules / Expectations
- trim spaces from key before lookup
- empty key after trim => fallback
- env var unset => fallback
- env var set but empty => fallback
- env var set and non-empty => return value

## Prior reading
- [os package](https://pkg.go.dev/os)
- [strings package](https://pkg.go.dev/strings)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: configuration defaults are part of your reliability contract.
- After this kata, you should be able to write safer startup config logic.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
