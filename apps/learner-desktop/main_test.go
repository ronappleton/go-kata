package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestResolveServerInvocation(t *testing.T) {
	t.Run("explicit server binary wins", func(t *testing.T) {
		inv, err := resolveServerInvocation("/repo", "127.0.0.1:17777", "/tmp/studio-server", "/tmp/bundled")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if inv.name != "/tmp/studio-server" {
			t.Fatalf("expected explicit binary, got %q", inv.name)
		}
		if inv.dir != "" {
			t.Fatalf("expected empty dir, got %q", inv.dir)
		}
	})

	t.Run("bundled server binary is second choice", func(t *testing.T) {
		inv, err := resolveServerInvocation("/repo", "127.0.0.1:17777", "", "/tmp/bundled")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if inv.name != "/tmp/bundled" {
			t.Fatalf("expected bundled binary, got %q", inv.name)
		}
	})

	t.Run("fallback is go run command", func(t *testing.T) {
		inv, err := resolveServerInvocation("/repo", "127.0.0.1:17777", "", "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if inv.name != "go" {
			t.Fatalf("expected go command, got %q", inv.name)
		}
		if inv.dir != "/repo" {
			t.Fatalf("expected repo dir, got %q", inv.dir)
		}
		joined := strings.Join(inv.args, " ")
		if !strings.Contains(joined, "run ./apps/learner-studio") {
			t.Fatalf("unexpected args: %v", inv.args)
		}
	})
}

func TestParseLaunchMode(t *testing.T) {
	mode, err := parseLaunchMode("auto")
	if err != nil || mode != launchModeAuto {
		t.Fatalf("expected auto mode, got mode=%q err=%v", mode, err)
	}

	mode, err = parseLaunchMode("EXTERNAL")
	if err != nil || mode != launchModeExternal {
		t.Fatalf("expected external mode, got mode=%q err=%v", mode, err)
	}

	if _, err := parseLaunchMode("unknown"); err == nil {
		t.Fatalf("expected error for unsupported mode")
	}
}

func TestShouldUseEmbeddedMode(t *testing.T) {
	useEmbedded, err := shouldUseEmbeddedMode(launchModeAuto)
	if err != nil {
		t.Fatalf("unexpected error for auto mode: %v", err)
	}
	if useEmbedded {
		t.Fatalf("expected auto to be false in non-webview build")
	}

	useEmbedded, err = shouldUseEmbeddedMode(launchModeExternal)
	if err != nil {
		t.Fatalf("unexpected error for external mode: %v", err)
	}
	if useEmbedded {
		t.Fatalf("expected external mode to disable embedded launch")
	}

	if _, err := shouldUseEmbeddedMode(launchModeEmbedded); err == nil {
		t.Fatalf("expected error for embedded mode without native webview support")
	}
}

func TestLaunchCandidates(t *testing.T) {
	url := "http://127.0.0.1:17777"

	linux := launchCandidates("linux", url)
	if len(linux) == 0 {
		t.Fatalf("expected linux candidates")
	}
	if !strings.Contains(strings.Join(linux[0].args, " "), "--app=") {
		t.Fatalf("expected app mode for linux, got %v", linux[0].args)
	}

	windows := launchCandidates("windows", url)
	if len(windows) < 2 {
		t.Fatalf("expected windows candidates")
	}
	if !strings.Contains(strings.Join(windows[0].args, " "), "--app=") {
		t.Fatalf("expected app mode for windows, got %v", windows[0].args)
	}

	darwin := launchCandidates("darwin", url)
	if len(darwin) < 2 {
		t.Fatalf("expected darwin candidates")
	}
	if darwin[0].name != "open" {
		t.Fatalf("expected open launcher for darwin, got %q", darwin[0].name)
	}
}

func TestWaitForReadyTimeout(t *testing.T) {
	start := time.Now()
	err := waitForReady("http://127.0.0.1:1/api/track", 350*time.Millisecond)
	if err == nil {
		t.Fatalf("expected timeout error")
	}
	if time.Since(start) < 300*time.Millisecond {
		t.Fatalf("waitForReady returned too early")
	}
}

func TestWaitForReadyRequiresHTTP200(t *testing.T) {
	listener := mustListenForTest(t, "tcp4", "127.0.0.1:0")

	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})}
	go func() {
		_ = server.Serve(listener)
	}()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
	}()

	err := waitForReady("http://"+listener.Addr().String(), 350*time.Millisecond)
	if err == nil {
		t.Fatalf("expected timeout when readiness endpoint does not return 200")
	}
	if !strings.Contains(err.Error(), "unexpected status") {
		t.Fatalf("expected unexpected status error detail, got %v", err)
	}
}

func TestResolveLaunchAddr(t *testing.T) {
	t.Run("uses requested address when available", func(t *testing.T) {
		freeListener := mustListenForTest(t, "tcp", "127.0.0.1:0")
		freeAddr := freeListener.Addr().String()
		_ = freeListener.Close()

		resolved, fallbackUsed, err := resolveLaunchAddr(freeAddr)
		if err != nil {
			t.Fatalf("resolve launch addr: %v", err)
		}
		if fallbackUsed {
			t.Fatalf("did not expect fallback for available address")
		}
		if resolved != freeAddr {
			t.Fatalf("expected same address, got %q (wanted %q)", resolved, freeAddr)
		}
	})

	t.Run("falls back when requested address is busy", func(t *testing.T) {
		busyListener := mustListenForTest(t, "tcp", "127.0.0.1:0")
		defer busyListener.Close()
		busyAddr := busyListener.Addr().String()

		resolved, fallbackUsed, err := resolveLaunchAddr(busyAddr)
		if err != nil {
			t.Fatalf("resolve launch addr fallback: %v", err)
		}
		if !fallbackUsed {
			t.Fatalf("expected fallback to be used")
		}
		if resolved == busyAddr {
			t.Fatalf("expected fallback address different from busy address %q", busyAddr)
		}

		checkListener := mustListenForTest(t, "tcp", resolved)
		_ = checkListener.Close()
	})

	t.Run("errors when host or port missing", func(t *testing.T) {
		if _, _, err := resolveLaunchAddr("127.0.0.1"); err == nil {
			t.Fatalf("expected error for malformed address")
		}
	})
}

func mustListenForTest(t *testing.T, network, addr string) net.Listener {
	t.Helper()

	listener, err := net.Listen(network, addr)
	if err == nil {
		return listener
	}

	if errors.Is(err, os.ErrPermission) || strings.Contains(strings.ToLower(err.Error()), "operation not permitted") {
		t.Skipf("skipping socket test; listen not permitted in this environment: %v", err)
	}

	t.Fatalf("listen %s %s: %v", network, addr, err)
	return nil
}
