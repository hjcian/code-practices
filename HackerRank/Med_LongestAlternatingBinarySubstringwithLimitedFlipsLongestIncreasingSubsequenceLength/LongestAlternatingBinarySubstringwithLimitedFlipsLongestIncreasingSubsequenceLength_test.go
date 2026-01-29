package longestalternatingbinarysubstringwithlimitedflipslongestincreasingsubsequencelength

import "testing"

func TestLongestAlternatingSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int32
		expected int32
	}{
		{
			name:     "no flips needed",
			s:        "0101",
			k:        0,
			expected: 4,
		},
		{
			name:     "one flip allowed",
			s:        "0011",
			k:        1,
			expected: 3,
		},
		{
			name:     "multiple flips needed",
			s:        "000111",
			k:        2,
			expected: 6,
		},
		{
			name:     "single character",
			s:        "0",
			k:        0,
			expected: 1,
		},
		{
			name:     "all zeros",
			s:        "0000",
			k:        1,
			expected: 3,
		},
		{
			name:     "all ones",
			s:        "1111",
			k:        1,
			expected: 3,
		},
		{
			name:     "alternating string",
			s:        "101010",
			k:        0,
			expected: 6,
		},
		{
			name:     "large k",
			s:        "0011",
			k:        10,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestAlternatingSubstring(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("longestAlternatingSubstring(%q, %d) = %d, want %d", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
