package catalog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// TrackInfo is a lightweight summary of a track for the track selector.
type TrackInfo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	KataCount   int    `json:"kata_count"`
}

// DiscoverTracks scans a directory for track.json files and returns summaries.
func DiscoverTracks(tracksDir string) ([]TrackInfo, error) {
	entries, err := os.ReadDir(tracksDir)
	if err != nil {
		return nil, fmt.Errorf("read tracks dir: %w", err)
	}

	var tracks []TrackInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		trackPath := filepath.Join(tracksDir, entry.Name(), "track.json")
		data, err := os.ReadFile(trackPath)
		if err != nil {
			continue
		}

		var raw struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Stages      []struct {
				Categories []struct {
					KataIDs []string `json:"kata_ids"`
				} `json:"categories"`
			} `json:"stages"`
		}
		if err := json.Unmarshal(data, &raw); err != nil {
			continue
		}

		kataCount := 0
		for _, s := range raw.Stages {
			for _, c := range s.Categories {
				kataCount += len(c.KataIDs)
			}
		}

		tracks = append(tracks, TrackInfo{
			ID:          raw.ID,
			Title:       raw.Title,
			Description: raw.Description,
			KataCount:   kataCount,
		})
	}

	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].ID < tracks[j].ID
	})

	return tracks, nil
}
