package addoverflowdetect

import (
	"fmt"
	"testing"
)

func Test_AddOK(t *testing.T) {
	tests := []struct {
		a    int8
		b    int8
		want bool
	}{
		{127, 127, false},
		{-128, -128, false},
		{127, 1, false},
		{-128, -1, false},

		{127, 0, true},
		{0, 127, true},
		{-128, 0, true},
		{0, -128, true},

		{123, -123, true},
		{-123, 123, true},
		{-10, 123, true},
		{-1, 127, true},
		{-128, 1, true},
		{-10, 10, true},
		{10, -10, true},
		{5, 5, true},
		{-5, -5, true},
	}
	for _, test := range tests {
		fmt.Println("============")
		got := addOK(test.a, test.b)
		if got != test.want {
			t.Errorf("Expect %v %v is %v", test.a, test.b, test.want)
		}
	}
}
