package mediumtreeconstructor

import "testing"

func TestTreeConstructor(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		// 1. For input []string {"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"} the output was incorrect. The correct output is true
		{
			name:     "test case 1",
			input:    []string{"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"},
			expected: "true",
		},
		// 2. For input []string {"(1,2)", "(3,2)", "(2,12)", "(5,2)"} the output was incorrect. The correct output is false
		{
			name:     "test case 2",
			input:    []string{"(1,2)", "(3,2)", "(2,12)", "(5,2)"},
			expected: "false",
		},
		// 3. For input []string {"(2,5)", "(2,6)"} the output was incorrect. The correct output is false
		{
			name:     "test case 3",
			input:    []string{"(2,5)", "(2,6)"},
			expected: "false",
		},
		// 4. For input []string {"(10,20)"} the output was incorrect. The correct output is true
		{
			name:     "test case 4",
			input:    []string{"(10,20)"},
			expected: "true",
		},
		// 5. For input []string {"(2,3)", "(1,2)", "(4,9)", "(9,3)", "(12,9)", "(6,4)"} the output was incorrect. The correct output is true
		{
			name:     "test case 5",
			input:    []string{"(2,3)", "(1,2)", "(4,9)", "(9,3)", "(12,9)", "(6,4)"},
			expected: "true",
		},
		// 6. For input []string {"(2,3)", "(1,2)", "(4,9)", "(9,3)", "(12,9)", "(6,4)", "(1,9)"} the output was incorrect. The correct output is false
		{
			name:     "test case 6",
			input:    []string{"(2,3)", "(1,2)", "(4,9)", "(9,3)", "(12,9)", "(6,4)", "(1,9)"},
			expected: "false",
		},
		// 7. For input []string {"(1,2)", "(2,4)", "(7,4)"} the output was incorrect. The correct output is true
		{
			name:     "test case 7",
			input:    []string{"(1,2)", "(2,4)", "(7,4)"},
			expected: "true",
		},
		// 8. For input []string {"(5,6)", "(6,3)", "(2,3)", "(12,5)"} the output was incorrect. The correct output is true
		{
			name:     "test case 8",
			input:    []string{"(5,6)", "(6,3)", "(2,3)", "(12,5)"},
			expected: "true",
		},
		// 9. For input []string {"(10,20)", "(20,50)"} the output was incorrect. The correct output is true
		{
			name:     "test case 9",
			input:    []string{"(10,20)", "(20,50)"},
			expected: "true",
		},
		{
			name:     "valid binary tree",
			input:    []string{"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"},
			expected: "true",
		},
		{
			name:     "invalid binary tree multiple parents",
			input:    []string{"(1,2)", "(3,2)", "(2,12)", "(5,2)"},
			expected: "false",
		},
		{
			name:     "single node",
			input:    []string{"(1,2)"},
			expected: "true",
		},
		{
			name:     "cycle in tree",
			input:    []string{"(1,2)", "(2,3)", "(3,1)"},
			expected: "false",
		},
		{
			name:     "node with two parents",
			input:    []string{"(1,2)", "(1,3)"},
			expected: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TreeConstructor(tt.input)
			if result != tt.expected {
				t.Errorf("TreeConstructor(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}
