package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var kataDirPattern = regexp.MustCompile(`^kata-(\d{3})-(.+)$`)

type Track struct {
	ID          string
	Title       string
	Description string
	Categories  []Category
}

type Category struct {
	ID           string
	Title        string
	Description  string
	LearningGoal string
	Katas        []Kata
}

type Kata struct {
	ID         string
	Slug       string
	Title      string
	Focus      string
	Dir        string
	ReadmePath string
	Signature  string
	Rules      []string
}

type trackConfig struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Categories  []categoryConfig `json:"categories"`
}

type categoryConfig struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	LearningGoal string        `json:"learning_goal"`
	KataRanges   []rangeConfig `json:"kata_ranges"`
	KataIDs      []string      `json:"kata_ids"`
}

type rangeConfig struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func LoadTrack(trackConfigPath, repoRoot string) (Track, error) {
	configData, err := os.ReadFile(trackConfigPath)
	if err != nil {
		return Track{}, fmt.Errorf("read track config: %w", err)
	}

	var cfg trackConfig
	if err := json.Unmarshal(configData, &cfg); err != nil {
		return Track{}, fmt.Errorf("parse track config: %w", err)
	}

	if cfg.ID == "" || len(cfg.Categories) == 0 {
		return Track{}, errors.New("track config must include id and categories")
	}

	kataDirs, err := scanKataDirs(repoRoot)
	if err != nil {
		return Track{}, err
	}

	track := Track{
		ID:          cfg.ID,
		Title:       cfg.Title,
		Description: cfg.Description,
		Categories:  make([]Category, 0, len(cfg.Categories)),
	}

	for _, categoryCfg := range cfg.Categories {
		ids, err := expandKataIDs(categoryCfg)
		if err != nil {
			return Track{}, fmt.Errorf("category %q: %w", categoryCfg.ID, err)
		}

		category := Category{
			ID:           categoryCfg.ID,
			Title:        categoryCfg.Title,
			Description:  categoryCfg.Description,
			LearningGoal: categoryCfg.LearningGoal,
			Katas:        make([]Kata, 0, len(ids)),
		}

		for _, id := range ids {
			dirName, ok := kataDirs[id]
			if !ok {
				return Track{}, fmt.Errorf("category %q references missing kata %s", categoryCfg.ID, id)
			}

			slug := kataDirPattern.FindStringSubmatch(dirName)[2]
			absDir := filepath.Join(repoRoot, dirName)
			readmePath := filepath.Join(absDir, "README.md")

			readmeMeta, err := parseReadmeMetadata(readmePath)
			if err != nil {
				return Track{}, fmt.Errorf("parse readme metadata for %s: %w", dirName, err)
			}

			category.Katas = append(category.Katas, Kata{
				ID:         id,
				Slug:       slug,
				Title:      readmeMeta.title,
				Focus:      readmeMeta.focus,
				Dir:        absDir,
				ReadmePath: readmePath,
				Signature:  readmeMeta.signature,
				Rules:      readmeMeta.rules,
			})
		}

		sort.Slice(category.Katas, func(i, j int) bool {
			return category.Katas[i].ID < category.Katas[j].ID
		})

		track.Categories = append(track.Categories, category)
	}

	return track, nil
}

func (t Track) FindCategory(categoryID string) (Category, bool) {
	for _, category := range t.Categories {
		if category.ID == categoryID {
			return category, true
		}
	}
	return Category{}, false
}

func (t Track) FindKata(rawID string) (Kata, Category, bool) {
	id, err := NormalizeKataID(rawID)
	if err != nil {
		return Kata{}, Category{}, false
	}

	for _, category := range t.Categories {
		for _, kata := range category.Katas {
			if kata.ID == id {
				return kata, category, true
			}
		}
	}

	return Kata{}, Category{}, false
}

func (t Track) AllKatas() []Kata {
	total := 0
	for _, category := range t.Categories {
		total += len(category.Katas)
	}

	katas := make([]Kata, 0, total)
	for _, category := range t.Categories {
		katas = append(katas, category.Katas...)
	}

	sort.Slice(katas, func(i, j int) bool {
		return katas[i].ID < katas[j].ID
	})
	return katas
}

func NormalizeKataID(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", errors.New("kata id is required")
	}

	if len(trimmed) == 3 {
		if _, err := strconv.Atoi(trimmed); err == nil {
			return trimmed, nil
		}
	}

	n, err := strconv.Atoi(trimmed)
	if err != nil {
		return "", fmt.Errorf("invalid kata id %q", raw)
	}
	if n < 1 || n > 999 {
		return "", fmt.Errorf("kata id out of range: %d", n)
	}
	return fmt.Sprintf("%03d", n), nil
}

func scanKataDirs(repoRoot string) (map[string]string, error) {
	entries, err := os.ReadDir(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("read repo root: %w", err)
	}

	out := make(map[string]string)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		matches := kataDirPattern.FindStringSubmatch(entry.Name())
		if len(matches) != 3 {
			continue
		}

		out[matches[1]] = entry.Name()
	}
	return out, nil
}

func expandKataIDs(cfg categoryConfig) ([]string, error) {
	seen := make(map[string]bool)
	ids := make([]string, 0, len(cfg.KataIDs))

	for _, rule := range cfg.KataRanges {
		if rule.Start <= 0 || rule.End <= 0 || rule.End < rule.Start {
			return nil, fmt.Errorf("invalid kata range %d-%d", rule.Start, rule.End)
		}
		for i := rule.Start; i <= rule.End; i++ {
			id := fmt.Sprintf("%03d", i)
			if !seen[id] {
				seen[id] = true
				ids = append(ids, id)
			}
		}
	}

	for _, raw := range cfg.KataIDs {
		id, err := NormalizeKataID(raw)
		if err != nil {
			return nil, err
		}
		if !seen[id] {
			seen[id] = true
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		return nil, errors.New("no kata ids configured")
	}

	sort.Strings(ids)
	return ids, nil
}

type readmeMetadata struct {
	title     string
	focus     string
	signature string
	rules     []string
}

func parseReadmeMetadata(readmePath string) (readmeMetadata, error) {
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return readmeMetadata{}, err
	}

	lines := strings.Split(string(data), "\n")
	meta := readmeMetadata{
		title: filepath.Base(filepath.Dir(readmePath)),
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "# Kata ") {
			meta.title = parseTitle(line)
		}
		if strings.HasPrefix(line, "**Focus:**") {
			meta.focus = strings.TrimSpace(strings.TrimPrefix(line, "**Focus:**"))
		}
	}

	if meta.focus == "" {
		meta.focus = "General Go practice"
	}

	meta.signature = extractSignature(lines)
	meta.rules = extractRules(lines)

	return meta, nil
}

func parseTitle(line string) string {
	parts := strings.SplitN(line, "—", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(strings.TrimPrefix(line, "#"))
}

func extractSignature(lines []string) string {
	inGoBlock := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "```go":
			inGoBlock = true
		case inGoBlock && trimmed == "```":
			return ""
		case inGoBlock && strings.HasPrefix(trimmed, "func "):
			return trimmed
		}
	}
	return ""
}

func extractRules(lines []string) []string {
	rules := []string{}
	inRules := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "## Rules / Expectations":
			inRules = true
		case inRules && strings.HasPrefix(trimmed, "## "):
			return rules
		case inRules && strings.HasPrefix(trimmed, "- "):
			rules = append(rules, strings.TrimSpace(strings.TrimPrefix(trimmed, "- ")))
		}
	}
	return rules
}
