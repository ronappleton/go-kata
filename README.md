# Golang Katas 1–100 (Zero → Hero)

Structured Go practice for developers coming from other languages.

This repo is designed to teach, not just test.
Each kata is a small contract:
- what behavior must be correct
- why that behavior matters in real software
- how to practice it with focused feedback

## Two learning modes

### 1) CLI mode (`learner-cli`)

Fast, focused, and scriptable:
- list categories and progress
- run a kata test loop
- generate AI marking packets

![Learner CLI in operation](assets/screenshots/cli-app-operation.svg)

### 2) Browser mode (`learner-studio`)

A learning IDE flow:
- category progress and milestones
- tabbed docs + workbench
- tests left, code right
- run tests, pass modal, AI marking
- next recommended kata guidance

![Learner Studio in operation](assets/screenshots/learner-studio-operation.svg)

## Quick start

```bash
# CLI
go run ./apps/learner-cli list
go run ./apps/learner-cli show --kata 001
go run ./apps/learner-cli run --kata 001
go run ./apps/learner-cli mark --kata 001

# Studio
go run ./apps/learner-studio
# open http://127.0.0.1:7777
```

## Learning workflow (recommended)

1. Pick one kata.
2. Read the kata README and learning goals.
3. Run tests once before coding.
4. Implement one behavior at a time.
5. Rerun tests after each small change.
6. If tests pass, request AI marking feedback.
7. Move to the recommended next kata.

## Operability and quality checks

Run all learner platform tests:

```bash
GOCACHE=$(pwd)/.gocache go test ./apps/learner-cli ./apps/learner-studio ./internal/learning/...
```

Run full operability smoke checks:

```bash
./scripts/test_operability.sh
```

## Project structure

- `kata-001-*` ... `kata-100-*`: self-contained kata folders (`kata.go`, `kata_test.go`, `README.md`, local `go.mod`)
- `tracks/go-core-100/`: category model and learning goals
- `internal/learning/`: shared learning engine (catalog, runner, progress, marking)
- `apps/learner-cli/`: terminal learning app
- `apps/learner-studio/`: browser learning app

## Current platform capabilities

- category-level progress percentages
- milestone messaging per category
- next recommended kata selection
- run result capture and focused failure insights
- AI marking packet generation for ChatGPT handoff
