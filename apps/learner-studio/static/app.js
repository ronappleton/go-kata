const state = {
  track: null,
  selectedKataID: null,
  selectedKata: null,
  nextRecommended: null,
  activeTab: "docs",
  activeEditor: "code",
  theme: "light",
  focusMode: false,
  dirty: false,
  formatting: false,
  autoFormatTimer: null,
  learn: {
    flashcards: [],
    quiz: [],
  },
  flash: {
    index: 0,
    showingBack: false,
  },
  quiz: {
    index: 0,
    answers: {},
  },
  search: {
    tests: { query: "", matches: [], index: -1 },
    code: { query: "", matches: [], index: -1 },
  },
};

const el = {
  trackSummary: document.getElementById("track-summary"),
  categoryList: document.getElementById("category-list"),
  kataTitle: document.getElementById("kata-title"),
  kataSubtitle: document.getElementById("kata-subtitle"),
  coachMessage: document.getElementById("coach-message"),
  nextRecommendation: document.getElementById("next-recommendation"),

  readmeView: document.getElementById("readme-view"),

  testsEditor: document.getElementById("tests-editor"),
  testsLines: document.getElementById("tests-lines"),
  testsHighlight: document.getElementById("tests-highlight"),
  testsSearchInput: document.getElementById("tests-search-input"),
  testsSearchPrev: document.getElementById("tests-search-prev"),
  testsSearchNext: document.getElementById("tests-search-next"),
  testsSearchClear: document.getElementById("tests-search-clear"),
  testsSearchCount: document.getElementById("tests-search-count"),

  codeEditor: document.getElementById("code-editor"),
  codeLines: document.getElementById("code-lines"),
  codeHighlight: document.getElementById("code-highlight"),
  codeSearchInput: document.getElementById("code-search-input"),
  codeSearchPrev: document.getElementById("code-search-prev"),
  codeSearchNext: document.getElementById("code-search-next"),
  codeSearchClear: document.getElementById("code-search-clear"),
  codeSearchCount: document.getElementById("code-search-count"),

  runOutput: document.getElementById("run-output"),
  failureInsights: document.getElementById("failure-insights"),
  dirtyIndicator: document.getElementById("dirty-indicator"),

  saveBtn: document.getElementById("save-btn"),
  runBtn: document.getElementById("run-btn"),
  markBtn: document.getElementById("mark-btn"),
  themeBtn: document.getElementById("theme-btn"),
  focusBtn: document.getElementById("focus-btn"),
  formatCodeBtn: document.getElementById("format-code-btn"),
  formatTestsBtn: document.getElementById("format-tests-btn"),

  tabDocs: document.getElementById("tab-docs"),
  tabWorkbench: document.getElementById("tab-workbench"),
  tabFlashcards: document.getElementById("tab-flashcards"),
  tabQuiz: document.getElementById("tab-quiz"),
  panelDocs: document.getElementById("panel-docs"),
  panelWorkbench: document.getElementById("panel-workbench"),
  panelFlashcards: document.getElementById("panel-flashcards"),
  panelQuiz: document.getElementById("panel-quiz"),

  flashPrevBtn: document.getElementById("flash-prev-btn"),
  flashFlipBtn: document.getElementById("flash-flip-btn"),
  flashNextBtn: document.getElementById("flash-next-btn"),
  flashProgress: document.getElementById("flash-progress"),
  flashSide: document.getElementById("flash-side"),
  flashText: document.getElementById("flash-text"),
  flashTag: document.getElementById("flash-tag"),

  quizPrevBtn: document.getElementById("quiz-prev-btn"),
  quizNextBtn: document.getElementById("quiz-next-btn"),
  quizResetBtn: document.getElementById("quiz-reset-btn"),
  quizProgress: document.getElementById("quiz-progress"),
  quizPrompt: document.getElementById("quiz-prompt"),
  quizOptions: document.getElementById("quiz-options"),
  quizFeedback: document.getElementById("quiz-feedback"),
  quizScore: document.getElementById("quiz-score"),

  resetTestsBtn: document.getElementById("reset-tests-btn"),

  passModal: document.getElementById("pass-modal"),
  passCloseBtn: document.getElementById("pass-close-btn"),
  passMarkBtn: document.getElementById("pass-mark-btn"),
  passNextHint: document.getElementById("pass-next-hint"),
  passNextBtn: document.getElementById("pass-next-btn"),

  markModal: document.getElementById("mark-modal"),
  markCloseBtn: document.getElementById("mark-close-btn"),
  markPath: document.getElementById("mark-path"),
  markPrompt: document.getElementById("mark-prompt"),
  copyPromptBtn: document.getElementById("copy-prompt-btn"),
  openChatGPTLink: document.getElementById("open-chatgpt-link"),
};

