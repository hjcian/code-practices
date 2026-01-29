package nextgreaterelementwithpositionoffset

import "testing"

func TestFindNextGreaterElementsWithDistance(t *testing.T) {
	tests := []struct {
		name     string
		readings []int32
		expected [][]int32
	}{
		{
			name:     "empty array",
			readings: []int32{},
			expected: [][]int32{},
		},
		{
			name:     "single element",
			readings: []int32{5},
			expected: [][]int32{{-1, -1}},
		},
		{
			name:     "two elements ascending",
			readings: []int32{1, 2},
			expected: [][]int32{{2, 1}, {-1, -1}},
		},
		{
			name:     "two elements descending",
			readings: []int32{2, 1},
			expected: [][]int32{{-1, -1}, {-1, -1}},
		},
		{
			name:     "multiple elements",
			readings: []int32{1, 5, 0, 3, 4, 5},
			expected: [][]int32{{5, 1}, {-1, -1}, {3, 1}, {4, 1}, {5, 1}, {-1, -1}},
		},
		{
			name:     "increasing sequence",
			readings: []int32{1, 2, 3, 4, 5},
			expected: [][]int32{{2, 1}, {3, 1}, {4, 1}, {5, 1}, {-1, -1}},
		},
		{
			name:     "decreasing sequence",
			readings: []int32{5, 4, 3, 2, 1},
			expected: [][]int32{{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findNextGreaterElementsWithDistance(tt.readings)
			if len(result) != len(tt.expected) {
				t.Errorf("got length %d, want %d", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i][0] != tt.expected[i][0] || result[i][1] != tt.expected[i][1] {
					t.Errorf("at index %d: got %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}
