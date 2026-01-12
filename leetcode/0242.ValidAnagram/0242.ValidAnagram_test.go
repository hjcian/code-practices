package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"Valid anagram", "anagram", "nagaram", true},
		{"Not anagram", "rat", "car", false},
		{"Empty strings", "", "", true},
		{"Single character match", "a", "a", true},
		{"Single character mismatch", "a", "b", false},
		{"Different lengths", "ab", "abc", false},
		{"Same string", "hello", "hello", true},
		{"Reversed string", "abc", "cba", true},
		{"Unicode characters", "你好", "好你", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isAnagram(test.s, test.t)
			if result != test.expected {
				t.Errorf("isAnagram(%q, %q) = %v, expected %v", test.s, test.t, result, test.expected)
			}
		})
	}
}