const TAB_META = {
  docs: { button: el.tabDocs, panel: el.panelDocs },
  workbench: { button: el.tabWorkbench, panel: el.panelWorkbench },
  flashcards: { button: el.tabFlashcards, panel: el.panelFlashcards },
  quiz: { button: el.tabQuiz, panel: el.panelQuiz },
};

const editors = {
  tests: { input: el.testsEditor, layer: el.testsHighlight, lines: el.testsLines },
  code: { input: el.codeEditor, layer: el.codeHighlight, lines: el.codeLines },
};

const searchUI = {
  tests: {
    input: el.testsSearchInput,
    prev: el.testsSearchPrev,
    next: el.testsSearchNext,
    clear: el.testsSearchClear,
    count: el.testsSearchCount,
  },
  code: {
    input: el.codeSearchInput,
    prev: el.codeSearchPrev,
    next: el.codeSearchNext,
    clear: el.codeSearchClear,
    count: el.codeSearchCount,
  },
};

const GO_KEYWORDS = /\b(package|import|func|return|if|else|for|range|switch|case|default|type|struct|interface|map|chan|select|go|defer|var|const|break|continue|fallthrough)\b/g;
const GO_TYPES = /\b(int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|uintptr|float32|float64|string|bool|byte|rune|error|any)\b/g;
const GO_NUMBERS = /\b(\d+(?:\.\d+)?)\b/g;

boot();

async function boot() {
  initTheme();
  initFocusMode();
  wireEvents();
  setupSearch("tests");
  setupSearch("code");
  setupEditor("tests");
  setupEditor("code");
  renderSearchCount("tests");
  renderSearchCount("code");
  renderFlashcard();
  renderQuiz();

  const savedTab = localStorage.getItem("learner-studio-tab") || "docs";
  setTab(savedTab);

  await refreshTrack();
  const first = firstKataID();
  if (first) {
    await loadKata(first);
  }
}

function wireEvents() {
  el.tabDocs.addEventListener("click", () => setTab("docs"));
  el.tabWorkbench.addEventListener("click", () => setTab("workbench"));
  el.tabFlashcards.addEventListener("click", () => setTab("flashcards"));
  el.tabQuiz.addEventListener("click", () => setTab("quiz"));

  el.saveBtn.addEventListener("click", onSave);
  el.runBtn.addEventListener("click", onRun);
  el.markBtn.addEventListener("click", onMark);
  el.themeBtn.addEventListener("click", toggleTheme);
  el.focusBtn.addEventListener("click", toggleFocusMode);

  el.formatCodeBtn.addEventListener("click", () => onFormat("code"));
  el.formatTestsBtn.addEventListener("click", () => onFormat("tests"));

  el.flashPrevBtn.addEventListener("click", () => moveFlashcard(-1));
  el.flashNextBtn.addEventListener("click", () => moveFlashcard(1));
  el.flashFlipBtn.addEventListener("click", flipFlashcard);

  el.quizPrevBtn.addEventListener("click", () => moveQuiz(-1));
  el.quizNextBtn.addEventListener("click", () => moveQuiz(1));
  el.quizResetBtn.addEventListener("click", resetQuizAnswers);

  el.passCloseBtn.addEventListener("click", () => toggleModal(el.passModal, false));
  el.passModal.addEventListener("click", (event) => {
    if (event.target === el.passModal) {
      toggleModal(el.passModal, false);
    }
  });
  el.passMarkBtn.addEventListener("click", async () => {
    toggleModal(el.passModal, false);
    await onMark();
  });
  el.passNextBtn.addEventListener("click", async () => {
    if (!state.nextRecommended?.kata_id) {
      return;
    }
    toggleModal(el.passModal, false);
    await loadKata(state.nextRecommended.kata_id);
    setTab("docs");
  });

  el.markCloseBtn.addEventListener("click", () => toggleModal(el.markModal, false));
  el.markModal.addEventListener("click", (event) => {
    if (event.target === el.markModal) {
      toggleModal(el.markModal, false);
    }
  });
  el.copyPromptBtn.addEventListener("click", copyPrompt);
  el.resetTestsBtn.addEventListener("click", onResetFromDisk);

  document.addEventListener("keydown", async (event) => {
    if (event.key === "Escape") {
      if (!el.markModal.classList.contains("hidden")) {
        toggleModal(el.markModal, false);
      }
      if (!el.passModal.classList.contains("hidden")) {
        toggleModal(el.passModal, false);
      }
      return;
    }

    if (!state.selectedKataID) {
      return;
    }

    const ctrlOrMeta = event.ctrlKey || event.metaKey;

    if (ctrlOrMeta && event.shiftKey && event.key.toLowerCase() === "m") {
      event.preventDefault();
      toggleFocusMode();
      return;
    }

    if (ctrlOrMeta && event.key.toLowerCase() === "s") {
      event.preventDefault();
      await onSave();
      return;
    }

    if (ctrlOrMeta && event.key === "Enter") {
      event.preventDefault();
      await onRun();
      return;
    }

    if (event.altKey && event.shiftKey && event.key.toLowerCase() === "f") {
      event.preventDefault();
      await onFormat("both");
      return;
    }

    if (state.activeTab === "workbench" && ctrlOrMeta && event.key.toLowerCase() === "f") {
      event.preventDefault();
      focusSearch(state.activeEditor || "code");
      return;
    }

    if (state.activeTab === "workbench" && event.key === "F3") {
      event.preventDefault();
      moveSearchMatch(state.activeEditor || "code", event.shiftKey ? -1 : 1);
    }
  });
}

