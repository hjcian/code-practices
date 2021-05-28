package removenthnodefromendoflist

import "testing"

// CreateList create a single linked list for unit test
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{nums[0], nil}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{nums[i], nil}
		curr = curr.Next
	}
	return head
}

// IsTheSame compares the list for unit testing
func IsTheSame(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		l1 *ListNode
		n  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Example 1",
			args{CreateList([]int{1, 2, 3, 4, 5}), 2},
			CreateList([]int{1, 2, 3, 5}),
		},
		{
			"Example 2",
			args{CreateList([]int{1}), 1},
			CreateList([]int{}),
		},
		{
			"Example 3",
			args{CreateList([]int{1, 2}), 1},
			CreateList([]int{1}),
		},
		{
			"Test (187 / 208 test cases passed.)",
			args{CreateList([]int{1, 2}), 2},
			CreateList([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.l1, tt.args.n); !IsTheSame(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
