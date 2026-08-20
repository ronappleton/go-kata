#!/usr/bin/env python3
"""Generate the curriculum viewer HTML from track.json and kata metadata."""
import json, os, html as htmlmod

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(SCRIPT_DIR)

# Load all tracks
tracks = []
for d in sorted(os.listdir(os.path.join(ROOT, "tracks"))):
    tp = os.path.join(ROOT, "tracks", d, "track.json")
    if os.path.isfile(tp):
        with open(tp) as f:
            tracks.append(json.load(f))

# Load kata metadata
kata_metas = {}
for d in sorted(os.listdir(os.path.join(ROOT, "katas"))):
    if not d.startswith("kata-"):
        continue
    jp = os.path.join(ROOT, "katas", d, "kata.json")
    if os.path.isfile(jp):
        with open(jp) as f:
            km = json.load(f)
            kid = km.get("id", "")
            kata_metas[kid] = km

# Collect flashcards and quiz questions
all_flashcards = []
all_quiz = []
for kid, km in sorted(kata_metas.items()):
    for fc in km.get("flashcards", []):
        all_flashcards.append({"kata_id": kid, "front": fc.get("front",""), "back": fc.get("back","")})
    for q in km.get("quiz_questions", []):
        all_quiz.append({"kata_id": kid, "question": q.get("question",""), "options": q.get("options",[]), "answer": q.get("answer","")})

# Level colors
LEVEL_COLORS = {"junior": "#22c55e", "mid": "#14b8a6", "senior": "#f59e0b", "lead": "#ef4444"}
LEVEL_BADGES = {"junior": "JUNIOR", "mid": "MID", "senior": "SENIOR", "lead": "LEAD"}

# Build track data as JSON for the JS
tracks_json = json.dumps(tracks)
katas_json = json.dumps(kata_metas)
flashcards_json = json.dumps(all_flashcards)
quiz_json = json.dumps(all_quiz)

total_katas = sum(
    len(cat.get("kata_ids", []))
    for t in tracks
    for s in t.get("stages", [])
    for cat in s.get("categories", [])
)
total_stages = sum(len(t.get("stages", [])) for t in tracks)
total_cats = sum(
    len(s.get("categories", []))
    for t in tracks
    for s in t.get("stages", [])
)

