package pascalstriangleii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0119_PascalsTriangleII(t *testing.T) {
	tests := []struct {
		name         string
		giveRowIndex int
		want         []int
	}{
		{
			name:         "0",
			giveRowIndex: 0,
			want:         []int{1},
		},
		{
			name:         "1",
			giveRowIndex: 1,
			want:         []int{1, 1},
		},
		{
			name:         "2",
			giveRowIndex: 2,
			want:         []int{1, 2, 1},
		},
		{
			name:         "3",
			giveRowIndex: 3,
			want:         []int{1, 3, 3, 1},
		},
		{
			name:         "4",
			giveRowIndex: 4,
			want:         []int{1, 4, 6, 4, 1},
		},
		{
			name:         "5",
			giveRowIndex: 5,
			want:         []int{1, 5, 10, 10, 5, 1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.giveRowIndex)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