function setupSearch(name) {
  const ui = searchUI[name];
  ui.input.addEventListener("input", () => {
    state.search[name].query = ui.input.value;
    refreshSearchState(name, { keepIndex: false });
    refreshHighlight(name);
  });

  ui.input.addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
      event.preventDefault();
      moveSearchMatch(name, event.shiftKey ? -1 : 1);
      return;
    }
    if (event.key === "Escape") {
      event.preventDefault();
      clearSearch(name);
      editors[name].input.focus();
    }
  });

  ui.prev.addEventListener("click", () => moveSearchMatch(name, -1));
  ui.next.addEventListener("click", () => moveSearchMatch(name, 1));
  ui.clear.addEventListener("click", () => {
    clearSearch(name);
    editors[name].input.focus();
  });
}

function setupEditor(name) {
  const editor = editors[name];

  editor.input.addEventListener("focus", () => {
    state.activeEditor = name;
  });

  editor.input.addEventListener("input", () => {
    if (!state.selectedKata) {
      return;
    }
    state.dirty = true;
    refreshSearchState(name);
    renderDirtyState();
    refreshHighlight(name);
  });

  editor.input.addEventListener("scroll", () => syncScroll(name));

  editor.input.addEventListener("keydown", (event) => {
    if (event.key !== "Tab" || event.ctrlKey || event.metaKey || event.altKey) {
      return;
    }
    event.preventDefault();
    const node = editor.input;
    const start = node.selectionStart;
    const end = node.selectionEnd;
    const current = node.value;
    node.value = `${current.slice(0, start)}\t${current.slice(end)}`;
    node.selectionStart = start + 1;
    node.selectionEnd = start + 1;
    node.dispatchEvent(new Event("input", { bubbles: true }));
  });

  editor.input.addEventListener("blur", () => {
    if (!state.selectedKata || !state.dirty) {
      return;
    }
    scheduleAutoFormat(name);
  });

  refreshHighlight(name);
}

function syncScroll(name) {
  const editor = editors[name];
  editor.layer.scrollTop = editor.input.scrollTop;
  editor.layer.scrollLeft = editor.input.scrollLeft;
  editor.lines.scrollTop = editor.input.scrollTop;
}

function initTheme() {
  const stored = localStorage.getItem("learner-studio-theme");
  const preferred = window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
  applyTheme(stored || preferred);
}

function applyTheme(theme) {
  state.theme = theme === "dark" ? "dark" : "light";
  document.body.setAttribute("data-theme", state.theme);
  localStorage.setItem("learner-studio-theme", state.theme);
  el.themeBtn.textContent = state.theme === "dark" ? "Light mode" : "Dark mode";
}

function toggleTheme() {
  applyTheme(state.theme === "dark" ? "light" : "dark");
}

function initFocusMode() {
  const stored = localStorage.getItem("learner-studio-focus-mode");
  applyFocusMode(stored === "1");
}

function applyFocusMode(enabled) {
  state.focusMode = Boolean(enabled);
  document.body.classList.toggle("focus-mode", state.focusMode);
  el.focusBtn.classList.toggle("active", state.focusMode);
  el.focusBtn.textContent = state.focusMode ? "Exit focus" : "Focus mode";
  localStorage.setItem("learner-studio-focus-mode", state.focusMode ? "1" : "0");

  if (state.focusMode) {
    setTab("workbench");
  }
}

