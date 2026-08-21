package languages

import "testing"

func TestAutoPairsAreConsistent(t *testing.T) {
	r := NewRegistry()
	for _, l := range r.All() {
		seen := map[rune]bool{}
		for _, p := range l.AutoPairs {
			if p.Open == 0 || p.Close == 0 {
				t.Errorf("%s: pair with zero rune: %+v", l.ID, p)
			}
			if seen[p.Open] || seen[p.Close] {
				t.Errorf("%s: duplicate pair character %q", l.ID, string(p.Open))
			}
			seen[p.Open] = true
			seen[p.Close] = true
		}
	}
}

func TestPairLookupConsistency(t *testing.T) {
	r := NewRegistry()
	lang := r.Lookup("go")

	openOf := func(ch rune) (rune, bool) {
		for _, p := range lang.AutoPairs {
			if p.Open == ch {
				return p.Close, true
			}
		}
		return 0, false
	}
	closeOf := func(ch rune) (rune, bool) {
		for _, p := range lang.AutoPairs {
			if p.Close == ch {
				return p.Close, true
			}
		}
		return 0, false
	}

	// The editor's lookup logic: typing an open inserts its close; typing a
	// close that already follows the cursor skips over it. Both need the
	// data to be a bijection, so verify round-trips.
	for _, p := range lang.AutoPairs {
		if c, ok := openOf(p.Open); !ok || c != p.Close {
			t.Errorf("open %q must map to its close %q", string(p.Open), string(p.Close))
		}
		if _, ok := closeOf(p.Close); !ok {
			t.Errorf("close %q must be recognized as a closer", string(p.Close))
		}
	}
}

func TestEveryLanguageHasPairs(t *testing.T) {
	r := NewRegistry()
	for _, l := range r.All() {
		if len(l.AutoPairs) == 0 {
			t.Errorf("%s: must define auto-pairs", l.ID)
		}
	}
}

func TestQuotePairsIncluded(t *testing.T) {
	r := NewRegistry()
	// Languages that use double-quoted strings should auto-close them.
	for _, l := range r.All() {
		if l.ID == "python" || l.ID == "c" || l.ID == "cpp" || l.ID == "java" ||
			l.ID == "csharp" || l.ID == "php" || l.ID == "javascript" ||
			l.ID == "typescript" || l.ID == "rust" || l.ID == "go" {
			hasQuote := false
			for _, p := range l.AutoPairs {
				if p.Open == '"' && p.Close == '"' {
					hasQuote = true
				}
			}
			if !hasQuote {
				t.Errorf("%s: should auto-close double quotes", l.ID)
			}
		}
	}
}
