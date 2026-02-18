package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/marking"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/runner"
)

const (
	defaultTrackConfigRelative = "tracks/go-core-100/track.json"
	defaultRunTimeoutSeconds   = 90
	defaultOutputTailLines     = 100
)

//go:embed static/*
var embeddedStatic embed.FS

type studioServer struct {
	repoRoot string
	track    catalog.Track
	store    *progress.Store
	files    fs.FS
	mu       sync.Mutex
}

type trackResponse struct {
	ID             string                     `json:"id"`
	Title          string                     `json:"title"`
	Description    string                     `json:"description"`
	OverallDone    int                        `json:"overall_done"`
	OverallTotal   int                        `json:"overall_total"`
	OverallPercent int                        `json:"overall_percent"`
	Categories     []trackCategorySummaryItem `json:"categories"`
}

type trackCategorySummaryItem struct {
	ID           string                 `json:"id"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	LearningGoal string                 `json:"learning_goal"`
	Done         int                    `json:"done"`
	Total        int                    `json:"total"`
	Percent      int                    `json:"percent"`
	Katas        []trackKataSummaryItem `json:"katas"`
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

type runRequest struct {
	KataID         string  `json:"kata_id"`
	Code           *string `json:"code,omitempty"`
	Tests          *string `json:"tests,omitempty"`
	SaveBeforeRun  bool    `json:"save_before_run"`
	TimeoutSeconds int     `json:"timeout_seconds"`
}

type runResponse struct {
	Passed      bool                  `json:"passed"`
	DurationMS  int64                 `json:"duration_ms"`
	FailedTests []string              `json:"failed_tests"`
	OutputTail  string                `json:"output_tail"`
	Progress    progress.KataProgress `json:"progress"`
}

type markRequest struct {
	KataID string `json:"kata_id"`
}

type markResponse struct {
	PromptPath string `json:"prompt_path"`
	Prompt     string `json:"prompt"`
	ChatGPTURL string `json:"chatgpt_url"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	repoRootFlag := flag.String("repo", ".", "Path to kata repository root")
	addrFlag := flag.String("addr", "127.0.0.1:7777", "HTTP listen address")
	flag.Parse()

	server, err := newStudioServer(*repoRootFlag)
	if err != nil {
		fatalf("init learner studio: %v", err)
	}

	fmt.Printf("Learner Studio running at http://%s\n", *addrFlag)
	if err := http.ListenAndServe(*addrFlag, server.routes()); err != nil {
		fatalf("listen: %v", err)
	}
}

func newStudioServer(repoRoot string) (*studioServer, error) {
	absRepoRoot, err := filepath.Abs(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve repo root: %w", err)
	}

	track, err := catalog.LoadTrack(filepath.Join(absRepoRoot, defaultTrackConfigRelative), absRepoRoot)
	if err != nil {
		return nil, err
	}

	store := progress.NewStore(filepath.Join(absRepoRoot, ".learning", "progress.json"))
	if _, err := store.Load(); err != nil {
		return nil, fmt.Errorf("load progress: %w", err)
	}

	staticFS, err := fs.Sub(embeddedStatic, "static")
	if err != nil {
		return nil, fmt.Errorf("load embedded static files: %w", err)
	}

	return &studioServer{
		repoRoot: absRepoRoot,
		track:    track,
		store:    store,
		files:    staticFS,
	}, nil
}

func (s *studioServer) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/app.js", s.handleStaticFile("app.js", "text/javascript; charset=utf-8"))
	mux.HandleFunc("/styles.css", s.handleStaticFile("styles.css", "text/css; charset=utf-8"))

	mux.HandleFunc("/api/track", s.handleTrack)
	mux.HandleFunc("/api/kata", s.handleKata)
	mux.HandleFunc("/api/save", s.handleSave)
	mux.HandleFunc("/api/run", s.handleRun)
	mux.HandleFunc("/api/mark", s.handleMark)

	return mux
}

func (s *studioServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	s.serveFile(w, "index.html", "text/html; charset=utf-8")
}

