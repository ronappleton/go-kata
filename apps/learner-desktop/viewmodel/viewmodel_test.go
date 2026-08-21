package viewmodel

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/katas"
)

// testCatalog sets up a minimal test track with embedded kata content.
func testCatalog(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	trackDir := filepath.Join(root, "tracks", "go-core-100")
	if err := os.MkdirAll(trackDir, 0o700); err != nil {
		t.Fatal(err)
	}

	trackJSON := `{
  "id": "go-core-100",
  "title": "Go Mastery",
  "stages": [
    {
      "id": "foundation",
      "title": "Foundation",
      "level": "junior",
      "categories": [
        {
          "id": "go-setup",
          "title": "Go Setup",
          "kata_ranges": [{"start": 0, "end": 2}]
        }
      ]
    }
  ]
}`
	if err := os.WriteFile(filepath.Join(trackDir, "track.json"), []byte(trackJSON), 0o600); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestNewViewModelDefaults(t *testing.T) {
	vm := New(desktopConfig{ContentRoot: t.TempDir()})
	if vm.Mode != ModeLinear {
		t.Fatalf("expected default mode Linear, got %v", vm.Mode)
	}
	if vm.FlashSide != "Front" {
		t.Fatalf("expected flashcard side Front, got %q", vm.FlashSide)
	}
}

func TestInitLoadsTrack(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	if len(vm.Track.Stages) == 0 {
		t.Fatal("expected at least one stage after Init")
	}
	if vm.Workspace == nil {
		t.Fatal("expected workspace manager to be set")
	}
	if vm.Progress == nil {
		t.Fatal("expected progress store to be set")
	}
}

func TestInitRejectsMissingContent(t *testing.T) {
	vm := New(desktopConfig{ContentRoot: t.TempDir()})
	if err := vm.Init(); err == nil {
		t.Fatal("expected Init to fail with missing content")
	}
}

func TestSelectKataPopulatesContent(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}

	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)

	if vm.Selected.ID != "000" {
		t.Fatalf("expected selected kata 000, got %q", vm.Selected.ID)
	}
	if vm.Readme == "" {
		t.Fatal("expected Readme to be populated")
	}
	if vm.Solution == "" {
		t.Fatal("expected Solution to be populated (should use starter)")
	}
}

func TestSelectKataPopulatesFlashcards(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)
	if len(vm.FlashDeck) == 0 {
		t.Fatal("expected flashcard deck to be populated")
	}
}

func TestSelectKataPopulatesQuizDeck(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)
	if len(vm.QuizDeck) == 0 {
		t.Fatal("expected quiz deck to be populated")
	}
}

func TestFlipFlashcard(t *testing.T) {
	vm := New(desktopConfig{})
	if vm.FlashSide != "Front" {
		t.Fatal("expected initial side Front")
	}
	vm.FlipFlashcard()
	if vm.FlashSide != "Back" {
		t.Fatal("expected side Back after flip")
	}
	vm.FlipFlashcard()
	if vm.FlashSide != "Front" {
		t.Fatal("expected side Front after second flip")
	}
}

func TestNextPrevFlashcard(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)
	n := len(vm.FlashDeck)
	if n == 0 {
		t.Fatal("expected flashcards")
	}
	vm.FlashDeckIdx = 0
	vm.PrevFlashcard()
	if vm.FlashDeckIdx != 0 {
		t.Fatal("PrevFlashcard should not go below 0")
	}
	vm.FlashDeckIdx = 0
	vm.NextFlashcard()
	if vm.FlashDeckIdx != 1 {
		t.Fatalf("expected index 1 after NextFlashcard, got %d", vm.FlashDeckIdx)
	}
	vm.FlashDeckIdx = n - 1
	vm.NextFlashcard()
	if vm.FlashDeckIdx != n-1 {
		t.Fatal("NextFlashcard should not exceed deck length")
	}
}

func TestRevealQuizAnswer(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)
	if vm.QuizAnswered {
		t.Fatal("expected not answered initially")
	}
	ans := vm.RevealQuizAnswer()
	if ans == "" {
		t.Fatal("expected non-empty answer")
	}
	if !vm.QuizAnswered {
		t.Fatal("expected answered after RevealQuizAnswer")
	}
}

func TestNextQuizQuestion(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)
	vm.QuizDeckIdx = 0
	vm.NextQuizQuestion()
	if vm.QuizDeckIdx != 1 {
		t.Fatalf("expected index 1, got %d", vm.QuizDeckIdx)
	}
	if vm.QuizAnswered {
		t.Fatal("expected quiz to reset answered state")
	}
}

