package sametree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "both nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "one nil",
			p:    &TreeNode{Val: 1},
			q:    nil,
			want: false,
		},
		{
			name: "same single node",
			p:    &TreeNode{Val: 1},
			q:    &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "different values",
			p:    &TreeNode{Val: 1},
			q:    &TreeNode{Val: 2},
			want: false,
		},
		{
			name: "same tree structure",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "different tree structure",
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
