package removenthnodefromendoflist

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Remove Nth Node From End of List.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev *ListNode
	candidate := head
	curr := head

	// 先移動跑針至與 candidate 相對位置為 n 的地方，下面以此跑針為主
	for i := 1; i < n; i++ {
		curr = curr.Next
	}

	// 移動跑針， candidate 也跟著移動。prev 用來記錄 candidate 的 prev node
	for curr.Next != nil {
		prev = candidate
		curr = curr.Next
		candidate = candidate.Next
	}

	if prev != nil {
		// 移動完了，檢查 prev。若不為 nil 表示大家都移動過，那麼就把 candicate node 刪除
		prev.Next = candidate.Next // delete candidate node
	} else {
		// 若 prev 仍然為 nil，表示後來沒有跑，那麼 candidate 就仍然是 head
		// 那麼就把 head 重新指到 Next
		head = head.Next // no prev means delete the head
	}

	return head
}
