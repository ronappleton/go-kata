# Kata 089 — Structured Logging (slog)

**Focus:** log/slog, structured logging, key-value pairs, log levels

## Your task
Implement structured logging with Go's `log/slog` package.

### Learning goal
- What you are practicing: structured logging with key-value pairs and log levels.
- Why this matters: structured logs are machine-parseable, searchable, and essential for production debugging.
- How this grows your Go skills: you learn the slog API that replaced most third-party logging in Go 1.21+.

### Tips
- Use `slog.Info("msg", "key", value)` for structured logging.
- Use `slog.NewJSONHandler` for JSON output.
- Always include context in log messages.

## Rules / Expectations
- LogRequest logs method, path, status as structured fields
- LogError logs message and error

## What this kata is about (and why it matters)
- Core lesson: structured logs > unstructured logs. Always.
- After this kata, you'll never use fmt.Println for logging again.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
