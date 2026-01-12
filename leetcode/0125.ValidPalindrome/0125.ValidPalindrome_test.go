package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid palindrome", "A man, a plan, a canal: Panama", true},
		{"Invalid palindrome", "race a car", false},
		{"Empty string", "", true},
		{"Single character", "a", true},
		{"Only spaces", "   ", true},
		{"Only punctuation", "!@#$%", true},
		{"Palindrome with numbers", "0P", false},
		{"Numbers palindrome", "12321", true},
		{"Mixed case palindrome", "Madam", true},
		{"Non-alphanumeric chars", "a.", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
