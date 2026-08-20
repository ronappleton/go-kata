package evaluator

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	DefaultCodeLimitBytes      int64 = 256 * 1024
	DefaultTestsLimitBytes     int64 = 256 * 1024
	DefaultOutputLimitBytes    int64 = 1024 * 1024
	DefaultWorkspaceLimitBytes int64 = 64 * 1024 * 1024
	DefaultWallTime                  = 30 * time.Second
	DefaultMemory                    = "512m"
	DefaultCPUs                      = "1"
	DefaultPIDs                int64 = 128
)

type Limits struct {
	CodeBytes      int64
	TestsBytes     int64
	OutputBytes    int64
	WorkspaceBytes int64
	WallTime       time.Duration
	Memory         string
	CPUs           string
	PIDs           int64
}

func DefaultLimits() Limits {
	return Limits{
		CodeBytes:      DefaultCodeLimitBytes,
		TestsBytes:     DefaultTestsLimitBytes,
		OutputBytes:    DefaultOutputLimitBytes,
		WorkspaceBytes: DefaultWorkspaceLimitBytes,
		WallTime:       DefaultWallTime,
		Memory:         DefaultMemory,
		CPUs:           DefaultCPUs,
		PIDs:           DefaultPIDs,
	}
}

type Request struct {
	KataID       string
	Module       string
	Code         string
	LearnerTests string
	TrustedTests string
}

type Status string

const (
	StatusPassed              Status = "passed"
	StatusFailed              Status = "failed"
	StatusCompileError        Status = "compile-error"
	StatusRuntimePanic        Status = "runtime-panic"
	StatusTimeout             Status = "timeout"
	StatusOutputLimit         Status = "output-limit"
	StatusSandboxError        Status = "sandbox-error"
	StatusPodmanUnavailable   Status = "podman-unavailable"
	StatusEvaluatorIncomplete Status = "evaluator-incomplete"
	StatusCancelled           Status = "cancelled"
)

type Result struct {
	RunID          string
	Status         Status
	Passed         bool
	Duration       time.Duration
	FailedTests    []string
	Output         string
	EvaluatorError string
}

type Runner struct {
	Image      string
	PodmanPath string
	Limits     Limits
	runCounter atomic.Uint64
}

func NewRunner(image string) (*Runner, error) {
	image = strings.TrimSpace(image)
	if image == "" {
		return nil, errors.New("runner image is required")
	}
	if !strings.Contains(image, "@sha256:") {
		return nil, fmt.Errorf("runner image must use an immutable digest: %q", image)
	}

	limits := DefaultLimits()
	if err := validateLimits(limits); err != nil {
		return nil, err
	}
	return &Runner{Image: image, PodmanPath: "podman", Limits: limits}, nil
}

