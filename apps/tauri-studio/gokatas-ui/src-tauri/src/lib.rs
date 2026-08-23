use std::process::{Command, Stdio};
use std::sync::{Mutex, Condvar};
use std::thread;
use std::time::Duration;
use tauri::{Emitter, Manager};
use serde_json::Value;

struct AppState {
    port: Mutex<u16>,
    ready: Condvar,
    client: reqwest::Client,
}

#[tauri::command]
fn get_port(state: tauri::State<AppState>) -> u16 {
    *state.port.lock().unwrap()
}

#[tauri::command]
async fn get_catalog(state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/catalog", port);
    state.client.get(&url).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn get_kata(id: String, state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/kata/{}", port, id);
    state.client.get(&url).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn save_kata(id: String, code: String, tests: String, state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/kata/{}/save", port, id);
    let body = serde_json::json!({ "code": code, "tests": tests });
    state.client.post(&url).json(&body).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn run_kata(id: String, code: String, tests: String, state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/kata/{}/run", port, id);
    let body = serde_json::json!({ "code": code, "tests": tests });
    state.client.post(&url).json(&body).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn get_progress(state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/progress", port);
    state.client.get(&url).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn get_status(state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/status", port);
    state.client.get(&url).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn sync_content(state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/sync", port);
    state.client.post(&url).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

#[tauri::command]
async fn lint_code(code: String, language: String, state: tauri::State<'_, AppState>) -> Result<Value, String> {
    let port = *state.port.lock().unwrap();
    let url = format!("http://127.0.0.1:{}/api/lint", port);
    let body = serde_json::json!({ "code": code, "language": language });
    state.client.post(&url).json(&body).send().await
        .map_err(|e| e.to_string())?
        .json::<Value>().await
        .map_err(|e| e.to_string())
}

/// Kill any stale sidecar processes from previous runs.
fn kill_stale_sidecars() {
    let _ = Command::new("pkill")
        .args(["-f", "gokatas-sidecar"])
        .stdout(Stdio::null())
        .stderr(Stdio::null())
        .spawn();
    thread::sleep(Duration::from_millis(500));
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .manage(AppState {
            port: Mutex::new(0),
            ready: Condvar::new(),
            client: reqwest::Client::builder()
                .timeout(Duration::from_secs(30))
                .build()
                .unwrap(),
        })
        .setup(|app| {
            let app_handle = app.handle().clone();

            // Kill stale sidecar processes first
            kill_stale_sidecars();

            // Find the sidecar binary
            let mut candidates = vec![
                std::path::PathBuf::from("/usr/bin/gokatas-sidecar"),
            ];
            if let Ok(res) = app.path().resource_dir() {
                candidates.push(res.join("sidecar").join("gokatas-sidecar"));
            }
            if let Ok(cwd) = std::env::current_dir() {
                candidates.push(cwd.join("src-tauri").join("sidecar").join("gokatas-sidecar"));
                candidates.push(cwd.join("sidecar").join("gokatas-sidecar"));
            }

            let sidecar_path = candidates.iter().find(|p| p.exists())
                .expect("gokatas-sidecar binary not found — build with: go build -o /usr/bin/gokatas-sidecar ./apps/tauri-studio/gokatas-ui/src-tauri/sidecar/");

            log::info!("Launching Go sidecar from: {:?}", sidecar_path);

            // Use port 0 so the sidecar picks a free port
            let mut child = Command::new(sidecar_path)
                .env("GOKATAS_PORT", "0")
                .stdout(Stdio::piped())
                .stderr(Stdio::piped())
                .spawn()
                .expect("failed to spawn Go sidecar");

            let stderr = child.stderr.take();
            let _state_handle = app.state::<AppState>();

            // Read stderr for port discovery
            thread::spawn(move || {
                if let Some(stderr) = stderr {
                    use std::io::BufRead;
                    let reader = std::io::BufReader::new(stderr);
                    for line in reader.lines() {
                        match line {
                            Ok(line) => {
                                log::info!("[sidecar] {}", line);
                                if line.starts_with("GOKATAS_PORT=") {
                                    if let Ok(port) = line.trim_start_matches("GOKATAS_PORT=").parse::<u16>() {
                                        log::info!("Go sidecar listening on port {}", port);
                                        if let Some(s) = app_handle.try_state::<AppState>() {
                                            let mut p = s.port.lock().unwrap();
                                            *p = port;
                                            s.ready.notify_all();
                                            let _ = app_handle.emit("sidecar-ready", port);
                                        }
                                    }
                                }
                            }
                            Err(_) => break,
                        }
                    }
                }
                log::warn!("Sidecar stderr reader exited — process ended");
            });

            // Wait up to 20 seconds for the sidecar to report its port
            {
                let state = app.state::<AppState>();
                let port = *state.port.lock().unwrap();
                if port == 0 {
                    log::info!("Waiting for sidecar to start...");
                    let port_lock = state.port.lock().unwrap();
                    let (port_guard, _timeout) = state.ready.wait_timeout(port_lock, Duration::from_secs(20)).unwrap();
                    if *port_guard == 0 {
                        log::warn!("Sidecar did not report port after 20s — falling back to 9100");
                        // Write via a fresh lock
                        drop(port_guard);
                        *state.port.lock().unwrap() = 9100;
                    }
                }
            }

            let final_port = *app.state::<AppState>().port.lock().unwrap();
            log::info!("Sidecar ready on port {}", final_port);

            // Monitor sidecar lifecycle
            thread::spawn(move || {
                let status = child.wait();
                log::warn!("Sidecar process exited: {:?}", status);
            });

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![
            get_port,
            get_catalog,
            get_kata,
            save_kata,
            run_kata,
            get_progress,
            get_status,
            sync_content,
            lint_code,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
