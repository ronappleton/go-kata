package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
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
	return server
}
