#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-$ROOT_DIR/.gocache}"

if ! go list github.com/webview/webview_go >/dev/null 2>&1; then
  cat >&2 <<'MSG'
Missing native webview dependency: github.com/webview/webview_go
Install it first:
  go get github.com/webview/webview_go@latest
MSG
  exit 1
fi

exec go run -tags desktop_webview ./apps/learner-desktop -mode embedded "$@"
