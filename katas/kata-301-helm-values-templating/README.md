# Kata 301 — Helm Values & Templating

**Focus:** values.yaml, Go templates, conditional rendering, loops

## Your task
Render a Kubernetes Deployment from Helm values.

### Learning goal
- What you are practicing: using Go templates with Helm values to generate Kubernetes manifests.
- Why this matters: Helm's power comes from templating — one chart deploys to dev, staging, and production with different values.
- How this grows your skills: you learn Go templating and Kubernetes resource modeling.

## Rules / Expectations
- Output is a valid Kubernetes Deployment
- Uses values from TemplateContext
- Includes app name, replicas, image, and port

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
