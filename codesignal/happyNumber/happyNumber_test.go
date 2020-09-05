package happynumber

import "testing"

func Test_happyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example",
			args{19},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := happyNumber(tt.args.n); got != tt.want {
				t.Errorf("happyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