function toggleFocusMode() {
  applyFocusMode(!state.focusMode);
}

function scheduleAutoFormat(target) {
  if (state.autoFormatTimer) {
    clearTimeout(state.autoFormatTimer);
  }

  state.autoFormatTimer = setTimeout(async () => {
    await formatEditors({ silent: true, target });
  }, 180);
}

function focusSearch(name) {
  const ui = searchUI[name] || searchUI.code;
  ui.input.focus();
  ui.input.select();
}

function clearSearch(name) {
  state.search[name].query = "";
  state.search[name].matches = [];
  state.search[name].index = -1;
  searchUI[name].input.value = "";
  renderSearchCount(name);
  refreshHighlight(name);
}

function refreshSearchState(name, { keepIndex = true } = {}) {
  const entry = state.search[name];
  const previousIndex = entry.index;
  entry.matches = findMatches(editors[name].input.value, entry.query);

  if (entry.matches.length === 0) {
    entry.index = -1;
  } else if (keepIndex && previousIndex >= 0) {
    entry.index = Math.min(previousIndex, entry.matches.length - 1);
  } else {
    entry.index = -1;
  }

  renderSearchCount(name);
}

function renderSearchCount(name) {
  const entry = state.search[name];
  const ui = searchUI[name];
  const total = entry.matches.length;
  const current = total > 0 && entry.index >= 0 ? entry.index + 1 : 0;
  ui.count.textContent = `${current}/${total}`;
  ui.prev.disabled = total === 0;
  ui.next.disabled = total === 0;
  ui.clear.disabled = entry.query.length === 0;
}

function moveSearchMatch(name, direction) {
  const entry = state.search[name];
  if (entry.matches.length === 0) {
    renderSearchCount(name);
    return;
  }

  if (entry.index < 0) {
    entry.index = direction >= 0 ? 0 : entry.matches.length - 1;
  } else {
    const total = entry.matches.length;
    entry.index = (entry.index + direction + total) % total;
  }

  const match = entry.matches[entry.index];
  const input = editors[name].input;
  state.activeEditor = name;
  input.focus();
  input.setSelectionRange(match.start, match.end);
  renderSearchCount(name);
}

function findMatches(source, query) {
  const normalized = query.trim();
  if (!normalized) {
    return [];
  }

  const lowerSource = source.toLowerCase();
  const lowerQuery = normalized.toLowerCase();
  const matches = [];
  let from = 0;

  while (from <= lowerSource.length) {
    const start = lowerSource.indexOf(lowerQuery, from);
    if (start === -1) {
      break;
    }
    matches.push({ start, end: start + normalized.length });
    from = start + Math.max(1, normalized.length);
  }

  return matches;
}

function setTab(requestedTab) {
  let tab = requestedTab;
  if (!TAB_META[tab]) {
    tab = "docs";
  }
  if (state.focusMode && tab !== "workbench") {
    tab = "workbench";
  }

  state.activeTab = tab;
  localStorage.setItem("learner-studio-tab", tab);

  Object.entries(TAB_META).forEach(([key, meta]) => {
    const active = key === tab;
    meta.button.classList.toggle("active", active);
    meta.panel.classList.toggle("active", active);
  });
}

async function refreshTrack() {
  state.track = await api("/api/track");
  state.nextRecommended = state.track.next_recommended || null;
  renderTrack();
}

function renderTrack() {
  if (!state.track) {
    return;
  }

  el.trackSummary.textContent = `${state.track.overall_done}/${state.track.overall_total} complete (${state.track.overall_percent}%)`;
  el.coachMessage.textContent = state.track.coach_message || "Keep sessions focused: one kata, one reflection.";

  if (state.track.next_recommended) {
    const next = state.track.next_recommended;
    el.nextRecommendation.textContent = `Recommended next: ${next.kata_id} - ${next.kata_title}. ${next.reason}`;
  } else {
    el.nextRecommendation.textContent = "Track complete. Pick a kata to refactor for readability and edge-case tests.";
  }

  el.categoryList.innerHTML = "";

  state.track.categories.forEach((category) => {
    const card = document.createElement("section");
    card.className = "category-card";

    const head = document.createElement("div");
    head.className = "category-head";
    head.innerHTML = `
      <h4>${escapeHTML(category.title)}</h4>
      <span class="category-progress">${category.done}/${category.total} (${category.percent}%)</span>
    `;
    card.appendChild(head);

    const milestone = document.createElement("p");
    milestone.className = "category-milestone";
    milestone.textContent = `${category.milestone_label}: ${category.milestone_message}`;
    card.appendChild(milestone);

    const list = document.createElement("div");
    list.className = "kata-list";

    category.katas.forEach((kata) => {
      const btn = document.createElement("button");
      btn.className = "kata-item";
      if (kata.completed) {
        btn.classList.add("done");
      }
      if (state.selectedKataID === kata.id) {
        btn.classList.add("active");
      }
      btn.innerHTML = `<strong>${kata.id}</strong> ${escapeHTML(kata.title)}`;
      btn.addEventListener("click", () => loadKata(kata.id));
      list.appendChild(btn);
    });

    card.appendChild(list);
    el.categoryList.appendChild(card);
  });
}

