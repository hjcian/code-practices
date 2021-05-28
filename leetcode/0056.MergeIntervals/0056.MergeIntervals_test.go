package mergeintervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Ex 1",
			args{[][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}},
			[][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}},
		},
		{
			"Ex 2",
			args{[][]int{[]int{1, 4}, []int{4, 5}}},
			[][]int{[]int{1, 5}},
		},
		{
			"test 61 / 169",
			args{[][]int{[]int{1, 4}, []int{2, 3}}},
			[][]int{[]int{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
