package findtheduplicatenumber

import "testing"

func Test_findDuplicate(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			"Example 1",
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			"Example 2",
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			"Example My 1",
			[]int{3, 3, 3, 4, 2},
			3,
		},
		{
			"Example My 1",
			[]int{1, 4, 3, 4, 4},
			4,
		},
	}
	for _, test := range testCases {
		t.Log("================================")
		got := findDuplicate(test.input)
		if got != test.expect {
			t.Errorf("[%v] expect '%v', got '%v' \n", test.name, test.expect, got)
		}
	}
}
