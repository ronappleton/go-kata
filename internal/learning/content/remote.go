package content

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// remoteStore fetches kata content from a remote HTTP source.
// Supports GitHub raw URLs, S3, or any static file server.
type remoteStore struct {
	baseURL    string
	httpClient *http.Client
}

// NewRemoteStore creates a content store backed by a remote HTTP server.
// baseURL example: "https://raw.githubusercontent.com/ronappleton/gokatas-content/main"
func NewRemoteStore(baseURL string) *remoteStore {
	return &remoteStore{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

const maxFetchAttempts = 4

// fetch retrieves path from the remote base URL, retrying transient failures
// (network errors, 5xx, 429) with exponential backoff. Permanent 4xx responses
// (e.g. a genuinely missing kata) are returned immediately.
func (r *remoteStore) fetch(ctx context.Context, path string) ([]byte, error) {
	return r.fetchURL(ctx, r.baseURL+"/"+path)
}

// fetchAttempt performs a single HTTP GET. The bool return reports whether the
// failure is permanent (a 4xx other than 429) and should not be retried.
func (r *remoteStore) fetchAttempt(ctx context.Context, url string) ([]byte, bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, true, fmt.Errorf("create request: %w", err)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("fetch %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		permanent := resp.StatusCode >= 400 && resp.StatusCode < 500 && resp.StatusCode != http.StatusTooManyRequests
		return nil, permanent, fmt.Errorf("fetch %s: status %d", url, resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, fmt.Errorf("read %s: %w", url, err)
	}
	return data, false, nil
}

// GetManifest fetches the manifest from the remote source.
func (r *remoteStore) GetManifest(ctx context.Context) (*Manifest, error) {
	data, err := r.fetch(ctx, "manifest.json")
	if err != nil {
		return nil, fmt.Errorf("fetch manifest: %w", err)
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parse manifest: %w", err)
	}
	return &m, nil
}

// GetTrack fetches a track definition from the remote source.
func (r *remoteStore) GetTrack(ctx context.Context, trackID string) (*TrackMeta, error) {
	data, err := r.fetch(ctx, "tracks/"+trackID+"/track.json")
	if err != nil {
		return nil, fmt.Errorf("fetch track %s: %w", trackID, err)
	}
	var t TrackMeta
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("parse track %s: %w", trackID, err)
	}
	return &t, nil
}

// GetKata fetches a kata's content from the remote source.
func (r *remoteStore) GetKata(ctx context.Context, trackID, kataID string) (*KataContent, error) {
	data, err := r.fetch(ctx, "tracks/"+trackID+"/katas/"+kataID+".json")
	if err != nil {
		return nil, fmt.Errorf("fetch kata %s/%s: %w", trackID, kataID, err)
	}
	var k KataContent
	if err := json.Unmarshal(data, &k); err != nil {
		return nil, fmt.Errorf("parse kata %s/%s: %w", trackID, kataID, err)
	}
	return &k, nil
}

// BaseURL returns the remote base URL.
func (r *remoteStore) BaseURL() string { return r.baseURL }
