package validateproperlynestedbrackets

import "testing"

func TestAreBracketsProperlyMatched(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single pair parentheses", "()", true},
		{"single pair brackets", "[]", true},
		{"single pair braces", "{}", true},
		{"nested brackets", "({[]})", true},
		{"multiple pairs", "()[](){}", true},
		{"unmatched opening", "(", false},
		{"unmatched closing", ")", false},
		{"mismatched pair", "(]", false},
		{"closing before opening", ")(", false},
		{"mixed valid with text", "a(b)c", true},
		{"complex nested", "{[()()]}", true},
		{"extra closing bracket", "()[]{}", true},
		{"extra opening bracket", "()[]{}(", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := areBracketsProperlyMatched(tt.input)
			if result != tt.expected {
				t.Errorf("areBracketsProperlyMatched(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
