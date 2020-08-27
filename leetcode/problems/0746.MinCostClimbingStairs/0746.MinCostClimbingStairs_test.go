package mincostclimbingstairs

import "testing"

func Test_mincostclimbingstairs(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			"Ex 1",
			[]int{10, 15, 20},
			15,
		},
		{
			"Ex 2",
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
	}
	for _, test := range testCases {
		got := minCostClimbingStairs(test.nums)
		if got != test.expect {
			t.Errorf("[%v] expect %v, got %v\n", test.name, test.expect, got)
		}
	}
}
