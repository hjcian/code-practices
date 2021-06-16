package fibonaccinumber

import "testing"

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			"ex1",
			2,
			1,
		},
		{
			"ex2",
			3,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
