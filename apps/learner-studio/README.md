# Learner Studio (Web MVP)

Local browser-based learning IDE for this kata repo.

## Features

- category + progress sidebar
- category milestone feedback
- tabbed layout:
  - `Documentation` tab (kata README)
  - `Workbench` tab (tests left, code right)
- `Run Tests` button wired to `go test -json`
- pass modal with `Ask AI to Mark` flow
- next recommended kata guidance after each run
- focused failure insights (expected vs actual cues)
- AI marking packet generation for ChatGPT handoff

## Run

From repo root:

```bash
go run ./apps/learner-studio
```

Then open:

```text
http://127.0.0.1:7777
```

## Notes

- Progress is stored in `.learning/progress.json`
- AI marking prompts are written to `.learning/marking/`
