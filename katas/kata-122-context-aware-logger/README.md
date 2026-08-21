# Kata 122 — Context-aware Logger

**Focus:** context, io, log

## Your task
Implement:

```go
type Level int; func NewLogger(w io.Writer) *Logger; func (l *Logger) Log(ctx context.Context, level Level, msg string); func WithRequestID(ctx context.Context, id string) context.Context
```

### Learning goal
- What you are building: type Level int; func NewLogger(w io.Writer) *Logger; func (l *Logger) Log(ctx context.Context, level Level, msg string); func WithRequestID(ctx context.Context, id string) context.Context as a reliable contract. Focus: context, io, log.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): Log writes "timestamp LEVEL [request_id] msg" (request_id only when present in ctx) with RFC3339 timestamps and LEVEL in INFO/WARN/ERROR.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- INFO/WARN/ERROR levels
- timestamped lines
- request id from context when present
- writes to the injected writer

## Prior reading
- [Go context package](https://pkg.go.dev/context)
- [Go log package](https://pkg.go.dev/log)

## What this kata is about (and why it matters)
- Core lesson: logs are data: keep them structured, timestamped, and context-aware.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
