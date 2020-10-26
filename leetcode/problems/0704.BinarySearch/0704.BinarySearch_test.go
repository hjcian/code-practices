package binarysearch

import (
	"testing"
)

func Test_search(t *testing.T) {
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
			"Ex1",
			args{[]int{-1, 0, 3, 5, 9, 12, 13}, 9},
			4,
		},
		{
			"Ex2",
			args{[]int{-1, 0, 3, 5, 9, 12, 13}, 2},
			-1,
		},
		{
			"Test (45 / 46 test cases passed.)",
			args{[]int{2, 5}, 2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
