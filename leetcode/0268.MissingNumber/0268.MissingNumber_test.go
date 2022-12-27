package missingnumber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0268_MissingNumber(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "sample 1",
			arr:  []int{3, 0, 1},
			want: 2,
		},
		{
			name: "sample 2",
			arr:  []int{0, 1},
			want: 2,
		},
		{
			name: "sample 3",
			arr:  []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}

	// solutions := []func(arr []int) int{
	// 	missingNumber,
	// }
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			got := missingNumber_bit_operation(tt.arr)
			require.Equal(t, tt.want, got)
		})
	}
}
