package binarytreepostordertraversal

import (
	. "codepractices/leetcode/helper"
	"fmt"
)

func postorderTraversal(root *TreeNode) []int {
	return iterative(root)
}

func recursive(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	recursive(root.Left, ret)
	recursive(root.Right, ret)
	*ret = append(*ret, root.Val)
}

func printStack(stack []*TreeNode) {
	s := ""
	for _, e := range stack {
		s += fmt.Sprintf("%v ", e.Val)
	}
	fmt.Println(s)
}

func iterative(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cnt := 0
	for root != nil || len(stack) > 0 {
		if cnt++; cnt > 10 {
			break
		}
		if root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			stack = append(stack, root)
			root = root.Left
			printStack(stack)
		} else {
			node := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			fmt.Println("	pop:", node.Val)
			if stack[len(stack)-1] == node.Right {
				stack = stack[0 : len(stack)-1]
				stack = append(stack, node)
				root = node.Right
			} else {
				res = append(res, node.Val)
			}
		}
	}
	return res
}
