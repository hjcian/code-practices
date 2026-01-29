package maxuniquesubstringlengthinasession

import "testing"

func TestMaxDistinctSubstringLengthInSessions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int32
	}{
		{"empty string", "", 0},
		{"single asterisk", "*", 0},
		{"single character", "a", 1},
		{"all unique", "abc", 3},
		{"with duplicate", "abca", 3},
		{"with asterisk reset", "ab*cd", 2},
		{"multiple resets", "a*b*c", 1},
		{"asterisk at end", "abcd*", 4},
		{"asterisk at begin", "*abcd", 4},
		{"complex case", "abcabcbb", 3},
		{"with asterisk middle", "abc*def", 3},
		{"repeated chars", "aaaa", 1},
		{"two char pattern 1 ", "abab", 2},
		{"two char pattern 2", "abba", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDistinctSubstringLengthInSessions(tt.input)
			if result != tt.expected {
				t.Errorf("input %q: got %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
