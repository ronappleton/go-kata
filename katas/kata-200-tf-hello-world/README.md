# Kata 200 — Terraform Hello World

**Focus:** terraform init, plan, apply, destroy, HCL basics

## Your task
Generate a basic Terraform configuration file.

### Learning goal
- What you are practicing: writing Terraform HCL and understanding the init/plan/apply/destroy lifecycle.
- Why this matters: Infrastructure as Code is how modern teams manage cloud resources. Terraform is the standard.
- How this grows your skills: you learn declarative infrastructure and the Terraform workflow.

### The Terraform Workflow
1. `terraform init` — initialize the provider
2. `terraform plan` — preview changes
3. `terraform apply` — apply changes
4. `terraform destroy` — tear down resources

## Rules / Expectations
- Output contains valid HCL resource block
- Includes resource type and name

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
