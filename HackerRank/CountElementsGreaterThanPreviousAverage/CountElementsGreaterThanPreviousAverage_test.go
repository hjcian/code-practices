package countelementsgreaterthanpreviousaverage

import "testing"

func TestCountResponseTimeRegressions(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected int32
	}{
		{
			name:     "empty slice",
			input:    []int32{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int32{5},
			expected: 0,
		},
		{
			name:     "increasing sequence",
			input:    []int32{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "all equal elements",
			input:    []int32{5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "decreasing sequence",
			input:    []int32{10, 5, 3, 1},
			expected: 0,
		},
		{
			name:     "mixed values",
			input:    []int32{1, 2, 1, 3, 2},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countResponseTimeRegressions(tt.input)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
