package longestpalindromicsubstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0005_LongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		name string
		give string
		want string
	}{
		{
			name: "ex1",
			give: "babad",
			want: "bab",
		},
		{
			name: "ex2",
			give: "cbbd",
			want: "bb",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.give)
			require.Equal(t, tt.want, got)
		})
	}
}
