package findintersection

import "testing"

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		// 1. For input []string {"2, 5, 7, 10, 11, 12", "2, 7, 8, 9, 10, 11, 15"} the output was incorrect. The correct output is 2,7,10,11
		{
			name:     "multiple intersections",
			input:    []string{"2, 5, 7, 10, 11, 12", "2, 7, 8, 9, 10, 11, 15"},
			expected: "2,7,10,11",
		},
		// 2. For input []string {"1, 5, 6, 7, 10, 11, 12", "5, 6, 8, 11, 17"} the output was incorrect. The correct output is 5,6,11
		{
			name:     "another multiple intersections",
			input:    []string{"1, 5, 6, 7, 10, 11, 12", "5, 6, 8, 11, 17"},
			expected: "5,6,11",
		},
		// 3. For input []string {"2, 3, 4", "3"} the output was incorrect. The correct output is 3
		{
			name:     "single element intersection",
			input:    []string{"2, 3, 4", "3"},
			expected: "3",
		},
		// 4. For input []string {"1, 2, 3, 4, 5", "6, 7, 8, 9, 10"} the output was incorrect. The correct output is false
		{
			name:     "no intersection case",
			input:    []string{"1, 2, 3, 4, 5", "6, 7, 8, 9, 10"},
			expected: "false",
		},
		// 5. For input []string {"1, 2, 4, 5, 6, 9", "2, 3, 4, 8, 10"} the output was incorrect. The correct output is 2,4
		{
			name:     "two element intersection",
			input:    []string{"1, 2, 4, 5, 6, 9", "2, 3, 4, 8, 10"},
			expected: "2,4",
		},
		// 6. For input []string {"5, 6, 9, 11, 12, 16", "4, 6, 7, 11, 16"} the output was incorrect. The correct output is 6,11,16
		{
			name:     "three element intersection",
			input:    []string{"5, 6, 9, 11, 12, 16", "4, 6, 7, 11, 16"},
			expected: "6,11,16",
		},
		// 7. For input []string {"1, 3, 9, 10, 17, 18", "1, 4, 9, 10"} the output was incorrect. The correct output is 1,9,10
		{
			name:     "intersection with multiple elements",
			input:    []string{"1, 3, 9, 10, 17, 18", "1, 4, 9, 10"},
			expected: "1,9,10",
		},
		// 8. For input []string {"11, 12, 14, 16, 20", "11, 12, 13, 18, 21"} the output was incorrect. The correct output is 11,12
		{
			name:     "intersection at start",
			input:    []string{"11, 12, 14, 16, 20", "11, 12, 13, 18, 21"},
			expected: "11,12",
		},
		// 9. For input []string {"21, 22, 23, 25, 27, 28", "21, 24, 25, 29"} the output was incorrect. The correct output is 21,25
		{
			name:     "intersection with gaps",
			input:    []string{"21, 22, 23, 25, 27, 28", "21, 24, 25, 29"},
			expected: "21,25",
		},

		{
			name:     "basic intersection",
			input:    []string{"1, 2, 3, 4", "2, 4, 6, 8"},
			expected: "2,4",
		},
		{
			name:     "no intersection",
			input:    []string{"1, 3, 5", "2, 4, 6"},
			expected: "false",
		},
		{
			name:     "single element intersection",
			input:    []string{"1, 2, 3", "3, 4, 5"},
			expected: "3",
		},
		{
			name:     "all elements intersect",
			input:    []string{"1, 2, 3", "1, 2, 3"},
			expected: "1,2,3",
		},
		{
			name:     "empty second array",
			input:    []string{"1, 2, 3", ""},
			expected: "false",
		},
		{
			name:     "strings intersection",
			input:    []string{"a, b, c", "b, c, d"},
			expected: "b,c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindIntersection(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
