# Kata 088 — YAML/TOML Parsing

**Focus:** encoding/json patterns, struct tags, config file parsing

## Your task
Implement:

```go
func ParseConfig(data string) (Config, error)
```

### Learning goal
- What you are practicing: parsing structured config data with struct tags.
- Why this matters: every Go application needs configuration. Understanding struct tags and config parsing patterns is essential.
- How this grows your Go skills: you learn how struct tags control serialization and how to design config types.

### Tips
- Use `encoding/json` for JSON parsing.
- Struct tags like `json:"name"` control field mapping.
- Always handle parse errors explicitly.

## Rules / Expectations
- valid JSON => Config with correct fields
- invalid JSON => error

## What this kata is about (and why it matters)
- Core lesson: struct tags are Go's way of controlling serialization.
- After this kata, you can parse any JSON/YAML/TOML config.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
