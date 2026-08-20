#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-$ROOT_DIR/.gocache}"

echo "==> Formatting check"
if [[ -n "$(gofmt -l apps internal)" ]]; then
  gofmt -l apps internal
  exit 1
fi

echo "==> Go tests"
go test ./...

echo "==> Go vet"
go vet ./...

echo "==> Native source/build prerequisites"
if pkg-config --exists gtk4; then
  mkdir -p "$ROOT_DIR/.build"
  go test -tags gtk4 ./apps/learner-desktop
  go build -tags gtk4 -o "$ROOT_DIR/.build/gokatas" ./apps/learner-desktop
else
  echo "GTK4 development files unavailable; native build deferred."
fi

echo "==> Sandbox prerequisites"
if command -v podman >/dev/null 2>&1; then
  podman info >/dev/null
else
  echo "Podman unavailable; sandbox integration deferred."
fi

echo "Operability checks passed."