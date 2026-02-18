package kata

import "strings"

// NormalizeUsername currently contains a bug: it misses normalization steps.
func NormalizeUsername(name string) string {
	return strings.ToLower(name)
}
