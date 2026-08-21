package catalog

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/content"
)

type fakeContentProvider struct {
	failKatas map[string]bool // kata IDs that should fail to load
}

func (f fakeContentProvider) GetManifest(context.Context) (*content.Manifest, error) { return nil, nil }
func (f fakeContentProvider) GetTrack(context.Context, string) (*content.TrackMeta, error) {
	return &content.TrackMeta{
		ID: "go-core-100", Title: "Go Mastery",
		Stages: []content.StageMeta{{ID: "foundation", Title: "Foundation", Level: "junior", Categories: []content.CategoryMeta{{ID: "basics", Title: "Basics", KataIDs: []string{"001", "002"}}}}},
	}, nil
}
func (f fakeContentProvider) GetKata(_ context.Context, _ string, kataID string) (*content.KataContent, error) {
	if f.failKatas[kataID] {
		return nil, fmt.Errorf("kata %s unavailable", kataID)
	}
	return &content.KataContent{ID: kataID, Slug: "build-greeting", KataGo: "package kata", Readme: "# Build Greeting", JSON: `{"id":"001","title":"Build Greeting","focus":"functions","signature":"func BuildGreeting(name string) string","evaluator_status":"ready"}`}, nil
}
func (fakeContentProvider) Sync(context.Context) (*content.SyncResult, error) { return nil, nil }
func (fakeContentProvider) IsRemote() bool                                    { return true }
func (fakeContentProvider) LastSync() (result time.Time)                      { return result }
func (fakeContentProvider) ContentDir() string                                { return "" }

func TestLoadTrackFromContent(t *testing.T) {
	track, err := LoadTrackFromContent(context.Background(), fakeContentProvider{}, "go-core-100")
	if err != nil {
		t.Fatal(err)
	}
	kata, _, ok := track.FindKata("001")
	if !ok || kata.Title != "Build Greeting" || kata.Content.Readme == "" {
		t.Fatalf("unexpected converted kata: %+v", kata)
	}
}

func TestLoadTrackFromContentToleratesMissingKatas(t *testing.T) {
	track, err := LoadTrackFromContent(context.Background(), fakeContentProvider{failKatas: map[string]bool{"002": true}}, "go-core-100")
	if err == nil {
		t.Fatal("expected partial-failure error to be reported")
	}
	if len(track.AllKatas()) != 1 {
		t.Fatalf("expected 1 surviving kata, got %d", len(track.AllKatas()))
	}
	if _, _, ok := track.FindKata("001"); !ok {
		t.Fatalf("kata 001 should still be present")
	}
}

func TestLoadTrackFromContentFailsWhenNothingLoads(t *testing.T) {
	_, err := LoadTrackFromContent(context.Background(), fakeContentProvider{failKatas: map[string]bool{"001": true, "002": true}}, "go-core-100")
	if err == nil {
		t.Fatal("expected error when no katas load")
	}
}
