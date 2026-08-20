# Kata 148 — Docker Multi-Stage Build

**Focus:** Dockerfile, multi-stage builds, Go binary, containerization

## Your task
Implement a Dockerfile generator for Go applications.

### Learning goal
- What you are practicing: creating efficient Docker images for Go applications.
- Why this matters: multi-stage builds create tiny production images (10MB vs 1GB+).
- How this grows your Go skills: you learn the deployment side of Go development.

### Tips
- Use `golang` as the build stage, `alpine` or `scratch` as runtime.
- Copy only the binary in the final stage.
- Use `CGO_ENABLED=0` for static binaries.

## Rules / Expectations
- Output contains FROM, COPY, RUN, EXPOSE
- Multi-stage (at least 2 FROM statements)
- Copies binary in final stage

## What this kata is about (and why it matters)
- Core lesson: containerization is a core DevOps skill for Go developers.
- After this kata, you can containerize any Go application efficiently.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
