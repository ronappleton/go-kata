package diagnostics

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Report struct {
	PodmanAvailable bool
	Rootless        bool
	ImageAvailable  bool
	Message         string
}

func Check(ctx context.Context, podmanPath, image string) Report {
	if strings.TrimSpace(podmanPath) == "" {
		podmanPath = "podman"
	}
	if _, err := exec.LookPath(podmanPath); err != nil {
		return Report{Message: "Podman is not installed."}
	}

	infoCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	info := exec.CommandContext(infoCtx, podmanPath, "info", "--format", "{{.Host.Security.Rootless}}")
	output, err := info.Output()
	if err != nil {
		return Report{Message: fmt.Sprintf("Podman is installed but rootless operation is unavailable: %v", err)}
	}
	rootless := strings.EqualFold(strings.TrimSpace(string(output)), "true")
	report := Report{PodmanAvailable: true, Rootless: rootless}
	if !rootless {
		report.Message = "Podman is available, but rootless mode is not active for this user."
		return report
	}

	image = strings.TrimSpace(image)
	if image == "" {
		report.Message = "Rootless Podman is ready. Configure a digest-pinned runner image."
		return report
	}
	imageCtx, imageCancel := context.WithTimeout(ctx, 3*time.Second)
	defer imageCancel()
	check := exec.CommandContext(imageCtx, podmanPath, "image", "exists", image)
	if err := check.Run(); err != nil {
		report.Message = "Rootless Podman is ready, but the configured runner image is not available."
		return report
	}
	report.ImageAvailable = true
	report.Message = "Podman and the configured runner image are ready."
	return report
}
