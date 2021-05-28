package combinationsum

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Ex 1",
			args{[]int{2, 3, 6, 7}, 7},
			[][]int{[]int{7}, []int{2, 2, 3}},
		},
		// {
		// 	"Ex 2",
		// 	args{[]int{2, 3, 5}, 8},
		// 	[][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
