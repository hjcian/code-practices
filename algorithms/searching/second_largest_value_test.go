package searching

import "testing"

func TestSecondLargestValue(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
		ok       bool
	}{
		{
			name:     "normal case",
			arr:      []int{1, 5, 3, 9, 2},
			expected: 5,
			ok:       true,
		},
		{
			name:     "normal case 2",
			arr:      []int{3, 3, 2, 1},
			expected: 2,
			ok:       true,
		},
		{
			name:     "normal case 3",
			arr:      []int{3, 3, 4, 2, 1},
			expected: 3,
			ok:       true,
		},
		{
			name:     "two elements",
			arr:      []int{10, 5},
			expected: 5,
			ok:       true,
		},
		{
			name:     "duplicates",
			arr:      []int{7, 7, 3, 2},
			expected: 3,
			ok:       true,
		},
		{
			name:     "single element",
			arr:      []int{5},
			expected: 0,
			ok:       false,
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: 0,
			ok:       false,
		},
		{
			name:     "negative numbers",
			arr:      []int{-1, -5, -3, -2},
			expected: -2,
			ok:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := secondLargestValue(tt.arr)
			if ok != tt.ok || (ok && got != tt.expected) {
				t.Errorf("secondLargestValue(%v) = %d, %v; want %d, %v", tt.arr, got, ok, tt.expected, tt.ok)
			}
		})
	}
}
