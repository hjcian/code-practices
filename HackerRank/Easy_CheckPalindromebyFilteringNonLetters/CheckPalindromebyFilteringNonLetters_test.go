package checkpalindromebyfilteringnonletters

import "testing"

func TestIsAlphabeticPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected bool
	}{
		{"Simple palindrome", "racecar", true},
		{"Palindrome with spaces", "race car", true},
		{"Palindrome with punctuation", "A man, a plan, a canal: Panama", true},
		{"Not a palindrome", "hello", false},
		{"Single letter", "a", true},
		{"Two same letters", "aa", true},
		{"Two different letters", "ab", false},
		{"Empty string", "", true},
		{"Numbers and letters palindrome", "A1b1A", true},
		{"Mixed case palindrome", "RaCeCaR", true},
		{"With special characters", "Was it a car or a cat I saw?", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAlphabeticPalindrome(tt.code)
			if result != tt.expected {
				t.Errorf("isAlphabeticPalindrome(%q) = %v, want %v", tt.code, result, tt.expected)
			}
		})
	}
}
