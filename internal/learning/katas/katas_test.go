package katas

import (
	"encoding/json"
	"testing"
)

func TestContentHasKatas(t *testing.T) {
	if len(Content) == 0 {
		t.Fatal("Content map is empty — embedded katas missing")
	}
}

func TestKataContentFields(t *testing.T) {
	for id, kc := range Content {
		if kc.ID != id {
			t.Errorf("kata %s: ID field %q mismatch", id, kc.ID)
		}
		if kc.Slug == "" {
			t.Errorf("kata %s: Slug is empty", id)
		}
		if kc.KataGo == "" {
			t.Errorf("kata %s: KataGo is empty", id)
		}
		if kc.Readme == "" {
			t.Errorf("kata %s: Readme is empty", id)
		}
	}
}

func TestKataJSONIsValidMeta(t *testing.T) {
	for id, kc := range Content {
		if kc.JSON == "" {
			t.Logf("kata %s: no JSON metadata (skipped)", id)
			continue
		}
		var meta KataMeta
		if err := json.Unmarshal([]byte(kc.JSON), &meta); err != nil {
			t.Errorf("kata %s: JSON parse error: %v", id, err)
			continue
		}
		if meta.ID != id {
			t.Errorf("kata %s: JSON id field %q mismatch", id, meta.ID)
		}
		if meta.Title == "" {
			t.Errorf("kata %s: JSON title is empty", id)
		}
		if meta.Focus == "" {
			t.Errorf("kata %s: JSON focus is empty", id)
		}
	}
}

func TestSpecificKnownKatas(t *testing.T) {
	// Kata 001 should always exist
	kc, ok := Content["001"]
	if !ok {
		t.Fatal("kata 001 not found in Content")
	}
	if kc.Slug != "build-greeting" {
		t.Errorf("kata 001 slug = %q, want build-greeting", kc.Slug)
	}
	if kc.KataGo == "" {
		t.Error("kata 001 KataGo is empty")
	}
}

func TestAllKatasHaveMatchingTestContent(t *testing.T) {
	for id, kc := range Content {
		if kc.KataGo != "" && kc.KataTest == "" {
			t.Errorf("kata %s: has KataGo but no KataTest", id)
		}
	}
}

func TestKataMetaMinimalFields(t *testing.T) {
	for id, kc := range Content {
		if kc.JSON == "" {
			continue
		}
		var meta KataMeta
		if err := json.Unmarshal([]byte(kc.JSON), &meta); err != nil {
			continue
		}
		if meta.Rules == nil {
			t.Errorf("kata %s: Rules is nil in JSON", id)
		}
		if meta.EvaluatorStatus == "" {
			t.Errorf("kata %s: EvaluatorStatus is empty in JSON", id)
		}
	}
}
