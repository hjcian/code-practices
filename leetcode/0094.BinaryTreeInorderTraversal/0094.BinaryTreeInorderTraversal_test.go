package binarytreeinordertraversal

import (
	"codepractices/leetcode/helper"
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  []int
	}{
		{
			"ex1",
			[]int{1, helper.NULL, 2, 3},
			[]int{1, 3, 2},
		},
		{
			"ex2",
			[]int{},
			[]int{},
		},
		{
			"ex3",
			[]int{1},
			[]int{1},
		},
		{
			"README sample",
			[]int{1, 2, 3, 4, 5},
			[]int{4, 2, 5, 1, 3},
		},
	}
	for _, tt := range tests {
		root := helper.BuildTree(tt.nodes)

		t.Run("recursive_"+tt.name, func(t *testing.T) {
			got := []int{}
			recursive(root, &got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}
		})

		t.Run("iterative_"+tt.name, func(t *testing.T) {
			if got := iterative(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