func (s *studioServer) handleStaticFile(name, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.serveFile(w, name, contentType)
	}
}

func (s *studioServer) serveFile(w http.ResponseWriter, name, contentType string) {
	data, err := fs.ReadFile(s.files, name)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write(data)
}

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

		resp.Categories = append(resp.Categories, item)
	}

	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleKata(w http.ResponseWriter, r *http.Request) {
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

	readme, err := os.ReadFile(kata.ReadmePath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read readme: %v", err))
		return
	}
	code, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata.go: %v", err))
		return
	}
	tests, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	resp := kataResponse{
		ID:        kata.ID,
		Title:     kata.Title,
		Focus:     kata.Focus,
		Signature: kata.Signature,
		Rules:     append([]string(nil), kata.Rules...),
		Category: kataCategoryReference{
			ID:    category.ID,
			Title: category.Title,
		},
		Readme:   string(readme),
		Code:     string(code),
		Tests:    string(tests),
		Progress: state.Attempts[kata.ID],
	}

	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req saveRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), []byte(req.Code), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
		return
	}
	if err := os.WriteFile(filepath.Join(kata.Dir, "kata_test.go"), []byte(req.Tests), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata_test.go: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "saved"})
}

func (s *studioServer) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req runRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	if req.TimeoutSeconds <= 0 {
		req.TimeoutSeconds = defaultRunTimeoutSeconds
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if req.SaveBeforeRun {
		if req.Code != nil {
			if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), []byte(*req.Code), 0o644); err != nil {
				writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
				return
			}
		}
		if req.Tests != nil {
			if err := os.WriteFile(filepath.Join(kata.Dir, "kata_test.go"), []byte(*req.Tests), 0o644); err != nil {
				writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata_test.go: %v", err))
				return
			}
		}
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(req.TimeoutSeconds)*time.Second)
	defer cancel()

	start := time.Now()
	runResult, _ := runner.RunKataTests(ctx, kata.Dir)
	duration := runResult.Elapsed
	if duration <= 0 {
		duration = time.Since(start)
	}

	attemptedAt := time.Now().UTC()
	state, err := s.store.RecordAttempt(kata.ID, progress.AttemptResult{
		Passed:      runResult.Passed,
		Duration:    duration,
		FailedTests: runResult.FailedTests,
		OutputTail:  tailLines(runResult.RawOutput, defaultOutputTailLines),
		RanAt:       attemptedAt,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("record progress: %v", err))
		return
	}

	resp := runResponse{
		Passed:      runResult.Passed,
		DurationMS:  duration.Milliseconds(),
		FailedTests: append([]string(nil), runResult.FailedTests...),
		OutputTail:  tailLines(runResult.RawOutput, defaultOutputTailLines),
		Progress:    state.Attempts[kata.ID],
	}
	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleMark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req markRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, category, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	kataCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata.go: %v", err))
		return
	}
	testCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	last := state.Attempts[kata.ID]
	prompt := marking.BuildPrompt(marking.PromptInput{
		TrackTitle:      s.track.Title,
		CategoryTitle:   category.Title,
		KataID:          kata.ID,
		KataTitle:       kata.Title,
		Focus:           kata.Focus,
		Signature:       kata.Signature,
		Rules:           kata.Rules,
		KataCode:        string(kataCode),
		TestCode:        string(testCode),
		LastRunResult:   last.LastResult,
		LastFailedTests: last.LastFailedTests,
		LastOutputTail:  last.LastOutputTail,
	})

	target := marking.DefaultPromptPath(s.repoRoot, kata.ID, time.Now().UTC())
	if err := marking.WritePrompt(target, prompt); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write prompt: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, markResponse{
		PromptPath: target,
		Prompt:     prompt,
		ChatGPTURL: "https://chatgpt.com/",
	})
}

func decodeJSON(r *http.Request, out any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(out); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}

func tailLines(input string, maxLines int) string {
	if maxLines <= 0 {
		return ""
	}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) <= maxLines {
		return strings.Join(lines, "\n")
	}
	return strings.Join(lines[len(lines)-maxLines:], "\n")
}

func fatalf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
