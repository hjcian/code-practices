package threesumclosest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				nums:   []int{-1, 2, -1, 4},
				target: 1,
			},
			2,
		},
		{
			"test 1",
			args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			0,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.args.nums, tt.args.target)
			t.Log(tt.args.nums, tt.args.target)
			t.Log(got)
			t.Log(tt.want)
			assert.EqualValues(t, tt.want, got, fmt.Sprintf("[%v] should euqal", tt.name))
		})
	}
}
