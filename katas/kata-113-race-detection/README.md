# Kata 113 — Race Detection

**Focus:** go test -race, data races, sync.Mutex, race detector

## Your task
Make the Counter type safe for concurrent use and verify with the race detector.

### Learning goal
- What you are practicing: identifying and fixing data races using Go's race detector.
- Why this matters: data races cause non-deterministic bugs that are nearly impossible to reproduce without the detector.
- How this grows your Go skills: you learn to use `go test -race` as a standard part of your workflow.

### Tips
- Use `sync.Mutex` to protect shared state.
- Run `go test -race ./...` to detect races.
- The race detector adds ~10x overhead — use it in tests, not production.

## Rules / Expectations
- Counter is safe for concurrent use
- 1000 goroutines each increment once => value is 1000
- Passes with -race flag

## What this kata is about (and why it matters)
- Core lesson: always run the race detector. Always.
- After this kata, `go test -race` will be part of your muscle memory.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test -race ./...
```
