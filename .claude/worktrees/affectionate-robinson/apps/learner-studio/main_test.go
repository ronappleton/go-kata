package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
)

func TestTrackEndpoint(t *testing.T) {
	server := mustTestServer(t)
	req := httptest.NewRequest(http.MethodGet, "/api/track", nil)
	rec := httptest.NewRecorder()

	server.routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var payload trackResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.ID != "go-core-100" {
		t.Fatalf("unexpected track id: %s", payload.ID)
	}
	if len(payload.Categories) == 0 {
		t.Fatalf("expected categories in track payload")
	}
	if payload.CoachMessage == "" {
		t.Fatalf("expected coach message")
	}
	if payload.NextRecommended == nil {
		t.Fatalf("expected next recommendation")
	}
	if payload.Categories[0].MilestoneLabel == "" {
		t.Fatalf("expected category milestone data")
	}
}

func TestKataEndpoint(t *testing.T) {
	server := mustTestServer(t)
	req := httptest.NewRequest(http.MethodGet, "/api/kata?id=001", nil)
	rec := httptest.NewRecorder()

	server.routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var payload kataResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.ID != "001" {
		t.Fatalf("unexpected kata id: %s", payload.ID)
	}
	if payload.Code == "" {
		t.Fatalf("expected kata code")
	}
	if payload.Tests == "" {
		t.Fatalf("expected kata tests")
	}
}

func TestPathwaysEndpoint(t *testing.T) {
	server := mustTestServer(t)
	req := httptest.NewRequest(http.MethodGet, "/api/pathways", nil)
	rec := httptest.NewRecorder()

	server.routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}

	var payload pathwaysResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(payload.Items) == 0 {
		t.Fatalf("expected pathways")
	}
	if payload.Items[0].ID == "" || payload.Items[0].Title == "" {
		t.Fatalf("expected pathway metadata")
	}
}

func TestLearnEndpoint(t *testing.T) {
	server := mustTestServer(t)
	req := httptest.NewRequest(http.MethodGet, "/api/learn?id=001", nil)
	rec := httptest.NewRecorder()

	server.routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}

	var payload learnResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.KataID != "001" {
		t.Fatalf("unexpected kata id: %s", payload.KataID)
	}
	if len(payload.Flashcards) == 0 {
		t.Fatalf("expected flashcards")
	}
	if len(payload.Quiz) == 0 {
		t.Fatalf("expected quiz questions")
	}

	for _, q := range payload.Quiz {
		if len(q.Options) < 2 {
			t.Fatalf("expected at least 2 options, got %d for %s", len(q.Options), q.ID)
		}
		if q.AnswerIndex < 0 || q.AnswerIndex >= len(q.Options) {
			t.Fatalf("answer index out of range for %s", q.ID)
		}
	}
}

func TestRunEndpointIncludesCoachingFields(t *testing.T) {
	server := mustTestServer(t)
	body := `{"kata_id":"001","save_before_run":false,"timeout_seconds":30}`
	req := httptest.NewRequest(http.MethodPost, "/api/run", strings.NewReader(body))
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	server.routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}

	var payload runResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload.CoachHint == "" {
		t.Fatalf("expected coach hint")
	}
	if payload.NextRecommended == nil {
		t.Fatalf("expected next recommendation")
	}
}

func TestResetBuggyEndpoint(t *testing.T) {
	server := mustTestServer(t)
	body := `{"kata_id":"131"}`
	req := httptest.NewRequest(http.MethodPost, "/api/reset-buggy", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}

	var payload formatResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload.Code == "" || payload.Tests == "" {
		t.Fatalf("expected code/tests in response")
	}
	if strings.Contains(payload.Code, "//go:build ignore") {
		t.Fatalf("expected stripped code, got build tags in payload")
	}
}

func TestFormatEndpoint(t *testing.T) {
	server := mustTestServer(t)
	body := `{"kata_id":"001","code":"package kata001\nfunc  X( )int{return 1}\n","tests":"package kata001\nimport \"testing\"\nfunc TestX( t *testing.T){}"}`
	req := httptest.NewRequest(http.MethodPost, "/api/format", strings.NewReader(body))
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	server.routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}

	var payload formatResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !strings.Contains(payload.Code, "func X() int") {
		t.Fatalf("expected formatted code, got:\n%s", payload.Code)
	}
	if !strings.Contains(payload.Tests, "func TestX(t *testing.T)") {
		t.Fatalf("expected formatted tests, got:\n%s", payload.Tests)
	}
}

func mustTestServer(t *testing.T) *studioServer {
	t.Helper()

	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("resolve repo root: %v", err)
	}

	server, err := newStudioServer(repoRoot)
	if err != nil {
		t.Fatalf("newStudioServer: %v", err)
	}
	server.store = progress.NewStore(filepath.Join(t.TempDir(), "progress.json"))
	return server
}

func TestExtractFailureInsights(t *testing.T) {
	output := strings.Join([]string{
		"--- FAIL: TestFizzBuzz (0.00s)",
		"    kata_test.go:20: got: Buzz",
		"    kata_test.go:21: want: Fizz",
	}, "\n")

	got := extractFailureInsights(output)
	if len(got) == 0 {
		t.Fatalf("expected insights, got none")
	}
	foundMismatch := false
	for _, item := range got {
		if item.Kind == "mismatch" && item.Expected != "" && item.Actual != "" {
			foundMismatch = true
			break
		}
	}
	if !foundMismatch {
		t.Fatalf("expected mismatch insight with expected/actual values, got %+v", got)
	}
}

func TestCategoryMilestone(t *testing.T) {
	label, message, next, remaining := categoryMilestone(5, 20)
	if label == "" || message == "" {
		t.Fatalf("expected non-empty milestone label/message")
	}
	if next <= 0 {
		t.Fatalf("expected positive next target, got %d", next)
	}
	if remaining < 0 {
		t.Fatalf("expected non-negative remaining, got %d", remaining)
	}
}
