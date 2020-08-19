package movezeroes

import (
	"reflect"
	"testing"
)

func Test_283(t *testing.T) {
	testSuits := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"Example", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Test 2", []int{1}, []int{1}},
		{"My Test", []int{1, 0, 0, 0, 3, 12}, []int{1, 3, 12, 0, 0, 0}},
	}
	for _, test := range testSuits {
		moveZeroes(test.input)
		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expect, test.input)
		}
	}
}
