package subtreeofanothertree

import "testing"

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    *TreeNode
		subRoot *TreeNode
		want    bool
	}{
		{
			name:    "both nil",
			root:    nil,
			subRoot: nil,
			want:    true,
		},
		{
			name:    "root nil",
			root:    nil,
			subRoot: &TreeNode{Val: 1},
			want:    false,
		},
		{
			name:    "subRoot nil",
			root:    &TreeNode{Val: 1},
			subRoot: nil,
			want:    false,
		},
		{
			name:    "identical single node",
			root:    &TreeNode{Val: 1},
			subRoot: &TreeNode{Val: 1},
			want:    true,
		},
		{
			name: "subRoot is root",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			want: true,
		},
		{
			name: "subRoot is right subtree",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			},
			subRoot: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			want: true,
		},
		{
			name:    "subRoot not in tree",
			root:    &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
			subRoot: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.root, tt.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
