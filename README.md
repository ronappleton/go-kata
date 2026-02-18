# Golang Katas 1–100 (Zero → Hero)

## How to use
1. Open a kata folder (kata-001 ... kata-100) in GoLand.
2. Read that kata’s `README.md`.
3. Implement the function(s) in `kata.go`.
4. Fill in tests in `kata_test.go` (remove `t.Skip(...)` and add real assertions).
5. Run:

```bash
go test ./...
```

## Marking
- Submit your solution (code + tests) for a kata and it will be marked and corrected.
- One kata at a time.

Each kata folder is self-contained with its own `go.mod`.

## Learning platform scaffold (new)

This repo now includes a learning-platform foundation:

- track/category metadata in `tracks/go-core-100/track.json`
- shared engine packages in `internal/learning/*`
- CLI learning app in `apps/learner-cli`
- browser learning studio in `apps/learner-studio`

Try it:

```bash
go run ./apps/learner-cli list
go run ./apps/learner-cli run --kata 001
go run ./apps/learner-studio
```
