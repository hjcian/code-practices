package rgbtohexconversion

import "testing"

func Test_rgb(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"sample 1",
			args{255, 255, 255},
			"FFFFFF",
		},
		{
			"sample 2",
			args{255, 255, 300},
			"FFFFFF",
		},
		{
			"sample 3",
			args{0, 0, 0},
			"000000",
		},
		{
			"sample 4",
			args{148, 0, 211},
			"9400D3",
		},
		{
			"sample 5",
			args{0, 0, -20},
			"000000",
		},
		{
			"sample 6",
			args{173, 255, 47},
			"ADFF2F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("rgb() = %v, want %v", got, tt.want)
			}
		})
	}
}
