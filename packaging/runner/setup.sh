#!/usr/bin/env bash
set -euo pipefail

IMAGE="${1:-${GOKATAS_RUNNER_IMAGE:-}}"
if [[ ! "$IMAGE" =~ @sha256:[0-9a-fA-F]{64}$ ]]; then
  echo "Usage: $0 registry/name@sha256:<64-hex-digest>" >&2
  exit 2
fi

if ! command -v podman >/dev/null 2>&1; then
  echo "Podman is required. Install rootless Podman before configuring the runner." >&2
  exit 1
fi

if ! podman info >/dev/null 2>&1; then
  echo "Rootless Podman is not available for this user." >&2
  exit 1
fi

if ! podman image exists "$IMAGE"; then
  echo "Pulling the digest-pinned runner image: $IMAGE"
  podman pull "$IMAGE"
fi
if ! podman image exists "$IMAGE"; then
  echo "Runner image was not available after pull: $IMAGE" >&2
  exit 1
fi

CONFIG_HOME="${XDG_CONFIG_HOME:-$HOME/.config}"
CONFIG_DIR="$CONFIG_HOME/gokatas"
mkdir -p "$CONFIG_DIR"
umask 077
printf '%s\n' "$IMAGE" > "$CONFIG_DIR/runner-image"
echo "Configured digest-pinned runner image in $CONFIG_DIR/runner-image"