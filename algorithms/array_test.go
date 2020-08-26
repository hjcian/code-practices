package arrayproblems

import (
	"reflect"
	"testing"
)

func Test_arrayproblems(t *testing.T) {
	t.Run("238. Product of Array Except Self", func(t *testing.T) {
		tests := []struct {
			input  []int
			output []int
		}{
			{
				[]int{1, 2, 3, 4},
				[]int{24, 12, 8, 6},
			},
		}
		for i, test := range tests {
			got := productExceptSelf(test.input)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("[case-%v] got %v, expect %v", i+1, got, test.output)
			}
		}
	})

	t.Run("78. Subsets", func(t *testing.T) {
		t.Skip()
		tests := []struct {
			input  []int
			output [][]int
		}{
			// {
			// 	[]int{1, 2, 3},
			// 	[][]int{[]int{3}, []int{1}, []int{2}, []int{1, 2, 3}, []int{1, 3}, []int{2, 3}, []int{1, 2}, []int{}},
			// },
			{
				[]int{9, 0, 3, 5, 7},
				[][]int{[]int{}},
			},
		}
		for i, test := range tests {
			got := subsets(test.input)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("[case-%v] got %v, expect %v", i+1, got, test.output)
			}
		}
	})

	t.Run("581. Shortest Unsorted Continuous Subarray", func(t *testing.T) {
		tests := []struct {
			input  []int
			output int
		}{
			{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
			{[]int{1, 2, 3, 4}, 0},
		}
		for i, test := range tests {
			got := findUnsortedSubarray(test.input)
			if got != test.output {
				t.Errorf("[case-%v] got %v, expect %v", i+1, got, test.output)
			}
		}
	})

	t.Run("1464. Maximum Product of Two Elements in an Array", func(t *testing.T) {
		tests := []struct {
			input  []int
			output int
		}{
			{[]int{3, 4, 5, 2}, 12},
			{[]int{1, 5, 4, 5}, 16},
			{[]int{3, 7}, 12},
		}
		for i, test := range tests {
			got := maxProduct(test.input)
			if got != test.output {
				t.Errorf("[case-%v] got %v, expect %v", i+1, got, test.output)
			}
		}
	})

}
