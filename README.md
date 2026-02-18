# GoKatas: From Zero to Senior in Go
[![Operability](https://github.com/ronappleton/golang-katas-1-100/actions/workflows/operability.yml/badge.svg)](https://github.com/ronappleton/golang-katas-1-100/actions/workflows/operability.yml)

This repository is a complete learning system, not just a pile of coding exercises.

It is designed to take a learner from first programming steps to senior-level Go habits using:
- a structured 140-kata curriculum
- staged pathways (No Knowledge -> Junior -> Mid -> Senior)
- multiple learning modes (build, debug, recall, quiz, reflect)

## Why this is different
Most kata repos only ask learners to "make tests green." This system teaches:
- what the kata is training
- why the skill matters in real Go work
- where the skill appears in services, I/O, and production debugging
- how each kata fits into a progression path

## Dark mode preview
| Documentation | Workbench |
|---|---|
| ![Studio docs dark](apps/learner-studio/assets/screenshots/learner-studio-docs-dark.svg) | ![Studio workbench dark](apps/learner-studio/assets/screenshots/learner-studio-workbench-dark.svg) |

## Use the tools
- CLI workflow: [apps/learner-cli/README.md](apps/learner-cli/README.md)
- Studio workflow: [apps/learner-studio/README.md](apps/learner-studio/README.md)
- Desktop workflow: [apps/learner-desktop/README.md](apps/learner-desktop/README.md)

## Quick start
```bash
# Terminal mode
go run ./apps/learner-cli list

# Studio mode
go run ./apps/learner-studio

# Desktop-launch mode
go run ./apps/learner-desktop

# Embedded native-webview mode (self-contained window)
./scripts/run_desktop_native.sh
```

## Curriculum at a glance
- Starter foundations (`001-020`)
- Well-known packages (`021-030`)
- Core + services + concurrency (`031-130`)
- Bug Fix Lab (`131-135`)
- Databases and I/O (`136-140`)

Full ordered list: [INDEX.md](INDEX.md)
