# GoKatas Native Desktop

This is the supported product entry point for Ubuntu 24.04 LTS amd64.

## Requirements

- GTK4 runtime
- rootless Podman
- digest-pinned Go runner image

## Development build

```bash
go build -tags gtk4 -o learner-desktop ./apps/learner-desktop
GOKATAS_RUNNER_IMAGE='registry.example/gokatas-runner@sha256:<64-hex-digest>' ./learner-desktop -content .
```

The application uses a native GTK4 window. It does not start a web server or launch a browser.

## Workspace locations

The application follows XDG locations:

- configuration: `$XDG_CONFIG_HOME/gokatas`
- learner workspaces: `$XDG_DATA_HOME/gokatas/workspaces`
- progress and run history: `$XDG_STATE_HOME/gokatas`
- cache: `$XDG_CACHE_HOME/gokatas`

## Sandboxing

Every run is delegated to a fresh rootless Podman container with no network, a read-only container root, read-only source mounts, dropped capabilities, no-new-privileges, and bounded CPU, memory, PID, output, disk, and wall-clock resources.