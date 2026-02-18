# Learner CLI (MVP)

Terminal-first learning workflow for the kata track.

## What it does

- loads category track metadata from `tracks/go-core-100/track.json`
- shows category/kata progress
- runs a kata's tests with `go test -json`
- records attempts in `.learning/progress.json`
- generates a structured AI marking prompt in `.learning/marking/`

## Run

From repo root:

```bash
go run ./apps/learner-cli list
go run ./apps/learner-cli katas --category foundations
go run ./apps/learner-cli run --kata 001
go run ./apps/learner-cli mark --kata 001
```

## Data files

- progress state: `.learning/progress.json`
- marking prompts: `.learning/marking/*.md`
