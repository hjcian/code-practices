package mergeandsortintervals

import "testing"

func TestMergeHighDefinitionIntervals(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int32
		expected [][]int32
	}{
		{
			name:     "empty intervals",
			input:    [][]int32{},
			expected: [][]int32{},
		},
		{
			name:     "single interval",
			input:    [][]int32{{1, 3}},
			expected: [][]int32{{1, 3}},
		},
		{
			name:     "no overlapping intervals",
			input:    [][]int32{{1, 2}, {3, 4}, {5, 6}},
			expected: [][]int32{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:     "overlapping intervals",
			input:    [][]int32{{1, 3}, {2, 6}, {8, 10}},
			expected: [][]int32{{1, 6}, {8, 10}},
		},
		{
			name:     "completely overlapping",
			input:    [][]int32{{1, 10}, {2, 5}, {3, 7}},
			expected: [][]int32{{1, 10}},
		},
		{
			name:     "touching intervals",
			input:    [][]int32{{1, 2}, {2, 3}},
			expected: [][]int32{{1, 3}},
		},
		{
			name:     "unsorted input",
			input:    [][]int32{{5, 6}, {1, 3}, {2, 4}},
			expected: [][]int32{{1, 4}, {5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeHighDefinitionIntervals(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("got %d intervals, want %d", len(result), len(tt.expected))
			}
			for i := range result {
				if result[i][0] != tt.expected[i][0] || result[i][1] != tt.expected[i][1] {
					t.Errorf("interval %d: got [%d, %d], want [%d, %d]", i, result[i][0], result[i][1], tt.expected[i][0], tt.expected[i][1])
				}
			}
		})
	}
}
