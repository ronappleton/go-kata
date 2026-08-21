package diagnostics

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// InstallPlan describes how to install or configure podman on this platform.
type InstallPlan struct {
	Platform     string   // "linux", "darwin", "windows"
	PackageMgr   string   // "apt", "dnf", "yum", "brew", "unknown"
	NeedsSudo    bool
	Commands     []string // ordered install/configure commands to run
	Notes        string   // extra guidance for the user
}

// DetectInstallPlan determines the platform-specific install commands for podman.
func DetectInstallPlan() InstallPlan {
	switch runtime.GOOS {
	case "linux":
		return detectLinux()
	case "darwin":
		return detectDarwin()
	case "windows":
		return detectWindows()
	default:
		return InstallPlan{Platform: runtime.GOOS, Notes: "Please install Podman manually for your platform."}
	}
}

func detectLinux() InstallPlan {
	mgr := detectLinuxPkgMgr()
	switch mgr {
	case "apt":
		return InstallPlan{
			Platform:   "linux",
			PackageMgr: "apt",
			NeedsSudo:  true,
			Commands:   []string{"apt-get update", "apt-get install -y podman"},
			Notes:      "Podman will be installed from the default Ubuntu/Debian repositories.",
		}
	case "dnf":
		return InstallPlan{
			Platform:   "linux",
			PackageMgr: "dnf",
			NeedsSudo:  true,
			Commands:   []string{"dnf install -y podman"},
			Notes:      "Podman will be installed from Fedora/RHEL repositories.",
		}
	case "yum":
		return InstallPlan{
			Platform:   "linux",
			PackageMgr: "yum",
			NeedsSudo:  true,
			Commands:   []string{"yum install -y podman"},
			Notes:      "Podman will be installed from CentOS/RHEL repositories.",
		}
	}
	return InstallPlan{
		Platform: "linux",
		Notes:    "Could not detect package manager. Install Podman manually: https://podman.io/getting-started/installation",
	}
}

func detectDarwin() InstallPlan {
	if _, err := exec.LookPath("brew"); err == nil {
		return InstallPlan{
			Platform:   "darwin",
			PackageMgr: "brew",
			NeedsSudo:  false,
			Commands:   []string{"brew install podman", "podman machine init", "podman machine start"},
			Notes:      "After installation, podman needs a macOS machine to run containers.",
		}
	}
	return InstallPlan{
		Platform: "darwin",
		Notes:    "Install Homebrew first: https://brew.sh, then run brew install podman.\nAlternatively, download Podman Desktop: https://podman-desktop.io",
	}
}

func detectWindows() InstallPlan {
	return InstallPlan{
		Platform: "windows",
		Notes:    "Download Podman Desktop from https://podman-desktop.io and follow the installer.",
	}
}

func detectLinuxPkgMgr() string {
	for _, mgr := range []string{"apt", "dnf", "yum"} {
		if _, err := exec.LookPath(mgr); err == nil {
			return mgr
		}
	}
	return "unknown"
}

// CanAutoInstall returns true if the platform supports unattended installation.
func (p InstallPlan) CanAutoInstall() bool {
	return len(p.Commands) > 0 && p.PackageMgr != "unknown"
}

// FormatDialogMessage returns a user-friendly message for the install dialog.
func (p InstallPlan) FormatDialogMessage() string {
	var sb strings.Builder
	sb.WriteString("Podman is required to run kata tests in a sandbox.\n\n")
	if len(p.Commands) > 0 {
		sb.WriteString(fmt.Sprintf("The following commands will be run:\n"))
		for _, cmd := range p.Commands {
			if p.NeedsSudo {
				sb.WriteString(fmt.Sprintf("  <b>sudo</b> %s\n", cmd))
			} else {
				sb.WriteString(fmt.Sprintf("  <b>%s</b>\n", cmd))
			}
		}
	}
	if p.Notes != "" {
		sb.WriteString("\n" + p.Notes)
	}
	return sb.String()
}

// IsPodmanInstalled checks if podman is on PATH.
func IsPodmanInstalled() bool {
	_, err := exec.LookPath("podman")
	return err == nil
}

// RunnerImageExists checks if the digest-pinned image exists locally.
func RunnerImageExists(image string) bool {
	image = strings.TrimSpace(image)
	if image == "" {
		return false
	}
	cmd := exec.Command("podman", "image", "exists", image)
	return cmd.Run() == nil
}

// RunnerImagePullAttempt tries to pull the runner image and returns success/error.
func RunnerImagePullAttempt(image string) error {
	image = strings.TrimSpace(image)
	if image == "" {
		return fmt.Errorf("no runner image configured")
	}
	cmd := exec.Command("podman", "pull", image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
