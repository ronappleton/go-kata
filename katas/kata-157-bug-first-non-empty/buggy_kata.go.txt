//go:build ignore
// +build ignore

package kata

// FirstNonEmpty currently returns the first item even if empty.
func FirstNonEmpty(values []string, fallback string) string {
	if len(values) == 0 {
		return fallback
	}
	return values[0]
}
