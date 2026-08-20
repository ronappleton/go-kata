//go:build gtk4

package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/workspace"
)

// ── Flashcards ──

func (n *nativeApp) buildFlashcards() *gtk.Widget {
	box := gtk.NewBox(gtk.OrientationVertical, 12)
	heading := gtk.NewLabel("Flashcards")
	heading.AddCSSClass("heading")
	heading.SetHAlign(gtk.AlignStart)
	n.flashSide = gtk.NewLabel("Front")
	n.flashSide.SetHAlign(gtk.AlignStart)
	n.flashText = gtk.NewLabel("Select a kata to begin recall.")
	n.flashText.AddCSSClass("body")
	n.flashText.SetWrap(true)
	n.flashText.SetHAlign(gtk.AlignStart)
	box.Append(heading)
	n.flashSide.AddCSSClass("flash-side")
	box.Append(n.flashSide)
	box.Append(n.flashText)

	// Confidence buttons
	controls := gtk.NewBox(gtk.OrientationHorizontal, 8)
	previous := gtk.NewButtonWithLabel("Previous")
	previous.ConnectClicked(func() {
		if n.flashDeckIdx > 0 {
			n.flashDeckIdx--
			n.flashSide.SetText("Front")
			n.updateFlashcard()
		}
	})
	flip := gtk.NewButtonWithLabel("Flip")
	flip.ConnectClicked(func() {
		if n.flashSide.Text() == "Front" {
			n.flashSide.SetText("Back")
		} else {
			n.flashSide.SetText("Front")
		}
		n.updateFlashcard()
	})
	next := gtk.NewButtonWithLabel("Next")
	next.ConnectClicked(func() {
		if n.flashDeckIdx < len(n.flashDeck)-1 {
			n.flashDeckIdx++
			n.flashSide.SetText("Front")
			n.updateFlashcard()
		}
	})
	controls.Append(previous)
	controls.Append(flip)
	controls.Append(next)
	box.Append(controls)

	// Confidence buttons for spaced repetition
	confBox := gtk.NewBox(gtk.OrientationHorizontal, 6)
	again := gtk.NewButtonWithLabel("Again")
	again.AddCSSClass("confidence-again")
	again.ConnectClicked(func() { n.flashConfidence(0) })
	hard := gtk.NewButtonWithLabel("Hard")
	hard.AddCSSClass("confidence-hard")
	hard.ConnectClicked(func() { n.flashConfidence(1) })
	good := gtk.NewButtonWithLabel("Good")
	good.AddCSSClass("confidence-good")
	good.ConnectClicked(func() { n.flashConfidence(2) })
	easy := gtk.NewButtonWithLabel("Easy")
	easy.AddCSSClass("confidence-easy")
	easy.ConnectClicked(func() { n.flashConfidence(3) })
	confBox.Append(again)
	confBox.Append(hard)
	confBox.Append(good)
	confBox.Append(easy)
	box.Append(confBox)

	return &box.Widget
}

func (n *nativeApp) buildFlashDeck() {
	n.flashDeck = nil
	n.flashDeckIdx = 0

	if n.mode == ModeReview {
		// Cross-kata deck: pull from all katas with flashcards
		for _, kata := range n.track.AllKatas() {
			if len(kata.Flashcards) > 0 {
				for _, fc := range kata.Flashcards {
					n.flashDeck = append(n.flashDeck, flashcardEntry{
						kataID: kata.ID,
						front:  fc.Front,
						back:   fc.Back,
					})
				}
			}
		}
		// Shuffle for review mode
		rand.Shuffle(len(n.flashDeck), func(i, j int) {
			n.flashDeck[i], n.flashDeck[j] = n.flashDeck[j], n.flashDeck[i]
		})
		// Limit to 20 cards
		if len(n.flashDeck) > 20 {
			n.flashDeck = n.flashDeck[:20]
		}
	} else {
		// Per-kata deck from current selection
		n.populateCurrentFlashDeck()
	}
}

func (n *nativeApp) populateCurrentFlashDeck() {
	n.flashDeck = nil
	n.flashDeckIdx = 0
	if n.selected.ID == "" {
		return
	}
	// Use pre-written flashcards from kata metadata
	if len(n.selected.Flashcards) > 0 {
		for _, fc := range n.selected.Flashcards {
			n.flashDeck = append(n.flashDeck, flashcardEntry{
				kataID: n.selected.ID,
				front:  fc.Front,
				back:   fc.Back,
			})
		}
		return
	}
	// Fallback: generate from rules
	if n.selected.Signature != "" {
		n.flashDeck = append(n.flashDeck, flashcardEntry{
			kataID: n.selected.ID,
			front:  "What function contract are you implementing?",
			back:   n.selected.Signature,
		})
	}
	if n.selected.Focus != "" {
		n.flashDeck = append(n.flashDeck, flashcardEntry{
			kataID: n.selected.ID,
			front:  "What skill is this kata training?",
			back:   n.selected.Focus,
		})
	}
	for _, rule := range n.selected.Rules {
		parts := strings.SplitN(rule, "=>", 2)
		front := "What behavior is required?"
		back := rule
		if len(parts) == 2 {
			front = "What should happen when " + strings.TrimSpace(parts[0]) + "?"
			back = strings.TrimSpace(parts[1])
		}
		n.flashDeck = append(n.flashDeck, flashcardEntry{
			kataID: n.selected.ID,
			front:  front,
			back:   back,
		})
	}
}

