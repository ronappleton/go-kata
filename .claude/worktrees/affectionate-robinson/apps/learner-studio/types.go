package main

import (
	"regexp"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
)

type pathwayDefinition struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Categories       []string `json:"categories"`
	RecommendedModes []string `json:"recommended_modes"`
	LevelOutcome     string   `json:"level_outcome"`
}

type pathwaysConfig struct {
	Pathways []pathwayDefinition `json:"pathways"`
}

type pathwayResponseItem struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	RecommendedModes []string `json:"recommended_modes"`
	LevelOutcome     string   `json:"level_outcome"`
	Done             int      `json:"done"`
	Total            int      `json:"total"`
	Percent          int      `json:"percent"`
	Status           string   `json:"status"`
	NextKataID       string   `json:"next_kata_id,omitempty"`
	NextKataTitle    string   `json:"next_kata_title,omitempty"`
}

type pathwaysResponse struct {
	Items []pathwayResponseItem `json:"items"`
}

type trackResponse struct {
	ID              string                     `json:"id"`
	Title           string                     `json:"title"`
	Description     string                     `json:"description"`
	OverallDone     int                        `json:"overall_done"`
	OverallTotal    int                        `json:"overall_total"`
	OverallPercent  int                        `json:"overall_percent"`
	CoachMessage    string                     `json:"coach_message"`
	NextRecommended *nextKataRecommendation    `json:"next_recommended,omitempty"`
	Categories      []trackCategorySummaryItem `json:"categories"`
}

type trackCategorySummaryItem struct {
	ID                string                 `json:"id"`
	Title             string                 `json:"title"`
	Description       string                 `json:"description"`
	LearningGoal      string                 `json:"learning_goal"`
	Done              int                    `json:"done"`
	Total             int                    `json:"total"`
	Percent           int                    `json:"percent"`
	MilestoneLabel    string                 `json:"milestone_label"`
	MilestoneMessage  string                 `json:"milestone_message"`
	NextTargetPercent int                    `json:"next_target_percent"`
	RemainingToNext   int                    `json:"remaining_to_next"`
	Katas             []trackKataSummaryItem `json:"katas"`
}

type trackKataSummaryItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Focus     string `json:"focus"`
	Completed bool   `json:"completed"`
}

type kataResponse struct {
	ID        string                `json:"id"`
	Title     string                `json:"title"`
	Focus     string                `json:"focus"`
	Signature string                `json:"signature"`
	Rules     []string              `json:"rules"`
	Category  kataCategoryReference `json:"category"`
	Readme    string                `json:"readme"`
	Code      string                `json:"code"`
	Tests     string                `json:"tests"`
	Progress  progress.KataProgress `json:"progress"`
}

type kataCategoryReference struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type saveRequest struct {
	KataID string `json:"kata_id"`
	Code   string `json:"code"`
	Tests  string `json:"tests"`
}

type resetBuggyRequest struct {
	KataID string `json:"kata_id"`
}

type runRequest struct {
	KataID         string  `json:"kata_id"`
	Code           *string `json:"code,omitempty"`
	Tests          *string `json:"tests,omitempty"`
	SaveBeforeRun  bool    `json:"save_before_run"`
	TimeoutSeconds int     `json:"timeout_seconds"`
}

type formatRequest struct {
	KataID string  `json:"kata_id"`
	Code   *string `json:"code,omitempty"`
	Tests  *string `json:"tests,omitempty"`
}

type formatResponse struct {
	Code  string `json:"code"`
	Tests string `json:"tests"`
}

type learnResponse struct {
	KataID     string             `json:"kata_id"`
	Flashcards []flashcardItem    `json:"flashcards"`
	Quiz       []quizQuestionItem `json:"quiz"`
}

type flashcardItem struct {
	ID    string `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	Tag   string `json:"tag,omitempty"`
}

type quizQuestionItem struct {
	ID          string   `json:"id"`
	Prompt      string   `json:"prompt"`
	Options     []string `json:"options"`
	AnswerIndex int      `json:"answer_index"`
	Explanation string   `json:"explanation"`
}

type runResponse struct {
	Passed          bool                    `json:"passed"`
	CompileErr      bool                    `json:"compile_err"`
	DurationMS      int64                   `json:"duration_ms"`
	FailedTests     []string                `json:"failed_tests"`
	OutputTail      string                  `json:"output_tail"`
	FailureInsights []failureInsight        `json:"failure_insights"`
	CoachHint       string                  `json:"coach_hint"`
	NextRecommended *nextKataRecommendation `json:"next_recommended,omitempty"`
	Progress        progress.KataProgress   `json:"progress"`
}

type markRequest struct {
	KataID string `json:"kata_id"`
}

type markResponse struct {
	PromptPath string `json:"prompt_path"`
	Prompt     string `json:"prompt"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type nextKataRecommendation struct {
	KataID        string `json:"kata_id"`
	KataTitle     string `json:"kata_title"`
	CategoryID    string `json:"category_id"`
	CategoryTitle string `json:"category_title"`
	Reason        string `json:"reason"`
}

type failureInsight struct {
	Kind     string `json:"kind"`
	Summary  string `json:"summary"`
	Expected string `json:"expected,omitempty"`
	Actual   string `json:"actual,omitempty"`
}

var expectedGotPattern = regexp.MustCompile(`(?i)expected[: ]+(.+?)[,; ]+got[: ]+(.+)$`)
var goTestPrefixPattern = regexp.MustCompile(`^[^:]+:\d+:\s*`)
