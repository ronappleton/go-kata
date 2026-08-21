package content

import (
	"context"
	"testing"
)

func TestProviderRequiresRemoteOrCache(t *testing.T) {
	p, err := NewProvider(ProviderConfig{ContentDir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := p.GetManifest(context.Background()); err == nil {
		t.Fatal("expected an unavailable curriculum error")
	}
}
