package firstfactorial

import "testing"

func TestFirstFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}

	for _, tt := range tests {
		result := FirstFactorial(tt.input)
		if result != tt.expected {
			t.Errorf("FirstFactorial(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}