# Build HTML
h = f'''<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>GoKatas — Curriculum Viewer</title>
<style>
* {{ margin:0; padding:0; box-sizing:border-box; }}
body {{ background:#0f172a; color:#e2e8f0; font-family:'Inter',-apple-system,system-ui,sans-serif; }}
.header {{ background:linear-gradient(135deg,#0f172a 0%,#1e293b 100%); padding:20px 24px; border-bottom:1px solid #334155; position:sticky; top:0; z-index:100; }}
.header h1 {{ font-size:24px; font-weight:700; }}
.header h1 span {{ color:#14b8a6; }}
.header .sub {{ color:#94a3b8; font-size:13px; margin-top:2px; }}
.stats {{ display:flex; gap:24px; margin-top:12px; flex-wrap:wrap; }}
.stat {{ text-align:center; }}
.stat .val {{ font-size:22px; font-weight:700; color:#14b8a6; }}
.stat .lbl {{ font-size:10px; text-transform:uppercase; letter-spacing:1px; color:#64748b; }}

.track-tabs {{ display:flex; gap:8px; margin-top:16px; flex-wrap:wrap; }}
.track-tab {{ padding:8px 16px; border-radius:8px; border:1px solid #334155; background:transparent; color:#94a3b8; cursor:pointer; font-size:13px; font-weight:500; transition:all .15s; }}
.track-tab:hover {{ background:#1e293b; color:#e2e8f0; }}
.track-tab.active {{ background:#14b8a6; color:#0f172a; border-color:#14b8a6; font-weight:600; }}

.stage-filters {{ display:flex; gap:6px; margin-top:12px; flex-wrap:wrap; }}
.stage-btn {{ padding:6px 14px; border-radius:6px; border:1px solid #334155; background:transparent; color:#94a3b8; cursor:pointer; font-size:12px; transition:all .15s; }}
.stage-btn:hover {{ background:#1e293b; color:#e2e8f0; }}
.stage-btn.active {{ background:#14b8a6; color:#0f172a; border-color:#14b8a6; font-weight:600; }}

.track-section {{ display:none; padding:16px 24px; }}
.track-section.active {{ display:block; }}

.track-header {{ display:flex; align-items:center; gap:12px; margin-bottom:20px; padding-bottom:16px; border-bottom:1px solid #1e293b; }}
.track-header h2 {{ font-size:20px; font-weight:600; }}
.track-desc {{ color:#94a3b8; font-size:13px; }}

.stage-section {{ margin-bottom:28px; }}
.stage-header {{ display:flex; align-items:center; gap:10px; margin-bottom:16px; }}
.stage-title {{ font-size:18px; font-weight:600; }}
.level-badge {{ padding:3px 10px; border-radius:12px; font-size:10px; font-weight:700; letter-spacing:1px; text-transform:uppercase; }}
.stage-desc {{ color:#94a3b8; font-size:13px; margin-bottom:14px; }}

.category {{ background:#1e293b; border-radius:12px; padding:16px 20px; margin-bottom:12px; border:1px solid #334155; }}
.cat-header {{ display:flex; justify-content:space-between; align-items:center; margin-bottom:10px; }}
.cat-title {{ font-size:14px; font-weight:600; color:#cbd5e1; text-transform:uppercase; letter-spacing:.5px; }}
.cat-count {{ font-size:12px; color:#64748b; }}

.kata-list {{ list-style:none; }}
.kata-item {{ display:flex; align-items:center; gap:10px; padding:6px 0; border-bottom:1px solid #0f172a; }}
.kata-item:last-child {{ border-bottom:none; }}
.kata-dot {{ width:8px; height:8px; border-radius:50%; background:#64748b; flex-shrink:0; }}
.kata-dot.ready {{ background:#f59e0b; }}
.kata-dot.done {{ background:#22c55e; }}
.kata-id {{ font-size:13px; font-weight:700; color:#14b8a6; font-family:monospace; min-width:28px; }}
.kata-title {{ font-size:13px; color:#e2e8f0; }}
.kata-tags {{ display:flex; gap:4px; margin-left:auto; }}
.kata-tag {{ font-size:9px; padding:2px 6px; border-radius:4px; background:#0f172a; color:#64748b; border:1px solid #334155; }}

.flashcard-section {{ margin-top:32px; padding:24px; background:#1e293b; border-radius:12px; border:1px solid #334155; }}
.flashcard-section h3 {{ font-size:16px; font-weight:600; margin-bottom:16px; color:#14b8a6; }}
.fc-card {{ background:#0f172a; border-radius:10px; padding:24px; text-align:center; min-height:120px; display:flex; flex-direction:column; align-items:center; justify-content:center; border:1px solid #334155; margin-bottom:12px; cursor:pointer; transition:all .2s; }}
.fc-card:hover {{ border-color:#14b8a6; }}
.fc-label {{ font-size:10px; text-transform:uppercase; letter-spacing:1px; color:#64748b; margin-bottom:8px; }}
.fc-text {{ font-size:15px; line-height:1.5; }}
.fc-kata {{ font-size:11px; color:#14b8a6; margin-top:8px; font-family:monospace; }}
.fc-controls {{ display:flex; gap:8px; justify-content:center; margin-top:12px; }}
.fc-btn {{ padding:8px 16px; border-radius:6px; border:1px solid #334155; background:#0f172a; color:#94a3b8; cursor:pointer; font-size:12px; transition:all .15s; }}
.fc-btn:hover {{ background:#1e293b; color:#e2e8f0; }}
.fc-btn.primary {{ background:#14b8a6; color:#0f172a; border-color:#14b8a6; font-weight:600; }}
.fc-btn.again {{ border-color:#ef4444; color:#ef4444; }}
.fc-btn.hard {{ border-color:#f59e0b; color:#f59e0b; }}
.fc-btn.good {{ border-color:#14b8a6; color:#14b8a6; }}
.fc-btn.easy {{ border-color:#22c55e; color:#22c55e; }}
.fc-counter {{ font-size:12px; color:#64748b; text-align:center; margin-top:8px; }}

.quiz-section {{ margin-top:20px; padding:24px; background:#1e293b; border-radius:12px; border:1px solid #334155; }}
.quiz-section h3 {{ font-size:16px; font-weight:600; margin-bottom:16px; color:#f59e0b; }}
.quiz-q {{ font-size:14px; margin-bottom:12px; color:#e2e8f0; }}
.quiz-opts {{ list-style:none; }}
.quiz-opt {{ padding:8px 12px; margin:4px 0; border-radius:6px; border:1px solid #334155; background:#0f172a; color:#94a3b8; cursor:pointer; font-size:13px; transition:all .15s; }}
.quiz-opt:hover {{ border-color:#14b8a6; color:#e2e8f0; }}
.quiz-opt.correct {{ border-color:#22c55e; color:#22c55e; background:#052e16; }}
.quiz-opt.wrong {{ border-color:#ef4444; color:#ef4444; background:#2d0a0a; }}
</style>
</head>
<body>
<div class="header">
  <h1>Go<span>Katas</span></h1>
  <div class="sub">Junior to Lead — Go Mastery + Addon Tracks</div>
  <div class="stats">
    <div class="stat"><div class="val">{total_katas}</div><div class="lbl">Katas</div></div>
    <div class="stat"><div class="val">{total_stages}</div><div class="lbl">Stages</div></div>
    <div class="stat"><div class="val">{total_cats}</div><div class="lbl">Categories</div></div>
    <div class="stat"><div class="val">{len(all_flashcards)}</div><div class="lbl">Flashcards</div></div>
    <div class="stat"><div class="val">{len(all_quiz)}</div><div class="lbl">Quiz Qs</div></div>
  </div>
  <div class="track-tabs" id="trackTabs"></div>
  <div class="stage-filters" id="stageFilters"></div>
</div>

<div id="trackContent"></div>

<div class="flashcard-section">
  <h3>🃏 Flashcard Review</h3>
  <div class="fc-card" id="fcCard" onclick="flipCard()">
    <div class="fc-label" id="fcLabel">FRONT</div>
    <div class="fc-text" id="fcText"></div>
    <div class="fc-kata" id="fcKata"></div>
  </div>
  <div class="fc-controls">
    <button class="fc-btn" onclick="prevCard()">Prev</button>
    <button class="fc-btn primary" onclick="flipCard()">Flip</button>
    <button class="fc-btn" onclick="nextCard()">Next</button>
    <button class="fc-btn again" onclick="confidence('again')">Again</button>
    <button class="fc-btn hard" onclick="confidence('hard')">Hard</button>
    <button class="fc-btn good" onclick="confidence('good')">Good</button>
    <button class="fc-btn easy" onclick="confidence('easy')">Easy</button>
  </div>
  <div class="fc-counter" id="fcCounter"></div>
</div>

<div class="quiz-section">
  <h3>📝 Quiz Mode</h3>
  <div class="quiz-q" id="quizQ"></div>
  <ul class="quiz-opts" id="quizOpts"></ul>
  <div class="fc-controls" style="margin-top:12px;">
    <button class="fc-btn" onclick="prevQuiz()">Prev</button>
    <button class="fc-btn primary" onclick="nextQuiz()">Next</button>
  </div>
  <div class="fc-counter" id="quizCounter"></div>
</div>

<div style="padding:16px 24px; text-align:center; color:#64748b; font-size:11px;">
  GoKatas — GTK4 + Go + Podman
</div>

<script>
const TRACKS = {tracks_json};
const KATAS = {katas_json};
const FLASHCARDS = {flashcards_json};
const QUIZ = {quiz_json};
const LEVEL_COLORS = {json.dumps(LEVEL_COLORS)};
const LEVEL_BADGES = {json.dumps(LEVEL_BADGES)};

let currentTrack = null;
let currentStage = null;
let fcIdx = 0;
let fcFlipped = false;
let quizIdx = 0;

function init() {{
  renderTrackTabs();
  selectTrack(TRACKS[0].id);
  renderFlashcard();
  renderQuiz();
}}

function renderTrackTabs() {{
  const el = document.getElementById('trackTabs');
  el.innerHTML = TRACKS.map(t => 
    `<button class="track-tab" data-id="${{t.id}}" onclick="selectTrack('${{t.id}}')">${{t.title}}</button>`
  ).join('');
}}

function selectTrack(id) {{
  currentTrack = id;
  currentStage = null;
  document.querySelectorAll('.track-tab').forEach(b => b.classList.toggle('active', b.dataset.id === id));
  renderStageFilters();
  renderTrackContent();
}}

function renderStageFilters() {{
  const track = TRACKS.find(t => t.id === currentTrack);
  const el = document.getElementById('stageFilters');
  if (!track) {{ el.innerHTML = ''; return; }}
  const stages = track.stages || [];
  el.innerHTML = `<button class="stage-btn active" onclick="filterStage(null)">All</button>` +
    stages.map(s => `<button class="stage-btn" data-stage="${{s.id}}" onclick="filterStage('${{s.id}}')">${{s.title}}</button>`).join('');
}}

function filterStage(id) {{
  currentStage = id;
  document.querySelectorAll('.stage-btn').forEach(b => b.classList.toggle('active', 
    (id === null && !b.dataset.stage) || b.dataset.stage === id));
  renderTrackContent();
}}

function renderTrackContent() {{
  const track = TRACKS.find(t => t.id === currentTrack);
  if (!track) return;
  const el = document.getElementById('trackContent');
  let html = `<div class="track-section active"><div class="track-header"><div><h2>${{track.title}}</h2><div class="track-desc">${{track.description}}</div></div></div>`;
  
  const stages = track.stages || [];
  for (const stage of stages) {{
    if (currentStage && stage.id !== currentStage) continue;
    const color = LEVEL_COLORS[stage.level] || '#94a3b8';
    const badge = LEVEL_BADGES[stage.level] || stage.level.toUpperCase();
    html += `<div class="stage-section">
      <div class="stage-header">
        <span class="stage-title">${{stage.title}}</span>
        <span class="level-badge" style="background:${{color}}22;color:${{color}}">${{badge}}</span>
      </div>
      <div class="stage-desc">${{stage.description}}</div>`;
    
    for (const cat of (stage.categories || [])) {{
      html += `<div class="category">
        <div class="cat-header">
          <span class="cat-title">${{cat.title}}</span>
          <span class="cat-count">${{cat.kata_ids.length}} katas</span>
        </div>
        <ul class="kata-list">`;
      for (const kid of cat.kata_ids) {{
        const k = KATAS[kid];
        const title = k ? k.title : `Kata ${{kid}}`;
        const tags = k && k.tags ? k.tags.slice(0,3) : [];
        html += `<li class="kata-item">
          <span class="kata-dot"></span>
          <span class="kata-id">${{kid}}</span>
          <span class="kata-title">${{title}}</span>
          <span class="kata-tags">${{tags.map(t => `<span class="kata-tag">${{t}}</span>`).join('')}}</span>
        </li>`;
      }}
      html += '</ul></div>';
    }}
    html += '</div>';
  }}
  html += '</div>';
  el.innerHTML = html;
}}

function renderFlashcard() {{
  if (!FLASHCARDS.length) return;
  const fc = FLASHCARDS[fcIdx];
  document.getElementById('fcLabel').textContent = fcFlipped ? 'BACK' : 'FRONT';
  document.getElementById('fcText').textContent = fcFlipped ? fc.back : fc.front;
  document.getElementById('fcKata').textContent = `[${{fc.kata_id}}]`;
  document.getElementById('fcCounter').textContent = `${{fcIdx + 1}} / ${{FLASHCARDS.length}}`;
  fcFlipped = false;
}}

function flipCard() {{
  fcFlipped = !fcFlipped;
  const fc = FLASHCARDS[fcIdx];
  document.getElementById('fcLabel').textContent = fcFlipped ? 'BACK' : 'FRONT';
  document.getElementById('fcText').textContent = fcFlipped ? fc.back : fc.front;
}}

function prevCard() {{ if (fcIdx > 0) {{ fcIdx--; renderFlashcard(); }} }}
function nextCard() {{ if (fcIdx < FLASHCARDS.length - 1) {{ fcIdx++; renderFlashcard(); }} }}
function confidence(level) {{ if (fcIdx < FLASHCARDS.length - 1) {{ fcIdx++; renderFlashcard(); }} }}

function renderQuiz() {{
  if (!QUIZ.length) return;
  const q = QUIZ[quizIdx];
  document.getElementById('quizQ').textContent = q.question;
  document.getElementById('quizOpts').innerHTML = q.options.map(o => 
    `<li class="quiz-opt" onclick="checkAnswer(this, '${{o}}', '${{q.answer}}')">${{o}}</li>`
  ).join('');
  document.getElementById('quizCounter').textContent = `${{quizIdx + 1}} / ${{QUIZ.length}}`;
}}

function checkAnswer(el, chosen, answer) {{
  document.querySelectorAll('.quiz-opt').forEach(o => {{
    o.classList.remove('correct', 'wrong');
    if (o.textContent === answer) o.classList.add('correct');
    if (o === el && chosen !== answer) o.classList.add('wrong');
  }});
}}

function prevQuiz() {{ if (quizIdx > 0) {{ quizIdx--; renderQuiz(); }} }}
function nextQuiz() {{ if (quizIdx < QUIZ.length - 1) {{ quizIdx++; renderQuiz(); }} }}

init();
</script>
</body>
</html>'''

with open(os.path.join(ROOT, "curriculum-viewer.html"), "w") as f:
    f.write(h)

print(f"Generated curriculum-viewer.html: {total_katas} katas, {len(all_flashcards)} flashcards, {len(all_quiz)} quiz Qs, {len(tracks)} tracks")
