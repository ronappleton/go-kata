package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/content"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/katas"
)

// LoadTrackFromContent builds the catalog model from a ContentManager. This is
// the production path: the catalog does not need embedded kata data or a
// package-relative filesystem layout.
//
// Individual kata failures are tolerated: a kata that cannot be fetched is
// skipped (with an error returned) rather than aborting the whole track, so a
// single transient network error does not blank the curriculum.
func LoadTrackFromContent(ctx context.Context, source content.ContentManager, trackID string) (Track, error) {
	meta, err := source.GetTrack(ctx, trackID)
	if err != nil {
		return Track{}, err
	}
	track := Track{ID: meta.ID, Title: meta.Title, Description: meta.Description}
	var skipped []string
	for _, stageMeta := range meta.Stages {
		stage := Stage{
			ID: stageMeta.ID, Title: stageMeta.Title, Level: stageMeta.Level,
			Description: stageMeta.Description,
		}
		for _, categoryMeta := range stageMeta.Categories {
			category := Category{
				ID: categoryMeta.ID, Title: categoryMeta.Title,
				Description: categoryMeta.Description, LearningGoal: categoryMeta.LearningGoal,
			}
			for _, kataID := range categoryMeta.KataIDs {
				kataContent, err := source.GetKata(ctx, meta.ID, kataID)
				if err != nil {
					skipped = append(skipped, kataID)
					continue
				}
				kata, err := kataFromContent(kataContent, stageMeta.ID, categoryMeta.ID)
				if err != nil {
					skipped = append(skipped, kataID)
					continue
				}
				category.Katas = append(category.Katas, kata)
			}
			if len(category.Katas) > 0 {
				stage.Categories = append(stage.Categories, category)
			}
		}
		if len(stage.Categories) > 0 {
			track.Stages = append(track.Stages, stage)
		}
	}
	if len(track.Stages) == 0 {
		return Track{}, fmt.Errorf("track %s contains no loadable katas", trackID)
	}
	if len(skipped) > 0 {
		return track, fmt.Errorf("track %s: %d kata(s) unavailable: %s", trackID, len(skipped), strings.Join(skipped, ", "))
	}
	return track, nil
}

func kataFromContent(raw *content.KataContent, stageID, categoryID string) (Kata, error) {
	if raw == nil || raw.ID == "" {
		return Kata{}, fmt.Errorf("kata content is empty")
	}
	var meta katas.KataMeta
	if err := json.Unmarshal([]byte(raw.JSON), &meta); err != nil {
		return Kata{}, fmt.Errorf("parse metadata for kata %s: %w", raw.ID, err)
	}
	if meta.ID == "" {
		meta.ID = raw.ID
	}
	if meta.Slug == "" {
		meta.Slug = raw.Slug
	}
	if meta.Language == "" {
		meta.Language = "go"
	}
	return Kata{
		ID: meta.ID, Slug: meta.Slug, Title: meta.Title, Language: meta.Language,
		Focus: meta.Focus,
		Signature: meta.Signature, Rules: meta.Rules, EvaluatorStatus: meta.EvaluatorStatus,
		Stage: stageID, Category: categoryID, Level: meta.Level, Tags: meta.Tags,
		Prerequisites: meta.Prerequisites, EstimatedMin: meta.EstimatedMinutes,
		Flashcards: meta.Flashcards, QuizQuestions: meta.QuizQuestions,
		Content: katas.KataContent{ID: raw.ID, Slug: raw.Slug, KataGo: raw.KataGo,
			KataTest: raw.KataTest, BuggyKata: raw.BuggyKata, Readme: raw.Readme, JSON: raw.JSON},
	}, nil
}
