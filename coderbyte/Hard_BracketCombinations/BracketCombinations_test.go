package hardbracketcombinations

import "testing"

func TestBracketCombinations(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"zero pairs", 0, 1},
		{"one pair", 1, 1},
		{"two pairs", 2, 2},
		{"three pairs", 3, 5},
		{"four pairs", 4, 14},
		{"five pairs", 5, 42},
		{"six pairs", 6, 132},
		{"seven pairs", 7, 429},
		{"eight pairs", 8, 1430},
		{"nine pairs", 9, 4862},
		{"ten pairs", 10, 16796},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BracketCombinations(tt.input)
			if result != tt.expected {
				t.Errorf("BracketCombinations(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
