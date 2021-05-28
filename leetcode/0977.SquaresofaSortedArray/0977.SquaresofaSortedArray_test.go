package squaresofasortedarray

import (
	"reflect"
	"testing"
)

func Test_squaresofasortedarray(t *testing.T) {
	testCases := []struct {
		name   string
		A      []int
		expect []int
	}{
		{
			"Ex 1",
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			"Ex 2",
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
		{
			"Test 46",
			[]int{-3, -3, -2, 1},
			[]int{1, 4, 9, 9},
		},
	}
	for _, test := range testCases {
		got := sortedSquares(test.A)
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("expected %v, got %v", test.expect, got)
		}
	}
}
