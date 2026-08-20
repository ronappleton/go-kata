#!/usr/bin/env python3
"""Create new tracks and AI literacy katas for the GoKatas curriculum expansion."""
import json, os
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
KATAS_DIR = ROOT / "katas"
TRACKS_DIR = ROOT / "tracks"

# ═══════════════════════════════════════════════
# AI LITERACY KATAS (added to Go Foundation)
# ═══════════════════════════════════════════════
AI_KATAS = [
    {
        "id": "160", "slug": "ai-prompt-engineering",
        "title": "AI Prompt Engineering for Developers",
        "focus": "effective prompting, context setting, constraint specification",
        "stage": "foundation", "category": "ai-literacy", "level": "junior",
        "tags": ["ai", "tooling"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

// PromptBuilder constructs effective AI prompts.
// Your task: implement the function below.
//
// A good prompt includes:
// 1. What you want (clear goal)
// 2. Constraints (language, style, requirements)
// 3. Context (existing code, patterns to follow)
//
// BuildPrompt takes a goal and constraints and returns a well-structured prompt.
func BuildPrompt(goal string, constraints []string) string {
\treturn ""
}''',
        "kata_test": '''package kata

import "testing"

func TestBuildPrompt(t *testing.T) {
\ttests := []struct {
\t\tname        string
\t\tgoal        string
\t\tconstraints []string
\t\twantMinLen  int
\t}{
\t\t{"simple goal", "sort a slice", nil, 10},
\t\t{"with constraints", "parse JSON", []string{"use encoding/json", "handle errors"}, 20},
\t}

\tfor _, tc := range tests {
\t\tt.Run(tc.name, func(t *testing.T) {
\t\t\tgot := BuildPrompt(tc.goal, tc.constraints)
\t\t\tif len(got) < tc.wantMinLen {
\t\t\t\tt.Errorf("prompt too short: %d chars, want >= %d", len(got), tc.wantMinLen)
\t\t\t}
\t\t})
\t}
}''',
        "readme": '''# Kata 160 — AI Prompt Engineering for Developers

**Focus:** effective prompting, context setting, constraint specification

## Your task
Implement a prompt builder that constructs effective AI prompts.

### Learning goal
- What you are practicing: writing clear, structured prompts that get useful AI responses.
- Why this matters: AI is only as good as your prompt. Vague prompts get vague answers. Specific prompts get specific, testable code.
- How this grows your Go skills: you learn to articulate requirements precisely — a skill that helps even when you're NOT using AI.

### The Developer's AI Contract
1. **Never copy-paste AI output without understanding it.** Read every line. If you can't explain it, you can't maintain it.
2. **Use AI for boilerplate, not for thinking.** AI is great at "write a struct with these fields." It's bad at "design the architecture."
3. **Always verify with tests.** AI-generated code that passes tests is trustworthy. AI-generated code without tests is a liability.
4. **Treat AI like a junior developer.** Review its code. Ask questions. Don't blindly merge.

### Tips
- Start with "I want to..." not "Can you..."
- Include the language, error handling style, and test expectations
- Show existing code patterns you want to follow
- Ask for explanations, not just code

## Rules / Expectations
- prompt includes the goal
- prompt includes constraints when provided
- prompt is clear and actionable

## What this kata is about (and why it matters)
- Core lesson: AI is a tool, not a replacement for thinking.
- After this kata, you'll prompt AI effectively and verify its output.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What makes a good AI prompt for code generation?", "back": "Clear goal, specific constraints, existing patterns to follow, and expected error handling."},
            {"front": "When should you NOT use AI for coding?", "back": "When you need to understand core concepts, when debugging complex issues, when making architecture decisions, or when learning a new language."},
            {"front": "What is the 'junior developer' rule for AI?", "back": "Treat AI output like junior dev code: review it, ask questions, don't blindly merge."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What should you do with AI-generated code before using it?", "options": ["Copy it directly", "Read and understand every line", "Trust it if it compiles", "Run it once and move on"], "answer": "Read and understand every line"},
            {"type": "multiple_choice", "question": "When is AI MOST useful for a developer?", "options": ["Designing system architecture", "Writing boilerplate and repetitive code", "Debugging complex race conditions", "Making product decisions"], "answer": "Writing boilerplate and repetitive code"},
        ],
    },
    {
        "id": "161", "slug": "ai-code-review",
        "title": "AI-Assisted Code Review",
        "focus": "using AI for code review, verifying AI suggestions, critical thinking",
        "stage": "foundation", "category": "ai-literacy", "level": "junior",
        "tags": ["ai", "testing"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

// ReviewWithAI takes code and returns suggested improvements.
// Your task: implement a function that demonstrates
// critical thinking about AI suggestions.
//
// The function should:
// 1. Identify what the code does
// 2. Flag potential issues (even if subtle)
// 3. Suggest improvements with reasoning
type ReviewResult struct {
\tIssues    []string
\tSuggestions []string
\tConfidence float64 // 0.0 to 1.0
}

// ReviewCode analyzes code and returns a structured review.
func ReviewCode(code string) ReviewResult {
\treturn ReviewResult{}
}''',
        "kata_test": '''package kata

import "testing"

func TestReviewCode(t *testing.T) {
\tcode := `func add(a, b int) int { return a + b }`
\tresult := ReviewCode(code)
\tif len(result.Issues) == 0 && len(result.Suggestions) == 0 {
\t\tt.Error("expected at least one issue or suggestion")
\t}
\tif result.Confidence < 0 || result.Confidence > 1 {
\t\tt.Errorf("confidence out of range: %f", result.Confidence)
\t}
}''',
        "readme": '''# Kata 161 — AI-Assisted Code Review

**Focus:** using AI for code review, verifying AI suggestions, critical thinking

## Your task
Implement a code review function that demonstrates critical thinking.

### Learning goal
- What you are practicing: using AI as a review partner while maintaining your own judgment.
- Why this matters: AI can catch issues you miss, but it also hallucinates. You need to know when to trust it and when to question it.
- How this grows your Go skills: you develop the habit of reviewing code critically — your own and AI's.

### The Verification Loop
1. **Ask AI to review your code** — it will find real issues and fake ones
2. **Verify each suggestion** — does it actually apply? Is it correct?
3. **Run the tests** — AI suggestions that break tests are wrong
4. **Apply your judgment** — some AI suggestions are technically correct but practically wrong

### Red Flags in AI Code Review
- Suggesting changes that would break existing tests
- Over-engineering simple solutions
- Ignoring the project's existing patterns
- Suggesting deprecated or insecure approaches
- Being confidently wrong about language features

## Rules / Expectations
- Returns structured review with issues and suggestions
- Confidence score reflects actual certainty
- Suggestions are actionable and specific

## What this kata is about (and why it matters)
- Core lesson: AI is a review partner, not a replacement for your judgment.
- After this kata, you'll use AI reviews critically and effectively.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is the verification loop for AI code review?", "back": "Ask AI → Verify each suggestion → Run tests → Apply your judgment."},
            {"front": "What are red flags in AI code review suggestions?", "back": "Breaking tests, over-engineering, ignoring patterns, deprecated approaches, confident wrong answers."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What should you do when AI suggests a code change?", "options": ["Apply it immediately", "Verify it against tests and existing patterns", "Ask a second AI for confirmation", "Ignore it if you didn't ask for it"], "answer": "Verify it against tests and existing patterns"},
        ],
    },
    {
        "id": "162", "slug": "ai-debugging-partner",
        "title": "AI as a Debugging Partner",
        "focus": "using AI to understand errors, not just fix them, root cause analysis",
        "stage": "practitioner", "category": "ai-literacy", "level": "mid",
        "tags": ["ai", "testing"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// DebugSession represents a structured debugging approach with AI.
type DebugSession struct {
\tError       string
\tContext     string
\tHypotheses  []string
\tTests       []string
\tConclusion  string
}

// StartDebugging takes an error and context, and generates
// a structured debugging plan. The goal is to UNDERSTAND
// the error, not just fix it.
func StartDebugging(errMsg, context string) DebugSession {
\treturn DebugSession{}
}''',
        "kata_test": '''package kata

import "testing"

func TestStartDebugging(t *testing.T) {
\tsession := StartDebugging("nil pointer dereference", "in user handler")
\tif len(session.Hypotheses) == 0 {
\t\tt.Error("expected at least one hypothesis")
\t}
\tif len(session.Tests) == 0 {
\t\tt.Error("expected at least one test")
\t}
\tif session.Conclusion != "" {
\t\tt.Error("conclusion should be empty until debugging is done")
\t}
}''',
        "readme": '''# Kata 162 — AI as a Debugging Partner

**Focus:** using AI to understand errors, not just fix them, root cause analysis

## Your task
Implement a structured debugging session that uses AI to understand errors.

### Learning goal
- What you are practicing: using AI to understand WHY code fails, not just HOW to fix it.
- Why this matters: AI can generate a fix in seconds, but if you don't understand the root cause, you'll hit the same bug again. Understanding > Fixing.
- How this grows your Go skills: you learn systematic debugging and root cause analysis.

### The Anti-Copy-Paste Debugging Rule
When you hit an error:
1. **Read the error message yourself first.** Understand what it says.
2. **Check the stack trace.** Where did it happen? What called what?
3. **Form a hypothesis.** "I think X is nil because Y."
4. **THEN ask AI** — but ask "WHY is this happening?" not "HOW do I fix this?"
5. **Verify the AI's explanation** matches your hypothesis
6. **Write a test** that reproduces the bug before fixing it

### Never Do This
- Copy-pasting AI's "fix" without understanding why it works
- Asking AI to "just fix it" without providing error context
- Skipping the hypothesis step — that's where learning happens

## Rules / Expectations
- Debugging session includes hypotheses
- Includes test suggestions
- Conclusion is empty until analysis is complete

## What this kata is about (and why it matters)
- Core lesson: AI helps you understand, not just fix. Understanding is the skill.
- After this kata, you'll debug with AI as a thinking partner, not a crutch.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is the anti-copy-paste debugging rule?", "back": "Read the error yourself → Check stack trace → Form hypothesis → THEN ask AI 'why' not 'how' → Verify explanation → Write reproduction test."},
            {"front": "Why is 'understanding > fixing' important with AI?", "back": "If you don't understand the root cause, you'll hit the same bug again. AI fixes are temporary without understanding."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "When debugging, what should you ask AI first?", "options": ["How do I fix this?", "Why is this happening?", "Can you rewrite this function?", "What's the best practice here?"], "answer": "Why is this happening?"},
        ],
    },
]

# ═══════════════════════════════════════════════
# TERRAFORM TRACK
# ═══════════════════════════════════════════════
TERRAFORM_KATAS = [
    {
        "id": "200", "slug": "tf-hello-world",
        "title": "Terraform Hello World",
        "focus": "terraform init, plan, apply, destroy, HCL basics",
        "stage": "foundation", "category": "tf-basics", "level": "junior",
        "tags": ["terraform", "iac"],
        "estimated_minutes": 20,
        "kata_go": '''package kata

// GenerateMainTF creates a basic Terraform main.tf file.
// Your task: implement the function below.
func GenerateMainTF(resourceType, resourceName string) string {
\treturn ""
}''',
        "kata_test": '''package kata

import (
\t"strings"
\t"testing"
)

func TestGenerateMainTF(t *testing.T) {
\ttf := GenerateMainTF("aws_instance", "web")
\tif !strings.Contains(tf, "resource") {
\t\tt.Error("expected Terraform resource block")
\t}
\tif !strings.Contains(tf, "aws_instance") {
\t\tt.Error("expected resource type")
\t}
\tif !strings.Contains(tf, "web") {
\t\tt.Error("expected resource name")
\t}
}''',
        "readme": '''# Kata 200 — Terraform Hello World

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
''',
        "flashcards": [
            {"front": "What are the 4 core Terraform commands?", "back": "init (setup), plan (preview), apply (execute), destroy (cleanup)"},
            {"front": "What is HCL?", "back": "HashiCorp Configuration Language — Terraform's declarative syntax for defining infrastructure."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which Terraform command previews changes before applying?", "options": ["terraform init", "terraform plan", "terraform apply", "terraform validate"], "answer": "terraform plan"},
        ],
    },
    {
        "id": "201", "slug": "tf-modules",
        "title": "Terraform Modules & Reusability",
        "focus": "modules, variables, outputs, reusability, composition",
        "stage": "foundation", "category": "tf-basics", "level": "junior",
        "tags": ["terraform", "iac"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// ModuleConfig represents a Terraform module configuration.
type ModuleConfig struct {
\tName       string
\tSource     string
\tVariables  map[string]string
\tOutputs    []string
}

// GenerateModule creates a reusable Terraform module.
func GenerateModule(config ModuleConfig) string {
\treturn ""
}''',
        "kata_test": '''package kata

import (
\t"strings"
\t"testing"
)

func TestGenerateModule(t *testing.T) {
\tconfig := ModuleConfig{
\t\tName:   "vpc",
\t\tSource: "./modules/vpc",
\t\tVariables: map[string]string{
\t\t\t"cidr": "10.0.0.0/16",
\t\t},
\t\tOutputs: []string{"vpc_id"},
\t}
\tresult := GenerateModule(config)
\tif !strings.Contains(result, "module") {
\t\tt.Error("expected module block")
\t}
\tif !strings.Contains(result, "vpc") {
\t\tt.Error("expected module name")
\t}
}''',
        "readme": '''# Kata 201 — Terraform Modules & Reusability

**Focus:** modules, variables, outputs, reusability, composition

## Your task
Generate a reusable Terraform module configuration.

### Learning goal
- What you are practicing: creating reusable Terraform modules with inputs and outputs.
- Why this matters: modules are how you share infrastructure patterns across teams and projects.
- How this grows your skills: you learn infrastructure composition and DRY principles.

## Rules / Expectations
- Output contains module block
- Includes source, variables, and outputs

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is a Terraform module?", "back": "A reusable container of Terraform configurations. Modules accept variables and produce outputs."},
            {"front": "Why use modules in Terraform?", "back": "DRY infrastructure. Share patterns across teams. Version and test infrastructure components."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What is the purpose of Terraform modules?", "options": ["Run commands faster", "Reuse infrastructure patterns", "Store state remotely", "Manage secrets"], "answer": "Reuse infrastructure patterns"},
        ],
    },
]

# ═══════════════════════════════════════════════
# HELM TRACK
# ═══════════════════════════════════════════════
HELM_KATAS = [
    {
        "id": "300", "slug": "helm-chart-basics",
        "title": "Helm Chart Basics",
        "focus": "Chart.yaml, values.yaml, templates, helm install, helm upgrade",
        "stage": "foundation", "category": "helm-basics", "level": "junior",
        "tags": ["helm", "kubernetes"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// ChartConfig represents a Helm chart configuration.
type ChartConfig struct {
\tName    string
\tVersion string
\tAppVersion string
\tDescription string
}

// GenerateChartYAML creates a Chart.yaml file.
func GenerateChartYAML(config ChartConfig) string {
\treturn ""
}''',
        "kata_test": '''package kata

import (
\t"strings"
\t"testing"
)

func TestGenerateChartYAML(t *testing.T) {
\tconfig := ChartConfig{
\t\tName:       "my-app",
\t\tVersion:    "1.0.0",
\t\tAppVersion: "1.0.0",
\t\tDescription: "A Helm chart",
\t}
\tyaml := GenerateChartYAML(config)
\tif !strings.Contains(yaml, "apiVersion") {
\t\tt.Error("expected apiVersion in Chart.yaml")
\t}
\tif !strings.Contains(yaml, "my-app") {
\t\tt.Error("expected chart name")
\t}
}''',
        "readme": '''# Kata 300 — Helm Chart Basics

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
''',
        "flashcards": [
            {"front": "What is Helm?", "back": "The package manager for Kubernetes. It packages, installs, and upgrades Kubernetes applications."},
            {"front": "What are the core Helm files?", "back": "Chart.yaml (metadata), values.yaml (configuration), templates/ (Kubernetes manifests with Go templating)."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "Which Helm command renders templates locally without deploying?", "options": ["helm install", "helm template", "helm upgrade", "helm lint"], "answer": "helm template"},
        ],
    },
    {
        "id": "301", "slug": "helm-values-templating",
        "title": "Helm Values & Templating",
        "focus": "values.yaml, Go templates, conditional rendering, loops",
        "stage": "foundation", "category": "helm-basics", "level": "junior",
        "tags": ["helm", "kubernetes"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// TemplateContext holds values for Helm template rendering.
type TemplateContext struct {
\tAppName   string
\tReplicas  int
\tImage     string
\tPort      int
\tEnableIngress bool
}

// RenderDeployment creates a Kubernetes Deployment manifest
// with Helm-style templating.
func RenderDeployment(ctx TemplateContext) string {
\treturn ""
}''',
        "kata_test": '''package kata

import (
\t"strings"
\t"testing"
)

func TestRenderDeployment(t *testing.T) {
\tctx := TemplateContext{
\t\tAppName:  "web",
\t\tReplicas: 3,
\t\tImage:    "nginx:latest",
\t\tPort:     80,
\t}
\tresult := RenderDeployment(ctx)
\tif !strings.Contains(result, "kind: Deployment") {
\t\tt.Error("expected Deployment kind")
\t}
\tif !strings.Contains(result, "web") {
\t\tt.Error("expected app name")
\t}
}''',
        "readme": '''# Kata 301 — Helm Values & Templating

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
''',
        "flashcards": [
            {"front": "How does Helm use Go templates?", "back": "Kubernetes manifests in templates/ use {{ .Values.variable }} to inject configuration from values.yaml."},
            {"front": "Why is values.yaml powerful?", "back": "One chart, many configurations. Override values per environment without changing templates."},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "How do you reference a value in a Helm template?", "options": ["${value}", "{{ .Values.value }}", "<< value >>", "%%value%%"], "answer": "{{ .Values.value }}"},
        ],
    },
]

# ═══════════════════════════════════════════════
# SECURITY & CVE TRACK
# ═══════════════════════════════════════════════
SECURITY_KATAS = [
    {
        "id": "400", "slug": "cve-analysis",
        "title": "CVE Analysis & Response",
        "focus": "CVE databases, vulnerability assessment, patching strategy, dependency scanning",
        "stage": "foundation", "category": "security-fundamentals", "level": "junior",
        "tags": ["security", "tooling"],
        "estimated_minutes": 25,
        "kata_go": '''package kata

// CVEInfo represents a parsed CVE entry.
type CVEInfo struct {
\tID          string
\tSeverity    string // "critical", "high", "medium", "low"
\tAffected    string // package name
\tFixedIn     string // version that fixes it
\tDescription string
}

// AnalyzeCVE parses a CVE description and determines severity.
// Your task: implement the function below.
func AnalyzeCVE(id, description string) CVEInfo {
\treturn CVEInfo{ID: id}
}''',
        "kata_test": '''package kata

import "testing"

func TestAnalyzeCVE(t *testing.T) {
\tinfo := AnalyzeCVE("CVE-2024-1234", "Remote code execution in parser")
\tif info.ID != "CVE-2024-1234" {
\t\tt.Errorf("expected ID CVE-2024-1234, got %s", info.ID)
\t}
\tif info.Severity == "" {
\t\tt.Error("expected severity to be set")
\t}
}''',
        "readme": '''# Kata 400 — CVE Analysis & Response

**Focus:** CVE databases, vulnerability assessment, patching strategy, dependency scanning

## Your task
Implement CVE analysis and severity assessment.

### Learning goal
- What you are practicing: understanding CVEs, assessing severity, and planning responses.
- Why this matters: CVEs are discovered daily. Knowing how to assess and respond is a critical developer skill.
- How this grows your skills: you learn security assessment and incident response.

### The CVE Response Workflow
1. **Discover** — `govulncheck ./...` or GitHub Dependabot alerts
2. **Assess** — Is this critical? High? Does it affect us?
3. **Verify** — Are we actually using the vulnerable code path?
4. **Patch** — Update the dependency, or mitigate if no fix exists
5. **Test** — Ensure the patch doesn't break anything
6. **Document** — Record the decision and rationale

## Rules / Expectations
- Parses CVE ID correctly
- Assigns severity based on description
- Returns structured CVE info

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
''',
        "flashcards": [
            {"front": "What is a CVE?", "back": "Common Vulnerabilities and Exposures — a standardized identifier for security vulnerabilities."},
            {"front": "What is the CVE response workflow?", "back": "Discover → Assess → Verify → Patch → Test → Document."},
            {"front": "What tool scans Go dependencies for vulnerabilities?", "back": "govulncheck ./... (from golang.org/x/vuln/cmd/govulncheck)"},
        ],
        "quiz_questions": [
            {"type": "multiple_choice", "question": "What should you do FIRST when a CVE is reported?", "options": ["Panic and update everything", "Assess if you're actually affected", "Ignore it until production breaks", "Rewrite the vulnerable code"], "answer": "Assess if you're actually affected"},
            {"type": "multiple_choice", "question": "Which tool scans Go dependencies for known vulnerabilities?", "options": ["go vet", "govulncheck", "golangci-lint", "go test"], "answer": "govulncheck"},
        ],
    },
]

def create_kata_dirs(katas, track_id):
    """Create kata directories with all content files."""
    created = 0
    for kata in katas:
        kata_dir = KATAS_DIR / f"kata-{kata['id']}-{kata['slug']}"
        if kata_dir.exists():
            print(f"  SKIP {kata_dir.name} (exists)")
            continue

        kata_dir.mkdir(parents=True, exist_ok=True)
        (kata_dir / "kata.go.txt").write_text(kata["kata_go"])
        (kata_dir / "kata_test.go.txt").write_text(kata["kata_test"])
        (kata_dir / "README.md").write_text(kata["readme"])

        meta = {
            "id": kata["id"], "slug": kata["slug"], "title": kata["title"],
            "focus": kata["focus"], "signature": "", "rules": [],
            "evaluator_status": "incomplete",
            "stage": kata["stage"], "category": kata["category"],
            "level": kata["level"], "tags": kata["tags"],
            "prerequisites": [], "estimated_minutes": kata["estimated_minutes"],
            "flashcards": kata["flashcards"], "quiz_questions": kata["quiz_questions"],
        }
        (kata_dir / "kata.json").write_text(json.dumps(meta, indent=2) + "\n")
        print(f"  CREATED {kata_dir.name}")
        created += 1
    return created

def create_track(track_id, title, description, stages):
    """Create a track directory with track.json."""
    track_dir = TRACKS_DIR / track_id
    track_dir.mkdir(parents=True, exist_ok=True)

    track = {
        "id": track_id,
        "title": title,
        "description": description,
        "stages": stages,
    }

    with open(track_dir / "track.json", "w") as f:
        json.dump(track, f, indent=2)
        f.write("\n")

    print(f"  Created track: {track_id}")

def main():
    print("=== GoKatas Curriculum Expansion ===\n")

    # 1. Create AI literacy katas
    print("1. Creating AI literacy katas...")
    created = create_kata_dirs(AI_KATAS, "go-core-100")
    print(f"   Created {created} AI katas\n")

    # 2. Create Terraform track
    print("2. Creating Terraform track...")
    create_track("terraform-100", "Infrastructure as Code: Terraform",
        "Learn Terraform from basics to production-ready infrastructure.",
        [{
            "id": "foundation", "title": "Terraform Foundations", "level": "junior",
            "description": "Learn HCL, core commands, and basic resource management.",
            "categories": [{
                "id": "tf-basics", "title": "Terraform Basics",
                "description": "HCL syntax, init/plan/apply/destroy, modules, and state.",
                "learning_goal": "Write and manage Terraform configurations with confidence.",
                "kata_ids": ["200", "201"],
            }],
        }])
    created = create_kata_dirs(TERRAFORM_KATAS, "terraform-100")
    print(f"   Created {created} Terraform katas\n")

    # 3. Create Helm track
    print("3. Creating Helm track...")
    create_track("helm-100", "Container Orchestration: Helm",
        "Learn Helm charts for Kubernetes application deployment.",
        [{
            "id": "foundation", "title": "Helm Foundations", "level": "junior",
            "description": "Chart structure, templating, values, and deployment lifecycle.",
            "categories": [{
                "id": "helm-basics", "title": "Helm Basics",
                "description": "Chart.yaml, values.yaml, Go templates, and helm commands.",
                "learning_goal": "Create and manage Helm charts for Kubernetes applications.",
                "kata_ids": ["300", "301"],
            }],
        }])
    created = create_kata_dirs(HELM_KATAS, "helm-100")
    print(f"   Created {created} Helm katas\n")

    # 4. Create Security track
    print("4. Creating Security & CVE track...")
    create_track("security-100", "Security & CVE Awareness",
        "Understand vulnerabilities, CVEs, and security-first development.",
        [{
            "id": "foundation", "title": "Security Fundamentals", "level": "junior",
            "description": "CVE analysis, dependency scanning, secure coding basics.",
            "categories": [{
                "id": "security-fundamentals", "title": "Security Fundamentals",
                "description": "CVE databases, vulnerability assessment, patching strategies.",
                "learning_goal": "Assess and respond to security vulnerabilities systematically.",
                "kata_ids": ["400"],
            }],
        }])
    created = create_kata_dirs(SECURITY_KATAS, "security-100")
    print(f"   Created {created} Security katas\n")

    # 5. Update the Go track to include AI literacy category
    print("5. Adding AI literacy to Go Foundation track...")
    track_path = TRACKS_DIR / "go-core-100" / "track.json"
    with open(track_path) as f:
        track = json.load(f)

    # Add ai-literacy category to foundation stage
    foundation = track["stages"][0]
    ai_category = {
        "id": "ai-literacy",
        "title": "AI-Augmented Development",
        "description": "Using AI as a tool without losing your ability to think. Learn when to use it, when not to, and how to verify its output.",
        "learning_goal": "Use AI effectively as a development partner while maintaining critical thinking and deep understanding.",
        "kata_ids": ["160", "161"],
    }
    foundation["categories"].append(ai_category)

    # Add AI debugging kata to practitioner
    practitioner = track["stages"][1]
    ai_mid = {
        "id": "ai-literacy-mid",
        "title": "AI at Scale",
        "description": "Advanced AI collaboration patterns for complex debugging and architecture.",
        "learning_goal": "Use AI as a thinking partner for complex problems without becoming dependent.",
        "kata_ids": ["162"],
    }
    practitioner["categories"].append(ai_mid)

    with open(track_path, "w") as f:
        json.dump(track, f, indent=2)
        f.write("\n")
    print("   Updated go-core-100 track.json\n")

    print("=== Expansion complete ===")
    print("Tracks: go-core-100 (163 katas), terraform-100, helm-100, security-100")

if __name__ == "__main__":
    main()
