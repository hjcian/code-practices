package easyquestionsmarks

import (
	"testing"
)

func TestQuestionsMarks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// 1. For input "aa6?9" the output was incorrect. The correct output is false
		{
			name:     "invalid case with only one question mark",
			input:    "aa6?9",
			expected: "false",
		},
		// 2. For input "acc?7??sss?3rr1??????5" the output was incorrect. The correct output is true
		{
			name:     "valid case with multiple characters and question marks",
			input:    "acc?7??sss?3rr1??????5",
			expected: "true",
		},
		// 3. For input "aaaaaaarrrrr??????" the output was incorrect. The correct output is false
		{
			name:     "invalid case with no digits",
			input:    "aaaaaaarrrrr??????",
			expected: "false",
		},
		// 4. For input "9???1???9???1???9" the output was incorrect. The correct output is true
		{
			name:     "valid case with multiple valid pairs",
			input:    "9???1???9???1???9",
			expected: "true",
		},
		// 5. For input "9???1???9??1???9" the output was incorrect. The correct output is false
		{
			name:     "invalid case with one invalid pair",
			input:    "9???1???9??1???9",
			expected: "false",
		},
		// 6. For input "4?5?4?aamm9" the output was incorrect. The correct output is false
		{
			name:     "invalid case with insufficient question marks",
			input:    "4?5?4?aamm9",
			expected: "false",
		},
		// 7. For input "5??aaaaaaaaaaaaaaaaaaa?5?5" the output was incorrect. The correct output is false
		{
			name:     "invalid case with insufficient question marks and extra characters",
			input:    "5??aaaaaaaaaaaaaaaaaaa?5?5",
			expected: "false",
		},
		// 8. For input "5??aaaaaaaaaaaaaaaaaaa?5?a??5" the output was incorrect. The correct output is true
		{
			name:     "valid case with sufficient question marks and extra characters",
			input:    "5??aaaaaaaaaaaaaaaaaaa?5?a??5",
			expected: "true",
		},
		// 9. For input "mbbv???????????4??????ddsdsdcc9?" the output was incorrect. The correct output is false
		{
			name:     "invalid case with insufficient question marks between digits",
			input:    "mbbv???????????4??????ddsdsdcc9?",
			expected: "false",
		},
		{
			name:     "valid case with 3 question marks between digits summing to 10",
			input:    "1???9",
			expected: "true",
		},
		{
			name:     "invalid case with 2 question marks",
			input:    "1??9",
			expected: "false",
		},
		{
			name:     "invalid case with 4 question marks",
			input:    "1????9",
			expected: "false",
		},
		{
			name:     "invalid case digits don't sum to 10",
			input:    "1???5",
			expected: "false",
		},
		{
			name:     "valid case with extra characters",
			input:    "a1???9b",
			expected: "true",
		},
		{
			name:     "valid case multiple pairs only first match returns true",
			input:    "2???7???3???7",
			expected: "true",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "false",
		},
		{
			name:     "no digits",
			input:    "abc???def",
			expected: "false",
		},
		{
			name:     "single digit",
			input:    "5",
			expected: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuestionsMarks(tt.input)
			if result != tt.expected {
				t.Errorf("QuestionsMarks(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
