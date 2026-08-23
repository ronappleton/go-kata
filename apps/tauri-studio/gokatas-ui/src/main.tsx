import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";

// Global error handler — render visible error instead of blank screen
class ErrorBoundary extends React.Component<
  { children: React.ReactNode },
  { error: string | null }
> {
  state = { error: null as string | null };
  static getDerivedStateFromError(err: any) {
    return { error: err?.message || String(err) };
  }
  render() {
    if (this.state.error) {
      return (
        <div style={{ padding: 40, color: "#ff4444", fontFamily: "monospace", background: "#0d1117", minHeight: "100vh" }}>
          <h2 style={{ color: "#ff6666" }}>GoKatas crashed</h2>
          <pre style={{ whiteSpace: "pre-wrap" }}>{this.state.error}</pre>
        </div>
      );
    }
    return this.props.children;
  }
}

console.log("[GoKatas] Starting app...");

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <ErrorBoundary>
      <App />
    </ErrorBoundary>
  </React.StrictMode>,
);
