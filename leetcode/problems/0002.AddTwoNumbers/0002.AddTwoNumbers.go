package addtwonumbers

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList create a single linked list for unit test
func CreateList(nums []int) *ListNode {
	head := &ListNode{nums[0], nil}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{nums[i], nil}
		curr = curr.Next
	}
	return head
}

// ShowList print the entire linked list for debugging
func ShowList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
	fmt.Println()
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

// Runtime: 12 ms, faster than 62.78% of Go online submissions for Add Two Numbers.
// Memory Usage: 5.2 MB, less than 100.00% of Go online submissions for Add Two Numbers.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// ShowList(l1)
	// ShowList(l2)
	p1 := l1
	p2 := l2
	var head *ListNode
	var prev *ListNode
	carry := 0
	for p1 != nil || p2 != nil || carry == 1 {
		sum := 0
		v1 := 0
		v2 := 0
		if p1 != nil {
			v1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v2 = p2.Val
			p2 = p2.Next
		}
		sum = v1 + v2 + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		node := &ListNode{sum, nil}
		if head == nil {
			head = node
			prev = head
		} else {
			prev.Next = node
			prev = prev.Next
		}
	}
	// ShowList(head)
	return head
}
