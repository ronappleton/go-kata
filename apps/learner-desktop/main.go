package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	defaultDesktopAddr    = "127.0.0.1:17777"
	readinessPath         = "/api/track"
	readinessWaitTimeout  = 25 * time.Second
	shutdownWaitTimeout   = 3 * time.Second
	serverBinaryName      = "learner-studio-server"
	windowsExecutableSuff = ".exe"
)

type launchSpec struct {
	name string
	args []string
}

func main() {
	repoFlag := flag.String("repo", ".", "Path to kata repository root")
	addrFlag := flag.String("addr", defaultDesktopAddr, "Learner Studio bind address")
	serverBinFlag := flag.String("server-bin", "", "Path to learner studio server binary (optional)")
	flag.Parse()

	repoRoot, err := filepath.Abs(*repoFlag)
	if err != nil {
		fatalf("resolve repo root: %v", err)
	}

	serverCmd, err := buildServerCommand(repoRoot, *addrFlag, strings.TrimSpace(*serverBinFlag))
	if err != nil {
		fatalf("prepare server command: %v", err)
	}
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr

	if err := serverCmd.Start(); err != nil {
		fatalf("start learner studio server: %v", err)
	}

	serverExited := make(chan error, 1)
	go func() {
		serverExited <- serverCmd.Wait()
	}()

	baseURL := "http://" + *addrFlag
	if err := waitForReady(baseURL+readinessPath, readinessWaitTimeout); err != nil {
		_ = stopServer(serverCmd, serverExited)
		fatalf("wait for learner studio: %v", err)
	}

	if err := launchDesktopWindow(baseURL); err != nil {
		_ = stopServer(serverCmd, serverExited)
		fatalf("launch desktop window: %v", err)
	}

	fmt.Printf("Learner Studio desktop launched (%s)\n", baseURL)
	fmt.Println("Press Ctrl+C to close.")

	signalCtx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopSignals()

	select {
	case err := <-serverExited:
		if err != nil {
			fatalf("learner studio server exited: %v", err)
		}
	case <-signalCtx.Done():
	}

	if err := stopServer(serverCmd, serverExited); err != nil {
		fatalf("stop learner studio server: %v", err)
	}
}

func buildServerCommand(repoRoot, addr, explicitServerBin string) (*exec.Cmd, error) {
	invocation, err := resolveServerInvocation(repoRoot, addr, explicitServerBin, findBundledServerBinary())
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(invocation.name, invocation.args...)
	if invocation.dir != "" {
		cmd.Dir = invocation.dir
	}
	return cmd, nil
}

type invocation struct {
	name string
	args []string
	dir  string
}

func resolveServerInvocation(repoRoot, addr, explicitServerBin, bundledServerBin string) (invocation, error) {
	if repoRoot == "" {
		return invocation{}, errors.New("repo root is required")
	}
	if addr == "" {
		return invocation{}, errors.New("addr is required")
	}

	if explicitServerBin != "" {
		return invocation{
			name: explicitServerBin,
			args: []string{"-repo", repoRoot, "-addr", addr},
		}, nil
	}

	if bundledServerBin != "" {
		return invocation{
			name: bundledServerBin,
			args: []string{"-repo", repoRoot, "-addr", addr},
		}, nil
	}

	return invocation{
		name: "go",
		args: []string{"run", "./apps/learner-studio", "-repo", repoRoot, "-addr", addr},
		dir:  repoRoot,
	}, nil
}

func findBundledServerBinary() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}

	dir := filepath.Dir(exePath)
	candidates := []string{serverBinaryName}
	if runtime.GOOS == "windows" {
		candidates = append(candidates, serverBinaryName+windowsExecutableSuff)
	}

	for _, candidate := range candidates {
		path := filepath.Join(dir, candidate)
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			continue
		}
		return path
	}

	return ""
}

func waitForReady(url string, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	client := &http.Client{Timeout: 800 * time.Millisecond}
	deadline := time.Now().Add(timeout)
	lastErr := errors.New("studio not ready yet")

	for time.Now().Before(deadline) {
		resp, err := client.Get(url)
		if err == nil {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 500 {
				return nil
			}
			lastErr = fmt.Errorf("unexpected status: %d", resp.StatusCode)
		} else {
			lastErr = err
		}
		time.Sleep(180 * time.Millisecond)
	}

	return fmt.Errorf("timeout waiting for %s: %w", url, lastErr)
}

func launchDesktopWindow(url string) error {
	candidates := launchCandidates(runtime.GOOS, url)
	attempted := make([]string, 0, len(candidates))

	for _, item := range candidates {
		if _, err := exec.LookPath(item.name); err != nil {
			continue
		}

		cmd := exec.Command(item.name, item.args...)
		if err := cmd.Start(); err == nil {
			return nil
		}
		attempted = append(attempted, item.name)
	}

	if len(attempted) == 0 {
		return errors.New("no supported desktop launcher found (install Chrome, Edge, Chromium, or xdg-open/open)")
	}

	return fmt.Errorf("failed to launch desktop window with: %s", strings.Join(attempted, ", "))
}

func launchCandidates(goos, url string) []launchSpec {
	appURL := "--app=" + url

	switch goos {
	case "darwin":
		return []launchSpec{
			{name: "open", args: []string{"-na", "Google Chrome", "--args", appURL, "--new-window"}},
			{name: "open", args: []string{"-na", "Microsoft Edge", "--args", appURL, "--new-window"}},
			{name: "open", args: []string{"-na", "Chromium", "--args", appURL, "--new-window"}},
			{name: "open", args: []string{url}},
		}
	case "windows":
		return []launchSpec{
			{name: "msedge.exe", args: []string{appURL, "--new-window"}},
			{name: "chrome.exe", args: []string{appURL, "--new-window"}},
			{name: "rundll32.exe", args: []string{"url.dll,FileProtocolHandler", url}},
		}
	default:
		return []launchSpec{
			{name: "google-chrome", args: []string{appURL, "--new-window"}},
			{name: "chromium", args: []string{appURL, "--new-window"}},
			{name: "chromium-browser", args: []string{appURL, "--new-window"}},
			{name: "microsoft-edge", args: []string{appURL, "--new-window"}},
			{name: "xdg-open", args: []string{url}},
		}
	}
}

func stopServer(serverCmd *exec.Cmd, exited <-chan error) error {
	if serverCmd == nil || serverCmd.Process == nil {
		return nil
	}

	if serverCmd.ProcessState != nil && serverCmd.ProcessState.Exited() {
		return nil
	}

	_ = serverCmd.Process.Signal(os.Interrupt)
	select {
	case <-exited:
		return nil
	case <-time.After(shutdownWaitTimeout):
	}

	if err := serverCmd.Process.Kill(); err != nil {
		return err
	}
	<-exited
	return nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
