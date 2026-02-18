const state = {
  track: null,
  selectedKataID: null,
  selectedKata: null,
  nextRecommended: null,
  activeTab: "docs",
  dirty: false,
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
  codeEditor: document.getElementById("code-editor"),
  runOutput: document.getElementById("run-output"),
  failureInsights: document.getElementById("failure-insights"),
  dirtyIndicator: document.getElementById("dirty-indicator"),
  saveBtn: document.getElementById("save-btn"),
  runBtn: document.getElementById("run-btn"),
  markBtn: document.getElementById("mark-btn"),
  tabDocs: document.getElementById("tab-docs"),
  tabWorkbench: document.getElementById("tab-workbench"),
  panelDocs: document.getElementById("panel-docs"),
  panelWorkbench: document.getElementById("panel-workbench"),
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

boot();

async function boot() {
  wireEvents();
  await refreshTrack();
  const first = firstKataID();
  if (first) {
    await loadKata(first);
  }
}

function wireEvents() {
  el.tabDocs.addEventListener("click", () => setTab("docs"));
  el.tabWorkbench.addEventListener("click", () => setTab("workbench"));
  el.saveBtn.addEventListener("click", onSave);
  el.runBtn.addEventListener("click", onRun);
  el.markBtn.addEventListener("click", onMark);
  el.passCloseBtn.addEventListener("click", () => toggleModal(el.passModal, false));
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
  el.copyPromptBtn.addEventListener("click", copyPrompt);
  el.resetTestsBtn.addEventListener("click", onResetFromDisk);

  el.codeEditor.addEventListener("input", () => {
    if (!state.selectedKata) {
      return;
    }
    state.dirty = true;
    renderDirtyState();
  });
}

function setTab(tab) {
  state.activeTab = tab;
  const docsActive = tab === "docs";
  el.tabDocs.classList.toggle("active", docsActive);
  el.tabWorkbench.classList.toggle("active", !docsActive);
  el.panelDocs.classList.toggle("active", docsActive);
  el.panelWorkbench.classList.toggle("active", !docsActive);
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
    el.nextRecommendation.textContent = `Recommended next: ${next.kata_id} — ${next.kata_title}. ${next.reason}`;
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
  state.dirty = false;

  el.kataTitle.textContent = `${kata.id} — ${kata.title}`;
  el.kataSubtitle.textContent = `${kata.category.title} · ${kata.focus}`;
  el.readmeView.textContent = kata.readme;
  el.testsEditor.value = kata.tests;
  el.codeEditor.value = kata.code;
  el.runOutput.textContent = formatProgress(kata.progress);
  renderFailureInsights([]);
  renderDirtyState();
  renderButtons(true);
  renderTrack();
}

function renderButtons(enabled) {
  el.saveBtn.disabled = !enabled;
  el.runBtn.disabled = !enabled;
  el.markBtn.disabled = !enabled;
}

function renderDirtyState() {
  el.dirtyIndicator.textContent = state.dirty ? "Unsaved changes" : "Saved";
  el.dirtyIndicator.classList.toggle("dirty", state.dirty);
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
    el.failureInsights.innerHTML = `<div class="insight-card"><strong>No focused issues yet</strong>Run tests to surface the first mismatch, then fix one thing at a time.</div>`;
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

  await api("/api/save", {
    method: "POST",
    body: JSON.stringify({
      kata_id: state.selectedKataID,
      code: el.codeEditor.value,
      tests: el.testsEditor.value,
    }),
  });

  state.dirty = false;
  renderDirtyState();
  el.runOutput.textContent = "Saved.\nRun tests to validate behavior.";
}

async function onRun() {
  ensureSelectedKata();

  el.runOutput.textContent = "Running tests...";

  const result = await api("/api/run", {
    method: "POST",
    body: JSON.stringify({
      kata_id: state.selectedKataID,
      code: el.codeEditor.value,
      tests: el.testsEditor.value,
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
    el.passNextHint.textContent = `Recommended next: ${state.nextRecommended.kata_id} — ${state.nextRecommended.kata_title}`;
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

async function onResetFromDisk() {
  ensureSelectedKata();
  await loadKata(state.selectedKataID);
  el.runOutput.textContent = "Reloaded from disk.";
  renderFailureInsights([]);
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
  return input
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
