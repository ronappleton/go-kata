package content

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestRemoteFetchRetriesTransientFailures(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&attempts, 1) <= 2 {
			http.Error(w, "boom", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	store := NewRemoteStore(srv.URL)
	data, err := store.fetch(context.Background(), "manifest.json")
	if err != nil {
		t.Fatalf("fetch after retries failed: %v", err)
	}
	if string(data) != `{"ok":true}` {
		t.Fatalf("unexpected body %q", data)
	}
	if attempts != 3 {
		t.Fatalf("expected 3 attempts (2 failures + 1 success), got %d", attempts)
	}
}

func TestRemoteFetchDoesNotRetryNotFound(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		http.NotFound(w, r)
	}))
	defer srv.Close()

	store := NewRemoteStore(srv.URL)
	if _, err := store.fetch(context.Background(), "tracks/x/katas/999.json"); err == nil {
		t.Fatal("expected error for missing kata")
	}
	if attempts != 1 {
		t.Fatalf("expected no retry on 404, got %d attempts", attempts)
	}
}

func TestRemoteFetchStopsWhenContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	store := NewRemoteStore("http://127.0.0.1:1")
	if _, err := store.fetch(ctx, "manifest.json"); err == nil {
		t.Fatal("expected context-cancelled error")
	}
}
