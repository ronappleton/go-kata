package updater

import (
	"archive/tar"
	"compress/gzip"
	"bytes"
	"os"
	"testing"
)

func TestNormalizeVersion(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"v1.2.3", "1.2.3"},
		{"1.2.3", "1.2.3"},
		{"v0.1.0", "0.1.0"},
		{"  v2.0.0  ", "v2.0.0"},
		{"", ""},
	}
	for _, tt := range tests {
		got := normalizeVersion(tt.input)
		if got != tt.want {
			t.Errorf("normalizeVersion(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestIsNewer(t *testing.T) {
	tests := []struct {
		latest, current string
		want            bool
	}{
		{"2.0.0", "1.0.0", true},
		{"1.1.0", "1.0.0", true},
		{"1.0.1", "1.0.0", true},
		{"1.0.0", "1.0.0", false},
		{"1.0.0", "2.0.0", false},
		{"1.0.0", "1.1.0", false},
		{"1.0.0", "1.0.1", false},
		{"10.0.0", "9.99.99", true},
		{"0.1.0", "0.0.9", true},
		{"", "", false},
		{"1.0", "1.0.0", false},
		{"1.0.0", "1.0", true},
	}
	for _, tt := range tests {
		got := isNewer(tt.latest, tt.current)
		if got != tt.want {
			t.Errorf("isNewer(%q, %q) = %v, want %v", tt.latest, tt.current, got, tt.want)
		}
	}
}

func TestSplitVersion(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"1.2.3", []int{1, 2, 3}},
		{"v1.2.3", []int{1, 2, 3}},
		{"0.1.0", []int{0, 1, 0}},
		{"1.2.3.4", []int{1, 2, 3, 4}},
		{"abc", nil},
		{"", nil},
	}
	for _, tt := range tests {
		got := splitVersion(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("splitVersion(%q) = %v (len %d), want %v", tt.input, got, len(got), tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("splitVersion(%q)[%d] = %d, want %d", tt.input, i, got[i], tt.want[i])
			}
		}
	}
}

func TestFindBinaryAsset(t *testing.T) {
	tests := []struct {
		name    string
		assets  []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		}
		wantURL string
	}{
		{
			name: "finds linux-amd64",
			assets: []struct {
				Name               string `json:"name"`
				BrowserDownloadURL string `json:"browser_download_url"`
			}{
				{Name: "gokatas-1.0.0-darwin-arm64.tar.gz", BrowserDownloadURL: "https://example.com/darwin"},
				{Name: "gokatas-1.0.0-linux-amd64.tar.gz", BrowserDownloadURL: "https://example.com/linux"},
			},
			wantURL: "https://example.com/linux",
		},
		{
			name: "no matching asset",
			assets: []struct {
				Name               string `json:"name"`
				BrowserDownloadURL string `json:"browser_download_url"`
			}{
				{Name: "gokatas-1.0.0-darwin-arm64.tar.gz", BrowserDownloadURL: "https://example.com/darwin"},
			},
			wantURL: "",
		},
		{
			name: "empty assets",
			assets: []struct {
				Name               string `json:"name"`
				BrowserDownloadURL string `json:"browser_download_url"`
			}{},
			wantURL: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findBinaryAsset(tt.assets)
			if got != tt.wantURL {
				t.Errorf("findBinaryAsset() = %q, want %q", got, tt.wantURL)
			}
		})
	}
}

func TestExtractBinaryFromTarball(t *testing.T) {
	// Create a tar.gz with a fake binary
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	content := []byte("#!/bin/sh\necho hello")
	tw.WriteHeader(&tar.Header{
		Name: "gokatas",
		Mode: 0755,
		Size: int64(len(content)),
	})
	tw.Write(content)
	tw.Close()
	gw.Close()

	// Extract to a real temp file
	tmpFile, err := os.CreateTemp(t.TempDir(), "test-binary-*")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = extractBinaryFromTarball(&buf, tmpFile, "linux")
	if err != nil {
		t.Fatalf("extractBinaryFromTarball: %v", err)
	}

	// Check the file has content
	stat, err := tmpFile.Stat()
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if stat.Size() == 0 {
		t.Error("extracted empty binary")
	}
}

func TestExtractBinaryRejectsNonBinary(t *testing.T) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	content := []byte("This is a readme")
	tw.WriteHeader(&tar.Header{
		Name: "README.txt",
		Mode: 0644,
		Size: int64(len(content)),
	})
	tw.Write(content)
	tw.Close()
	gw.Close()

	tmpFile, err := os.CreateTemp(t.TempDir(), "test-binary-*")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = extractBinaryFromTarball(&buf, tmpFile, "linux")
	if err == nil {
		t.Error("expected error for non-binary tarball")
	}
}

func TestNewManager(t *testing.T) {
	m := NewManager("v1.2.3")
	if m.CurrentVersion() != "1.2.3" {
		t.Errorf("CurrentVersion() = %q, want 1.2.3", m.CurrentVersion())
	}

	m2 := NewManager("0.1.0")
	if m2.CurrentVersion() != "0.1.0" {
		t.Errorf("CurrentVersion() = %q, want 0.1.0", m2.CurrentVersion())
	}
}

func TestReleaseConstants(t *testing.T) {
	if ReleaseRepo == "" {
		t.Error("ReleaseRepo is empty")
	}
	if ReleaseURL == "" {
		t.Error("ReleaseURL is empty")
	}
	if ReleaseURL != "https://api.github.com/repos/"+ReleaseRepo+"/releases/latest" {
		t.Errorf("ReleaseURL doesn't match repo: %s", ReleaseURL)
	}
}