function firstKataID() {
  if (!state.track || !state.track.categories.length) {
    return null;
  }
  const firstCategory = state.track.categories[0];
  if (!firstCategory.katas.length) {
    return null;
  }
  return firstCategory.katas[0].id;
}

async function loadKata(kataID) {
  const kata = await api(`/api/kata?id=${encodeURIComponent(kataID)}`);
  state.selectedKataID = kata.id;
  state.selectedKata = kata;
  state.activeEditor = "code";
  state.dirty = false;

  el.kataTitle.textContent = `${kata.id} - ${kata.title}`;
  el.kataSubtitle.textContent = `${kata.category.title} - ${kata.focus}`;
  el.readmeView.textContent = kata.readme;
  setEditorValue("tests", kata.tests);
  setEditorValue("code", kata.code);
  clearSearch("tests");
  clearSearch("code");
  el.runOutput.textContent = formatProgress(kata.progress);
  renderFailureInsights([]);
  renderDirtyState();
  renderButtons(true);
  renderTrack();

  await loadLearningModes(kata.id);
}

async function loadLearningModes(kataID) {
  try {
    const payload = await api(`/api/learn?id=${encodeURIComponent(kataID)}`);
    state.learn.flashcards = Array.isArray(payload.flashcards) ? payload.flashcards : [];
    state.learn.quiz = Array.isArray(payload.quiz) ? payload.quiz : [];
  } catch (error) {
    state.learn.flashcards = [];
    state.learn.quiz = [];
    el.runOutput.textContent = `Learning modes unavailable: ${error.message || String(error)}`;
  }

  state.flash.index = 0;
  state.flash.showingBack = false;
  state.quiz.index = 0;
  state.quiz.answers = {};
  renderFlashcard();
  renderQuiz();
}

function renderButtons(enabled) {
  el.saveBtn.disabled = !enabled;
  el.runBtn.disabled = !enabled;
  el.markBtn.disabled = !enabled;
  el.formatCodeBtn.disabled = !enabled;
  el.formatTestsBtn.disabled = !enabled;
  el.focusBtn.disabled = false;
}

function renderDirtyState() {
  el.dirtyIndicator.textContent = state.dirty ? "Unsaved changes" : "Saved";
  el.dirtyIndicator.classList.toggle("dirty", state.dirty);
}

function renderFlashcard() {
  const cards = state.learn.flashcards || [];
  const total = cards.length;

  if (total === 0) {
    el.flashProgress.textContent = "0/0";
    el.flashSide.textContent = "Front";
    el.flashText.textContent = "Select a kata to load flashcards.";
    el.flashTag.textContent = "";
    el.flashPrevBtn.disabled = true;
    el.flashNextBtn.disabled = true;
    el.flashFlipBtn.disabled = true;
    return;
  }

  if (state.flash.index >= total) {
    state.flash.index = total - 1;
  }
  if (state.flash.index < 0) {
    state.flash.index = 0;
  }

  const card = cards[state.flash.index];
  const showingBack = state.flash.showingBack;
  el.flashProgress.textContent = `${state.flash.index + 1}/${total}`;
  el.flashSide.textContent = showingBack ? "Back" : "Front";
  el.flashText.textContent = showingBack ? card.back : card.front;
  el.flashTag.textContent = card.tag ? `Tag: ${card.tag}` : "";

  el.flashPrevBtn.disabled = total <= 1;
  el.flashNextBtn.disabled = total <= 1;
  el.flashFlipBtn.disabled = false;
}

function moveFlashcard(direction) {
  const cards = state.learn.flashcards || [];
  if (cards.length === 0) {
    return;
  }
  state.flash.index = (state.flash.index + direction + cards.length) % cards.length;
  state.flash.showingBack = false;
  renderFlashcard();
}