func TestSetMode(t *testing.T) {
	vm := New(desktopConfig{})
	modeChanged := false
	vm.CB.OnModeChanged = func(mode LearningMode) {
		modeChanged = true
		if mode != ModeReview {
			t.Fatalf("expected Review mode, got %v", mode)
		}
	}
	vm.SetMode(ModeReview)
	if vm.Mode != ModeReview {
		t.Fatalf("expected mode Review, got %v", vm.Mode)
	}
	if !modeChanged {
		t.Fatal("expected OnModeChanged to be called")
	}
}

func TestSetStatusCallsCallback(t *testing.T) {
	vm := New(desktopConfig{})
	var got string
	vm.CB.OnStatusChanged = func(msg string) {
		got = msg
	}
	vm.setStatus("hello")
	if got != "hello" {
		t.Fatalf("expected status callback with %q, got %q", "hello", got)
	}
}

func TestSidebarItemsReturnsStages(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	items := vm.SidebarItems()
	if len(items) == 0 {
		t.Fatal("expected at least one sidebar stage")
	}
	if items[0].ID != "foundation" {
		t.Fatalf("expected first stage ID 'foundation', got %q", items[0].ID)
	}
	if len(items[0].Katas) == 0 {
		t.Fatal("expected at least one kata in the foundation stage")
	}
}

func TestProgressBarFraction(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	stage := vm.Track.Stages[0]
	completed, total := vm.StageProgress(stage)
	if total == 0 {
		t.Fatal("expected total > 0")
	}
	if completed != 0 {
		t.Fatalf("expected 0 completed, got %d", completed)
	}
}

func TestLearningModeString(t *testing.T) {
	if ModeLinear.String() != "Linear" {
		t.Fatal("expected Linear")
	}
	if ModeADHD.String() != "ADHD" {
		t.Fatal("expected ADHD")
	}
	if ModeReview.String() != "Review" {
		t.Fatal("expected Review")
	}
}

func TestSaveCurrentRequiresKata(t *testing.T) {
	vm := New(desktopConfig{})
	// Should not panic when no kata is selected
	vm.SaveCurrent()
}

func TestCurrentFlashDeckReturnsCorrectMode(t *testing.T) {
	root := testCatalog(t)
	vm := New(desktopConfig{ContentRoot: root})
	if err := vm.Init(); err != nil {
		t.Fatal(err)
	}
	kata, _, ok := vm.Track.FindKata("000")
	if !ok {
		t.Fatal("expected to find kata 000")
	}
	vm.SelectKata(kata)
	vm.SetMode(ModeReview)
	vm.BuildReviewDecks()
	deck := vm.CurrentFlashDeck()
	if len(deck) == 0 {
		// Review deck may be empty if katas have no flashcards
		// That's acceptable
	}
	vm.SetMode(ModeLinear)
	deck = vm.CurrentFlashDeck()
	if len(deck) == 0 {
		t.Fatal("expected flashcard deck from current kata")
	}
}

func TestQuizDeckGenerationFromRules(t *testing.T) {
	kata := catalog.Kata{
		ID:       "999",
		Title:    "Test Kata",
		Signature: "func Test() string",
		Rules:    []string{"input => expected output"},
		Content: katas.KataContent{
			KataGo: "package kata",
			KataTest: "package kata",
			Readme: "# Test",
			JSON: `{"id":"999","title":"Test","rules":["input => expected output"]}`,
		},
	}
	vm := New(desktopConfig{})
	vm.Selected = kata
	vm.populateCurrentQuizDeck()
	if len(vm.QuizDeck) != 1 {
		t.Fatalf("expected 1 quiz question from rules, got %d", len(vm.QuizDeck))
	}
	if !strings.Contains(vm.QuizDeck[0].Answer, "expected output") {
		t.Fatalf("expected answer to contain 'expected output', got %q", vm.QuizDeck[0].Answer)
	}
}

func TestFlashcardDeckGenerationFromRules(t *testing.T) {
	kata := catalog.Kata{
		ID:        "999",
		Title:     "Test Kata",
		Signature: "func Test() string",
		Focus:     "testing",
		Rules:     []string{"input => expected output"},
		Content: katas.KataContent{
			KataGo: "package kata",
			KataTest: "package kata",
			Readme: "# Test",
			JSON: `{"id":"999","title":"Test","rules":["input => expected output"]}`,
		},
	}
	vm := New(desktopConfig{})
	vm.Selected = kata
	vm.populateCurrentFlashDeck()
	// 1 signature + 1 focus + 1 rule = 3
	if len(vm.FlashDeck) != 3 {
		t.Fatalf("expected 3 flashcards from rules, got %d", len(vm.FlashDeck))
	}
}
