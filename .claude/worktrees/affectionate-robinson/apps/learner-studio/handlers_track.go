package main

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
)

func (s *studioServer) handleTrack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	resp := trackResponse{
		ID:          s.track.ID,
		Title:       s.track.Title,
		Description: s.track.Description,
		Categories:  make([]trackCategorySummaryItem, 0, len(s.track.Categories)),
	}

	allIDs := make([]string, 0, len(s.track.AllKatas()))
	for _, kata := range s.track.AllKatas() {
		allIDs = append(allIDs, kata.ID)
	}
	resp.OverallDone = progress.CompletedCount(state, allIDs)
	resp.OverallTotal = len(allIDs)
	if resp.OverallTotal > 0 {
		resp.OverallPercent = int((float64(resp.OverallDone) / float64(resp.OverallTotal)) * 100)
	}
	resp.CoachMessage = coachMessage(resp.OverallDone, resp.OverallTotal)
	resp.NextRecommended = s.nextRecommendation(state)

	for _, category := range s.track.Categories {
		item := trackCategorySummaryItem{
			ID:           category.ID,
			Title:        category.Title,
			Description:  category.Description,
			LearningGoal: category.LearningGoal,
			Katas:        make([]trackKataSummaryItem, 0, len(category.Katas)),
		}

		ids := make([]string, 0, len(category.Katas))
		for _, kata := range category.Katas {
			ids = append(ids, kata.ID)
			item.Katas = append(item.Katas, trackKataSummaryItem{
				ID:        kata.ID,
				Title:     kata.Title,
				Focus:     kata.Focus,
				Completed: state.Attempts[kata.ID].Passes > 0,
			})
		}

		item.Done = progress.CompletedCount(state, ids)
		item.Total = len(category.Katas)
		if item.Total > 0 {
			item.Percent = int((float64(item.Done) / float64(item.Total)) * 100)
		}
		item.MilestoneLabel, item.MilestoneMessage, item.NextTargetPercent, item.RemainingToNext = categoryMilestone(item.Done, item.Total)

		resp.Categories = append(resp.Categories, item)
	}

	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handlePathways(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	items := make([]pathwayResponseItem, 0, len(s.pathways))
	for _, pathway := range s.pathways {
		ids := s.pathwayKataIDs(pathway)
		done := progress.CompletedCount(state, ids)
		total := len(ids)
		percent := 0
		if total > 0 {
			percent = int((float64(done) / float64(total)) * 100)
		}

		item := pathwayResponseItem{
			ID:               pathway.ID,
			Title:            pathway.Title,
			Description:      pathway.Description,
			RecommendedModes: append([]string(nil), pathway.RecommendedModes...),
			LevelOutcome:     pathway.LevelOutcome,
			Done:             done,
			Total:            total,
			Percent:          percent,
			Status:           pathwayStatus(done, total),
		}

		if next, ok := s.nextIncompleteFromIDs(state, ids); ok {
			item.NextKataID = next.ID
			item.NextKataTitle = next.Title
		}

		items = append(items, item)
	}

	writeJSON(w, http.StatusOK, pathwaysResponse{Items: items})
}

func pathwayStatus(done, total int) string {
	if total == 0 {
		return "not-started"
	}
	if done <= 0 {
		return "not-started"
	}
	if done >= total {
		return "completed"
	}
	return "in-progress"
}

func (s *studioServer) pathwayKataIDs(pathway pathwayDefinition) []string {
	ids := make([]string, 0, 32)
	seen := make(map[string]bool)

	for _, categoryID := range pathway.Categories {
		category, ok := s.track.FindCategory(categoryID)
		if !ok {
			continue
		}
		for _, kata := range category.Katas {
			if seen[kata.ID] {
				continue
			}
			seen[kata.ID] = true
			ids = append(ids, kata.ID)
		}
	}

	sort.Strings(ids)
	return ids
}

func (s *studioServer) nextIncompleteFromIDs(state progress.State, ids []string) (catalog.Kata, bool) {
	index := make(map[string]catalog.Kata, len(ids))
	for _, kata := range s.track.AllKatas() {
		index[kata.ID] = kata
	}

	for _, id := range ids {
		if state.Attempts[id].Passes > 0 {
			continue
		}
		if kata, ok := index[id]; ok {
			return kata, true
		}
	}
	return catalog.Kata{}, false
}

func (s *studioServer) nextRecommendation(state progress.State) *nextKataRecommendation {
	for _, category := range s.track.Categories {
		for _, kata := range category.Katas {
			if state.Attempts[kata.ID].Passes > 0 {
				continue
			}
			return &nextKataRecommendation{
				KataID:        kata.ID,
				KataTitle:     kata.Title,
				CategoryID:    category.ID,
				CategoryTitle: category.Title,
				Reason:        fmt.Sprintf("Next best step: %s in %s. %s", kata.Title, category.Title, category.LearningGoal),
			}
		}
	}
	return nil
}

func coachMessage(done, total int) string {
	if total == 0 {
		return "Track ready. Start with one kata and complete the full loop: read, code, test, reflect."
	}
	if done == 0 {
		return "Start small: complete one kata end-to-end before moving on."
	}
	if done >= total {
		return "Track complete. Revisit a weak category and refactor one kata for readability and tests."
	}
	return fmt.Sprintf("You have %d/%d complete. Keep focus: one kata per session, then review what you learned.", done, total)
}

func categoryMilestone(done, total int) (label, message string, nextTargetPercent, remainingToNext int) {
	if total <= 0 {
		return "No katas", "This category has no katas configured.", 0, 0
	}

	percent := int((float64(done) / float64(total)) * 100)
	thresholds := []int{25, 50, 75, 100}
	nextTargetPercent = 100
	for _, threshold := range thresholds {
		if percent < threshold {
			nextTargetPercent = threshold
			break
		}
	}

	targetCount := int((float64(nextTargetPercent) / 100.0) * float64(total))
	if targetCount < 1 {
		targetCount = 1
	}
	if targetCount > total {
		targetCount = total
	}
	remainingToNext = targetCount - done
	if remainingToNext < 0 {
		remainingToNext = 0
	}

	switch {
	case percent == 0:
		label = "Start line"
		message = "Complete your first kata in this category to establish momentum."
	case percent < 25:
		label = "Early momentum"
		message = "Good start. Keep the same category to build pattern recognition."
	case percent < 50:
		label = "Building consistency"
		message = "You are stacking wins. Tighten tests around edge cases."
	case percent < 75:
		label = "Solid core"
		message = "Halfway and stable. Focus on failure paths and clarity."
	case percent < 100:
		label = "Finishing strong"
		message = "You are close. Polish behavior contracts and naming quality."
	default:
		label = "Category complete"
		message = "Great work. Move to the next category and keep the same discipline."
		nextTargetPercent = 100
		remainingToNext = 0
	}

	return label, message, nextTargetPercent, remainingToNext
}
