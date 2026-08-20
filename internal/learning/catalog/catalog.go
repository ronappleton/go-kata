package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/katas"
)

// Stage represents a learning stage (Foundation, Practitioner, Senior, Lead).
type Stage struct {
	ID          string
	Title       string
	Level       string // "junior", "mid", "senior", "lead"
	Description string
	Categories  []Category
}

// Track is the top-level learning track.
type Track struct {
	ID          string
	Title       string
	Description string
	Stages      []Stage
	// Legacy: flat categories (for backward compat with old track.json)
	Categories []Category
}

// Category holds a group of katas within a stage.
type Category struct {
	ID           string
	Title        string
	Description  string
	LearningGoal string
	Katas        []Kata
}

// Kata is a single exercise with all its metadata and content.
type Kata struct {
	ID              string
	Slug            string
	Title           string
	Focus           string
	Signature       string
	Rules           []string
	EvaluatorStatus string
	Stage           string   // "foundation", "practitioner", "senior", "lead"
	Category        string   // category ID within the stage
	Level           string   // "junior", "mid", "senior", "lead"
	Tags            []string // topic tags for cross-kata queries
	Prerequisites   []string // kata IDs that should be completed first
	EstimatedMin    int      // estimated minutes to complete
	Flashcards      []katas.Flashcard
	QuizQuestions   []katas.QuizQuestion
	Content         katas.KataContent
}

// ── JSON config types ──

type trackConfig struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Stages      []stageConfig    `json:"stages"`
	Categories  []categoryConfig `json:"categories"` // legacy flat format
}

type stageConfig struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Level       string           `json:"level"`
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

// LoadTrack reads a track.json and returns a fully populated Track.
// Supports both the new 4-stage format and the legacy flat-categories format.
func LoadTrack(trackConfigPath string) (Track, error) {
	configData, err := os.ReadFile(trackConfigPath)
	if err != nil {
		return Track{}, fmt.Errorf("read track config: %w", err)
	}

	var cfg trackConfig
	if err := json.Unmarshal(configData, &cfg); err != nil {
		return Track{}, fmt.Errorf("parse track config: %w", err)
	}

	track := Track{
		ID:          cfg.ID,
		Title:       cfg.Title,
		Description: cfg.Description,
	}

	// New format: stages with nested categories
	if len(cfg.Stages) > 0 {
		track.Stages = make([]Stage, 0, len(cfg.Stages))
		for _, stageCfg := range cfg.Stages {
			stage, err := buildStage(stageCfg)
			if err != nil {
				return Track{}, fmt.Errorf("stage %q: %w", stageCfg.ID, err)
			}
			track.Stages = append(track.Stages, stage)
		}
		return track, nil
	}

	// Legacy format: flat categories
	if len(cfg.Categories) == 0 {
		return Track{}, errors.New("track config must include stages or categories")
	}

	track.Categories = make([]Category, 0, len(cfg.Categories))
	for _, catCfg := range cfg.Categories {
		cat, err := buildCategory(catCfg)
		if err != nil {
			return Track{}, fmt.Errorf("category %q: %w", catCfg.ID, err)
		}
		track.Categories = append(track.Categories, cat)
	}

	return track, nil
}

func buildStage(cfg stageConfig) (Stage, error) {
	stage := Stage{
		ID:          cfg.ID,
		Title:       cfg.Title,
		Level:       cfg.Level,
		Description: cfg.Description,
		Categories:  make([]Category, 0, len(cfg.Categories)),
	}

	for _, catCfg := range cfg.Categories {
		cat, err := buildCategory(catCfg)
		if err != nil {
			return Stage{}, err
		}
		stage.Categories = append(stage.Categories, cat)
	}

	return stage, nil
}

