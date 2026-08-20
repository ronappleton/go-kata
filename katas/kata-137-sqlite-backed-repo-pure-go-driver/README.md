# Kata 108 — SQLite-backed Repo (pure Go driver)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata78() error
```

### Learning goal
- What you are building: Build `func Kata78() error` as a reliable contract. Focus: real-world Go integration and testing.
- Why this matters in real projects: Data bugs persist. Transaction-safe behavior protects your system.
- How this grows your Go skills: You practice repository boundaries and transaction-aware failure handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: match the full README contract; include tests for happy path and edge cases; and keep API and error style idiomatic Go.

### Tips
- Test success and rollback paths equally.
- Keep persistence details out of business logic.
- Assert transaction outcomes, not just returned rows.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go database/sql package](https://pkg.go.dev/database/sql)
- [modernc.org/sqlite driver docs](https://pkg.go.dev/modernc.org/sqlite)
- [SQLite transaction docs](https://www.sqlite.org/lang_transaction.html)

## What this kata is about (and why it matters)
- Core lesson: protect data with clear transaction boundaries.
- After this kata, you should be able to explain transaction outcomes for success and failure paths.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
