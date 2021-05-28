package maximumproductofthreenumbers

import "testing"

func Test_maximumproductofthreenumbers(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			"Ex 1",
			[]int{1, 2, 3},
			6,
		},
		{
			"Ex 2",
			[]int{1, 2, 3, 4},
			24,
		},
		{
			"Ex 3",
			[]int{1, 2, 3, 4},
			24,
		},
	}
	for _, test := range testCases {
		got := maximumProduct(test.nums)
		if got != test.expect {
			t.Errorf("[%v] expect %v, got %v\n", test.name, test.expect, got)
		}
	}
}
