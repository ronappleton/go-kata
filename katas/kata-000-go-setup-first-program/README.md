# Kata 000 — Go Setup & First Program

**Focus:** go mod init, main package, go run, go build

## Your task
Implement:

```go
func InitModule(name string) string
```

### Learning goal
- What you are practicing: understanding the Go project layout, module system, and basic toolchain commands.
- Why this matters: every Go project starts with a module. Understanding `go mod init`, `go run`, and `go build` is essential before writing any code.
- How this grows your Go skills: you learn the development workflow that every Go developer uses daily.

### Tips
- Run `go help mod` to see available commands.
- Try creating a real module with `go mod init example.com/myapp`.

## Rules / Expectations
- simple module => "example.com/myapp"
- with path => "github.com/user/repo"

## What this kata is about (and why it matters)
- Core lesson: every Go project needs a module. Understanding the toolchain is step one.
- After this kata, you should be comfortable setting up new Go projects.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
