package checkfornonidenticalstringrotation

import "testing"

func TestIsNonTrivialRotation(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "identical strings",
			s1:       "hello",
			s2:       "hello",
			expected: false,
		},
		{
			name:     "valid rotation",
			s1:       "hello",
			s2:       "elloh",
			expected: true,
		},
		{
			name:     "valid rotation multiple steps",
			s1:       "hello",
			s2:       "lohel",
			expected: true,
		},
		{
			name:     "not a rotation",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "single character",
			s1:       "a",
			s2:       "a",
			expected: false,
		},
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			expected: false,
		},
		{
			name:     "different lengths",
			s1:       "abc",
			s2:       "abcd",
			expected: false,
		},
		{
			name:     "two character rotation",
			s1:       "ab",
			s2:       "ba",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isNonTrivialRotation(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("isNonTrivialRotation(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
