package leetcode

import "testing"

func Test_219(t *testing.T) {
	type Input struct {
		nums []int
		k    int
	}
	testSuits := []struct {
		name   string
		input  Input
		expect bool
	}{
		{
			"Example 1",
			Input{
				[]int{1, 2, 3, 1},
				3,
			},
			true,
		},
		{
			"Example 2",
			Input{
				[]int{1, 0, 1, 1},
				1,
			},
			true,
		},
		{
			"Example 3",
			Input{
				[]int{1, 2, 3, 1, 2, 3},
				2,
			},
			false,
		},
	}

	for _, test := range testSuits {
		got := containsNearbyDuplicate(test.input.nums, test.input.k)
		if got != test.expect {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, got)
		}
	}
}
