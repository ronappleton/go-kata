//go:build integration

package evaluator

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func podmanRunner(t *testing.T) *Runner {
	t.Helper()
	image := strings.TrimSpace(os.Getenv("GOKATAS_RUNNER_IMAGE"))
	if image == "" {
		t.Skip("GOKATAS_RUNNER_IMAGE is not configured")
	}
	runner, err := NewRunner(image)
	if err != nil {
		t.Fatal(err)
	}
	runner.Limits.WallTime = 30 * time.Second
	return runner
}

func TestPodmanEvaluatorPassAndFail(t *testing.T) {
	runner := podmanRunner(t)

	trusted := `package kata

import "testing"

func TestGreeting(t *testing.T) {
	if Greeting() != "hello" {
		t.Fatalf("Greeting mismatch")
	}
}
`
	pass := runner.Run(context.Background(), Request{
		KataID:       "integration-pass",
		Module:       "kataintegration",
		Code:         "package kata\nfunc Greeting() string { return \"hello\" }\n",
		TrustedTests: trusted,
	})
	if !pass.Passed || pass.Status != StatusPassed {
		t.Fatalf("expected pass, got %+v", pass)
	}

	fail := runner.Run(context.Background(), Request{
		KataID:       "integration-fail",
		Module:       "kataintegration",
		Code:         "package kata\nfunc Greeting() string { return \"wrong\" }\n",
		TrustedTests: trusted,
	})
	if fail.Passed || fail.Status != StatusFailed {
		t.Fatalf("expected trusted test failure, got %+v", fail)
	}
}

func TestPodmanEvaluatorSandboxIsolation(t *testing.T) {
	runner := podmanRunner(t)

	code := `package kata

import (
	"net"
	"os"
	"time"
)

// DialProbe reports whether an external network dial succeeded.
func DialProbe() bool {
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 2*time.Second)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

// WriteProbe attempts to write to the read-only container root.
func WriteProbe() error {
	return os.WriteFile("/escape-attempt.txt", []byte("x"), 0o600)
}
`
	trusted := `package kata

import "testing"

func TestNetworkIsBlocked(t *testing.T) {
	if DialProbe() {
		t.Fatal("network access should be blocked")
	}
}

func TestRootFilesystemIsReadOnly(t *testing.T) {
	if err := WriteProbe(); err == nil {
		t.Fatal("container root should be read-only")
	}
}
`
	res := runner.Run(context.Background(), Request{
		KataID:       "integration-isolation",
		Module:       "kataintegration",
		Code:         code,
		TrustedTests: trusted,
	})
	if !res.Passed || res.Status != StatusPassed {
		t.Fatalf("expected sandbox to hold, got %+v", res)
	}
}
