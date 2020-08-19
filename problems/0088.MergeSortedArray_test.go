package leetcode

import (
	"reflect"
	"testing"
)

func Test_88(t *testing.T) {
	type Input struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	testSuits := []struct {
		name   string
		input  Input
		expect []int
	}{
		{
			"Test 11",
			Input{
				[]int{4, 5, 6, 0, 0, 0},
				3,
				[]int{1, 2, 3},
				3,
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
		// {
		// 	"Test 1",
		// 	Input{
		// 		[]int{1},
		// 		1,
		// 		[]int{},
		// 		0,
		// 	},
		// 	[]int{1},
		// },
		// {
		// 	"Example",
		// 	Input{
		// 		[]int{1, 2, 3, 0, 0, 0},
		// 		3,
		// 		[]int{2, 5, 6},
		// 		3,
		// 	},
		// 	[]int{1, 2, 2, 3, 5, 6},
		// },
	}
	for _, test := range testSuits {
		merge(test.input.nums1, test.input.m, test.input.nums2, test.input.n)
		if !reflect.DeepEqual(test.input.nums1, test.expect) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, test.input.nums1)
		}
	}
}
