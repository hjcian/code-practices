package maximumnumberofnonoverlappingintervals

import "testing"

func TestMaximizeNonOverlappingMeetings(t *testing.T) {
	tests := []struct {
		name     string
		meetings [][]int32
		expected int32
	}{
		{
			name:     "no meetings",
			meetings: [][]int32{},
			expected: 1,
		},
		{
			name:     "single meeting",
			meetings: [][]int32{{1, 2}},
			expected: 1,
		},
		{
			name:     "two non-overlapping meetings",
			meetings: [][]int32{{1, 2}, {3, 4}},
			expected: 2,
		},
		{
			name:     "two overlapping meetings",
			meetings: [][]int32{{1, 3}, {2, 4}},
			expected: 1,
		},
		{
			name:     "multiple meetings with selection needed",
			meetings: [][]int32{{1, 2}, {2, 3}, {3, 4}},
			expected: 3,
		},
		{
			name:     "complex case",
			meetings: [][]int32{{1, 4}, {2, 5}, {3, 6}, {5, 7}},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximizeNonOverlappingMeetings(tt.meetings)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
