package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	bucket := make([]*ListNode, 0)
	for list1 != nil || list2 != nil {
		if list1 == nil {
			bucket = append(bucket, list2)
			list2 = list2.Next
		} else if list2 == nil {
			bucket = append(bucket, list1)
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			bucket = append(bucket, list1)
			list1 = list1.Next
		} else {
			bucket = append(bucket, list2)
			list2 = list2.Next
		}
	}
	if len(bucket) == 0 {
		return nil
	}

	for i := 0; i < len(bucket); i++ {
		if i == len(bucket)-1 {
			bucket[i].Next = nil
			break
		}
		bucket[i].Next = bucket[i+1]
	}
	return bucket[0]
}
