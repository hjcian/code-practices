package binarytreeinordertraversal

import (
	. "codepractices/leetcode/helper"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	recursive(root, &ret)
	return ret
}

func recursive(node *TreeNode, ret *[]int) {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
	// Memory Usage: 2.1 MB, less than 51.94% of Go online submissions for Binary Tree Inorder Traversal.

	if node == nil {
		return
	}
	recursive(node.Left, ret)
	*ret = append(*ret, node.Val)
	recursive(node.Right, ret)
}

func iterative(root *TreeNode) []int {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
	// Memory Usage: 2.1 MB, less than 51.94% of Go online submissions for Binary Tree Inorder Traversal.
	stack := []*TreeNode{}
	ret := []int{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ret = append(ret, node.Val)
			root = node.Right
		}
	}
	return ret
}
