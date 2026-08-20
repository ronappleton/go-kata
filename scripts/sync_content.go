//go:build ignore

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Manifest struct {
	Version   string       `json:"version"`
	MinAppVer string       `json:"min_app_version"`
	Tracks    []TrackEntry `json:"tracks"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type TrackEntry struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	KataCount int    `json:"kata_count"`
	Checksum  string `json:"checksum,omitempty"`
}

type TrackMeta struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Stages      []StageMeta   `json:"stages"`
}

type StageMeta struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Level       string         `json:"level"`
	Description string         `json:"description"`
	Categories  []CategoryMeta `json:"categories"`
}

type CategoryMeta struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	LearningGoal string   `json:"learning_goal,omitempty"`
	KataIDs      []string `json:"kata_ids"`
}

type KataContent struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Version   string `json:"version"`
	KataGo    string `json:"kata_go"`
	KataTest  string `json:"kata_test"`
	BuggyKata string `json:"buggy_kata,omitempty"`
	Readme    string `json:"readme"`
	JSON      string `json:"json"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: go run sync_content.go <repo-root> <output-dir>\n")
		fmt.Fprintf(os.Stderr, "  repo-root:   path to the go-kata repository\n")
		fmt.Fprintf(os.Stderr, "  output-dir:  path to write the content repo files\n")
		os.Exit(1)
	}

	repoRoot := os.Args[1]
	outputDir := os.Args[2]

	// Ensure output directory exists
	os.MkdirAll(outputDir, 0o755)
	os.MkdirAll(filepath.Join(outputDir, "tracks"), 0o755)

	// Read all track.json files
	entries, err := os.ReadDir(filepath.Join(repoRoot, "tracks"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "read tracks dir: %v\n", err)
		os.Exit(1)
	}

	var manifestTracks []TrackEntry

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		trackID := entry.Name()
		trackJSONPath := filepath.Join(repoRoot, "tracks", trackID, "track.json")

		trackData, err := os.ReadFile(trackJSONPath)
		if err != nil {
			fmt.Printf("skip %s: %v\n", trackID, err)
			continue
		}

		// Parse track to get kata IDs
		var rawTrack struct {
			ID    string `json:"id"`
			Title string `json:"title"`
			Stages []struct {
				Categories []struct {
					KataIDs []string `json:"kata_ids"`
				} `json:"categories"`
			} `json:"stages"`
		}
		if err := json.Unmarshal(trackData, &rawTrack); err != nil {
			fmt.Printf("skip %s: parse error: %v\n", trackID, err)
			continue
		}

		// Collect all kata IDs
		var allKataIDs []string
		for _, stage := range rawTrack.Stages {
			for _, cat := range stage.Categories {
				allKataIDs = append(allKataIDs, cat.KataIDs...)
			}
		}

		// Write track metadata to output
		trackOutputDir := filepath.Join(outputDir, "tracks", trackID)
		os.MkdirAll(filepath.Join(trackOutputDir, "katas"), 0o755)

		// Convert track.json to our format
		var fullTrack struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Stages      []struct {
				ID          string `json:"id"`
				Title       string `json:"title"`
				Level       string `json:"level"`
				Description string `json:"description"`
				Categories  []struct {
					ID           string   `json:"id"`
					Title        string   `json:"title"`
					Description  string   `json:"description"`
					LearningGoal string   `json:"learning_goal"`
					KataIDs      []string `json:"kata_ids"`
				} `json:"categories"`
			} `json:"stages"`
		}
		json.Unmarshal(trackData, &fullTrack)

		trackMeta := TrackMeta{
			ID:          fullTrack.ID,
			Title:       fullTrack.Title,
			Description: fullTrack.Description,
		}
		for _, s := range fullTrack.Stages {
			sm := StageMeta{
				ID:          s.ID,
				Title:       s.Title,
				Level:       s.Level,
				Description: s.Description,
			}
			for _, c := range s.Categories {
				sm.Categories = append(sm.Categories, CategoryMeta{
					ID:           c.ID,
					Title:        c.Title,
					Description:  c.Description,
					LearningGoal: c.LearningGoal,
					KataIDs:      c.KataIDs,
				})
			}
			trackMeta.Stages = append(trackMeta.Stages, sm)
		}

		trackMetaJSON, _ := json.MarshalIndent(trackMeta, "", "  ")
		os.WriteFile(filepath.Join(trackOutputDir, "track.json"), trackMetaJSON, 0o644)

		// Write each kata
		kataCount := 0
		for _, kataID := range allKataIDs {
			kataDir := findKataDir(repoRoot, kataID)
			if kataDir == "" {
				fmt.Printf("  skip kata %s: directory not found\n", kataID)
				continue
			}

			kataGo := readFileIfExists(filepath.Join(kataDir, "kata.go.txt"))
			kataTest := readFileIfExists(filepath.Join(kataDir, "kata_test.go.txt"))
			buggyKata := readFileIfExists(filepath.Join(kataDir, "buggy_kata.go.txt"))
			readme := readFileIfExists(filepath.Join(kataDir, "README.md"))
			jsonData := readFileIfExists(filepath.Join(kataDir, "kata.json"))

			// Extract title from JSON
			var meta struct {
				Title string `json:"title"`
				Slug  string `json:"slug"`
			}
			json.Unmarshal([]byte(jsonData), &meta)

			// Extract version from JSON
			var versionMeta struct {
				Version string `json:"version"`
			}
			json.Unmarshal([]byte(jsonData), &versionMeta)
			if versionMeta.Version == "" {
				versionMeta.Version = "1.0.0"
			}

			kata := KataContent{
				ID:        kataID,
				Slug:      meta.Slug,
				Title:     meta.Title,
				Version:   versionMeta.Version,
				KataGo:    kataGo,
				KataTest:  kataTest,
				BuggyKata: buggyKata,
				Readme:    readme,
				JSON:      jsonData,
			}

			kataJSON, _ := json.MarshalIndent(kata, "", "  ")
			os.WriteFile(filepath.Join(trackOutputDir, "katas", kataID+".json"), kataJSON, 0o644)
			kataCount++
		}

		// Calculate checksum for the track
		checksum := fmt.Sprintf("%x", sha256.Sum256(trackMetaJSON))

		manifestTracks = append(manifestTracks, TrackEntry{
			ID:        trackID,
			Title:     rawTrack.Title,
			KataCount: kataCount,
			Checksum:  "sha256:" + checksum[:16],
		})

		fmt.Printf("track %s: %d katas\n", trackID, kataCount)
	}

	// Sort tracks by ID
	sort.Slice(manifestTracks, func(i, j int) bool {
		return manifestTracks[i].ID < manifestTracks[j].ID
	})

	// Write manifest
	manifest := Manifest{
		Version:   "1.0.0",
		MinAppVer: "0.1.0",
		Tracks:    manifestTracks,
		UpdatedAt: time.Now(),
	}

	manifestJSON, _ := json.MarshalIndent(manifest, "", "  ")
	os.WriteFile(filepath.Join(outputDir, "manifest.json"), manifestJSON, 0o644)

	// Write README
	var readme strings.Builder
	readme.WriteString("# GoKatas Content Repository\n\n")
	readme.WriteString("This repository contains the kata content for the GoKatas learning platform.\n\n")
	readme.WriteString("## Tracks\n\n")
	for _, t := range manifestTracks {
		readme.WriteString(fmt.Sprintf("- **%s** — %d katas\n", t.Title, t.KataCount))
	}
	readme.WriteString(fmt.Sprintf("\nLast updated: %s\n", manifest.UpdatedAt.Format(time.RFC3339)))
	os.WriteFile(filepath.Join(outputDir, "README.md"), []byte(readme.String()), 0o644)

	fmt.Printf("\nWrote %d tracks, %d total katas to %s\n", len(manifestTracks), len(manifestTracks), outputDir)
}

func findKataDir(repoRoot, kataID string) string {
	entries, err := os.ReadDir(filepath.Join(repoRoot, "katas"))
	if err != nil {
		return ""
	}
	prefix := "kata-" + strings.TrimLeft(kataID, "0") + "-"
	// Also try zero-padded
	if len(kataID) < 3 {
		prefix = "kata-" + strings.Repeat("0", 3-len(kataID)) + kataID + "-"
	} else {
		prefix = "kata-" + kataID + "-"
	}
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), prefix) {
			return filepath.Join(repoRoot, "katas", e.Name())
		}
	}
	return ""
}

func readFileIfExists(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}
