package longestincreasingsubsequencelength

import "testing"

func TestComputeLongestIncreasingSubsequenceLength(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		quality  []int32
		expected int32
	}{
		{
			name:     "empty array",
			n:        0,
			quality:  []int32{},
			expected: 0,
		},
		{
			name:     "single element",
			n:        1,
			quality:  []int32{5},
			expected: 1,
		},
		{
			name:     "all increasing",
			n:        4,
			quality:  []int32{1, 2, 3, 4},
			expected: 4,
		},
		{
			name:     "all decreasing",
			n:        4,
			quality:  []int32{4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "mixed sequence",
			n:        6,
			quality:  []int32{3, 10, 2, 1, 20, 25},
			expected: 4,
		},
		{
			name:     "duplicates",
			n:        5,
			quality:  []int32{1, 3, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "two elements increasing",
			n:        2,
			quality:  []int32{1, 2},
			expected: 2,
		},
		{
			name:     "two elements decreasing",
			n:        2,
			quality:  []int32{2, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeLongestIncreasingSubsequenceLength(tt.n, tt.quality)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
