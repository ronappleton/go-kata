package content

import (
	"context"
	"time"
)

// ContentManager provides access to kata content from multiple sources.
// Resolution order: local cache → remote fetch. Production callers should
// provide a remote URL and rely on the local cache after first sync.
type ContentManager interface {
	// GetManifest returns the current manifest of available tracks.
	GetManifest(ctx context.Context) (*Manifest, error)

	// GetTrack returns full track metadata including kata IDs.
	GetTrack(ctx context.Context, trackID string) (*TrackMeta, error)

	// GetKata returns the content for a specific kata.
	GetKata(ctx context.Context, trackID, kataID string) (*KataContent, error)

	// Sync downloads any updated content from the remote source.
	Sync(ctx context.Context) (*SyncResult, error)

	// IsRemote returns true if remote content is available.
	IsRemote() bool

	// LastSync returns when content was last synced.
	LastSync() time.Time

	// ContentDir returns the local cache directory.
	ContentDir() string
}

// Manifest describes available content and versions.
type Manifest struct {
	Version   string       `json:"version"`
	MinAppVer string       `json:"min_app_version"`
	Tracks    []TrackEntry `json:"tracks"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// TrackEntry is a lightweight track summary in the manifest.
type TrackEntry struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	KataCount int    `json:"kata_count"`
	Checksum  string `json:"checksum,omitempty"`
}

// TrackMeta is the full track definition with categories and kata IDs.
type TrackMeta struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Stages      []StageMeta `json:"stages"`
}

// StageMeta is a stage within a track.
type StageMeta struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Level       string         `json:"level"`
	Description string         `json:"description"`
	Categories  []CategoryMeta `json:"categories"`
}

// CategoryMeta is a category within a stage.
type CategoryMeta struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	LearningGoal string   `json:"learning_goal,omitempty"`
	KataIDs      []string `json:"kata_ids"`
}

// KataContent holds all files for a single kata.
type KataContent struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Version   string `json:"version"` // Semantic version (e.g., "1.0.0")
	KataGo    string `json:"kata_go"`
	KataTest  string `json:"kata_test"`
	BuggyKata string `json:"buggy_kata,omitempty"`
	Readme    string `json:"readme"`
	JSON      string `json:"json"` // Full metadata JSON
}

// SyncResult describes what changed during a sync.
type SyncResult struct {
	Updated  int
	Added    int
	Removed  int
	Failed   []string
	SyncedAt time.Time
}
