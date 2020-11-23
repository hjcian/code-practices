package besttravel

import "testing"

func TestChooseBestSum(t *testing.T) {
	type args struct {
		t  int
		k  int
		ls []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"failed 1",
			args{331, 1, []int{91, 74, 73, 85, 73, 81, 87}},
			91,
		},
		{
			"test 1",
			args{163, 3, []int{50, 55, 56, 57, 58}},
			163,
		},
		{
			"test 2",
			args{163, 3, []int{50}},
			-1,
		},
		{
			"test 3",
			args{230, 3, []int{91, 74, 73, 85, 73, 81, 87}},
			228,
		},
		{
			"test 4",
			args{331, 2, []int{91, 74, 73, 85, 73, 81, 87}},
			178,
		},
		{
			"test 5",
			args{331, 4, []int{91, 74, 73, 85, 73, 81, 87}},
			331,
		},
		{
			"test 6",
			args{331, 5, []int{91, 74, 73, 85, 73, 81, 87}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChooseBestSum(tt.args.t, tt.args.k, tt.args.ls); got != tt.want {
				t.Errorf("ChooseBestSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
