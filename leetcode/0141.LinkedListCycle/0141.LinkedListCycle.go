package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 16 ms, faster than 42.97% of Go online submissions for Linked List Cycle.
// Memory Usage: 4.4 MB, less than 31.46% of Go online submissions for Linked List Cycle.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	firstPtr := head
	secondPtr := head.Next

	for firstPtr != nil && secondPtr != nil {
		if firstPtr == secondPtr {
			return true
		}
		if secondPtr.Next == nil {
			return false
		}
		firstPtr = firstPtr.Next
		secondPtr = secondPtr.Next.Next
	}
	return false
}
