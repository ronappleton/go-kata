package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
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

func main() {
	repoFlag := flag.String("repo", ".", "Path to kata repository root")
	addrFlag := flag.String("addr", defaultDesktopAddr, "Learner Studio bind address")
	serverBinFlag := flag.String("server-bin", "", "Path to learner studio server binary (optional)")
	flag.Parse()

	if !embeddedModeAvailable() {
		fatalf("This binary was built without native webview support.\nRebuild with: go build -tags desktop_webview ./apps/learner-desktop")
	}

	repoRoot, err := filepath.Abs(*repoFlag)
	if err != nil {
		fatalf("resolve repo root: %v", err)
	}

	resolvedAddr, fallbackUsed, err := resolveLaunchAddr(*addrFlag)
	if err != nil {
		fatalf("resolve desktop address: %v", err)
	}
	if fallbackUsed {
		fmt.Printf("Requested address %q is unavailable. Using %q instead.\n", *addrFlag, resolvedAddr)
	}

	serverCmd, err := buildServerCommand(repoRoot, resolvedAddr, strings.TrimSpace(*serverBinFlag))
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

	baseURL := "http://" + resolvedAddr
	if err := waitForReady(baseURL+readinessPath, readinessWaitTimeout); err != nil {
		_ = stopServer(serverCmd, serverExited)
		fatalf("wait for learner studio: %v", err)
	}

	fmt.Printf("Learner Studio desktop launching (%s)\n", baseURL)
	if err := openEmbeddedWindow(baseURL); err != nil {
		_ = stopServer(serverCmd, serverExited)
		fatalf("launch embedded webview: %v", err)
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
	prepareCommand(cmd)
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
			if resp.StatusCode == http.StatusOK {
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

func stopServer(serverCmd *exec.Cmd, exited <-chan error) error {
	if serverCmd == nil || serverCmd.Process == nil {
		return nil
	}

	if serverCmd.ProcessState != nil && serverCmd.ProcessState.Exited() {
		return nil
	}

	_ = sendInterrupt(serverCmd)
	select {
	case <-exited:
		return nil
	case <-time.After(shutdownWaitTimeout):
	}

	if err := forceKill(serverCmd); err != nil && !errors.Is(err, os.ErrProcessDone) && !isNoSuchProcess(err) {
		return err
	}
	<-exited
	return nil
}

func resolveLaunchAddr(requested string) (resolved string, fallbackUsed bool, err error) {
	addr := strings.TrimSpace(requested)
	if addr == "" {
		return "", false, errors.New("address is required")
	}
	if _, _, splitErr := net.SplitHostPort(addr); splitErr != nil {
		return "", false, fmt.Errorf("address must include host and port: %w", splitErr)
	}

	if isAddrAvailable(addr) {
		return addr, false, nil
	}

	host, _, _ := net.SplitHostPort(addr)
	if host == "" {
		host = "127.0.0.1"
	}
	fallback, err := findFreeAddr(host)
	if err != nil {
		return "", false, err
	}
	return fallback, true, nil
}

func isAddrAvailable(addr string) bool {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	_ = listener.Close()
	return true
}

func findFreeAddr(host string) (string, error) {
	listener, err := net.Listen("tcp", net.JoinHostPort(host, "0"))
	if err != nil {
		return "", fmt.Errorf("allocate fallback port on %s: %w", host, err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	return net.JoinHostPort(host, strconv.Itoa(port)), nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
