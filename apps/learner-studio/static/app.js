const state = {
  track: null,
  selectedKataID: null,
  selectedKata: null,
  activeTab: "docs",
  dirty: false,
};

const el = {
  trackSummary: document.getElementById("track-summary"),
  categoryList: document.getElementById("category-list"),
  kataTitle: document.getElementById("kata-title"),
  kataSubtitle: document.getElementById("kata-subtitle"),
  readmeView: document.getElementById("readme-view"),
  testsEditor: document.getElementById("tests-editor"),
  codeEditor: document.getElementById("code-editor"),
  runOutput: document.getElementById("run-output"),
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
  renderTrack();
}

function renderTrack() {
  if (!state.track) {
    return;
  }

  el.trackSummary.textContent = `${state.track.overall_done}/${state.track.overall_total} complete (${state.track.overall_percent}%)`;
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
  if (result.failed_tests && result.failed_tests.length > 0) {
    lines.push(`Failing tests: ${result.failed_tests.join(", ")}`);
  }
  if (result.output_tail) {
    lines.push("");
    lines.push(result.output_tail);
  }
  el.runOutput.textContent = lines.join("\n");

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
