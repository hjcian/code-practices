package arraypartitioni

import "testing"

func Test_arraypartitioni(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			"Example",
			[]int{1, 4, 3, 2},
			4,
		},
	}
	for _, test := range testCases {
		got := arrayPairSum(test.nums)
		if got != test.expect {
			t.Errorf("[%v] expect %v, got %v \n", test.name, test.expect, got)
		}
	}
}
