package longestsubstringwithoutrepeatingcharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{"abcabcbb"},
			3,
		},
		{
			"Example 2",
			args{"bbbbbbb"},
			1,
		},
		{
			"Example 3",
			args{"pwwkew"},
			3,
		},
		{
			"Example 4",
			args{""},
			0,
		},
		{
			"Test (179 / 987 test cases passed.)",
			args{"dvdf"},
			3,
		},
		{
			"Test (171 / 987 test cases passed.)",
			args{"tmmzuxt"},
			5,
		},
		{
			"Test (698 / 987 test cases passed.)",
			args{"abcabcbb"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
