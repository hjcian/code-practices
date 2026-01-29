package longestincreasingsubsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"increasing sequence", []int{1, 2, 3, 4, 5}, 5},
		{"decreasing sequence", []int{5, 4, 3, 2, 1}, 1},
		{"example 1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"example 2", []int{0, 1, 0, 4, 4, 4, 3, 5, 9, 7, 6}, 5},
		{"duplicates", []int{1, 1, 1, 1}, 1},
		{"mixed", []int{3, 1, 4, 1, 5, 9, 2, 6}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLIS(tt.nums)
			if result != tt.expected {
				t.Errorf("lengthOfLIS(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
