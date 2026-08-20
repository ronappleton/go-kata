package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
)

func (s *studioServer) handleLearn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	rawID := strings.TrimSpace(r.URL.Query().Get("id"))
	if rawID == "" {
		writeError(w, http.StatusBadRequest, "missing id query parameter")
		return
	}

	kata, category, ok := s.track.FindKata(rawID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	writeJSON(w, http.StatusOK, learnResponse{
		KataID:     kata.ID,
		Flashcards: s.buildFlashcards(kata, category),
		Quiz:       s.buildQuiz(kata, 8),
	})
}

func (s *studioServer) buildFlashcards(kata catalog.Kata, category catalog.Category) []flashcardItem {
	cards := []flashcardItem{
		{
			ID:    "signature",
			Front: fmt.Sprintf("Which function contract are you implementing in kata %s?", kata.ID),
			Back:  valueOr(kata.Signature, "Use the README task section to confirm the exact function signature."),
			Tag:   "signature",
		},
		{
			ID:    "focus",
			Front: "What specific skill is this kata training?",
			Back:  valueOr(kata.Focus, "Read the Focus line in the README."),
			Tag:   "focus",
		},
		{
			ID:    "goal",
			Front: fmt.Sprintf("Why is this kata in the %q category?", category.Title),
			Back:  valueOr(category.LearningGoal, "Practice one behavior contract at a time."),
			Tag:   "goal",
		},
	}

	ruleIdx := 1
	for _, rawRule := range kata.Rules {
		rule := strings.TrimSpace(rawRule)
		if rule == "" {
			continue
		}

		front := "Which behavior is explicitly required by this kata?"
		back := rule
		if left, right, ok := splitRule(rule); ok {
			front = fmt.Sprintf("When %s, what should happen?", left)
			back = right
		}

		cards = append(cards, flashcardItem{
			ID:    fmt.Sprintf("rule-%d", ruleIdx),
			Front: front,
			Back:  back,
			Tag:   "rule",
		})
		ruleIdx++
	}

	cards = append(cards, flashcardItem{
		ID:    "completion",
		Front: "How do you know this kata is done?",
		Back:  "All README rules pass in tests, including edge and invalid-input cases.",
		Tag:   "completion",
	})

	return cards
}

func (s *studioServer) buildQuiz(kata catalog.Kata, maxQuestions int) []quizQuestionItem {
	if maxQuestions <= 0 {
		maxQuestions = 6
	}

	outcomes, statements := s.rulePools(kata.ID)
	questions := make([]quizQuestionItem, 0, maxQuestions)

	for _, rawRule := range kata.Rules {
		if len(questions) >= maxQuestions {
			break
		}

		rule := strings.TrimSpace(rawRule)
		if rule == "" {
			continue
		}

		q := quizQuestionItem{
			ID:          fmt.Sprintf("q-%d", len(questions)+1),
			Explanation: rule,
		}

		var correct string
		var pool []string
		if left, right, ok := splitRule(rule); ok {
			q.Prompt = fmt.Sprintf("For kata %s, choose the correct outcome when %s.", kata.ID, left)
			correct = right
			pool = filterPool(outcomes, right)
		} else {
			q.Prompt = fmt.Sprintf("Which statement is part of kata %s's required behavior?", kata.ID)
			correct = rule
			pool = filterPool(statements, rule)
		}

		options, answerIdx := buildOptionsDeterministic(correct, pool, len(questions))
		q.Options = options
		q.AnswerIndex = answerIdx
		questions = append(questions, q)
	}

	return questions
}

func (s *studioServer) rulePools(excludeKataID string) (outcomes []string, statements []string) {
	outcomeSet := make(map[string]struct{})
	statementSet := make(map[string]struct{})

	for _, kata := range s.track.AllKatas() {
		if kata.ID == excludeKataID {
			continue
		}
		for _, rawRule := range kata.Rules {
			rule := strings.TrimSpace(rawRule)
			if rule == "" {
				continue
			}
			statementSet[rule] = struct{}{}
			if _, right, ok := splitRule(rule); ok {
				outcomeSet[right] = struct{}{}
			}
		}
	}

	outcomes = mapKeysSorted(outcomeSet)
	statements = mapKeysSorted(statementSet)
	return outcomes, statements
}

func splitRule(rule string) (left, right string, ok bool) {
	parts := strings.SplitN(rule, "=>", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	left = strings.TrimSpace(parts[0])
	right = strings.TrimSpace(parts[1])
	if left == "" || right == "" {
		return "", "", false
	}
	return left, right, true
}

func mapKeysSorted(set map[string]struct{}) []string {
	out := make([]string, 0, len(set))
	for key := range set {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

func filterPool(pool []string, exclude string) []string {
	exclude = strings.TrimSpace(exclude)
	filtered := make([]string, 0, len(pool))
	seen := make(map[string]struct{})

	for _, item := range pool {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" || trimmed == exclude {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		filtered = append(filtered, trimmed)
	}
	return filtered
}

func buildOptionsDeterministic(correct string, pool []string, salt int) ([]string, int) {
	correct = strings.TrimSpace(correct)
	if correct == "" {
		correct = "Use the README rule as the expected behavior."
	}

	wrong := make([]string, 0, 3)
	seen := map[string]bool{correct: true}

	for _, candidate := range pool {
		if seen[candidate] {
			continue
		}
		seen[candidate] = true
		wrong = append(wrong, candidate)
		if len(wrong) == 3 {
			break
		}
	}

	fallbacks := []string{
		"Return the zero value and continue.",
		"Panic when this case is hit.",
		"Skip this input without changing output.",
		"Return nil regardless of input.",
	}
	for _, fallback := range fallbacks {
		if len(wrong) == 3 {
			break
		}
		if seen[fallback] {
			continue
		}
		seen[fallback] = true
		wrong = append(wrong, fallback)
	}

	options := append(append([]string{}, wrong...), correct)
	if len(options) == 1 {
		return options, 0
	}

	shift := (len(correct) + salt) % len(options)
	if shift != 0 {
		options = append(options[shift:], options[:shift]...)
	}

	answerIndex := 0
	for idx, option := range options {
		if option == correct {
			answerIndex = idx
			break
		}
	}

	return options, answerIndex
}
