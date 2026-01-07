package easyfirstreverse

import "testing"

func TestFirstReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// 1. For input "I Love Coding" the output was incorrect. The correct output is gnidoC evoL I
		{"I Love Coding", "I Love Coding", "gnidoC evoL I"},
		// 2. For input "h333llLo" the output was incorrect. The correct output is oLll333h
		{"With numbers", "h333llLo", "oLll333h"},
		// 3. For input "Yo0" the output was incorrect. The correct output is 0oY
		{"Mixed case and number", "Yo0", "0oY"},
		// 4. For input "thisiscool" the output was incorrect. The correct output is loocsisiht
		{
			"Lowercase letters", "thisiscool", "loocsisiht",
		},
		// 5. For input "commacomma!" the output was incorrect. The correct output is !ammocammoc
		{
			"With punctuation", "commacomma!", "!ammocammoc",
		},
		// 6. For input "123456789" the output was incorrect. The correct output is 987654321
		{
			"Numbers", "123456789", "987654321",
		},
		// 7. For input "lettersz!23z" the output was incorrect. The correct output is z32!zsrettel
		{
			"Mixed characters", "lettersz!23z", "z32!zsrettel",
		},
		// 8. For input "aq" the output was incorrect. The correct output is qa
		{
			"Two letters", "aq", "qa",
		},
		// 9. For input "b" the output was incorrect. The correct output is b
		{
			"I Love Coding", "I Love Coding",
			"gnidoC evoL I",
		},
		{"Simple string", "hello", "olleh"},
		{"Single character", "a", "a"},
		{"Empty string", "", ""},
		{"Two characters", "ab", "ba"},
		{"With spaces", "hello world", "dlrow olleh"},
		{"With special chars", "a!b@c", "c@b!a"},
		{"Numbers", "12345", "54321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstReverse(tt.input)
			if result != tt.expected {
				t.Errorf("FirstReverse(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
