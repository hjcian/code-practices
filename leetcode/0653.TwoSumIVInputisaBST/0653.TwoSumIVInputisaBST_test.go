package twosumivinputisabst

import "testing"

func TestFindTarget(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		k    int
		want bool
	}{
		{
			name: "two sum exists",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
			k:    9,
			want: true,
		},
		{
			name: "two sum does not exist",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
			k:    28,
			want: false,
		},
		{
			name: "empty tree",
			root: nil,
			k:    5,
			want: false,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 5},
			k:    10,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTarget(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
