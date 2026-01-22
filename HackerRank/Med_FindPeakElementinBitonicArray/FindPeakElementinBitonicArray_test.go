package findpeakelementinbitonicarray

import "testing"

func TestFindPeakIndex(t *testing.T) {
	tests := []struct {
		name     string
		counts   []int32
		expected int32
	}{
		{
			name:     "peak in middle",
			counts:   []int32{1, 2, 3, 4, 5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "peak at end",
			counts:   []int32{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "peak at start",
			counts:   []int32{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "single element",
			counts:   []int32{5},
			expected: 0,
		},
		{
			name:     "two elements ascending",
			counts:   []int32{1, 5},
			expected: 1,
		},
		{
			name:     "two elements descending",
			counts:   []int32{5, 1},
			expected: 0,
		},
		{
			name:     "large values with peak",
			counts:   []int32{10, 20, 30, 40, 50, 40, 30, 20, 10},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakIndex(tt.counts)
			if result != tt.expected {
				t.Errorf("findPeakIndex(%v) = %d, want %d", tt.counts, result, tt.expected)
			}
		})
	}
}
