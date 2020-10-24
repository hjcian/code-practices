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
			"Example 2",
			args{"abc", "abc"},
			3,
		},
		{
			"Example 3",
			args{"abc", "def"},
			0,
		},
		{
			"Test 18 / 43 test cases passed.",
			args{"mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"},
			6,
		},
		{
			"Discussion: C++ with picture, O(nm)",
			args{"ace", "xabccde"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
			if got := longestCommonSubsequenceRecursive(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceRecursive() = %v, want %v", got, tt.want)
			}
		})
	}

}
