use std::process::Command;
use std::sync::Mutex;
use std::thread;
use std::time::Duration;
use tauri::Manager;

struct SidecarState {
    port: Mutex<u16>,
}

#[tauri::command]
fn get_port(state: tauri::State<SidecarState>) -> u16 {
    let port = *state.port.lock().unwrap();
    port
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .manage(SidecarState { port: Mutex::new(0) })
        .setup(|app| {
            let app_handle = app.handle().clone();

            // Find the sidecar binary — check multiple locations
            let candidates = vec![
                app.path().resource_dir().unwrap().join("sidecar").join("gokatas-sidecar"),
                std::path::PathBuf::from("/usr/bin/gokatas-sidecar"),
                std::env::current_dir().unwrap().join("src-tauri").join("sidecar").join("gokatas-sidecar"),
            ];
            let sidecar_path = candidates.into_iter().find(|p| p.exists())
                .expect("gokatas-sidecar binary not found — install the sidecar or run from the project directory");

            log::info!("Launching Go sidecar from: {:?}", sidecar_path);

            let mut child = Command::new(&sidecar_path)
                .env("GOKATAS_PORT", "9100")
                .stdout(std::process::Stdio::piped())
                .stderr(std::process::Stdio::piped())
                .spawn()
                .expect("failed to spawn Go sidecar");

            let stderr = child.stderr.take();
            thread::spawn(move || {
                if let Some(stderr) = stderr {
                    use std::io::BufRead;
                    let reader = std::io::BufReader::new(stderr);
                    for line in reader.lines() {
                        match line {
                            Ok(line) => {
                                if line.starts_with("GOKATAS_PORT=") {
                                    if let Ok(port) = line.trim_start_matches("GOKATAS_PORT=").parse::<u16>() {
                                        log::info!("Go sidecar listening on port {}", port);
                                        if let Some(state) = app_handle.try_state::<SidecarState>() {
                                            let mut p = state.port.lock().unwrap();
                                            *p = port;
                                        }
                                    }
                                } else {
                                    log::info!("[sidecar] {}", line);
                                }
                            }
                            Err(_) => break,
                        }
                    }
                }
            });

            thread::sleep(Duration::from_millis(800));
            {
                let state = app.state::<SidecarState>();
                let port = *state.port.lock().unwrap();
                if port == 0 {
                    let mut p = state.port.lock().unwrap();
                    *p = 9100;
                }
            }

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![get_port])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
