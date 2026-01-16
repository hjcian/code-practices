package lrucache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "example from comments",
			input:    []string{"A", "B", "C", "D", "A", "E", "D", "Z"},
			expected: "C-A-E-D-Z",
		},
		{
			name:     "failing test case",
			input:    []string{"J", "M", "R", "A", "B", "C", "R", "M", "F", "C"},
			expected: "B-R-M-F-C",
		},
		{
			name:     "single element",
			input:    []string{"A"},
			expected: "A",
		},
		{
			name:     "no duplicates under 5",
			input:    []string{"A", "B", "C"},
			expected: "A-B-C",
		},
		{
			name:     "fill to capacity",
			input:    []string{"A", "B", "C", "D", "E"},
			expected: "A-B-C-D-E",
		},
		{
			name:     "exceed capacity",
			input:    []string{"A", "B", "C", "D", "E", "F"},
			expected: "B-C-D-E-F",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LRUCache(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
