package evaluator

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestNewRunnerRequiresDigestPinnedImage(t *testing.T) {
	if _, err := NewRunner("golang:1.24"); err == nil {
		t.Fatal("expected floating image tag to be rejected")
	}

	runner, err := NewRunner("registry.example/gokatas-runner@sha256:" + strings.Repeat("a", 64))
	if err != nil {
		t.Fatalf("expected digest image to be accepted: %v", err)
	}
	if runner.Limits.WallTime != DefaultWallTime {
		t.Fatalf("unexpected default wall time: %s", runner.Limits.WallTime)
	}
}

func TestValidateRequestSeparatesIncompleteEvaluator(t *testing.T) {
	limits := DefaultLimits()
	err := validateRequest(Request{KataID: "001", Code: "package kata"}, limits)
	if err == nil || !strings.Contains(err.Error(), "trusted evaluator") {
		t.Fatalf("expected incomplete evaluator error, got %v", err)
	}
}

func TestPodmanArgsUseIsolationAndReadOnlyInputs(t *testing.T) {
	runner, err := NewRunner("registry.example/gokatas-runner@sha256:" + strings.Repeat("b", 64))
	if err != nil {
		t.Fatal(err)
	}
	args := strings.Join(runner.podmanArgs("run-id", "/tmp/run", "/tmp/trusted", "/tmp/learner"), " ")
	for _, required := range []string{
		"--network=none",
		"--read-only",
		"--cap-drop=ALL",
		"--security-opt=no-new-privileges",
		"--pids-limit",
		"--memory",
		"--cpus",
		"dst=/workspace/kata.go,ro",
		"dst=/workspace/kata_test.go,ro",
		"dst=/workspace/learner_test.go,ro",
		"GOPROXY=off",
		"--pull=never",
	} {
		if !strings.Contains(args, required) {
			t.Errorf("Podman args missing %q: %s", required, args)
		}
	}
}

func TestPodmanArgsOmitsEmptyLearnerTests(t *testing.T) {
	runner, err := NewRunner("registry.example/gokatas-runner@sha256:" + strings.Repeat("c", 64))
	if err != nil {
		t.Fatal(err)
	}
	args := strings.Join(runner.podmanArgs("run-id", "/tmp/run", "/tmp/trusted", ""), " ")
	if strings.Contains(args, "learner_test.go") {
		t.Fatalf("empty learner tests should not mount learner_test.go: %s", args)
	}
}

func TestCappedBufferCancelsAtLimit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buffer := cappedBuffer{limit: 4, cancel: cancel}
	if _, err := buffer.Write([]byte("12345")); err != nil {
		t.Fatal(err)
	}
	if !buffer.exceeded {
		t.Fatal("expected output limit to be recorded")
	}
	if buffer.String() != "1234" {
		t.Fatalf("unexpected capped output %q", buffer.String())
	}
	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
		t.Fatal("expected cancellation")
	}
}

func TestClassifyFailure(t *testing.T) {
	cases := []struct {
		output string
		want   Status
	}{
		{"panic: bad thing", StatusRuntimePanic},
		{"kata [build failed]", StatusCompileError},
		{"--- FAIL: TestThing", StatusFailed},
	}
	for _, tc := range cases {
		if got := classifyFailure(tc.output); got != tc.want {
			t.Fatalf("classifyFailure(%q) = %q, want %q", tc.output, got, tc.want)
		}
	}
}
