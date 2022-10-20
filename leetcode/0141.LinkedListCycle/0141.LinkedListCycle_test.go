package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_0141_LinkedListCycle(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		pos   int // pos is used to denote the index of the node that tail's next pointer is connected to.
		want  bool
	}{
		{
			name:  "example 1",
			input: []int{3, 2, 0, -4},
			pos:   1,
			want:  true,
		},
		{
			name:  "example 2",
			input: []int{1, 2},
			pos:   0,
			want:  true,
		},
		{
			name:  "example 3",
			input: []int{1},
			pos:   -1,
			want:  false,
		},
		{
			name:  "null head case",
			input: []int{},
			pos:   -1,
			want:  false,
		},
		{
			name:  "two nodes",
			input: []int{1, 2},
			pos:   -1,
			want:  false,
		},
		{
			name:  "all one nodes",
			input: []int{1, 1, 1, 1},
			pos:   -1,
			want:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			head := MakeLinkedList(tt.input, tt.pos)
			printLinkedList(t, head, 10)
			require.Equal(t, tt.want, hasCycle(head))
		})
	}
}

func printLinkedList(t *testing.T, head *ListNode, limit int) {
	current := head
	for i := 0; i < limit && current != nil; i++ {
		t.Log(i, current.Val)
		current = current.Next
	}
}

func MakeLinkedList(nums []int, pos int) (head *ListNode) {
	if len(nums) == 0 {
		return nil
	}
	tail := &ListNode{Val: nums[len(nums)-1]}
	// initialize from last number
	head = tail
	for i := len(nums) - 2; i >= 0; i-- {
		node := &ListNode{Val: nums[i]}
		node.Next = head
		head = node
		if i == pos {
			// link tail to this node
			tail.Next = node
		}
	}
	return
}
