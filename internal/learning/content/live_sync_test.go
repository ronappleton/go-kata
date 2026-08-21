package content

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

// TestLiveRemoteSync exercises the real remote content source. It is skipped
// unless GOKATAS_LIVE_SYNC=1 so normal CI runs stay offline.
func TestLiveRemoteSync(t *testing.T) {
	if os.Getenv("GOKATAS_LIVE_SYNC") != "1" {
		t.Skip("set GOKATAS_LIVE_SYNC=1 to run the live remote sync test")
	}
	p, err := NewProvider(ProviderConfig{
		ContentDir: t.TempDir(),
		RemoteURL:  "https://raw.githubusercontent.com/ronappleton/gokatas-content/main",
	})
	if err != nil {
		t.Fatal(err)
	}
	var lastDone, lastTotal int
	p.SetProgress(func(done, total int) { lastDone, lastTotal = done, total })

	start := time.Now()
	res, err := p.Sync(context.Background())
	if err != nil {
		t.Fatalf("sync failed: %v", err)
	}
	t.Logf("sync: added=%d failed=%d in %s (progress saw %d/%d)", res.Added, len(res.Failed), time.Since(start).Round(time.Millisecond), lastDone, lastTotal)
	if res.Added == 0 {
		t.Fatal("expected katas to be downloaded")
	}
	if len(res.Failed) > 0 {
		t.Fatalf("unexpected failures: %v", res.Failed[:min(5, len(res.Failed))])
	}
	if lastTotal == 0 || lastDone != lastTotal {
		t.Fatalf("progress should reach %d/%d", lastTotal, lastTotal)
	}
	m, err := p.GetManifest(context.Background())
	if err != nil || len(m.Tracks) == 0 {
		t.Fatalf("manifest: %v tracks=%d", err, len(m.Tracks))
	}
	fmt.Printf("tracks: %d\n", len(m.Tracks))

	start = time.Now()
	res2, err := p.Sync(context.Background())
	if err != nil {
		t.Fatalf("resync failed: %v", err)
	}
	t.Logf("resync: added=%d failed=%d in %s", res2.Added, len(res2.Failed), time.Since(start).Round(time.Millisecond))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
