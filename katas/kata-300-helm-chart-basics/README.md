# Kata 300 — Helm Chart Basics

**Focus:** Chart.yaml, values.yaml, templates, helm install, helm upgrade

## Your task
Generate a Helm Chart.yaml configuration.

### Learning goal
- What you are practicing: creating Helm charts for Kubernetes application deployment.
- Why this matters: Helm is the package manager for Kubernetes. Every production K8s app uses Helm charts.
- How this grows your skills: you learn Kubernetes packaging and templating.

### The Helm Workflow
1. `helm create` — scaffold a chart
2. `helm template` — render templates locally
3. `helm install` — deploy to cluster
4. `helm upgrade` — update deployment
5. `helm rollback` — revert changes

## Rules / Expectations
- Output contains valid Chart.yaml structure
- Includes apiVersion, name, version, description

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
