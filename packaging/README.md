# Ubuntu packaging

The supported release artifact is a Debian package for Ubuntu 24.04 LTS amd64.

## Build

Install the build dependencies, then run:

```bash
sudo apt install build-essential pkg-config libgtk-4-dev podman dpkg
./packaging/build_deb.sh 0.1.0
```

The package is written to `dist/`.

## Install

```bash
sudo apt install ./dist/gokatas_0.1.0_amd64.deb
```

The first-run setup must configure a digest-pinned runner image:

```bash
/usr/lib/gokatas/setup-runner.sh registry.example/gokatas-runner@sha256:<64-hex-digest>
```

The setup command uses rootless Podman as the logged-in user. It does not configure Docker or require a privileged container daemon.