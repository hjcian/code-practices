package transposematrix

import (
	"reflect"
	"testing"
)

func Test_transposematrix(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]int
		expect [][]int
	}{
		{
			"Ex 1",
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{1, 4, 7},
				[]int{2, 5, 8},
				[]int{3, 6, 9},
			},
		},
		{
			"Ex 2",
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
			},
			[][]int{
				[]int{1, 4},
				[]int{2, 5},
				[]int{3, 6},
			},
		},
	}
	for _, test := range testCases {
		got := transpose(test.matrix)
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("[%v] expect %v, got %v\n", test.name, test.expect, got)
		}
	}
}
