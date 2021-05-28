package toeplitzmatrix

import "testing"

func Test_toeplitzmatrix(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]int
		expect bool
	}{
		{
			"Ex 1",
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 1, 2, 3},
				[]int{9, 5, 1, 2},
			},
			true,
		},
		{
			"Ex 2",
			[][]int{
				[]int{1, 2},
				[]int{2, 2},
			},
			false,
		},
	}
	for _, test := range testCases {
		got := isToeplitzMatrix(test.matrix)
		if got != test.expect {
			t.Errorf("[%v] expect %v, got %v\n", test.name, test.expect, got)
		}
	}
}
