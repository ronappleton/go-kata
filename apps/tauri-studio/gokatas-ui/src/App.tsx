import { useState, useEffect, useCallback, useRef } from "react";
import Editor from "@monaco-editor/react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { catalog, kata, progress, syncApi } from "./api";
import type { Track, KataDetail, KataSummary, ProgressState, StageSummary } from "./types";
import "./index.css";

// ── Main App ──

export default function App() {
  const [track, setTrack] = useState<Track | null>(null);
  const [selectedKata, setSelectedKata] = useState<KataDetail | null>(null);
  const [progressState, setProgressState] = useState<ProgressState>({ attempts: {} });
  const [activeTab, setActiveTab] = useState<"docs" | "workbench" | "output">("docs");
  const [code, setCode] = useState("");
  const [tests, setTests] = useState("");
  const [output, setOutput] = useState("");
  const [running, setRunning] = useState(false);
  const [syncing, setSyncing] = useState(false);
  const [statusMsg, setStatusMsg] = useState("Starting…");

  // Load catalog on mount
  useEffect(() => {
    let cancelled = false;
    async function load() {
      try {
        const t = await catalog.get();
        if (!cancelled) {
          setTrack(t);
          setStatusMsg(`Ready · ${t.kataCount} katas`);
        }
      } catch (e) {
        if (!cancelled) setStatusMsg("Waiting for curriculum…");
      }
      // Load progress
      try {
        const p = await progress.get();
        if (!cancelled) setProgressState(p);
      } catch {}
    }
    // Retry until sidecar is ready
    const interval = setInterval(() => {
      load().then(() => {
        if (track === null) return; // keep retrying
      });
    }, 1000);
    load();
    return () => { cancelled = true; clearInterval(interval); };
  }, []);

  // Select a kata
  const selectKata = useCallback(async (k: KataSummary) => {
    try {
      setStatusMsg(`Loading ${k.id}…`);
      const detail = await kata.get(k.id);
      setSelectedKata(detail);
      setCode(detail.content.kataGo || "");
      setTests(detail.content.kataTest || "");
      setOutput("");
      setActiveTab("docs");
      setStatusMsg(`${k.id} — ${k.title}`);
    } catch (e: any) {
      setStatusMsg(`Error: ${e.message}`);
    }
  }, []);

  // Save
  const save = useCallback(async () => {
    if (!selectedKata) return;
    try {
      await kata.save(selectedKata.kata.id, { code, tests });
      setStatusMsg("Saved");
    } catch (e: any) {
      setStatusMsg(`Save failed: ${e.message}`);
    }
  }, [selectedKata, code, tests]);

  // Run
  const run = useCallback(async () => {
    if (!selectedKata || running) return;
    setRunning(true);
    setOutput("Running in sandbox…");
    setActiveTab("output");
    try {
      await kata.save(selectedKata.kata.id, { code, tests });
      const result = await kata.run(selectedKata.kata.id, { code, tests });
      const lines = [
        result.status.toUpperCase(),
        `Duration: ${result.duration}`,
      ];
      if (result.failedTests.length > 0) {
        lines.push(`Failing: ${result.failedTests.join(", ")}`);
      }
      if (result.evaluatorError) {
        lines.push(`Error: ${result.evaluatorError}`);
      }
      if (result.output) {
        lines.push("", result.output);
      }
      setOutput(lines.join("\n"));
      setStatusMsg(result.passed ? "Passed ✓" : "Failed ✗");
      // Refresh progress
      const p = await progress.get();
      setProgressState(p);
    } catch (e: any) {
      setOutput(`Error: ${e.message}`);
    } finally {
      setRunning(false);
    }
  }, [selectedKata, code, tests, running]);

  // Sync
  const sync = useCallback(async () => {
    setSyncing(true);
    setStatusMsg("Syncing…");
    try {
      await syncApi.trigger();
      // Wait and reload
      await new Promise((r) => setTimeout(r, 2000));
      const t = await catalog.get();
      setTrack(t);
      setStatusMsg(`Curriculum ready · ${t.kataCount} katas`);
    } catch (e: any) {
      setStatusMsg(`Sync failed: ${e.message}`);
    } finally {
      setSyncing(false);
    }
  }, []);

  return (
    <div className="flex h-screen overflow-hidden" style={{ background: "var(--color-bg)" }}>
      {/* ── Sidebar ── */}
      <aside
        className="flex flex-col border-r shrink-0 overflow-hidden"
        style={{ width: 360, minWidth: 280, background: "var(--color-surface)", borderColor: "var(--color-border)" }}
      >
        {/* Brand header */}
        <div className="px-4 pt-4 pb-3 flex items-center justify-between">
          <div>
            <span className="text-xl font-extrabold" style={{ color: "var(--color-accent)" }}>Go</span>
            <span className="text-xl font-extrabold" style={{ color: "var(--color-text)" }}>Katas</span>
            <div className="text-xs" style={{ color: "var(--color-text-dim)" }}>Junior to Lead</div>
          </div>
          <button
            onClick={sync}
            disabled={syncing}
            className="text-xs px-3 py-1.5 rounded-lg font-semibold"
            style={{ background: "var(--color-surface-hi)", color: "var(--color-text-dim)", border: "1px solid var(--color-border)" }}
          >
            {syncing ? "Syncing…" : "↻ Sync"}
          </button>
        </div>

        {/* Stage list */}
        <div className="flex-1 overflow-y-auto px-3 pb-4">
          {!track ? (
            <div className="p-6 text-sm" style={{ color: "var(--color-text-faint)" }}>
              {statusMsg}
            </div>
          ) : (
            track.stages.map((stage) => (
              <StageBlock
                key={stage.id}
                stage={stage}
                selectedId={selectedKata?.kata.id}
                progress={progressState}
                onSelect={selectKata}
              />
            ))
          )}
        </div>
      </aside>

      {/* ── Main content ── */}
      <main className="flex-1 flex flex-col overflow-hidden">
        {/* Title bar */}
        <div className="px-5 pt-4 pb-2 flex items-center justify-between" style={{ borderBottom: "1px solid var(--color-border)" }}>
          <div className="flex-1 min-w-0">
            <h1 className="text-lg font-extrabold truncate" style={{ color: "var(--color-text)" }}>
              {selectedKata ? `${selectedKata.kata.id} — ${selectedKata.kata.title}` : "Select a kata"}
            </h1>
            {selectedKata && (
              <p className="text-xs truncate" style={{ color: "var(--color-text-dim)" }}>
                {selectedKata.kata.focus} · {selectedKata.kata.signature} · evaluator: {selectedKata.kata.evaluatorStatus}
              </p>
            )}
          </div>
          <div className="flex items-center gap-2 ml-4">
            {/* Status */}
            <span className="text-xs px-3 py-1 rounded-full" style={{ background: "var(--color-surface-hi)", color: "var(--color-text-dim)", border: "1px solid var(--color-border)" }}>
              {statusMsg}
            </span>
          </div>
        </div>

        {/* Tab bar */}
        <div className="px-5 py-2 flex gap-2" style={{ borderBottom: "1px solid var(--color-border)" }}>
          {(["docs", "workbench", "output"] as const).map((tab) => (
            <button
              key={tab}
              className={`tab-btn ${activeTab === tab ? "active" : ""}`}
              onClick={() => setActiveTab(tab)}
            >
              {tab === "docs" ? "📖 Readme" : tab === "workbench" ? "⌨ Code" : "▶ Output"}
            </button>
          ))}
          {activeTab === "workbench" && (
            <div className="flex-1 flex justify-end gap-2">
              <button
                onClick={save}
                className="text-xs px-4 py-1.5 rounded-lg font-semibold"
                style={{ background: "var(--color-surface-hi)", color: "var(--color-text)", border: "1px solid var(--color-border)" }}
              >
                Save
              </button>
              <button
                onClick={run}
                disabled={running || selectedKata?.kata.evaluatorStatus !== "ready"}
                className="text-xs px-4 py-1.5 rounded-lg font-bold"
                style={{ background: running ? "var(--color-surface-hi)" : "var(--color-accent)", color: running ? "var(--color-text-faint)" : "#071014", border: "1px solid transparent" }}
              >
                {running ? "Running…" : "▶ Run in sandbox"}
              </button>
            </div>
          )}
        </div>

        {/* Tab content */}
        <div className="flex-1 overflow-hidden">
          {activeTab === "docs" && <DocsTab kata={selectedKata} />}
          {activeTab === "workbench" && (
            <WorkbenchTab code={code} tests={tests} onCodeChange={setCode} onTestsChange={setTests} />
          )}
          {activeTab === "output" && <OutputTab output={output} running={running} />}
        </div>
      </main>
    </div>
  );
}

