# Kata 149 — CI/CD Pipeline

**Focus:** GitHub Actions, CI/CD, automated testing, deployment

## Your task
Implement a CI/CD pipeline configuration generator.

### Learning goal
- What you are practicing: designing automated build/test/deploy pipelines.
- Why this matters: CI/CD is how modern teams ship code. Every Go project needs it.
- How this grows your Go skills: you learn GitHub Actions, automated testing, and deployment workflows.

### Tips
- Include: checkout, build, test, lint, deploy steps.
- Use matrix builds for multiple Go versions.
- Cache dependencies for faster builds.

## Rules / Expectations
- Pipeline has a name
- Pipeline has at least one step
- Steps include test and build

## What this kata is about (and why it matters)
- Core lesson: automate everything. Manual processes break.
- After this kata, you can set up CI/CD for any Go project.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
