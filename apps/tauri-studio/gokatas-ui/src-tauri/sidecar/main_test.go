package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/content"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/workspace"
)

func newTestServer(t *testing.T) *server {
	t.Helper()
	tmpDir := t.TempDir()

	paths, err := workspace.ResolvePaths("gokatas-test")
	if err != nil {
		t.Fatalf("resolve paths: %v", err)
	}
	// Override paths to use temp dir
	paths.Data = tmpDir + "/data"
	paths.State = tmpDir + "/state"
	paths.Config = tmpDir + "/config"
	paths.Cache = tmpDir + "/cache"

	ws := workspace.NewManager(paths)
	_ = ws.Ensure()

	provider, err := content.NewProvider(content.ProviderConfig{
		ContentDir: tmpDir + "/content",
	})
	if err != nil {
		t.Fatalf("create provider: %v", err)
	}

	// Load track from the project's track.json
	track, err := catalog.LoadTrack("../../../../../tracks/go-core-100/track.json")
	if err != nil {
		t.Fatalf("load track: %v", err)
	}

	s := &server{
		contentDir: tmpDir + "/content",
		track:      track,
		ws:         ws,
		provider:   provider,
		progress:   progress.NewStore(paths.State + "/progress.json"),
	}
	return s
}

func TestHandleStatus(t *testing.T) {
	s := newTestServer(t)

	req := httptest.NewRequest("GET", "/api/status", nil)
	w := httptest.NewRecorder()
	s.handleStatus(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if _, ok := result["kataCount"]; !ok {
		t.Error("missing kataCount")
	}
	if _, ok := result["stageCount"]; !ok {
		t.Error("missing stageCount")
	}
	if _, ok := result["syncReady"]; !ok {
		t.Error("missing syncReady")
	}
	if _, ok := result["hasRunner"]; !ok {
		t.Error("missing hasRunner")
	}
}

func TestHandleCatalog(t *testing.T) {
	s := newTestServer(t)

	req := httptest.NewRequest("GET", "/api/catalog", nil)
	w := httptest.NewRecorder()
	s.handleCatalog(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("catalog = %d, want 200", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	kataCount, _ := result["kataCount"].(float64)
	if int(kataCount) == 0 {
		t.Error("catalog returned 0 katas")
	}

	stages, ok := result["stages"].([]interface{})
	if !ok || len(stages) == 0 {
		t.Error("catalog returned no stages")
	}
}

func TestHandleGetKataFound(t *testing.T) {
	s := newTestServer(t)

	req := httptest.NewRequest("GET", "/api/kata/001", nil)
	req.SetPathValue("id", "001")
	w := httptest.NewRecorder()
	s.handleGetKata(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("getKata = %d, want 200", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if _, ok := result["kata"]; !ok {
		t.Error("missing kata field")
	}
	if _, ok := result["content"]; !ok {
		t.Error("missing content field")
	}
}

func TestHandleGetKataNotFound(t *testing.T) {
	s := newTestServer(t)

	req := httptest.NewRequest("GET", "/api/kata/99999", nil)
	req.SetPathValue("id", "99999")
	w := httptest.NewRecorder()
	s.handleGetKata(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("getKata = %d, want 404", w.Code)
	}
}

func TestHandleProgressEmpty(t *testing.T) {
	s := newTestServer(t)

	req := httptest.NewRequest("GET", "/api/progress", nil)
	w := httptest.NewRecorder()
	s.handleProgress(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("progress = %d, want 200", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if _, ok := result["attempts"]; !ok {
		t.Error("missing attempts field")
	}
}

func TestHandleLintGo(t *testing.T) {
	s := newTestServer(t)

	body := `{"code": "package main\n\nfunc main() {}\n", "language": "go"}`
	req := httptest.NewRequest("POST", "/api/lint", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleLint(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("lint = %d, want 200", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if _, ok := result["diagnostics"]; !ok {
		t.Error("missing diagnostics field")
	}
}

func TestHandleLintBadGo(t *testing.T) {
	s := newTestServer(t)

	// Malformed Go code (missing closing brace)
	body := `{"code": "package main\n\nfunc main() {\n", "language": "go"}`
	req := httptest.NewRequest("POST", "/api/lint", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleLint(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("lint = %d, want 200", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	diags, _ := result["diagnostics"].([]interface{})
	if len(diags) == 0 {
		t.Error("expected diagnostics for malformed Go code")
	}
}

func TestHandleLintBadJSON(t *testing.T) {
	s := newTestServer(t)

	req := httptest.NewRequest("POST", "/api/lint", strings.NewReader("not json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.handleLint(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("lint = %d, want 400", w.Code)
	}
}

func TestHandleSaveKata(t *testing.T) {
	s := newTestServer(t)

	body := `{"code": "package kata\nfunc Hello() string { return \"hi\" }", "tests": "package kata\nfunc TestHello(t *testing.T) {}"}`
	req := httptest.NewRequest("POST", "/api/kata/001/save", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", "001")
	w := httptest.NewRecorder()
	s.handleSaveKata(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("save = %d, want 200", w.Code)
	}

	var result map[string]string
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if result["status"] != "saved" {
		t.Errorf("status = %q, want saved", result["status"])
	}
}

func TestHandleRunNoRunner(t *testing.T) {
	s := newTestServer(t)
	// runner is nil by default

	body := `{"code": "package kata", "tests": "package kata"}`
	req := httptest.NewRequest("POST", "/api/kata/001/run", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", "001")
	w := httptest.NewRecorder()
	s.handleRunKata(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("run = %d, want 503", w.Code)
	}
}

func TestCORSMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := corsMiddleware(mux)

	// OPTIONS preflight
	req := httptest.NewRequest("OPTIONS", "/api/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("CORS OPTIONS = %d, want 204", w.Code)
	}
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("missing CORS Allow-Origin header")
	}
}

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	writeJSON(w, map[string]string{"hello": "world"})

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
}

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
	writeError(w, http.StatusBadRequest, "bad request")

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
}
