package content

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// localStore provides access to kata content cached on the filesystem.
type localStore struct {
	dir string // ~/.local/share/gokatas/content
}

// NewLocalStore creates a content store backed by the local filesystem.
func NewLocalStore(contentDir string) (*localStore, error) {
	if err := os.MkdirAll(contentDir, 0o755); err != nil {
		return nil, fmt.Errorf("create content dir: %w", err)
	}
	return &localStore{dir: contentDir}, nil
}

func (l *localStore) manifestPath() string    { return filepath.Join(l.dir, "manifest.json") }
func (l *localStore) trackPath(id string) string {
	return filepath.Join(l.dir, "tracks", id, "track.json")
}
func (l *localStore) kataPath(trackID, kataID string) string {
	return filepath.Join(l.dir, "tracks", trackID, "katas", kataID+".json")
}

func (l *localStore) HasManifest() bool {
	_, err := os.Stat(l.manifestPath())
	return err == nil
}

func (l *localStore) GetManifest() (*Manifest, error) {
	data, err := os.ReadFile(l.manifestPath())
	if err != nil {
		return nil, fmt.Errorf("read manifest: %w", err)
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parse manifest: %w", err)
	}
	return &m, nil
}

func (l *localStore) SaveManifest(m *Manifest) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.manifestPath(), data, 0o644)
}

func (l *localStore) GetTrack(trackID string) (*TrackMeta, error) {
	data, err := os.ReadFile(l.trackPath(trackID))
	if err != nil {
		return nil, fmt.Errorf("read track %s: %w", trackID, err)
	}
	var t TrackMeta
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("parse track %s: %w", trackID, err)
	}
	return &t, nil
}

func (l *localStore) SaveTrack(trackID string, t *TrackMeta) error {
	dir := filepath.Dir(l.trackPath(trackID))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.trackPath(trackID), data, 0o644)
}

func (l *localStore) GetKata(trackID, kataID string) (*KataContent, error) {
	data, err := os.ReadFile(l.kataPath(trackID, kataID))
	if err != nil {
		return nil, fmt.Errorf("read kata %s/%s: %w", trackID, kataID, err)
	}
	var k KataContent
	if err := json.Unmarshal(data, &k); err != nil {
		return nil, fmt.Errorf("parse kata %s/%s: %w", trackID, kataID, err)
	}
	return &k, nil
}

func (l *localStore) SaveKata(trackID string, k *KataContent) error {
	dir := filepath.Dir(l.kataPath(trackID, k.ID))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.kataPath(trackID, k.ID), data, 0o644)
}

func (l *localStore) LastSyncTime() time.Time {
	info, err := os.Stat(l.manifestPath())
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

func (l *localStore) Dir() string { return l.dir }
