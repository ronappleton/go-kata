// Package main is the Go HTTP API sidecar for the GoKatas Tauri desktop app.
//
// It exposes the existing catalog, content, workspace, evaluator, and progress
// packages over a lightweight HTTP API on localhost. The Tauri shell process
// launches this binary on startup and the React frontend communicates via fetch.
//
// Endpoints:
//
//	GET  /api/catalog          → track with stages, categories, katas
//	GET  /api/kata/:id         → kata content (readme, starter, tests, metadata)
//	POST /api/kata/:id/save    → save user solution + learner tests
//	POST /api/kata/:id/run     → run in sandbox, return result
//	GET  /api/progress         → user progress state
//	GET  /api/status           → app status (sync state, runner health)
//	POST /api/sync             → trigger content sync with progress via SSE
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/content"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/evaluator"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/workspace"
)

type server struct {
	contentDir string
	contentURL string
	track      catalog.Track
	provider   content.ContentManager
	paths      workspace.Paths
	ws         *workspace.Manager
	progress   *progress.Store
	runner     *evaluator.Runner

	mu sync.RWMutex
}

type apiError struct {
	Error string `json:"error"`
}

func main() {
	port := 0
	if p := os.Getenv("GOKATAS_PORT"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			port = v
		}
	}

	contentDir := contentCacheDir()
	contentURL := os.Getenv("GOKATAS_CONTENT_URL")
	if contentURL == "" {
		contentURL = "https://raw.githubusercontent.com/ronappleton/gokatas-content/main"
	}

	s := &server{contentDir: contentDir, contentURL: contentURL}

	// Init workspace
	paths, err := workspace.ResolvePaths("gokatas")
	if err != nil {
		log.Fatalf("resolve paths: %v", err)
	}
	s.paths = paths
	s.ws = workspace.NewManager(paths)
	if err := s.ws.Ensure(); err != nil {
		log.Fatalf("ensure workspace: %v", err)
	}
	s.progress = progress.NewStore(filepath.Join(paths.State, "progress.json"))

	// Init content provider
	provider, err := content.NewProvider(content.ProviderConfig{
		ContentDir: contentDir,
		RemoteURL:  contentURL,
	})
	if err != nil {
		log.Fatalf("create provider: %v", err)
	}
	s.provider = provider

	// Sync content in background
	go s.syncContent()

	// Init evaluator
	image := os.Getenv("GOKATAS_RUNNER_IMAGE")
	if image == "" {
		image = defaultRunnerImage()
	}
	if image != "" {
		s.runner, _ = evaluator.NewRunner(image)
	}

	// Find a free port
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	actualPort := listener.Addr().(*net.TCPAddr).Port

	// Print port for Tauri to discover
	fmt.Fprintf(os.Stderr, "GOKATAS_PORT=%d\n", actualPort)
	fmt.Println(actualPort)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/catalog", s.handleCatalog)
	mux.HandleFunc("GET /api/kata/{id}", s.handleGetKata)
	mux.HandleFunc("POST /api/kata/{id}/save", s.handleSaveKata)
	mux.HandleFunc("POST /api/kata/{id}/run", s.handleRunKata)
	mux.HandleFunc("GET /api/progress", s.handleProgress)
	mux.HandleFunc("GET /api/status", s.handleStatus)
	mux.HandleFunc("POST /api/sync", s.handleSync)
	mux.HandleFunc("GET /api/sync/stream", s.handleSyncStream)

	handler := corsMiddleware(mux)

	srv := &http.Server{Handler: handler, ReadTimeout: 30 * time.Second, WriteTimeout: 5 * time.Minute}
	log.Printf("GoKatas API listening on :%d", actualPort)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("serve: %v", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *server) syncContent() {
	ctx := context.Background()
	_, err := s.provider.Sync(ctx)
	if err != nil {
		log.Printf("sync failed: %v", err)
	}
	s.loadTrack()
}

func (s *server) loadTrack() {
	ctx := context.Background()
	manifest, err := s.provider.GetManifest(ctx)
	if err != nil || len(manifest.Tracks) == 0 {
		log.Printf("loadTrack: no manifest: %v", err)
		return
	}
	track, err := catalog.LoadTrackFromContent(ctx, s.provider, manifest.Tracks[0].ID)
	if err != nil {
		log.Printf("loadTrack: %v", err)
	}
	s.mu.Lock()
	s.track = track
	s.mu.Unlock()
	log.Printf("loadTrack: %d stages, %d katas", len(track.Stages), len(track.AllKatas()))
}

// ── Handlers ──

