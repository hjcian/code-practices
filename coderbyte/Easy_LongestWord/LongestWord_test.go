package easylongestword

import "testing"

func TestLongestWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// 1. For input "this is some sort of sentence" the output was incorrect. The correct output is sentence
		{"this is some sort of sentence", "this is some sort of sentence", "sentence"},
		// 2. For input "longest word!!" the output was incorrect. The correct output is longest
		{"longest word!!", "longest word!!", "longest"},
		// 3. For input "coderbyte" the output was incorrect. The correct output is coderbyte
		{"coderbyte", "coderbyte", "coderbyte"},
		// 4. For input "a beautiful sentence^&!" the output was incorrect. The correct output is beautiful
		{"a beautiful sentence^&!", "a beautiful sentence^&!", "beautiful"},
		// 5. For input "oxford press" the output was incorrect. The correct output is oxford
		{"oxford press", "oxford press", "oxford"},
		// 6. For input "123456789 98765432" the output was incorrect. The correct output is 123456789
		{"123456789 98765432", "123456789 98765432", "123456789"},
		// 7. For input "letter after letter!!" the output was incorrect. The correct output is letter
		{"letter after letter!!", "letter after letter!!", "letter"},
		// 8. For input "a b c dee" the output was incorrect. The correct output is dee
		{"a b c dee", "a b c dee", "dee"},
		// 9. For input "a confusing /:sentence:/[ this is not!!!!!!!~" the output was incorrect. The correct output is confusing
		{"a confusing /:sentence:/[ this is not!!!!!!!~", "a confusing /:sentence:/[ this is not!!!!!!!~", "confusing"},
		{"fun&!! time", "fun&!! time", "time"},
		{"I love dogs", "I love dogs", "love"},
		{"a b c", "a b c", "a"},
		{"123 456 789", "123 456 789", "123"},
		{"the quick brown fox", "the quick brown fox", "quick"},
		{"", "", ""},
		{"a", "a", "a"},
		{"!@#$%", "!@#$%", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestWord(tt.input)
			if result != tt.expected {
				t.Errorf("LongestWord(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
