package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort20260106(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty array",
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
			name:     "random order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:     "duplicates",
			input:    []int{5, 2, 8, 2, 9, 1, 5},
			expected: []int{1, 2, 2, 5, 5, 8, 9},
		},
		{
			name:     "negative numbers",
			input:    []int{-3, -1, -4, -1, -5, 9, 2, 6},
			expected: []int{-5, -4, -3, -1, -1, 2, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original test case
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			result := quickSort_20260106(inputCopy)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("quickSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		left          int
		right         int
		expectedIndex int
	}{
		{
			name:          "simple partition",
			input:         []int{3, 1, 4, 2},
			left:          0,
			right:         3,
			expectedIndex: 1,
		},
		{
			name:          "all smaller than pivot",
			input:         []int{1, 2, 3, 4, 5},
			left:          0,
			right:         4,
			expectedIndex: 4,
		},
		{
			name:          "all larger than pivot",
			input:         []int{5, 4, 3, 2, 1},
			left:          0,
			right:         4,
			expectedIndex: 0,
		},
		{
			name:          "duplicated elements",
			input:         []int{-1, -4, -1, -3},
			left:          0,
			right:         3,
			expectedIndex: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			pivotIndex := partition_20260106(inputCopy, tt.left, tt.right)

			if pivotIndex != tt.expectedIndex {
				t.Errorf("partition(%v, %d, %d) returned index %d (%v), want %d",
					tt.input, tt.left, tt.right, pivotIndex, inputCopy, tt.expectedIndex)
			}

			// Verify that all elements before pivot are smaller
			pivot := inputCopy[pivotIndex]
			for i := tt.left; i < pivotIndex; i++ {
				if inputCopy[i] > pivot {
					t.Errorf("Element at index %d (%d) is greater than pivot (%d)",
						i, inputCopy[i], pivot)
				}
			}

			// Verify that all elements after pivot are greater or equal
			for i := pivotIndex + 1; i <= tt.right; i++ {
				if inputCopy[i] < pivot {
					t.Errorf("Element at index %d (%d) is less than pivot (%d)",
						i, inputCopy[i], pivot)
				}
			}
		})
	}
}
