//go:build ignore
// +build ignore

package kata

// SumPositive currently sums all values, including negatives.
func SumPositive(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
