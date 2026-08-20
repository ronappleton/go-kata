package katas

// Flashcard is a single flashcard with front/back sides.
type Flashcard struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

// QuizQuestion is a quiz question (multiple choice, fill-in-blank, etc.).
type QuizQuestion struct {
	Type     string   `json:"type"` // "multiple_choice", "fill_blank", "code_complete"
	Question string   `json:"question"`
	Options  []string `json:"options,omitempty"`
	Answer   string   `json:"answer"`
}

// KataMeta is the machine-readable metadata for a kata, matching kata.json schema.
type KataMeta struct {
	ID               string         `json:"id"`
	Slug             string         `json:"slug"`
	Title            string         `json:"title"`
	Focus            string         `json:"focus"`
	Signature        string         `json:"signature"`
	Rules            []string       `json:"rules"`
	EvaluatorStatus  string         `json:"evaluator_status"`
	Stage            string         `json:"stage"`
	Category         string         `json:"category"`
	Level            string         `json:"level"`
	Tags             []string       `json:"tags"`
	Prerequisites    []string       `json:"prerequisites"`
	EstimatedMinutes int            `json:"estimated_minutes"`
	Flashcards       []Flashcard    `json:"flashcards"`
	QuizQuestions    []QuizQuestion `json:"quiz_questions"`
}