func (n *nativeApp) flashcardCount() int {
	return len(n.flashDeck)
}

func (n *nativeApp) updateFlashcard() {
	if n.flashText == nil || n.flashDeckIdx >= len(n.flashDeck) {
		return
	}
	entry := n.flashDeck[n.flashDeckIdx]
	if n.flashSide.Text() == "Back" {
		n.flashText.SetText(entry.back)
	} else {
		n.flashText.SetText(fmt.Sprintf("[%s] %s", entry.kataID, entry.front))
	}
}

func (n *nativeApp) flashConfidence(level int) {
	// In review mode, advance to next card after rating
	if n.mode == ModeReview && n.flashDeckIdx < len(n.flashDeck)-1 {
		n.flashDeckIdx++
		n.flashSide.SetText("Front")
		n.updateFlashcard()
	}
}

// ── Quiz ──

func (n *nativeApp) buildQuiz() *gtk.Widget {
	box := gtk.NewBox(gtk.OrientationVertical, 12)
	heading := gtk.NewLabel("Quiz")
	heading.AddCSSClass("heading")
	heading.SetHAlign(gtk.AlignStart)
	n.quizText = gtk.NewLabel("Select a kata to begin the knowledge check.")
	n.quizText.AddCSSClass("body")
	n.quizText.SetWrap(true)
	n.quizText.SetHAlign(gtk.AlignStart)
	n.quizFeedback = gtk.NewLabel("")
	n.quizFeedback.SetWrap(true)
	n.quizFeedback.SetHAlign(gtk.AlignStart)
	n.quizFeedback.AddCSSClass("quiz-feedback")
	box.Append(heading)
	box.Append(n.quizText)
	box.Append(n.quizFeedback)

	// Quiz controls
	controls := gtk.NewBox(gtk.OrientationHorizontal, 8)
	reveal := gtk.NewButtonWithLabel("Reveal answer")
	reveal.ConnectClicked(func() {
		if len(n.quizDeck) == 0 || n.quizDeckIdx >= len(n.quizDeck) {
			return
		}
		entry := n.quizDeck[n.quizDeckIdx]
		n.quizFeedback.SetText("Answer: " + entry.answer)
		n.quizAnswered = true
	})
	next := gtk.NewButtonWithLabel("Next question")
	next.ConnectClicked(func() {
		if n.quizDeckIdx < len(n.quizDeck)-1 {
			n.quizDeckIdx++
			n.quizFeedback.SetText("")
			n.quizAnswered = false
			n.updateQuiz()
		}
	})
	controls.Append(reveal)
	controls.Append(next)
	box.Append(controls)

	// Progress indicator
	progressLabel := gtk.NewLabel("")
	progressLabel.AddCSSClass("footnote")
	progressLabel.SetHAlign(gtk.AlignStart)
	n.quizText.SetMarginBottom(4)

	box.Append(progressLabel)
	return &box.Widget
}

func (n *nativeApp) buildQuizDeck() {
	n.quizDeck = nil
	n.quizDeckIdx = 0
	n.quizAnswered = false

	if n.mode == ModeReview {
		// Cross-kata quiz: pull from all katas
		for _, kata := range n.track.AllKatas() {
			if len(kata.QuizQuestions) > 0 {
				for _, q := range kata.QuizQuestions {
					n.quizDeck = append(n.quizDeck, quizEntry{
						kataID:   kata.ID,
						question: q.Question,
						options:  q.Options,
						answer:   q.Answer,
					})
				}
			}
		}
		rand.Shuffle(len(n.quizDeck), func(i, j int) {
			n.quizDeck[i], n.quizDeck[j] = n.quizDeck[j], n.quizDeck[i]
		})
		if len(n.quizDeck) > 15 {
			n.quizDeck = n.quizDeck[:15]
		}
	} else {
		n.populateCurrentQuizDeck()
	}
}

func (n *nativeApp) populateCurrentQuizDeck() {
	n.quizDeck = nil
	n.quizDeckIdx = 0
	n.quizAnswered = false
	if n.selected.ID == "" {
		return
	}
	// Use pre-written quiz questions
	if len(n.selected.QuizQuestions) > 0 {
		for _, q := range n.selected.QuizQuestions {
			n.quizDeck = append(n.quizDeck, quizEntry{
				kataID:   n.selected.ID,
				question: q.Question,
				options:  q.Options,
				answer:   q.Answer,
			})
		}
		return
	}
	// Fallback: generate from rules
	for _, rule := range n.selected.Rules {
		parts := strings.SplitN(rule, "=>", 2)
		question := "Which behavior is required by the kata?"
		answer := rule
		if len(parts) == 2 {
			question = "What is the correct outcome when " + strings.TrimSpace(parts[0]) + "?"
			answer = strings.TrimSpace(parts[1])
		}
		n.quizDeck = append(n.quizDeck, quizEntry{
			kataID:   n.selected.ID,
			question: question,
			answer:   answer,
		})
	}
}

