package longestarithmeticsubsequencewithgivendifference

import "testing"

func TestFindLongestArithmeticProgression(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int32
		k        int32
		expected int32
	}{
		{
			name:     "empty array",
			arr:      []int32{},
			k:        1,
			expected: 0,
		},
		{
			name:     "single element",
			arr:      []int32{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "two elements with difference k",
			arr:      []int32{1, 2},
			k:        1,
			expected: 2,
		},
		{
			name:     "arithmetic sequence (3->6)",
			arr:      []int32{1, 3, 6, 7},
			k:        3,
			expected: 2,
		},
		{
			name:     "longer arithmetic sequence",
			arr:      []int32{9, 4, 7, 2, 10},
			k:        5,
			expected: 2,
		},
		{
			name:     "arithmetic progression (1->4)",
			arr:      []int32{1, 2, 4, 8},
			k:        3,
			expected: 2,
		},
		{
			name:     "unsorted array with progression",
			arr:      []int32{20, 1, 15, 3, 10, 5},
			k:        5,
			expected: 4,
		},
		{
			name:     "large k",
			arr:      []int32{20, 1, 15, 3, 10, 5},
			k:        50,
			expected: 1,
		},
		{
			name:     "desc array",
			arr:      []int32{5, 4, 3, 2, 1},
			k:        1,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLongestArithmeticProgression(tt.arr, tt.k)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