function flipFlashcard() {
  if (!state.learn.flashcards || state.learn.flashcards.length === 0) {
    return;
  }
  state.flash.showingBack = !state.flash.showingBack;
  renderFlashcard();
}

function renderQuiz() {
  const questions = state.learn.quiz || [];
  const total = questions.length;

  if (total === 0) {
    el.quizProgress.textContent = "0/0";
    el.quizPrompt.textContent = "Select a kata to load quiz questions.";
    el.quizOptions.innerHTML = "";
    el.quizFeedback.textContent = "";
    el.quizScore.textContent = "";
    el.quizPrevBtn.disabled = true;
    el.quizNextBtn.disabled = true;
    el.quizResetBtn.disabled = true;
    return;
  }

  if (state.quiz.index >= total) {
    state.quiz.index = total - 1;
  }
  if (state.quiz.index < 0) {
    state.quiz.index = 0;
  }

  const question = questions[state.quiz.index];
  const selected = state.quiz.answers[state.quiz.index];

  el.quizProgress.textContent = `${state.quiz.index + 1}/${total}`;
  el.quizPrompt.textContent = question.prompt;
  el.quizOptions.innerHTML = "";

  question.options.forEach((option, idx) => {
    const btn = document.createElement("button");
    btn.className = "quiz-option";
    btn.textContent = option;

    if (selected !== undefined) {
      if (idx === question.answer_index) {
        btn.classList.add("correct");
      }
      if (selected === idx && idx !== question.answer_index) {
        btn.classList.add("wrong");
      }
    }

    btn.addEventListener("click", () => {
      state.quiz.answers[state.quiz.index] = idx;
      renderQuiz();
    });

    el.quizOptions.appendChild(btn);
  });

  if (selected === undefined) {
    el.quizFeedback.textContent = "Choose one option, then move to the next question.";
  } else if (selected === question.answer_index) {
    el.quizFeedback.textContent = `Correct. ${question.explanation}`;
  } else {
    const correct = question.options[question.answer_index] || "(missing option)";
    el.quizFeedback.textContent = `Not yet. Correct answer: ${correct}. ${question.explanation}`;
  }

  const answered = Object.keys(state.quiz.answers).length;
  let correctCount = 0;
  Object.entries(state.quiz.answers).forEach(([idxRaw, selectedIdx]) => {
    const idx = Number(idxRaw);
    if (questions[idx] && questions[idx].answer_index === selectedIdx) {
      correctCount += 1;
    }
  });

  el.quizScore.textContent = `Score: ${correctCount}/${total} correct (${answered} answered)`;
  el.quizPrevBtn.disabled = state.quiz.index <= 0;
  el.quizNextBtn.disabled = state.quiz.index >= total - 1;
  el.quizResetBtn.disabled = false;
}

function moveQuiz(direction) {
  if (!state.learn.quiz || state.learn.quiz.length === 0) {
    return;
  }
  state.quiz.index += direction;
  if (state.quiz.index < 0) {
    state.quiz.index = 0;
  }
  if (state.quiz.index >= state.learn.quiz.length) {
    state.quiz.index = state.learn.quiz.length - 1;
  }
  renderQuiz();
}

function resetQuizAnswers() {
  state.quiz.answers = {};
  state.quiz.index = 0;
  renderQuiz();
}

function formatProgress(progress) {
  if (!progress || !progress.last_result) {
    return "No run yet.";
  }
  const lines = [];
  lines.push(`Last result: ${progress.last_result.toUpperCase()}`);
  lines.push(`Attempts: ${progress.attempts} | Passes: ${progress.passes}`);
  if (progress.last_duration_ms > 0) {
    lines.push(`Duration: ${progress.last_duration_ms}ms`);
  }
  if (progress.last_failed_tests && progress.last_failed_tests.length > 0) {
    lines.push(`Failing tests: ${progress.last_failed_tests.join(", ")}`);
  }
  if (progress.last_output_tail) {
    lines.push("");
    lines.push(progress.last_output_tail);
  }
  return lines.join("\n");
}

