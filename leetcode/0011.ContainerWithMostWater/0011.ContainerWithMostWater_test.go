package containerwithmostwater

import "testing"

func Test_containerwithmostwater(t *testing.T) {
	testCases := []struct {
		name   string
		height []int
		expect int
	}{
		{
			"Ex",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
	}
	for _, test := range testCases {
		got := maxArea(test.height)
		if got != test.expect {
			t.Errorf("[%v] expect %v, got %v\n", test.name, test.expect, got)
		}
	}
}
