package medimummatrixpath

import "testing"

func TestMatrixPath(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected bool
	}{
		{
			name: "simple path exists",
			matrix: [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 1, 1},
			},
			expected: true,
		},
		{
			name: "no path blocked",
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 1},
				{0, 0, 1},
			},
			expected: false,
		},
		{
			name: "single cell",
			matrix: [][]int{
				{1},
			},
			expected: true,
		},
		{
			name: "start is zero",
			matrix: [][]int{
				{0, 1},
				{1, 1},
			},
			expected: false,
		},
		{
			name: "end is zero",
			matrix: [][]int{
				{1, 1},
				{1, 0},
			},
			expected: false,
		},
		{
			name: "path with backtrack",
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: true,
		},
		{
			name: "all ones",
			matrix: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MatrixPath(tt.matrix)
			if result != tt.expected {
				t.Errorf("MatrixPath() = %v, want %v", result, tt.expected)
			}
		})
	}
}
