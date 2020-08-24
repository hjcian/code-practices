package reshapethematrix

import (
	"reflect"
	"testing"
)

func Test_reshapethematrix(t *testing.T) {
	testCases := []struct {
		name   string
		nums   [][]int
		r      int
		c      int
		expect [][]int
	}{
		{
			"Example 1",
			[][]int{[]int{1, 2}, []int{3, 4}},
			1,
			4,
			[][]int{[]int{1, 2, 3, 4}},
		},
		{
			"Example 2",
			[][]int{[]int{1, 2}, []int{3, 4}},
			2,
			4,
			[][]int{[]int{1, 2}, []int{3, 4}},
		},
	}
	for _, test := range testCases {
		got := matrixReshape(test.nums, test.r, test.c)
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("[%v] expect '%v', got '%v' \n", test.name, test.expect, got)
		}
	}
}
