package twosum

import (
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	type Input struct {
		nums   []int
		target int
	}
	testSuits := []struct {
		name   string
		input  Input
		expect []int
	}{
		{
			"Example",
			Input{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			"Test Error 1",
			Input{[]int{3, 2, 4}, 6},
			[]int{1, 2},
		},
		{
			"Test Error 2",
			Input{[]int{3, 3}, 6},
			[]int{0, 1},
		},
	}
	t.Run("first try", func(t *testing.T) {
		for _, test := range testSuits {
			got := twoSum(test.input.nums, test.input.target)
			if !reflect.DeepEqual(got, test.expect) {
				t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, got)
			}
		}
	})
	t.Run("twoSum20210528", func(t *testing.T) {
		for _, test := range testSuits {
			got := twoSum20210528(test.input.nums, test.input.target)
			if !reflect.DeepEqual(got, test.expect) {
				t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, got)
			}
		}
	})
	t.Run("twoSum20220820", func(t *testing.T) {
		for _, test := range testSuits {
			got := twoSum20220820(test.input.nums, test.input.target)
			if !reflect.DeepEqual(got, test.expect) {
				t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, got)
			}
		}
	})
}
