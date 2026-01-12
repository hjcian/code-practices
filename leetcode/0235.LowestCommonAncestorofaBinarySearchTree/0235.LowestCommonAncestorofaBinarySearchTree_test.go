package lowestcommonancestorofabinarysearchtree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		p        *TreeNode
		q        *TreeNode
		expected int
	}{
		{
			name: "both nodes on right",
			root: &TreeNode{Val: 6,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{Val: 4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{Val: 8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p:        &TreeNode{Val: 2},
			q:        &TreeNode{Val: 4},
			expected: 2,
		},
		{
			name:     "nodes on different sides",
			root:     &TreeNode{Val: 6, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 8}},
			p:        &TreeNode{Val: 2},
			q:        &TreeNode{Val: 8},
			expected: 6,
		},
		{
			name:     "one node is ancestor",
			root:     &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}, Right: &TreeNode{Val: 8}},
			p:        &TreeNode{Val: 6},
			q:        &TreeNode{Val: 2},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if result == nil || result.Val != tt.expected {
				t.Errorf("expected %d, got %v", tt.expected, result)
			}
		})
	}
}