// ── Sidebar Stage Block ──

function StageBlock({
  stage,
  selectedId,
  progress,
  onSelect,
}: {
  stage: StageSummary;
  selectedId?: string;
  progress: ProgressState;
  onSelect: (k: KataSummary) => void;
}) {
  const [collapsed, setCollapsed] = useState(false);
  const totalKatas = stage.categories.reduce((sum, c) => sum + c.katas.length, 0);
  const completedKatas = stage.categories.reduce(
    (sum, c) => sum + c.katas.filter((k) => (progress.attempts[k.id]?.passes ?? 0) > 0).length,
    0
  );
  const pct = totalKatas > 0 ? completedKatas / totalKatas : 0;

  return (
    <div className="mb-2">
      {/* Stage header */}
      <button
        onClick={() => setCollapsed(!collapsed)}
        className="w-full flex items-center gap-2 py-2 px-1 text-left"
      >
        <span className="text-xs font-extrabold uppercase" style={{ color: "var(--color-text)" }}>
          {stage.title}
        </span>
        <span
          className="text-[10px] font-bold px-2 py-0.5 rounded-full"
          style={{
            background: stage.level === "junior" ? "var(--color-success)" : stage.level === "senior" ? "var(--color-warning)" : stage.level === "lead" ? "var(--color-danger)" : "var(--color-accent)",
            color: "#071014",
          }}
        >
          {stage.level.toUpperCase()}
        </span>
        {/* Progress bar */}
        <div className="flex-1 mx-2 h-1 rounded-full overflow-hidden" style={{ background: "var(--color-border)" }}>
          <div className="h-full rounded-full transition-all" style={{ width: `${pct * 100}%`, background: "var(--color-accent)" }} />
        </div>
        <span className="text-[10px]" style={{ color: "var(--color-text-faint)" }}>
          {collapsed ? "▶" : "▼"}
        </span>
      </button>

      {/* Categories + katas */}
      {!collapsed &&
        stage.categories.map((cat) => (
          <div key={cat.id} className="ml-2">
            <div className="text-[10px] font-bold uppercase px-3 py-1" style={{ color: "var(--color-text-faint)" }}>
              {cat.title}
            </div>
            {cat.katas.map((k) => (
              <KataRow
                key={k.id}
                kata={k}
                active={k.id === selectedId}
                completed={(progress.attempts[k.id]?.passes ?? 0) > 0}
                onClick={() => onSelect(k)}
              />
            ))}
          </div>
        ))}
    </div>
  );
}

