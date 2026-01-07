package easycodelandusernamevalidation

import "testing"

// 1. The username is between 4 and 25 characters.
// 2. It must start with a letter.
// 3. It can only contain letters, numbers, and the underscore character.
// 4. It cannot end with an underscore character.

func TestCodelandUsernameValidation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Valid cases
		{"User1", "true"},
		{"User_123", "true"},
		{"a1234", "true"},
		{"ValidUsername", "true"},

		// Too short
		{"abc", "false"},
		{"a1", "false"},

		// Too long
		{"abcdefghijklmnopqrstuvwxyz", "false"},

		// Ends with underscore
		{"User_", "false"},
		{"abc_", "false"},

		// Doesn't start with letter
		{"1User", "false"},
		{"_User", "false"},

		// Contains invalid characters
		{"User@123", "false"},
		{"User-Name", "false"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := CodelandUsernameValidation(tt.input)
			if result != tt.expected {
				t.Errorf("CodelandUsernameValidation(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
