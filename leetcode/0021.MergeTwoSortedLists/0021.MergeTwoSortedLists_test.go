package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sliceToListNodes(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}

	cursor := head
	for _, value := range values[1:] {
		cursor.Next = &ListNode{
			Val: value,
		}
		cursor = cursor.Next
	}

	return head
}

func listNodesToSlice(head *ListNode) []int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

func Test_0021_MergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		name      string
		giveList1 []int
		giveList2 []int
		wantList  []int
	}{
		{
			name:      "Example 1",
			giveList1: []int{1, 2, 4},
			giveList2: []int{1, 3, 4},
			wantList:  []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			list1 := sliceToListNodes(tt.giveList1)
			list2 := sliceToListNodes(tt.giveList2)
			gotListNodes := mergeTwoLists(list1, list2)
			gotList := listNodesToSlice(gotListNodes)
			assert.Equal(t, tt.wantList, gotList)
		})
	}
}
