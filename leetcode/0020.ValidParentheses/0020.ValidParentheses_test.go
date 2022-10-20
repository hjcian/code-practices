package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0020_ValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "()",
			want: true,
		},
		{
			name: "example 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "example 3",
			s:    "(]",
			want: false,
		},
		{
			name: "my example 4",
			s:    "{([])}",
			want: true,
		},
		{
			name: "my example 5",
			s:    "{([)]}",
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, isValid(tt.s))
		})
	}
}
