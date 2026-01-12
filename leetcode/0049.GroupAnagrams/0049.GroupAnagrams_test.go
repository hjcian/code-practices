package groupanagrams

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "very long input",
			input: []string{"aaaaaaaaaaab", "aaaaaaaaaaa", "aaaaaaaaaaac", "aaaaaaaaaaad", "aaaaaaaaaaae", "aaaaaaaaaaaf", "aaaaaaaaaaag", "aaaaaaaaaaah", "aaaaaaaaaaai", "aaaaaaaaaaaj"},
			expected: [][]string{
				{"aaaaaaaaaaab"},
				{"aaaaaaaaaaa"},
				{"aaaaaaaaaaac"},
				{"aaaaaaaaaaad"},
				{"aaaaaaaaaaae"},
				{"aaaaaaaaaaaf"},
				{"aaaaaaaaaaag"},
				{"aaaaaaaaaaah"},
				{"aaaaaaaaaaai"},
				{"aaaaaaaaaaaj"},
			},
		},
		{
			name:  "example1",
			input: []string{"eat", "tea", "ate", "eta", "tan", "nat"},
			expected: [][]string{
				{"eat", "tea", "ate", "eta"},
				{"tan", "nat"},
			},
		},
		{
			name:     "single_word",
			input:    []string{"hello"},
			expected: [][]string{{"hello"}},
		},
		{
			name:     "all_anagrams",
			input:    []string{"ab", "ba"},
			expected: [][]string{{"ab", "ba"}},
		},
		{
			name:     "no_anagrams",
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "empty_string",
			input:    []string{""},
			expected: [][]string{{""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d groups, got %d", len(tt.expected), len(result))
			}
			for _, group := range result {
				slices.Sort(group)
			}
			for _, exp := range tt.expected {
				slices.Sort(exp)
				found := false
				for _, res := range result {
					if slices.Equal(res, exp) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected group %v not found in result", exp)
				}
			}
		})
	}
}
