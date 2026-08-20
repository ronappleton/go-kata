package content

import "encoding/json"

// KataVersion extracts the version from kata metadata JSON.
// Returns "1.0.0" as default if no version is specified.
func KataVersion(jsonData string) string {
	if jsonData == "" {
		return "1.0.0"
	}
	var v struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal([]byte(jsonData), &v); err != nil || v.Version == "" {
		return "1.0.0"
	}
	return v.Version
}

// VersionChanged returns true if the kata version has changed.
func VersionChanged(oldVersion, newVersion string) bool {
	return oldVersion != "" && newVersion != "" && oldVersion != newVersion
}

// FormatVersionStatus returns a human-readable status for version comparison.
func FormatVersionStatus(completedVersion, currentVersion string) string {
	if completedVersion == "" {
		return "new"
	}
	if completedVersion == currentVersion {
		return "completed"
	}
	return "updated"
}
