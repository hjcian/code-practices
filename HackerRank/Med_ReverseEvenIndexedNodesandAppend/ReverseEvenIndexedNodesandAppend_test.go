package reverseevenindexednodesandappend

import "testing"

func TestExtractAndAppendSponsoredNodes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected []int32
	}{
		{
			name:     "single node",
			input:    []int32{1},
			expected: []int32{1},
		},
		{
			name:     "two nodes",
			input:    []int32{1, 2},
			expected: []int32{2, 1},
		},
		{
			name:     "three nodes",
			input:    []int32{1, 2, 3},
			expected: []int32{2, 3, 1},
		},
		{
			name:     "four nodes",
			input:    []int32{1, 2, 3, 4},
			expected: []int32{2, 4, 3, 1},
		},
		{
			name:     "five nodes",
			input:    []int32{1, 2, 3, 4, 5},
			expected: []int32{2, 4, 5, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SinglyLinkedList{}
			for _, val := range tt.input {
				list.insertNodeIntoSinglyLinkedList(val)
			}

			result := extractAndAppendSponsoredNodes(list.head)

			var output []int32
			for node := result; node != nil; node = node.next {
				output = append(output, node.data)
			}

			if len(output) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, want %d", len(output), len(tt.expected))
			}

			for i, v := range output {
				if v != tt.expected[i] {
					t.Errorf("index %d: got %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}