function renderFailureInsights(insights) {
  if (!insights || insights.length === 0) {
    el.failureInsights.innerHTML = `<div class="insight-card"><strong>No focused issues yet</strong><div>Run tests to surface the first mismatch, then fix one thing at a time.</div></div>`;
    return;
  }

  el.failureInsights.innerHTML = "";
  insights.forEach((item) => {
    const card = document.createElement("div");
    card.className = "insight-card";

    let pair = "";
    if (item.expected || item.actual) {
      pair = `
        <div class="pair">
          ${item.expected ? `<div><b>Expected:</b> ${escapeHTML(item.expected)}</div>` : ""}
          ${item.actual ? `<div><b>Actual:</b> ${escapeHTML(item.actual)}</div>` : ""}
        </div>
      `;
    }

    card.innerHTML = `
      <strong>${escapeHTML(formatInsightKind(item.kind))}</strong>
      <div>${escapeHTML(item.summary || "Behavior mismatch detected.")}</div>
      ${pair}
    `;
    el.failureInsights.appendChild(card);
  });
}

function formatInsightKind(kind) {
  switch (kind) {
    case "mismatch":
      return "Value mismatch";
    case "test":
      return "Failing test";
    case "panic":
      return "Runtime panic";
    default:
      return "Feedback";
  }
}

async function onSave() {
  ensureSelectedKata();
  await formatEditors({ silent: true, target: "both" });

  await api("/api/save", {
    method: "POST",
    body: JSON.stringify({
      kata_id: state.selectedKataID,
      code: editors.code.input.value,
      tests: editors.tests.input.value,
    }),
  });

  state.dirty = false;
  renderDirtyState();
  el.runOutput.textContent = "Saved and formatted. Run tests to validate behavior.";
}

async function onRun() {
  ensureSelectedKata();
  await formatEditors({ silent: true, target: "both" });

  el.runOutput.textContent = "Running tests...";

  const result = await api("/api/run", {
    method: "POST",
    body: JSON.stringify({
      kata_id: state.selectedKataID,
      code: editors.code.input.value,
      tests: editors.tests.input.value,
      save_before_run: true,
      timeout_seconds: 90,
    }),
  });

  state.dirty = false;
  renderDirtyState();

  const lines = [];
  lines.push(result.passed ? "PASS" : "FAIL");
  lines.push(`Duration: ${result.duration_ms}ms`);
  if (result.coach_hint) {
    lines.push(`Coach hint: ${result.coach_hint}`);
  }
  if (result.failed_tests && result.failed_tests.length > 0) {
    lines.push(`Failing tests: ${result.failed_tests.join(", ")}`);
  }
  if (result.output_tail) {
    lines.push("");
    lines.push(result.output_tail);
  }
  el.runOutput.textContent = lines.join("\n");
  renderFailureInsights(result.failure_insights || []);

  state.nextRecommended = result.next_recommended || state.nextRecommended;
  if (state.nextRecommended) {
    el.passNextHint.textContent = `Recommended next: ${state.nextRecommended.kata_id} - ${state.nextRecommended.kata_title}`;
    el.passNextBtn.disabled = false;
  } else {
    el.passNextHint.textContent = "You completed the current track. Great work.";
    el.passNextBtn.disabled = true;
  }

  await refreshTrack();
  if (result.passed) {
    toggleModal(el.passModal, true);
  }
}

async function onMark() {
  ensureSelectedKata();

  const payload = await api("/api/mark", {
    method: "POST",
    body: JSON.stringify({ kata_id: state.selectedKataID }),
  });

  el.markPath.textContent = `Saved: ${payload.prompt_path}`;
  el.markPrompt.value = payload.prompt;
  el.openChatGPTLink.href = payload.chatgpt_url || "https://chatgpt.com/";
  toggleModal(el.markModal, true);
}

async function onFormat(side) {
  ensureSelectedKata();
  const target = side === "both" ? "both" : side;
  const changed = await formatEditors({ silent: false, target });
  if (changed) {
    state.dirty = true;
    renderDirtyState();
    el.runOutput.textContent = side === "both"
      ? "Formatted code and tests."
      : `Formatted ${side}.`;
  } else {
    el.runOutput.textContent = "Already formatted.";
  }
}

async function onResetFromDisk() {
  ensureSelectedKata();
  await loadKata(state.selectedKataID);
  el.runOutput.textContent = "Reloaded from disk.";
  renderFailureInsights([]);
}

