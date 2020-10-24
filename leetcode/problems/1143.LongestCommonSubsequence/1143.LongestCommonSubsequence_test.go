package longestcommonsubsequence

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{"abcde", "ace"},
			3,
		},
		{
			"Test 18 / 43 test cases passed.",
			args{"mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
