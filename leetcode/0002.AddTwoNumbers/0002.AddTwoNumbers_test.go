package addtwonumbers

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Example 1",
			args{CreateList([]int{2, 4, 3}), CreateList([]int{5, 6, 4})},
			CreateList([]int{7, 0, 8}),
		},
		{
			"Example 2",
			args{CreateList([]int{0}), CreateList([]int{0})},
			CreateList([]int{0}),
		},
		{
			"Example 3",
			args{CreateList([]int{9, 9, 9, 9, 9, 9, 9}), CreateList([]int{9, 9, 9, 9})},
			CreateList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !IsTheSame(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
