package threesum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Example",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			"Test 309/313",
			args{[]int{0, 0, 0}},
			[][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)
			require.EqualValues(t, tt.want, got, fmt.Sprintf("[%v] should equal", tt.name))
		})
	}
}
