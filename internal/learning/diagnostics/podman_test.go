package diagnostics

import (
	"context"
	"strings"
	"testing"
)

func TestCheckReportsMissingPodman(t *testing.T) {
	report := Check(context.Background(), "/definitely/missing/podman", "")
	if report.PodmanAvailable {
		t.Fatal("expected Podman to be unavailable")
	}
	if !strings.Contains(report.Message, "not installed") {
		t.Fatalf("unexpected message: %q", report.Message)
	}
}
