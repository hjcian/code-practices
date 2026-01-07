package mediumminwindowsubstring

import "testing"

func TestMinWindowSubstring(t *testing.T) {
	tests := []struct {
		name     string
		strArr   []string
		expected string
	}{
		// 1. For input []string {"aabdccdbcacd", "aad"} the output was incorrect. The correct output is aabd
		{
			name:     "additional test case 0",
			strArr:   []string{"aabdccdbcacd", "aad"},
			expected: "aabd",
		},
		// 2. For input []string {"ahffaksfajeeubsne", "jefaa"} the output was incorrect. The correct output is aksfaje
		{
			name:     "additional test case 1",
			strArr:   []string{"ahffaksfajeeubsne", "jefaa"},
			expected: "aksfaje",
		},
		// 3. For input []string {"aaffhkksemckelloe", "fhea"} the output was incorrect. The correct output is affhkkse
		{
			name:     "additional test case 2",
			strArr:   []string{"aaffhkksemckelloe", "fhea"},
			expected: "affhkkse",
		},
		// 4. For input []string {"aaaaaaaaa", "a"} the output was incorrect. The correct output is a
		{
			name:     "additional test case 3",
			strArr:   []string{"aaaaaaaaa", "a"},
			expected: "a",
		},
		// 5. For input []string {"aaffsfsfasfasfasfasfasfacasfafe", "fafe"} the output was incorrect. The correct output is fafe
		{
			name:     "additional test case 4",
			strArr:   []string{"aaffsfsfasfasfasfasfasfacasfafe", "fafe"},
			expected: "fafe",
		},
		// 6. For input []string {"aaffsfsfasfasfasfasfasfacasfafe", "fafsf"} the output was incorrect. The correct output is affsf
		{
			name:     "additional test case 5",
			strArr:   []string{"aaffsfsfasfasfasfasfasfacasfafe", "fafsf"},
			expected: "affsf",
		},
		// 7. For input []string {"vvavereveaevafefaef", "vvev"} the output was incorrect. The correct output is vvave
		{
			name:     "additional test case 6",
			strArr:   []string{"vvavereveaevafefaef", "vvev"},
			expected: "vvave",
		},
		// 8. For input []string {"caae", "cae"} the output was incorrect. The correct output is caae
		{
			name:     "additional test case 7",
			strArr:   []string{"caae", "cae"},
			expected: "caae",
		},

		// 9. For input []string {"cccaabbbbrr", "rbac"} the output was incorrect. The correct output is caabbbbr

		{
			name:     "additional test case 8",
			strArr:   []string{"cccaabbbbrr", "rbac"},
			expected: "caabbbbr",
		},
		// Input: []string {"ahffaksfajeeubsne", "jefaa"}
		// Output: aksfaje
		{
			name:     "additional test case 1",
			strArr:   []string{"ahffaksfajeeubsne", "jefaa"},
			expected: "aksfaje",
		},
		// Input: []string {"aaffhkksemckelloe", "fhea"}
		// Output: affhkkse
		{
			name:     "additional test case 2",
			strArr:   []string{"aaffhkksemckelloe", "fhea"},
			expected: "affhkkse",
		},
		{
			name:     "example 1",
			strArr:   []string{"aaabaaddae", "aed"},
			expected: "dae",
		},
		{
			name:     "example 2",
			strArr:   []string{"aabdccdbcacd", "aad"},
			expected: "aabd",
		},
		{
			name:     "single character",
			strArr:   []string{"abcdef", "a"},
			expected: "a",
		},
		{
			name:     "all characters needed",
			strArr:   []string{"abcdef", "abcdef"},
			expected: "abcdef",
		},
		{
			name:     "window at end",
			strArr:   []string{"xyzabc", "abc"},
			expected: "abc",
		},
		{
			name:     "multiple valid windows",
			strArr:   []string{"aabbcc", "abc"},
			expected: "abbc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinWindowSubstring(tt.strArr)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
