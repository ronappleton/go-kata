package updater

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	// ReleaseRepo is the GitHub repo to check for releases.
	ReleaseRepo = "ronappleton/go-kata"
	// ReleaseURL is the GitHub API endpoint for latest release.
	ReleaseURL = "https://api.github.com/repos/" + ReleaseRepo + "/releases/latest"
)

// UpdateInfo holds details about an available update.
type UpdateInfo struct {
	Version     string    // e.g. "v1.2.3"
	ReleaseURL  string    // HTML URL for the release page
	BinaryURL   string    // Direct download URL for the binary tarball
	Body        string    // Release notes
	PublishedAt time.Time // When the release was published
}

// Manager handles checking for and applying updates.
type Manager struct {
	currentVersion string
	httpClient     *http.Client
}

// NewManager creates an updater manager for the given current version.
func NewManager(currentVersion string) *Manager {
	return &Manager{
		currentVersion: strings.TrimPrefix(currentVersion, "v"),
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// CurrentVersion returns the running version.
func (m *Manager) CurrentVersion() string {
	return m.currentVersion
}

// CheckForUpdate queries GitHub for the latest release.
// Returns nil if the current version is up-to-date.
func (m *Manager) CheckForUpdate(ctx context.Context) (*UpdateInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ReleaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "GoKatas-Desktop/"+m.currentVersion)

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch releases: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// No releases yet
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var release struct {
		TagName   string    `json:"tag_name"`
		HTMLURL   string    `json:"html_url"`
		Body      string    `json:"body"`
		Published time.Time `json:"published_at"`
		Assets    []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("parse release: %w", err)
	}

	latest := normalizeVersion(release.TagName)
	current := normalizeVersion(m.currentVersion)

	if latest == "" || current == "" {
		return nil, fmt.Errorf("invalid version: latest=%q current=%q", release.TagName, m.currentVersion)
	}

	if !isNewer(latest, current) {
		return nil, nil // Up to date
	}

	// Find the right binary asset for this platform
	binaryURL := findBinaryAsset(release.Assets)
	if binaryURL == "" {
		return nil, fmt.Errorf("no binary asset found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	return &UpdateInfo{
		Version:     release.TagName,
		ReleaseURL:  release.HTMLURL,
		BinaryURL:   binaryURL,
		Body:        release.Body,
		PublishedAt: release.Published,
	}, nil
}

// DownloadAndReplace downloads the update and atomically replaces the running binary.
// It returns the path to the new binary, or an error.
func (m *Manager) DownloadAndReplace(ctx context.Context, info *UpdateInfo) (string, error) {
	// Get the path of the running executable
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("get executable path: %w", err)
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return "", fmt.Errorf("resolve executable symlink: %w", err)
	}

	// Download to a temp file in the same directory (for atomic rename)
	dir := filepath.Dir(exePath)
	tmpFile, err := os.CreateTemp(dir, ".gokatas-update-*.tmp")
	if err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer func() {
		tmpFile.Close()
		os.Remove(tmpPath) // Clean up on failure
	}()

	// Download the tarball
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, info.BinaryURL, nil)
	if err != nil {
		return "", fmt.Errorf("create download request: %w", err)
	}

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download update: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download returned status %d", resp.StatusCode)
	}

	// Extract the binary from the tarball
	if err := extractBinaryFromTarball(resp.Body, tmpFile, runtime.GOOS); err != nil {
		return "", fmt.Errorf("extract binary: %w", err)
	}

	// Make executable
	if err := tmpFile.Chmod(0755); err != nil {
		return "", fmt.Errorf("chmod: %w", err)
	}
	tmpFile.Close()

	// Atomic replace: rename temp -> current
	if err := os.Rename(tmpPath, exePath); err != nil {
		return "", fmt.Errorf("replace binary: %w", err)
	}

	return exePath, nil
}

// Restart replaces the current process with the new binary.
func Restart(newBinaryPath string) error {
	argv := os.Args
	return syscallExec(newBinaryPath, argv, os.Environ())
}

// findBinaryAsset finds the download URL for the current platform's binary.
func findBinaryAsset(assets []struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}) string {
	// Expected naming patterns:
	//   gokatas-VERSION-linux-amd64.tar.gz
	//   gokatas-VERSION-darwin-arm64.tar.gz
	//   gokatas-VERSION-windows-amd64.zip
	suffix := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	for _, a := range assets {
		name := strings.ToLower(a.Name)
		if strings.Contains(name, suffix) && strings.HasSuffix(name, ".tar.gz") {
			return a.BrowserDownloadURL
		}
	}
	// Fallback: try without arch (e.g. "darwin-universal")
	ossuffix := runtime.GOOS
	for _, a := range assets {
		name := strings.ToLower(a.Name)
		if strings.Contains(name, ossuffix) && strings.HasSuffix(name, ".tar.gz") {
			return a.BrowserDownloadURL
		}
	}
	return ""
}

// extractBinaryFromTarball reads a tar.gz stream and writes the first executable to out.
func extractBinaryFromTarball(r io.Reader, out *os.File, goos string) error {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("gzip reader: %w", err)
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar read: %w", err)
		}

		name := filepath.Base(header.Name)

		// Look for the binary (not .txt, not .json, not README)
		if goos == "windows" {
			if !strings.HasSuffix(name, ".exe") {
				continue
			}
		} else {
			// Skip non-binary files
			ext := filepath.Ext(name)
		if ext != "" {
			// Has extension — skip known non-binary files
			if ext == ".txt" || ext == ".json" || ext == ".md" || ext == ".yml" || ext == ".yaml" {
				continue
			}
		}
			// Also skip known non-binary names
			if name == "README" || name == "LICENSE" || name == "manifest.json" {
				continue
			}
		}

		// Skip directories
		if header.Typeflag == tar.TypeDir {
			continue
		}

		// Found a candidate — extract it
		if _, err := io.Copy(out, tr); err != nil {
			return fmt.Errorf("extract %s: %w", name, err)
		}
		return nil
	}
	return fmt.Errorf("no binary found in tarball")
}

// normalizeVersion strips leading "v" and normalizes.
func normalizeVersion(v string) string {
	v = strings.TrimPrefix(v, "v")
	v = strings.TrimSpace(v)
	return v
}

// isNewer returns true if latest > current using simple semver comparison.
func isNewer(latest, current string) bool {
	lParts := splitVersion(latest)
	cParts := splitVersion(current)

	for i := 0; i < 3; i++ {
		if i >= len(lParts) {
			return false
		}
		if i >= len(cParts) {
			return true
		}
		if lParts[i] > cParts[i] {
			return true
		}
		if lParts[i] < cParts[i] {
			return false
		}
	}
	return false
}

func splitVersion(v string) []int {
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	var nums []int
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			break
		}
		nums = append(nums, n)
	}
	return nums
}
