package searchinsertposition

import "testing"

func Test_searchInsert(t *testing.T) {
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
			"Ex 1",
			args{[]int{1, 3, 5, 6}, 5},
			2,
		},
		{
			"Ex 2",
			args{[]int{1, 3, 5, 6}, 2},
			1,
		},
		{
			"Ex 3",
			args{[]int{1, 3, 5, 6}, 7},
			4,
		},
		{
			"Ex 4",
			args{[]int{1, 3, 5, 6}, 0},
			0,
		},
		{
			"Ex 5",
			args{[]int{1}, 0},
			0,
		},
		{
			"9/62 test cases passed",
			args{[]int{1}, 1},
			0,
		},
		{
			"Runtime Error",
			args{[]int{1, 3}, 4},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