func (r *Runner) Run(ctx context.Context, req Request) (result Result) {
	started := time.Now()
	result = Result{RunID: r.nextRunID(), Status: StatusSandboxError}
	defer func() { result.Duration = time.Since(started) }()

	if err := validateRequest(req, r.Limits); err != nil {
		result.EvaluatorError = err.Error()
		if errors.Is(err, errIncompleteEvaluator) {
			result.Status = StatusEvaluatorIncomplete
		}
		return result
	}
	if r.PodmanPath == "" {
		r.PodmanPath = "podman"
	}
	if _, err := exec.LookPath(r.PodmanPath); err != nil {
		result.Status = StatusPodmanUnavailable
		result.EvaluatorError = fmt.Sprintf("podman is unavailable: %v", err)
		return result
	}

	workspace, err := os.MkdirTemp("", "gokatas-run-")
	if err != nil {
		result.EvaluatorError = fmt.Sprintf("create run workspace: %v", err)
		return result
	}
	defer os.RemoveAll(workspace)
	if err := os.Chmod(workspace, 0o700); err != nil {
		result.EvaluatorError = fmt.Sprintf("secure run workspace: %v", err)
		return result
	}

	trustedPath := filepath.Join(workspace, ".trusted-kata_test.go")
	if err := os.WriteFile(trustedPath, []byte(req.TrustedTests), 0o444); err != nil {
		result.EvaluatorError = fmt.Sprintf("write trusted evaluator: %v", err)
		return result
	}
	// These files are bind-mounted read-only into the container and live in a
	// disposable 0o700 workspace. Use world-readable modes because the container
	// process may run as a remapped UID that is not the file owner.
	if err := os.WriteFile(filepath.Join(workspace, "kata.go"), []byte(req.Code), 0o444); err != nil {
		result.EvaluatorError = fmt.Sprintf("write learner code: %v", err)
		return result
	}
	learnerTestsPath := ""
	if strings.TrimSpace(req.LearnerTests) != "" {
		learnerTestsPath = filepath.Join(workspace, "learner_test.go")
		if err := os.WriteFile(learnerTestsPath, []byte(req.LearnerTests), 0o444); err != nil {
			result.EvaluatorError = fmt.Sprintf("write learner tests: %v", err)
			return result
		}
	}
	module := req.Module
	if module == "" {
		module = "kata"
	}
	goMod := "module " + module + "\n\ngo 1.22\n"
	if err := os.WriteFile(filepath.Join(workspace, "go.mod"), []byte(goMod), 0o444); err != nil {
		result.EvaluatorError = fmt.Sprintf("write module metadata: %v", err)
		return result
	}

	runCtx, cancel := context.WithTimeout(ctx, r.Limits.WallTime)
	defer cancel()
	args := r.podmanArgs(result.RunID, workspace, trustedPath, learnerTestsPath)
	cmd := exec.CommandContext(runCtx, r.PodmanPath, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	var output cappedBuffer
	output.limit = r.Limits.OutputBytes
	output.cancel = cancel
	cmd.Stdout = &output
	cmd.Stderr = &output

	cmdErr := cmd.Run()
	result.Output = output.String()
	result.FailedTests = parseFailedTests(result.Output)
	defer r.removeContainer(result.RunID)

	if runCtx.Err() == context.DeadlineExceeded {
		result.Status = StatusTimeout
		return result
	}
	if ctx.Err() != nil {
		result.Status = StatusCancelled
		return result
	}
	if output.exceeded {
		result.Status = StatusOutputLimit
		return result
	}
	if cmdErr != nil {
		result.Status = classifyFailure(result.Output)
		result.EvaluatorError = cmdErr.Error()
		return result
	}
	if len(result.FailedTests) > 0 || strings.Contains(result.Output, `"Action":"fail","Package"`) {
		result.Status = StatusFailed
		return result
	}

	result.Status = StatusPassed
	result.Passed = true
	return result
}

func (r *Runner) podmanArgs(runID, workspace, trustedPath, learnerTestsPath string) []string {
	args := []string{
		"run", "--rm",
		"--name", runID,
		"--pull=never",
		"--network=none",
		"--read-only",
		"--cap-drop=ALL",
		"--security-opt=no-new-privileges",
		"--userns=keep-id",
		"--memory", r.Limits.Memory,
		"--cpus", r.Limits.CPUs,
		"--pids-limit", strconv.FormatInt(r.Limits.PIDs, 10),
		"--ulimit", "core=0",
		"--ulimit", "nofile=256:256",
		"--tmpfs", "/tmp:rw,nosuid,nodev,size=128m",
		"--tmpfs", "/workspace:rw,nosuid,nodev,size=" + strconv.FormatInt(r.Limits.WorkspaceBytes, 10),
		"--mount", "type=bind,src=" + filepath.Join(workspace, "kata.go") + ",dst=/workspace/kata.go,ro",
		"--mount", "type=bind,src=" + filepath.Join(workspace, "go.mod") + ",dst=/workspace/go.mod,ro",
		"--mount", "type=bind,src=" + trustedPath + ",dst=/workspace/kata_test.go,ro",
	}
	if learnerTestsPath != "" {
		args = append(args, "--mount", "type=bind,src="+learnerTestsPath+",dst=/workspace/learner_test.go,ro")
	}
	args = append(args,
		"--workdir", "/workspace",
		"--env", "HOME=/tmp",
		"--env", "GOCACHE=/tmp/go-cache",
		"--env", "GOMODCACHE=/tmp/go-mod-cache",
		"--env", "GOTOOLCHAIN=local",
		"--env", "GOPROXY=off",
		"--env", "GOSUMDB=off",
		"--env", "GOFLAGS=-mod=readonly",
		"--entrypoint", "go",
		r.Image,
		"test", "-vet=off", "-json", "./...",
	)
	return args
}

func (r *Runner) removeContainer(runID string) {
	if r.PodmanPath == "" {
		return
	}
	cmd := exec.Command(r.PodmanPath, "rm", "--force", runID)
	_ = cmd.Run()
}

func (r *Runner) nextRunID() string {
	counter := r.runCounter.Add(1)
	var random [4]byte
	_, _ = rand.Read(random[:])
	return fmt.Sprintf("gokatas-%d-%x-%d", os.Getpid(), random, counter)
}

func validateRequest(req Request, limits Limits) error {
	if strings.TrimSpace(req.KataID) == "" {
		return errors.New("kata id is required")
	}
	if strings.TrimSpace(req.Code) == "" {
		return errors.New("learner code is required")
	}
	if strings.TrimSpace(req.TrustedTests) == "" || isIncompleteEvaluator(req.TrustedTests) {
		return errIncompleteEvaluator
	}
	if int64(len(req.Code)) > limits.CodeBytes {
		return fmt.Errorf("learner code exceeds %d bytes", limits.CodeBytes)
	}
	if int64(len(req.LearnerTests)) > limits.TestsBytes {
		return fmt.Errorf("learner tests exceed %d bytes", limits.TestsBytes)
	}
	if int64(len(req.TrustedTests)) > limits.TestsBytes {
		return fmt.Errorf("trusted tests exceed %d bytes", limits.TestsBytes)
	}
	return validateLimits(limits)
}

var errIncompleteEvaluator = errors.New("trusted evaluator is incomplete")

func isIncompleteEvaluator(source string) bool {
	text := strings.TrimSpace(source)
	if text == "" || !strings.Contains(text, "t.Skip(") {
		return false
	}
	for _, marker := range []string{"t.Run(", "t.Error(", "t.Fatal(", "t.Fail(", "t.Errorf(", "t.Fatalf("} {
		if strings.Contains(text, marker) {
			return false
		}
	}
	return true
}

func validateLimits(limits Limits) error {
	if limits.CodeBytes <= 0 || limits.TestsBytes <= 0 || limits.OutputBytes <= 0 || limits.WorkspaceBytes <= 0 {
		return errors.New("evaluator limits must be positive")
	}
	if limits.WallTime <= 0 || limits.WallTime > 10*time.Minute {
		return errors.New("evaluator wall time must be between 1 second and 10 minutes")
	}
	if limits.Memory == "" || limits.CPUs == "" || limits.PIDs <= 0 {
		return errors.New("evaluator resource limits are incomplete")
	}
	return nil
}

type cappedBuffer struct {
	bytes.Buffer
	limit    int64
	exceeded bool
	cancel   context.CancelFunc
}

func (b *cappedBuffer) Write(p []byte) (int, error) {
	remaining := b.limit - int64(b.Len())
	if remaining <= 0 {
		b.exceeded = true
		if b.cancel != nil {
			b.cancel()
		}
		return len(p), nil
	}
	if int64(len(p)) > remaining {
		_, _ = b.Buffer.Write(p[:remaining])
		b.exceeded = true
		if b.cancel != nil {
			b.cancel()
		}
		return len(p), nil
	}
	return b.Buffer.Write(p)
}

func classifyFailure(output string) Status {
	lower := strings.ToLower(output)
	switch {
	case strings.Contains(lower, "panic:"):
		return StatusRuntimePanic
	case strings.Contains(lower, "build failed"), strings.Contains(lower, "cannot find package"), strings.Contains(lower, "undefined:"):
		return StatusCompileError
	default:
		return StatusFailed
	}
}

func parseFailedTests(output string) []string {
	lines := strings.Split(output, "\n")
	seen := make(map[string]bool)
	failed := make([]string, 0, 8)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.Contains(line, `"Action":"fail"`) || !strings.Contains(line, `"Test":`) {
			continue
		}
		const marker = `"Test":"`
		start := strings.Index(line, marker)
		if start < 0 {
			continue
		}
		start += len(marker)
		end := strings.IndexByte(line[start:], '"')
		if end < 0 {
			continue
		}
		name := line[start : start+end]
		if !seen[name] {
			seen[name] = true
			failed = append(failed, name)
		}
	}
	return failed
}
