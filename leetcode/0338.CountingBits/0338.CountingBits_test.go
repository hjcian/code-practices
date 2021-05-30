package countingbits

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Ex 1",
			args{2},
			[]int{0, 1, 1},
		},
		{
			"Ex 2",
			args{5},
			[]int{0, 1, 1, 2, 1, 2},
		},
	}
	t.Run("Basic solution", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("countBits() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("Kernighan's Algorithm", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := countBitsByKernighansAlgorithm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("countBits() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
