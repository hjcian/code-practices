package comparebstsforequalvaluesbutdifferentstructure

import "testing"

func TestVerifySameMultisetDifferentStructure(t *testing.T) {
	tests := []struct {
		name     string
		root1    []int32
		root2    []int32
		expected bool
	}{
		{
			name:     "same multiset different structure",
			root1:    []int32{4, 2, 5, 1, 3, 100001, 100001},
			root2:    []int32{3, 1, 5, 100001, 2, 4, 100001},
			expected: true,
		},
		{
			name:     "same multiset same structure",
			root1:    []int32{1, 100001, 100001},
			root2:    []int32{1, 100001, 100001},
			expected: false,
		},
		{
			name:     "different multisets",
			root1:    []int32{1, 2, 3},
			root2:    []int32{1, 2, 4},
			expected: false,
		},
		{
			name:     "empty first array",
			root1:    []int32{},
			root2:    []int32{1, 100001, 100001},
			expected: false,
		},
		{
			name:     "empty second array",
			root1:    []int32{1, 100001, 100001},
			root2:    []int32{},
			expected: false,
		},
		{
			name:     "single node same",
			root1:    []int32{5},
			root2:    []int32{5},
			expected: false,
		},
		{
			name:     "different lengths same values",
			root1:    []int32{2, 1, 3, 100001, 100001, 100001, 100001},
			root2:    []int32{1, 100001, 2, 100001, 3},
			expected: true,
		},
		{
			name:     "identical serialization including sentinels",
			root1:    []int32{4, 2, 5, 1, 3, 100001, 100001},
			root2:    []int32{4, 2, 5, 1, 3, 100001, 100001},
			expected: false,
		},
		{
			name:     "all sentinels (both empty trees)",
			root1:    []int32{100001, 100001, 100001},
			root2:    []int32{100001, 100001, 100001},
			expected: false,
		},
		{
			name:     "single node vs padded sentinels",
			root1:    []int32{5},
			root2:    []int32{5, 100001, 100001},
			expected: false,
		},
		{
			name:     "duplicates same multiset different shape",
			root1:    []int32{2, 2, 100001}, // root 2 with left child 2
			root2:    []int32{2, 100001, 2}, // root 2 with right child 2
			expected: true,
		},
		{
			name:     "both empty arrays",
			root1:    []int32{},
			root2:    []int32{},
			expected: false,
		},
		{
			name:     "same structure extra trailing sentinels",
			root1:    []int32{1, 100001, 100001},
			root2:    []int32{1, 100001, 100001, 100001, 100001},
			expected: false,
		},
		{
			name:     "root is sentinel (both empty expressed differently)",
			root1:    []int32{100001},
			root2:    []int32{100001, 100001},
			expected: false,
		},
		{
			name:     "empty array vs only sentinels",
			root1:    []int32{},
			root2:    []int32{100001, 100001},
			expected: false,
		},
		{
			name:     "all same values left chain vs right chain",
			root1:    []int32{3, 100001, 3, 100001, 3},          // 3 -> right 3 -> right 3
			root2:    []int32{3, 3, 100001, 3, 100001, 100001}, // 3 -> left 3, left child has right 3
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := verifySameMultisetDifferentStructure(tt.root1, tt.root2)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
