#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if ! pkg-config --exists gtk4; then
  echo "GTK4 development files are required. Install libgtk-4-dev and pkg-config." >&2
  exit 1
fi

if [[ -z "${GOKATAS_RUNNER_IMAGE:-}" && ! -f "${XDG_CONFIG_HOME:-$HOME/.config}/gokatas/runner-image" && ! -f /etc/gokatas/runner-image ]]; then
  echo "No digest-pinned runner image is configured." >&2
  echo "Run packaging/runner/setup.sh after preparing a local Podman image." >&2
  exit 1
fi

exec go run -tags gtk4 ./apps/learner-desktop -content "$ROOT_DIR" "$@"