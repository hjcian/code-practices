package mediumbracketmatcher

import "testing"

func TestBracketMatcher(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// coderbytes test cases
		// 1. For input "hello()" the output was incorrect. The correct output is 1
		{
			"hello()", "1",
		},
		// 2. For input "one(bracket)" the output was incorrect. The correct output is 1
		{
			"one(bracket)", "1",
		},
		// 3. For input "coder(b)(y)(t)(e))" the output was incorrect. The correct output is 0
		{
			"coder(b)(y)(t)(e))", "0",
		},
		// 4. For input "()coderbyte() yes()" the output was incorrect. The correct output is 1
		{
			"()coderbyte() yes()", "1",
		},
		// 5. For input "dogs and cats" the output was incorrect. The correct output is 1
		{
			"dogs and cats", "1",
		},
		// 6. For input "01()01()01()" the output was incorrect. The correct output is 1
		{
			"01()01()01()", "1",
		},
		// 7. For input "the color re(d))()(()" the output was incorrect. The correct output is 0
		{
			"the color re(d))()(()", "0",
		},
		// 8. For input "letter(s) gal(o)(r)((e)" the output was incorrect. The correct output is 0
		{
			"letter(s) gal(o)(r)((e)", "0",
		},
		// 9. For input "three let(t)ers" the output was incorrect. The correct output is 1
		{
			"three let(t)ers", "1",
		},
		// 10. For input "twice thri(c)(e)()()" the output was incorrect. The correct output is 1
		{
			"twice thri(c)(e)()()", "1",
		},
		// Auto-generated test cases
		{"(hello (world))", "1"},
		{"((hello (world))", "0"},
		{"(coder)(byte))", "0"},
		{"(c(oder)) b(yte)", "1"},
		{"", "1"},
		{"hello world", "1"},
		{"()", "1"},
		{"((()))", "1"},
		{"(())", "1"},
		{"())(", "0"},
		{")hello(", "0"},
		{"(a(b)c)", "1"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := BracketMatcher(tt.input)
			if result != tt.expected {
				t.Errorf("BracketMatcher(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
