package content

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/katas"
)

// embeddedStore provides access to the compiled-in kata content.
// This is the fallback when no remote or local content is available.
type embeddedStore struct{}

// NewEmbeddedStore creates a content store backed by compiled-in katas.
func NewEmbeddedStore() *embeddedStore {
	return &embeddedStore{}
}

// GetManifest builds a manifest from the embedded content.
func (e *embeddedStore) GetManifest(ctx context.Context) (*Manifest, error) {
	ids := katas.IDs()
	// Group by track (IDs < 200 = go-core-100, 200-299 = terraform, etc.)
	trackCounts := map[string]int{
		"go-core-100":   0,
		"terraform-100": 0,
		"helm-100":      0,
		"security-100":  0,
	}
	for _, id := range ids {
		switch {
		case id < "200":
			trackCounts["go-core-100"]++
		case id < "300":
			trackCounts["terraform-100"]++
		case id < "400":
			trackCounts["helm-100"]++
		default:
			trackCounts["security-100"]++
		}
	}

	trackTitles := map[string]string{
		"go-core-100":   "Go Mastery: Junior to Lead",
		"terraform-100": "Infrastructure as Code: Terraform",
		"helm-100":      "Container Orchestration: Helm",
		"security-100":  "Security & CVE Awareness",
	}

	var tracks []TrackEntry
	for id, count := range trackCounts {
		if count > 0 {
			tracks = append(tracks, TrackEntry{
				ID:        id,
				Title:     trackTitles[id],
				KataCount: count,
			})
		}
	}

	return &Manifest{
		Version: "embedded",
		Tracks:  tracks,
	}, nil
}

// GetTrack builds a track definition from embedded content.
func (e *embeddedStore) GetTrack(ctx context.Context, trackID string) (*TrackMeta, error) {
	// For embedded mode, we return a basic track structure.
	// The catalog package handles the full track.json parsing.
	return nil, fmt.Errorf("embedded store: use catalog.LoadTrack for full track data")
}

// GetKata returns content for a specific kata from embedded data.
func (e *embeddedStore) GetKata(ctx context.Context, trackID, kataID string) (*KataContent, error) {
	content, ok := katas.Content[kataID]
	if !ok {
		return nil, fmt.Errorf("kata %s not found", kataID)
	}

	// Parse the JSON metadata to extract fields
	var meta struct {
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}
	if err := json.Unmarshal([]byte(content.JSON), &meta); err != nil {
		meta.Title = "Kata " + kataID
	}

	return &KataContent{
		ID:        content.ID,
		Slug:      content.Slug,
		Title:     meta.Title,
		KataGo:    content.KataGo,
		KataTest:  content.KataTest,
		BuggyKata: content.BuggyKata,
		Readme:    content.Readme,
		JSON:      content.JSON,
	}, nil
}

// IDs returns all embedded kata IDs.
func (e *embeddedStore) IDs() []string {
	return katas.IDs()
}