// ── Sidebar Kata Row ──

function KataRow({
  kata: k,
  active,
  completed,
  onClick,
}: {
  kata: KataSummary;
  active: boolean;
  completed: boolean;
  onClick: () => void;
}) {
  const dotColor = completed ? "var(--color-success)" : k.evaluatorStatus === "ready" ? "var(--color-warning)" : "var(--color-text-faint)";

  return (
    <button
      className={`kata-row w-full text-left flex items-center gap-2 ${active ? "active" : ""}`}
      onClick={onClick}
    >
      <span className="text-[10px]" style={{ color: dotColor }}>●</span>
      <span className="truncate flex-1">{k.id} {k.title}</span>
    </button>
  );
}

// ── Docs Tab (Markdown) ──

function DocsTab({ kata: detail }: { kata: KataDetail | null }) {
  if (!detail) {
    return (
      <div className="p-8 text-center" style={{ color: "var(--color-text-faint)" }}>
        <div className="text-4xl mb-4">📖</div>
        <div className="text-lg font-semibold mb-2">Select a kata from the curriculum</div>
        <div className="text-sm">Choose a kata from the sidebar to see its instructions.</div>
      </div>
    );
  }

  return (
    <div className="h-full overflow-y-auto">
      <div className="p-6 max-w-3xl">
        <div className="markdown-body">
          <ReactMarkdown remarkPlugins={[remarkGfm]}>
            {detail.content.readme || "No readme available."}
          </ReactMarkdown>
        </div>
      </div>
    </div>
  );
}

