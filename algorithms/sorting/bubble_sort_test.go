package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
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
			name:     "unsorted",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "duplicates",
			input:    []int{2, 2, 1, 1},
			expected: []int{1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			result := bubbleSort(tt.input)
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