func buildCategory(cfg categoryConfig) (Category, error) {
	ids, err := expandKataIDs(cfg)
	if err != nil {
		return Category{}, fmt.Errorf("category %q: %w", cfg.ID, err)
	}

	cat := Category{
		ID:           cfg.ID,
		Title:        cfg.Title,
		Description:  cfg.Description,
		LearningGoal: cfg.LearningGoal,
		Katas:        make([]Kata, 0, len(ids)),
	}

	for _, id := range ids {
		kata, err := loadKata(id, cfg.ID)
		if err != nil {
			return Category{}, err
		}
		cat.Katas = append(cat.Katas, kata)
	}

	sort.Slice(cat.Katas, func(i, j int) bool {
		return cat.Katas[i].ID < cat.Katas[j].ID
	})

	return cat, nil
}

func loadKata(id, categoryID string) (Kata, error) {
	content, ok := katas.Content[id]
	if !ok {
		return Kata{}, fmt.Errorf("category %q references missing embedded kata %s", categoryID, id)
	}

	var meta katas.KataMeta
	if err := json.Unmarshal([]byte(content.JSON), &meta); err != nil {
		return Kata{}, fmt.Errorf("parse embedded metadata for kata %s: %w", id, err)
	}

	return Kata{
		ID:              id,
		Slug:            content.Slug,
		Title:           meta.Title,
		Focus:           meta.Focus,
		Signature:       meta.Signature,
		Rules:           meta.Rules,
		EvaluatorStatus: meta.EvaluatorStatus,
		Stage:           meta.Stage,
		Category:        meta.Category,
		Level:           meta.Level,
		Tags:            meta.Tags,
		Prerequisites:   meta.Prerequisites,
		EstimatedMin:    meta.EstimatedMinutes,
		Flashcards:      meta.Flashcards,
		QuizQuestions:   meta.QuizQuestions,
		Content:         content,
	}, nil
}

// ── Query methods ──

// FindStage returns a stage by ID.
func (t Track) FindStage(stageID string) (Stage, bool) {
	for _, stage := range t.Stages {
		if stage.ID == stageID {
			return stage, true
		}
	}
	return Stage{}, false
}

// FindCategory returns a category by ID, searching through all stages.
func (t Track) FindCategory(categoryID string) (Category, bool) {
	// Search in staged format
	for _, stage := range t.Stages {
		for _, cat := range stage.Categories {
			if cat.ID == categoryID {
				return cat, true
			}
		}
	}
	// Search in legacy flat format
	for _, cat := range t.Categories {
		if cat.ID == categoryID {
			return cat, true
		}
	}
	return Category{}, false
}

// FindKata returns a kata and its containing category by kata ID.
func (t Track) FindKata(rawID string) (Kata, Category, bool) {
	id, err := NormalizeKataID(rawID)
	if err != nil {
		return Kata{}, Category{}, false
	}

	// Search in staged format
	for _, stage := range t.Stages {
		for _, cat := range stage.Categories {
			for _, kata := range cat.Katas {
				if kata.ID == id {
					return kata, cat, true
				}
			}
		}
	}
	// Search in legacy flat format
	for _, cat := range t.Categories {
		for _, kata := range cat.Katas {
			if kata.ID == id {
				return kata, cat, true
			}
		}
	}

	return Kata{}, Category{}, false
}

// AllKatas returns every kata in the track, sorted by ID.
func (t Track) AllKatas() []Kata {
	var all []Kata

	// Collect from staged format
	for _, stage := range t.Stages {
		for _, cat := range stage.Categories {
			all = append(all, cat.Katas...)
		}
	}
	// Collect from legacy flat format
	for _, cat := range t.Categories {
		all = append(all, cat.Katas...)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].ID < all[j].ID
	})
	return all
}

// Stages returns all stages. If using legacy format, returns nil.
func (t Track) StagesList() []Stage {
	return t.Stages
}

// NormalizeKataID converts a raw kata ID string to a zero-padded 3-digit ID.
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
	if n < 0 || n > 999 {
		return "", fmt.Errorf("kata id out of range: %d", n)
	}
	return fmt.Sprintf("%03d", n), nil
}

func expandKataIDs(cfg categoryConfig) ([]string, error) {
	seen := make(map[string]bool)
	ids := make([]string, 0, len(cfg.KataIDs))

	for _, rule := range cfg.KataRanges {
		if rule.Start < 0 || rule.End < 0 || rule.End < rule.Start {
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
