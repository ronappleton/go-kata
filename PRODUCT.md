# Product

<!-- impeccable:product-schema 1 -->

## Platform

multi-platform desktop

> Native GTK4 desktop application supporting Linux (amd64), macOS (arm64),
> and Windows (amd64). Built with Go + gotk4 CGO bindings. Each platform
> requires its native GTK4 development libraries at build time.

## Stack

Go 1.25+ with `github.com/diamondburned/gotk4` (GTK4/GLib cgo bindings).
Evaluation runs learner code in a rootless Podman container from a
digest-pinned Go runner image (`--network=none`, read-only root, dropped
capabilities, no-new-privileges, CPU/memory/PID/output/time limits).

Content is distributed via a separate GitHub repository
(`gokatas-content`) and fetched on first launch with local caching.
The production package does not bundle mutable curriculum; cached content
provides offline support after the first successful sync.

## Users

- **Primary:** Go learners at any level — from first-day beginners
  through experienced developers reaching for lead-level architecture
  and leadership skills.
- **Secondary:** Teams using GoKatas as a shared training platform with
  consistent curriculum and progress tracking.

## Product Purpose

Teach Go through a structured, multi-stage curriculum of katas. The learner
writes a solution, and the app executes their submitted code and tests in an
isolated sandbox, comparing it against immutable trusted evaluator tests to
produce honest pass/fail feedback.

The curriculum covers four progression stages:
- **Foundation** (Junior): Go setup, language basics, data structures, error handling, testing fundamentals, AI literacy
- **Practitioner** (Mid): Standard library deep dive, web/networking, real-world packages (slog, cobra, testify, pgx), concurrency patterns
- **Senior**: Architecture & design, observability, performance, security, advanced data structures
- **Lead**: Code quality & review, build & deploy, design patterns, leadership & communication, bug fix lab

Plus addon tracks for Terraform, Helm, and Security/CVE awareness.

## Positioning

A desktop-native, offline-first Go practice workbench. Unlike web REPLs,
it runs on the learner's own machine as a first-class GTK4 application while
still isolating untrusted code in a disposable container. No browser, no
server, no cloud dependency for the core loop.

Three learning modes serve different cognitive styles:
- **Linear**: Sequential progression with prerequisite gates
- **ADHD**: Non-linear navigation with break reminders and quick wins
- **Review**: Cross-kata flashcard decks and quizzes with spaced repetition

## Operating Context

- **Linux**: Ubuntu 24.04 LTS+ (amd64), rootless Podman, GTK4
- **macOS**: 13.0+ (arm64 Apple Silicon), rootless Podman, GTK4 via Homebrew
- **Windows**: Windows 10+ (amd64), Podman Desktop, GTK4 via MSYS2 MinGW64
- Curriculum content is fetched from `gokatas-content` repo on first launch
- Learner work and progress live under XDG directories
  (`~/.local/share/gokatas`, `~/.local/state/gokatas`)

## Capabilities and Constraints

- 4-stage curriculum sidebar with progress bars and level badges
- Documentation view, solution + learner-test editors
- Sandboxed run/save via rootless Podman containers
- Flashcards with confidence tracking (Again/Hard/Good/Easy)
- Multiple choice and fill-in-the-blank quizzes
- Cross-kata review decks with spaced repetition
- Bug hunt mode for debugging practice
- Reflection mode for learning journaling
- Track selector for switching between Go Mastery, Terraform, Helm, Security
- Learning mode selector (Linear/ADHD/Review)

- Trusted evaluator tests are immutable and separated from learner-authored
  tests. Evaluators are classified `ready`, `incomplete`, or `missing`; only
  `ready` evaluators can run or mark a kata complete.
- Kata versioning: when a kata is updated, users who completed an older
  version see "needs re-completion" for the new version.
- Constraint: the runner image reference must be digest-pinned (`@sha256:`);
  runtime never pulls an image on the learner's behalf.
- Constraint: submitted code is never executed outside the container, and the
  container has no host filesystem access beyond the per-run read-only inputs.

## Brand Commitments

- Name: **GoKatas** (established in the codebase: app id, window title,
  desktop entry).
- Logo: **none confirmed (inferred)** — a fitting identity may be designed.
- Voice: plain, encouraging, developer-to-developer.

## Evidence on Hand

- 180 kata directories under `katas/` with READMEs, starter code, metadata,
  and flashcards/quizzes embedded in `kata.json`.
- 4 curriculum tracks under `tracks/` (Go Mastery, Terraform, Helm, Security).
- Remote content repository at `github.com/ronappleton/gokatas-content`.
- No testimonials, usage data, or press — future work must not fabricate these.

## Product Principles

1. **Safety is the feature.** Untrusted learner code never executes on the host.
2. **Honest feedback over gamification.** A kata without a real evaluator is
   shown as not-yet-runnable, never as "complete".
3. **Native first.** The app should feel like a proper desktop tool, not
   a ported web page.
4. **The learner owns their work.** Solutions, tests, and reflections live in
   the user's XDG directories and are never hidden or locked.
5. **Offline by default.** The core learn-and-run loop requires no network.
6. **Multi-platform.** Linux, macOS, and Windows users get the same experience.
7. **Content updates without app updates.** New katas and tracks ship via the
   content repository, not app releases.

## Accessibility & Inclusion

- Full keyboard operability, visible focus, and legible contrast in the
  shipped dark theme.
- ADHD mode with break reminders, non-linear navigation, and visual progress.
- Learning modes accommodate different cognitive styles (linear vs flexible).
