package diagnostics

import (
	"runtime"
	"strings"
	"testing"
)

func TestDetectInstallPlanLinux(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping Linux test on " + runtime.GOOS)
	}
	plan := DetectInstallPlan()
	if plan.Platform != "linux" {
		t.Fatalf("expected linux platform, got %q", plan.Platform)
	}
	// On CI, we might not have a known package manager
	if plan.CanAutoInstall() {
		if len(plan.Commands) == 0 {
			t.Fatal("expected non-empty commands for auto-install")
		}
	}
	if plan.FormatDialogMessage() == "" {
		t.Fatal("expected non-empty dialog message")
	}
}

func TestDetectInstallPlanDarwin(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("skipping Darwin test on " + runtime.GOOS)
	}
	plan := DetectInstallPlan()
	if plan.Platform != "darwin" {
		t.Fatalf("expected darwin platform, got %q", plan.Platform)
	}
	if !strings.Contains(plan.Notes, "podman") {
		t.Fatalf("expected notes to mention podman, got %q", plan.Notes)
	}
}

func TestDetectInstallPlanWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping Windows test on " + runtime.GOOS)
	}
	plan := DetectInstallPlan()
	if plan.Platform != "windows" {
		t.Fatalf("expected windows platform, got %q", plan.Platform)
	}
	if !strings.Contains(plan.Notes, "Podman Desktop") {
		t.Fatalf("expected notes to mention Podman Desktop, got %q", plan.Notes)
	}
}

func TestIsPodmanInstalled(t *testing.T) {
	// Just verify it doesn't panic
	_ = IsPodmanInstalled()
}

func TestRunnerImageExistsEmpty(t *testing.T) {
	if RunnerImageExists("") {
		t.Fatal("expected false for empty image")
	}
}

func TestDetectInstallPlanAlwaysReturnsMessage(t *testing.T) {
	plan := DetectInstallPlan()
	if plan.FormatDialogMessage() == "" {
		t.Fatal("expected non-empty dialog message for any platform")
	}
}
