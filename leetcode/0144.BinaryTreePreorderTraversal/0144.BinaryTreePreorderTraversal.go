package binarytreepreordertraversal

import (
	. "codepractices/leetcode/helper"
)

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	recursive(root, &ret)
	return ret
}

func recursive(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	recursive(root.Left, ret)
	recursive(root.Right, ret)
}

func iterative(root *TreeNode) []int {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
	// Memory Usage: 2.1 MB, less than 49.39% of Go online submissions for Binary Tree Preorder Traversal.
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			res = append(res, root.Val)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
	}

	return res
}
