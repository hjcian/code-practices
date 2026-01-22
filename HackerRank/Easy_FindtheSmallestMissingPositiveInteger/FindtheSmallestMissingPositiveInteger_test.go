package findthesmallestmissingpositiveinteger

import "testing"

func TestFindSmallestMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected int32
	}{
		{
			name:     "empty slice",
			input:    []int32{},
			expected: 1,
		},
		{
			name:     "example case",
			input:    []int32{3, 4, -1, 1},
			expected: 2,
		},
		{
			name:     "duplicates",
			input:    []int32{1, 1},
			expected: 2,
		},
		{
			name:     "consecutive sequence",
			input:    []int32{1, 2, 3},
			expected: 4,
		},
		{
			name:     "with gaps and negatives",
			input:    []int32{0, 10, 2, -10, 1},
			expected: 3,
		},
		{
			name:     "single element",
			input:    []int32{1},
			expected: 2,
		},
		{
			name:     "missing one at start",
			input:    []int32{2, 3, 4},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSmallestMissingPositive(tt.input)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