func (n *nativeApp) quizCard() (prompt, answer string) {
	if n.quizDeckIdx >= len(n.quizDeck) {
		return "No quiz questions available.", ""
	}
	entry := n.quizDeck[n.quizDeckIdx]
	prefix := ""
	if n.mode == ModeReview {
		prefix = fmt.Sprintf("[%s] ", entry.kataID)
	}
	if len(entry.options) > 0 {
		optionsStr := strings.Join(entry.options, "  |  ")
		return prefix + entry.question + "\n\n" + optionsStr, entry.answer
	}
	return prefix + entry.question, entry.answer
}

func (n *nativeApp) updateQuiz() {
	if n.quizText == nil {
		return
	}
	prompt, _ := n.quizCard()
	total := len(n.quizDeck)
	current := n.quizDeckIdx + 1
	n.quizText.SetText(fmt.Sprintf("%s\n\n— Question %d of %d —", prompt, current, total))
}

// ── Bug Hunt ──

func (n *nativeApp) buildBugHunt() *gtk.Widget {
	box := gtk.NewBox(gtk.OrientationVertical, 12)
	heading := gtk.NewLabel("Bug Hunt")
	heading.AddCSSClass("heading")
	heading.SetHAlign(gtk.AlignStart)
	n.bugText = gtk.NewLabel("Select a kata to inspect its evaluator status.")
	n.bugText.AddCSSClass("body")
	n.bugText.SetWrap(true)
	n.bugText.SetHAlign(gtk.AlignStart)
	box.Append(heading)
	box.Append(n.bugText)
	return &box.Widget
}

// ── Reflection ──

func (n *nativeApp) buildReflection() *gtk.Widget {
	box := gtk.NewBox(gtk.OrientationVertical, 8)
	heading := gtk.NewLabel("Reflection")
	heading.AddCSSClass("heading")
	heading.SetHAlign(gtk.AlignStart)
	prompt := gtk.NewLabel("What failed first, what fixed it, and what will you carry into the next kata?")
	prompt.AddCSSClass("body")
	prompt.SetWrap(true)
	prompt.SetHAlign(gtk.AlignStart)
	n.reflection = gtk.NewTextBuffer(nil)
	view := gtk.NewTextViewWithBuffer(n.reflection)
	view.SetWrapMode(gtk.WrapWordChar)
	view.SetVExpand(true)
	view.AddCSSClass("editor")
	scroll := gtk.NewScrolledWindow()
	scroll.SetChild(view)
	scroll.SetVExpand(true)
	save := gtk.NewButtonWithLabel("Save reflection")
	save.ConnectClicked(func() {
		if n.selected.ID == "" || n.workspace == nil {
			return
		}
		root, err := n.workspace.Workspace(n.selected.ID)
		if err != nil {
			n.setStatus(fmt.Sprintf("Reflection save failed: %v", err))
			return
		}
		if err := workspace.AtomicWrite(filepath.Join(root, "reflection.md"), []byte(n.bufferText(n.reflection)), 0o600); err != nil {
			n.setStatus(fmt.Sprintf("Reflection save failed: %v", err))
			return
		}
		n.setStatus("Reflection saved to the user workspace.")
	})
	box.Append(heading)
	box.Append(prompt)
	box.Append(scroll)
	box.Append(save)
	return &box.Widget
}

// ── Mode update ──

func (n *nativeApp) updateModes() {
	// Rebuild flashcard and quiz decks for current kata
	if n.mode == ModeReview {
		n.buildFlashDeck()
		n.buildQuizDeck()
	} else {
		n.populateCurrentFlashDeck()
		n.populateCurrentQuizDeck()
	}
	n.flashIndex = 0
	n.quizIndex = 0

	// Update flashcard display
	if n.flashDeckIdx < len(n.flashDeck) {
		n.flashSide.SetText("Front")
		n.updateFlashcard()
	} else if n.flashText != nil {
		n.flashText.SetText("No flashcards available for this kata.")
	}

	// Update quiz display
	if n.quizDeckIdx < len(n.quizDeck) {
		n.updateQuiz()
	} else if n.quizText != nil {
		n.quizText.SetText("No quiz questions available for this kata.")
	}

	// Bug hunt
	if n.bugText != nil {
		n.bugText.SetText(fmt.Sprintf("%s is in the %s evaluator state.\n\nUse Workbench to reproduce failures, then run the trusted evaluator in the sandbox.", n.selected.Title, n.selected.EvaluatorStatus))
	}

	// Reflection
	if n.reflection != nil && n.workspace != nil {
		root, err := n.workspace.Workspace(n.selected.ID)
		if err == nil {
			data, err := os.ReadFile(filepath.Join(root, "reflection.md"))
			if err == nil {
				n.reflection.SetText(string(data))
			} else {
				n.reflection.SetText("")
			}
		}
	}
}
