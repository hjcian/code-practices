package removeduplicatesfromsortedarray

import "testing"

func Test_26(t *testing.T) {
	testSuits := []struct {
		name      string
		input     []int
		expectLen int
	}{
		{"Example 1", []int{1, 1, 2}, 2},
		{"Example 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, test := range testSuits {
		got := removeDuplicates(test.input)
		if got != test.expectLen {
			t.Errorf("[%v] expect '%v', got '%v' ", test.name, test.expectLen, got)
		}
	}
}
