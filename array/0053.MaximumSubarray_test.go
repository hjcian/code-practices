package leetcode

import "testing"

func Test_53(t *testing.T) {
	type Input struct {
		nums []int
	}
	testSuits := []struct {
		name   string
		input  Input
		expect int
	}{
		{
			"Example 1",
			Input{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		},
	}
	for _, test := range testSuits {
		got := maxSubArray(test.input.nums)
		if got != test.expect {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, got)
		}
	}
}
