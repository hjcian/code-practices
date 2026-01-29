package subarrayswithgivensumandboundedmaximum

import "testing"

func TestCountSubarraysWithSumAndMaxAtMost(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int32
		k        int64
		M        int64
		expected int64
	}{
		{
			name:     "empty array",
			nums:     []int32{},
			k:        5,
			M:        2,
			expected: 0,
		},
		{
			name:     "single element within bounds",
			nums:     []int32{1},
			k:        1,
			M:        2,
			expected: 1,
		},
		{
			name:     "single element exceeds max",
			nums:     []int32{5},
			k:        5,
			M:        2,
			expected: 0,
		},
		{
			name:     "basic case with multiple subarrays",
			nums:     []int32{1, 2, 3},
			k:        3,
			M:        2,
			expected: 1,
		},
		{
			name:     "all elements exceed max",
			nums:     []int32{5, 6, 7},
			k:        10,
			M:        2,
			expected: 0,
		},
		{
			name:     "subarrays with delimiter",
			nums:     []int32{1, 2, 3, 5, 1, 2},
			k:        3,
			M:        2,
			expected: 2,
		},
		{
			name:     "larger array",
			nums:     []int32{1, 1, 1, 1},
			k:        2,
			M:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countSubarraysWithSumAndMaxAtMost(tt.nums, tt.k, tt.M)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
