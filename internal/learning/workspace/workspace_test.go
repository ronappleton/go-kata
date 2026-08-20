package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolvePathsUsesXDGEnvironment(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "/tmp/gokatas-config")
	t.Setenv("XDG_DATA_HOME", "/tmp/gokatas-data")
	t.Setenv("XDG_STATE_HOME", "/tmp/gokatas-state")
	t.Setenv("XDG_CACHE_HOME", "/tmp/gokatas-cache")

	paths, err := ResolvePaths("gokatas")
	if err != nil {
		t.Fatal(err)
	}
	if paths.Data != "/tmp/gokatas-data/gokatas" || paths.State != "/tmp/gokatas-state/gokatas" {
		t.Fatalf("unexpected XDG paths: %+v", paths)
	}
}

func TestWorkspaceRejectsTraversal(t *testing.T) {
	manager := NewManager(Paths{Data: t.TempDir()})
	if _, err := manager.Workspace("../outside"); err == nil {
		t.Fatal("expected traversal ID to be rejected")
	}
}

func TestAtomicWriteReplacesFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "state.json")
	if err := AtomicWrite(path, []byte("first"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := AtomicWrite(path, []byte("second"), 0o600); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(data)) != "second" {
		t.Fatalf("unexpected data %q", data)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("unexpected mode %o", info.Mode().Perm())
	}
}