func (s *server) handleCatalog(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	track := s.track
	s.mu.RUnlock()

	type kataSummary struct {
		ID              string   `json:"id"`
		Slug            string   `json:"slug"`
		Title           string   `json:"title"`
		Focus           string   `json:"focus"`
		Signature       string   `json:"signature"`
		EvaluatorStatus string   `json:"evaluatorStatus"`
		Level           string   `json:"level"`
		Language        string   `json:"language"`
		Tags            []string `json:"tags"`
	}
	type categorySummary struct {
		ID           string         `json:"id"`
		Title        string         `json:"title"`
		LearningGoal string         `json:"learningGoal"`
		Katas        []kataSummary  `json:"katas"`
	}
	type stageSummary struct {
		ID          string            `json:"id"`
		Title       string            `json:"title"`
		Level       string            `json:"level"`
		Categories  []categorySummary `json:"categories"`
	}

	result := struct {
		ID          string          `json:"id"`
		Title       string          `json:"title"`
		Description string          `json:"description"`
		Stages      []stageSummary  `json:"stages"`
		KataCount   int             `json:"kataCount"`
	}{
		ID:          track.ID,
		Title:       track.Title,
		Description: track.Description,
		KataCount:   len(track.AllKatas()),
	}

	for _, stage := range track.Stages {
		ss := stageSummary{ID: stage.ID, Title: stage.Title, Level: stage.Level}
		for _, cat := range stage.Categories {
			cs := categorySummary{ID: cat.ID, Title: cat.Title, LearningGoal: cat.LearningGoal}
			for _, k := range cat.Katas {
				cs.Katas = append(cs.Katas, kataSummary{
					ID: k.ID, Slug: k.Slug, Title: k.Title, Focus: k.Focus,
					Signature: k.Signature, EvaluatorStatus: k.EvaluatorStatus,
					Level: k.Level, Language: k.Language, Tags: k.Tags,
				})
			}
			ss.Categories = append(ss.Categories, cs)
		}
		result.Stages = append(result.Stages, ss)
	}

	writeJSON(w, result)
}

func (s *server) handleGetKata(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.mu.RLock()
	track := s.track
	s.mu.RUnlock()

	kata, _, ok := track.FindKata(id)
	if !ok {
		writeError(w, http.StatusNotFound, fmt.Sprintf("kata %s not found", id))
		return
	}

	writeJSON(w, map[string]interface{}{
		"kata":    kata,
		"content": kata.Content,
	})
}

func (s *server) handleSaveKata(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var body struct {
		Code   string `json:"code"`
		Tests  string `json:"tests"`
		Source string `json:"sourceFilename"`
		Test   string `json:"testFilename"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	sourceFile := body.Source
	if sourceFile == "" {
		sourceFile = "kata.go"
	}
	testFile := body.Test
	if testFile == "" {
		testFile = "kata_test.go"
	}

	if err := s.ws.SaveSolutionAs(id, body.Code, evaluator.DefaultCodeLimitBytes, sourceFile); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := s.ws.SaveLearnerTestsAs(id, body.Tests, evaluator.DefaultTestsLimitBytes, testFile); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, map[string]string{"status": "saved"})
}

func (s *server) handleRunKata(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if s.runner == nil {
		writeError(w, http.StatusServiceUnavailable, "sandbox runner not configured")
		return
	}

	var body struct {
		Code   string `json:"code"`
		Tests  string `json:"tests"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	s.mu.RLock()
	track := s.track
	s.mu.RUnlock()

	kata, _, ok := track.FindKata(id)
	if !ok {
		writeError(w, http.StatusNotFound, fmt.Sprintf("kata %s not found", id))
		return
	}

	result := s.runner.Run(context.Background(), evaluator.Request{
		KataID:       id,
		Module:       "kata" + id,
		Code:         body.Code,
		LearnerTests: body.Tests,
		TrustedTests: kata.Content.KataTest,
	})

	writeJSON(w, result)
}

func (s *server) handleProgress(w http.ResponseWriter, r *http.Request) {
	state, err := s.progress.Load()
	if err != nil {
		writeJSON(w, map[string]interface{}{"attempts": map[string]interface{}{}})
		return
	}
	writeJSON(w, state)
}

func (s *server) handleStatus(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	track := s.track
	s.mu.RUnlock()

	writeJSON(w, map[string]interface{}{
		"kataCount":  len(track.AllKatas()),
		"stageCount": len(track.Stages),
		"hasRunner":  s.runner != nil,
		"syncReady":  len(track.AllKatas()) > 0,
		"contentDir": s.contentDir,
	})
}

func (s *server) handleSync(w http.ResponseWriter, r *http.Request) {
	go s.syncContent()
	writeJSON(w, map[string]string{"status": "syncing"})
}

func (s *server) handleSyncStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	syncResult, err := s.provider.Sync(ctx)
	if err != nil {
		fmt.Fprintf(w, "data: {\"error\":%q}\n\n", err.Error())
		flusher.Flush()
		return
	}

	s.loadTrack()

	data, _ := json.Marshal(syncResult)
	fmt.Fprintf(w, "data: %s\n\n", data)
	flusher.Flush()
}

// ── Helpers ──

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(apiError{Error: msg})
}

func contentCacheDir() string {
	base, err := os.UserCacheDir()
	if err != nil {
		base = filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return filepath.Join(base, "gokatas", "content")
}

func defaultRunnerImage() string {
	if image := strings.TrimSpace(os.Getenv("GOKATAS_RUNNER_IMAGE")); image != "" {
		return image
	}
	configHome := strings.TrimSpace(os.Getenv("XDG_CONFIG_HOME"))
	if configHome == "" {
		home, _ := os.UserHomeDir()
		configHome = filepath.Join(home, ".config")
	}
	for _, path := range []string{
		filepath.Join(configHome, "gokatas", "runner-image"),
		"/etc/gokatas/runner-image",
	} {
		data, err := os.ReadFile(path)
		if err == nil {
			if s := strings.TrimSpace(string(data)); s != "" {
				return s
			}
		}
	}
	return ""
}
