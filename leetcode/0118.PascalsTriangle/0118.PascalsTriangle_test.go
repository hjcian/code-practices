package pascalstriangle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0118_PascalsTriangle(t *testing.T) {
	tests := []struct {
		name        string
		giveNumRows int
		want        [][]int
	}{
		{
			name:        "0 rows",
			giveNumRows: 0,
			want:        nil,
		},
		{
			name:        "1 rows",
			giveNumRows: 1,
			want:        [][]int{{1}},
		},
		{
			name:        "2 rows",
			giveNumRows: 2,
			want:        [][]int{{1}, {1, 1}},
		},
		{
			name:        "3 rows",
			giveNumRows: 3,
			want:        [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name:        "4 rows",
			giveNumRows: 4,
			want:        [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			name:        "5 rows",
			giveNumRows: 5,
			want:        [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.giveNumRows)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
