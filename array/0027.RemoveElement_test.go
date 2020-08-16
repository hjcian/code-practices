package leetcode

import "testing"

func Test_27(t *testing.T) {
	type Input struct {
		nums []int
		val  int
	}
	testSuits := []struct {
		name      string
		input     Input
		expectLen int
	}{
		{
			"Example 1",
			Input{[]int{3, 2, 2, 3}, 3},
			2},
		{
			"Example 2",
			Input{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			5},
	}

	for _, test := range testSuits {
		got := removeElement(test.input.nums, test.input.val)
		if got != test.expectLen {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expectLen, got)
		}
	}
}
