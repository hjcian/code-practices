package onepassremovalofkthnodefromend

import (
	"testing"
)

func TestRemoveKthNodeFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		list     []int32
		k        int32
		expected []int32
	}{
		{
			name:     "remove from middle",
			list:     []int32{1, 2, 3, 4, 5},
			k:        2,
			expected: []int32{1, 2, 4, 5},
		},
		{
			name:     "remove head",
			list:     []int32{1, 2, 3},
			k:        2,
			expected: []int32{2, 3},
		},
		{
			name:     "remove tail",
			list:     []int32{1, 2, 3},
			k:        0,
			expected: []int32{1, 2},
		},
		{
			name:     "single element",
			list:     []int32{1},
			k:        0,
			expected: []int32{},
		},
		{
			name:     "k greater than length",
			list:     []int32{1, 2, 3},
			k:        3,
			expected: []int32{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SinglyLinkedList{}
			for _, val := range tt.list {
				list.insertNodeIntoSinglyLinkedList(val)
			}

			result := removeKthNodeFromEnd(list.head, tt.k)

			var got []int32
			for node := result; node != nil; node = node.next {
				got = append(got, node.data)
			}

			if !equal(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}

func equal(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
