package stringtointeger_atoi_

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"positive number", "42", 42},
		{"negative number", "-42", -42},
		{"leading whitespace", "   -42", -42},
		{"trailing characters", "4193 with words", 4193},
		{"no digits", "words and 987", 0},
		{"overflow positive", "9223372036854775808", 2147483647},
		{"overflow negative", "-9223372036854775809", -2147483648},
		{"empty string", "", 0},
		{"only whitespace", "   ", 0},
		{"plus sign", "+123", 123},
		{"leading zeros", "00000-42", 0},
		{"single zero", "0", 0},

		{"lc wrong 1", "21474836460", 2147483647},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := myAtoi(tt.input)
			if result != tt.expected {
				t.Errorf("myAtoi(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
