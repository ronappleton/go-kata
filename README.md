# GoKatas

A native desktop Go learning environment. Structured curriculum from junior to lead, with sandboxed code execution, flashcards, quizzes, and multiple learning modes.

## Features

- **180 katas** across 4 learning tracks
- **4-stage curriculum**: Foundation → Practitioner → Senior → Lead
- **Sandboxed execution** via rootless Podman containers
- **3 learning modes**: Linear (sequential), ADHD (non-linear), Review (spaced repetition)
- **Flashcards & quizzes** with confidence tracking
- **Multi-platform**: Linux, macOS, Windows
- **Content updates** via remote repository (no rebuild needed)

## Supported Platforms

| Platform | Architecture | Build Requirement |
|----------|-------------|-------------------|
| Linux | amd64 | `libgtk-4-dev` |
| macOS | arm64 (Apple Silicon) | `brew install gtk4` |
| Windows | amd64 | MSYS2 MinGW64 + `pacman -S mingw-w64-x86_64-gtk4` |

## Quick Start

### Linux (Ubuntu/Debian)

```bash
make install-deps   # install GTK4 dev libraries
make build           # build desktop app
make test            # run tests
```

### macOS

```bash
make install-deps-macos   # install GTK4 via Homebrew
make build                # build desktop app
```

### Windows

```bash
# Open MSYS2 MinGW64 shell
make install-deps-windows   # install GTK4 via pacman
make build                  # build desktop app
```

### Docker (Linux only)

```bash
make docker-build    # reproducible Linux build in container
```

## Running

```bash
# From source
./dist/bin/gokatas -content .

# Installed (via .deb or manual install)
gokatas
```

The app requires rootless Podman and a digest-pinned runner image:

```bash
podman info
export GOKATAS_RUNNER_IMAGE='registry.example/gokatas-runner@sha256:<digest>'
```

## Architecture

```
apps/learner-desktop/     GTK4 native desktop application
internal/
  learning/
    catalog/              Curriculum loading (multi-track, 4-stage)
    content/              Remote + local + embedded content provider
    evaluator/            Rootless Podman code execution
    katas/                Embedded kata data (generated)
    marking/              Review packet generation
    progress/             Versioned progress state
    runner/               Kata test runner
    workspace/            XDG workspaces, atomic persistence
tracks/                   Curriculum configurations
  go-core-100/            Go Mastery: Junior to Lead (175 katas)
  terraform-100/          Infrastructure as Code (2 katas)
  helm-100/               Container Orchestration (2 katas)
  security-100/           Security & CVE Awareness (1 kata)
katas/                    Kata source content (180 directories)
scripts/                  Code generation and content sync
build/                    Docker build files
```

## Curriculum

### Go Mastery Track (175 katas)

| Stage | Level | Categories | Topics |
|-------|-------|------------|--------|
| **Foundation** | Junior | 6 | Setup, Language Basics, Data Structures, Error Handling, Testing, AI Literacy |
| **Practitioner** | Mid | 6 | StdLib, Web, Data Formats, Real-World Packages, Concurrency, AI at Scale |
| **Senior** | Senior | 5 | Architecture, Observability, Performance, Security, Advanced Data Structures |
| **Lead** | Lead | 5 | Code Quality, Build & Deploy, Patterns, Leadership, Bug Fix Lab |

### Addon Tracks

| Track | Description |
|-------|-------------|
| **Terraform** | Infrastructure as Code fundamentals |
| **Helm** | Kubernetes chart authoring |
| **Security** | CVE awareness and analysis |

## Build System

| Command | Description |
|---------|-------------|
| `make build` | Build for current platform |
| `make build-linux-amd64` | Build for Linux amd64 |
| `make build-darwin-arm64` | Build for macOS arm64 |
| `make build-windows-amd64` | Build for Windows amd64 |
| `make docker-build` | Build inside Docker (Linux) |
| `make test` | Run all tests |
| `make test-race` | Run tests with race detector |
| `make test-cover` | Run tests with coverage report |
| `make package` | Package current platform |
| `make package-darwin` | Package as macOS .app bundle |
| `make package-windows` | Package as Windows zip |
| `make clean` | Remove build artifacts |
| `make help` | Show all targets |

## Content Repository

Katas are distributed via a separate content repository:

- **Repository**: [github.com/ronappleton/gokatas-content](https://github.com/ronappleton/gokatas-content)
- **Raw URL**: `https://raw.githubusercontent.com/ronappleton/gokatas-content/main`

The app fetches content on first launch, caches locally, and checks for updates periodically. Embedded fallback works offline.

### Updating Content

```bash
# Regenerate embedded data after editing katas
make gen-katas

# Generate content repo for remote distribution
make sync-content
```

## Verification

```bash
make test           # run all tests
make vet            # run go vet
make lint           # run vet + gofmt
```

## License

See LICENSE file.
