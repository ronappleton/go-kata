# Runner image

The evaluator requires an immutable image reference:

```text
registry.example/gokatas-runner@sha256:<64 hex characters>
```

Builds must use the pinned Go toolchain in `Containerfile`, inspect the resulting image, and publish the digest in the Debian package configuration. Runtime execution uses `--pull=never`, so the application never lets a learner run trigger an image pull.

For local development, build an image and use the digest returned by Podman:

```bash
podman build --build-arg GO_BASE_IMAGE='docker.io/library/golang@sha256:<base-image-digest>' -t localhost/gokatas-runner:dev -f packaging/runner/Containerfile packaging/runner
podman image inspect localhost/gokatas-runner:dev --format '{{index .RepoDigests 0}}'
export GOKATAS_RUNNER_IMAGE='localhost/gokatas-runner@sha256:<digest>'
```

A release pipeline must publish and record the digest rather than using the development tag.