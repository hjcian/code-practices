package kdiffpairsinanarray

import "testing"

func Test_kdiffpairsinanarray(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		k      int
		expect int
	}{
		{
			"Ex 1",
			[]int{3, 1, 4, 1, 5},
			2,
			2,
		},
		{
			"Ex 2",
			[]int{1, 2, 3, 4, 5},
			1,
			4,
		},
		{
			"Ex 3",
			[]int{1, 3, 1, 5, 4},
			0,
			1,
		},
		{
			"Test 61",
			[]int{6, 3, 5, 7, 2, 3, 3, 8, 2, 4},
			2,
			5,
		},
	}
	for _, test := range testCases {
		t.Logf("===== [%v] ===== \n", test.name)
		got := findPairs(test.nums, test.k)
		if got != test.expect {
			t.Errorf("[%v] expect %v, got %v \n", test.name, test.expect, got)
		}
	}
}