async function formatEditors({ silent = false, target = "both" } = {}) {
  if (!state.selectedKataID || state.formatting) {
    return false;
  }

  const formatCode = target === "both" || target === "code";
  const formatTests = target === "both" || target === "tests";
  if (!formatCode && !formatTests) {
    return false;
  }

  const payload = { kata_id: state.selectedKataID };
  if (formatCode) {
    payload.code = editors.code.input.value;
  }
  if (formatTests) {
    payload.tests = editors.tests.input.value;
  }

  state.formatting = true;
  try {
    const response = await api("/api/format", {
      method: "POST",
      body: JSON.stringify(payload),
    });

    let changed = false;
    if (formatCode && typeof response.code === "string" && response.code !== editors.code.input.value) {
      setEditorValue("code", response.code);
      changed = true;
    }
    if (formatTests && typeof response.tests === "string" && response.tests !== editors.tests.input.value) {
      setEditorValue("tests", response.tests);
      changed = true;
    }

    return changed;
  } catch (error) {
    if (!silent) {
      el.runOutput.textContent = `Format error: ${error.message || String(error)}`;
    }
    return false;
  } finally {
    state.formatting = false;
  }
}

function setEditorValue(name, value) {
  const editor = editors[name];
  editor.input.value = value || "";
  refreshSearchState(name);
  refreshHighlight(name);
}

function refreshHighlight(name) {
  const editor = editors[name];
  const source = editor.input.value || "";
  const query = state.search[name].query;
  editor.layer.innerHTML = `${highlightGo(source, query)}\n`;
  editor.lines.textContent = `${buildLineNumbers(source)}\n`;
  syncScroll(name);
}

function highlightGo(source, query = "") {
  if (!source) {
    return "";
  }

  let escaped = escapeHTML(source).replace(/\r/g, "");
  const tokens = [];

  escaped = escaped.replace(/(`[^`]*`|"(?:\\.|[^"\\])*"|\/\/[^\n]*)/g, (match) => {
    const type = match.startsWith("//") ? "comment" : "string";
    const idx = tokens.push({ type, text: match }) - 1;
    return `@@TOK_${idx}_@@`;
  });

  escaped = escaped
    .replace(GO_KEYWORDS, '<span class="tok-keyword">$1</span>')
    .replace(GO_TYPES, '<span class="tok-type">$1</span>')
    .replace(GO_NUMBERS, '<span class="tok-number">$1</span>');

  escaped = escaped.replace(/@@TOK_(\d+)_@@/g, (_, rawIndex) => {
    const token = tokens[Number(rawIndex)];
    if (!token) {
      return "";
    }
    if (token.type === "comment") {
      return `<span class="tok-comment">${token.text}</span>`;
    }
    return `<span class="tok-string">${token.text}</span>`;
  });

  return highlightSearchInHTML(escaped, query);
}

function highlightSearchInHTML(html, query) {
  const normalized = query.trim();
  if (!normalized) {
    return html;
  }

  const escapedQuery = escapeHTML(normalized);
  const matcher = new RegExp(escapeRegex(escapedQuery), "gi");
  const parts = html.split(/(<[^>]+>)/g);
  for (let i = 0; i < parts.length; i++) {
    const part = parts[i];
    if (!part || part.startsWith("<")) {
      continue;
    }
    parts[i] = part.replace(matcher, (match) => `<span class="tok-find">${match}</span>`);
  }
  return parts.join("");
}

function buildLineNumbers(source) {
  const lines = source.split("\n").length;
  const nums = [];
  for (let i = 1; i <= lines; i++) {
    nums.push(String(i));
  }
  return nums.join("\n");
}

function escapeRegex(input) {
  return input.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
}

function toggleModal(node, visible) {
  node.classList.toggle("hidden", !visible);
}

async function copyPrompt() {
  if (!el.markPrompt.value) {
    return;
  }
  await navigator.clipboard.writeText(el.markPrompt.value);
  el.copyPromptBtn.textContent = "Copied";
  setTimeout(() => {
    el.copyPromptBtn.textContent = "Copy Prompt";
  }, 1200);
}

function ensureSelectedKata() {
  if (!state.selectedKataID) {
    throw new Error("No kata selected");
  }
}

async function api(path, options = {}) {
  const response = await fetch(path, {
    headers: {
      "Content-Type": "application/json",
    },
    ...options,
  });

  const isJSON = response.headers.get("content-type")?.includes("application/json");
  const body = isJSON ? await response.json() : await response.text();
  if (!response.ok) {
    const message = body?.error || body || `Request failed (${response.status})`;
    throw new Error(message);
  }
  return body;
}

function escapeHTML(input) {
  return String(input)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;");
}

window.addEventListener("error", (evt) => {
  el.runOutput.textContent = `Error: ${evt.message}`;
});

window.addEventListener("unhandledrejection", (evt) => {
  el.runOutput.textContent = `Error: ${evt.reason?.message || String(evt.reason)}`;
});
