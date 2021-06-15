package sortcolors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			"ex1",
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			"ex2",
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
		{
			"ex3",
			[]int{0},
			[]int{0},
		},
		{
			"ex4",
			[]int{1},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			assert.Equal(t, tt.expect, tt.nums)
		})
	}
}
