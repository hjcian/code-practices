package climbingstairs

import "testing"

func Test_climbStairs(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			"ex1",
			2,
			2,
		},
		{
			"ex2",
			3,
			3,
		},
		{
			"4",
			4,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
