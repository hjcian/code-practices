package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "common prefix",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "no common prefix",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "single string",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "identical strings",
			input:    []string{"test", "test", "test"},
			expected: "test",
		},
		{
			name:     "common prefix single char",
			input:    []string{"a", "ac", "ab"},
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
