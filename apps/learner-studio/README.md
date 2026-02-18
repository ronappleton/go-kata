# Learner Studio (Web MVP)

Local browser-based learning IDE for this kata repo.

## Features

- category + progress sidebar
- category milestone feedback
- tabbed layout:
  - `Documentation` tab (kata README)
  - `Workbench` tab (tests left, code right)
- theme toggle with persisted dark mode
- syntax-highlighted Go editors (tests + solution panels)
- auto-format on blur for each editor plus manual format buttons
- save/run keyboard shortcuts for fast practice loops
- `Run Tests` button wired to `go test -json`
- pass modal with `Ask AI to Mark` flow
- next recommended kata guidance after each run
- focused failure insights (expected vs actual cues)
- AI marking packet generation for ChatGPT handoff

## Workflow Tips

- Edit tests and solution side-by-side to keep behavior contracts explicit.
- Use `Run Tests` often; the `Failure Insights` panel is designed for one-fix-at-a-time learning.
- Use `Format` before reviewing code quality so feedback focuses on logic, not spacing.
- Generate an AI marking packet after a pass to get reflection prompts, not just a pass/fail.

## Shortcuts

- `Ctrl/Cmd+S`: save current kata files
- `Ctrl/Cmd+Enter`: run tests
- `Alt+Shift+F`: format both editors
- `Tab`: insert a tab in editor
- `Esc`: close open modal

## API Endpoints (UI)

- `GET /api/track`: category progress + recommendation
- `GET /api/kata?id=<id>`: load kata readme/code/tests/progress
- `POST /api/format`: gofmt current editor content
- `POST /api/save`: persist `kata.go` and `kata_test.go`
- `POST /api/run`: run tests and return coaching/failure insights
- `POST /api/mark`: generate prompt packet for AI review

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
