#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-$ROOT_DIR/.gocache}"

echo "==> Running learner platform tests"
go test ./apps/learner-studio ./apps/learner-desktop ./internal/learning/...

echo "==> Running desktop launcher compile smoke check"
go run ./apps/learner-desktop -h >/dev/null 2>&1

echo "==> Checking native launcher script exists"
test -x ./scripts/run_desktop_native.sh

echo "Operability checks passed."
