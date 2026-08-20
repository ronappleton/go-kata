package workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var safeID = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]*$`)

type Paths struct {
	Config string
	Data   string
	State  string
	Cache  string
}

func ResolvePaths(appName string) (Paths, error) {
	appName = strings.Trim(strings.TrimSpace(appName), "/")
	if appName == "" || strings.Contains(appName, string(filepath.Separator)) {
		return Paths{}, errors.New("invalid application name")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return Paths{}, fmt.Errorf("resolve home directory: %w", err)
	}
	configHome := valueOrEnv("XDG_CONFIG_HOME", filepath.Join(home, ".config"))
	dataHome := valueOrEnv("XDG_DATA_HOME", filepath.Join(home, ".local", "share"))
	stateHome := valueOrEnv("XDG_STATE_HOME", filepath.Join(home, ".local", "state"))
	cacheHome := valueOrEnv("XDG_CACHE_HOME", filepath.Join(home, ".cache"))

	return Paths{
		Config: filepath.Join(configHome, appName),
		Data:   filepath.Join(dataHome, appName),
		State:  filepath.Join(stateHome, appName),
		Cache:  filepath.Join(cacheHome, appName),
	}, nil
}

func valueOrEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" || !filepath.IsAbs(value) {
		return fallback
	}
	return value
}

type Manager struct {
	Paths Paths
}

func NewManager(paths Paths) *Manager {
	return &Manager{Paths: paths}
}

func (m *Manager) Ensure() error {
	for _, path := range []string{m.Paths.Config, m.Paths.Data, m.Paths.State, m.Paths.Cache} {
		if err := os.MkdirAll(path, 0o700); err != nil {
			return fmt.Errorf("create application directory %q: %w", path, err)
		}
	}
	return nil
}

func (m *Manager) Workspace(kataID string) (string, error) {
	if !safeID.MatchString(kataID) {
		return "", fmt.Errorf("invalid kata id %q", kataID)
	}
	root := filepath.Join(m.Paths.Data, "workspaces", kataID)
	if err := os.MkdirAll(root, 0o700); err != nil {
		return "", fmt.Errorf("create kata workspace: %w", err)
	}
	return root, nil
}

func (m *Manager) ReadSolution(kataID string) (string, error) {
	root, err := m.Workspace(kataID)
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath.Join(root, "solution.go"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	return string(data), nil
}

func (m *Manager) ReadLearnerTests(kataID string) (string, error) {
	root, err := m.Workspace(kataID)
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath.Join(root, "learner_test.go"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	return string(data), nil
}

func (m *Manager) SaveSolution(kataID, source string, maxBytes int64) error {
	if int64(len(source)) > maxBytes {
		return fmt.Errorf("solution exceeds %d bytes", maxBytes)
	}
	root, err := m.Workspace(kataID)
	if err != nil {
		return err
	}
	return AtomicWrite(filepath.Join(root, "solution.go"), []byte(source), 0o600)
}

func (m *Manager) SaveLearnerTests(kataID, source string, maxBytes int64) error {
	if int64(len(source)) > maxBytes {
		return fmt.Errorf("learner tests exceed %d bytes", maxBytes)
	}
	root, err := m.Workspace(kataID)
	if err != nil {
		return err
	}
	return AtomicWrite(filepath.Join(root, "learner_test.go"), []byte(source), 0o600)
}

func AtomicWrite(path string, data []byte, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".atomic-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if err := tmp.Chmod(mode); err != nil {
		_ = tmp.Close()
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}

func AtomicJSONWrite(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return AtomicWrite(path, data, 0o600)
}
