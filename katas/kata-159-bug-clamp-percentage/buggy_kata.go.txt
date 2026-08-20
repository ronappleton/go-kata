//go:build ignore
// +build ignore

package kata

// ClampPercentage currently fails to clamp negative values.
func ClampPercentage(v int) int {
	if v > 100 {
		return 100
	}
	return v
}
