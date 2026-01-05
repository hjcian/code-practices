package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "unsorted array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "duplicates",
			input:    []int{3, 3, 1, 1, 2},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input) // to avoid modifying the original input
			result := insertionSort(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("input: %v, got %v, want %v", input, result, tt.expected)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("input: %v, got %v, want %v", input, result, tt.expected)
				}
			}
		})
	}
}
