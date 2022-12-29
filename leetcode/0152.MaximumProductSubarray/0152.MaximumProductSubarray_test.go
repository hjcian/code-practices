package maximumproductsubarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0152_MaximumProductSubarray(t *testing.T) {
	tests := []struct {
		name     string
		giveNums []int
		want     int
	}{
		// {
		// 	name:     "ex1",
		// 	giveNums: []int{2, 3, -2, 4},
		// 	want:     6,
		// },
		// {
		// 	name:     "ex2",
		// 	giveNums: []int{-2, 0, -1},
		// 	want:     0,
		// },
		// {
		// 	name:     "ex3",
		// 	giveNums: []int{1, 2, 3},
		// 	want:     6,
		// },
		// {
		// 	name:     "ex4",
		// 	giveNums: []int{1, -1, 3},
		// 	want:     3,
		// },
		// {
		// 	name:     "ex5",
		// 	giveNums: []int{1, 2, 3, 0, 1, 5, 0, 4, 1},
		// 	want:     6,
		// },
		{
			name:     "ex6",
			giveNums: []int{-2, -3, -1, 2, -4, -3, -1},
			want:     144,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := maxProduct(tt.giveNums)
			require.Equal(t, tt.want, got)
		})
	}
}
