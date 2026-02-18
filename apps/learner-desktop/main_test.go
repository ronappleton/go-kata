package main

import (
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