// ── Workbench Tab (Monaco Editor) ──

function WorkbenchTab({
  code,
  tests,
  onCodeChange,
  onTestsChange,
}: {
  code: string;
  tests: string;
  onCodeChange: (v: string) => void;
  onTestsChange: (v: string) => void;
}) {
  const [editorTab, setEditorTab] = useState<"solution" | "tests">("solution");

  return (
    <div className="h-full flex flex-col">
      {/* Editor tabs */}
      <div className="flex gap-2 px-4 py-2" style={{ borderBottom: "1px solid var(--color-border)" }}>
        <button
          className={`tab-btn ${editorTab === "solution" ? "active" : ""}`}
          onClick={() => setEditorTab("solution")}
        >
          Solution (kata.go)
        </button>
        <button
          className={`tab-btn ${editorTab === "tests" ? "active" : ""}`}
          onClick={() => setEditorTab("tests")}
        >
          Learner Tests (kata_test.go)
        </button>
      </div>

      {/* Monaco editor */}
      <div className="flex-1">
        <Editor
          height="100%"
          language="go"
          theme="vs-dark"
          value={editorTab === "solution" ? code : tests}
          onChange={(v) => (editorTab === "solution" ? onCodeChange(v || "") : onTestsChange(v || ""))}
          options={{
            fontSize: 14,
            fontFamily: "'JetBrains Mono', 'Fira Code', monospace",
            minimap: { enabled: false },
            padding: { top: 12, bottom: 12 },
            scrollBeyondLastLine: false,
            lineNumbers: "on",
            tabSize: 4,
            automaticLayout: true,
          }}
        />
      </div>
    </div>
  );
}

// ── Output Tab ──

function OutputTab({ output, running }: { output: string; running: boolean }) {
  const preRef = useRef<HTMLPreElement>(null);

  useEffect(() => {
    if (preRef.current) {
      preRef.current.scrollTop = preRef.current.scrollHeight;
    }
  }, [output]);

  if (!output) {
    return (
      <div className="p-8 text-center" style={{ color: "var(--color-text-faint)" }}>
        <div className="text-4xl mb-4">▶</div>
        <div className="text-lg font-semibold mb-2">Output</div>
        <div className="text-sm">Run your solution to see results here.</div>
      </div>
    );
  }

  return (
    <div className="h-full overflow-hidden">
      <pre
        ref={preRef}
        className="h-full overflow-y-auto p-4 text-sm font-mono"
        style={{ background: "var(--color-code-bg)", color: "var(--color-text-dim)" }}
      >
        {output}
        {running && <span className="animate-pulse">▌</span>}
      </pre>
    </div>
  );
}
