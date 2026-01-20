package twosumivinputisabst

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, k int, record map[int]interface{}) bool {
	if root == nil {
		return false
	}
	_, ok := record[k-root.Val]
	if ok {
		return true
	}
	record[root.Val] = struct{}{}
	return dfs(root.Left, k, record) || dfs(root.Right, k, record)
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func rinorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	rinorder(root.Right, arr)
	*arr = append(*arr, root.Val)
	rinorder(root.Left, arr)
}

func hashmapSol_20260119(root *TreeNode, k int) bool {
	record := make(map[int]any, 0)
	return dfs(root, k, record)
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	iIter := Constructor(root, false)
	rIter := Constructor(root, true)

	left := iIter.Next()
	right := rIter.Next()
	for iIter.HasNext() && rIter.HasNext() {
		if left+right == k {
			return true
		}
		if left+right > k {
			right = rIter.Next()
		} else {
			left = iIter.Next()
		}
	}

	return false
}

type BSTIterator struct {
	Stack []*TreeNode
	Rev   bool
}

func Constructor(root *TreeNode, rev bool) *BSTIterator {
	b := &BSTIterator{
		Stack: make([]*TreeNode, 0),
		Rev:   rev,
	}
	b.pushAll(root)
	return b
}

func (b *BSTIterator) HasNext() bool {
	return len(b.Stack) > 0
}

func (b *BSTIterator) Next() int {
	node := b.Stack[len(b.Stack)-1]
	b.Stack = b.Stack[:len(b.Stack)-1]
	if b.Rev {
		b.pushAll(node.Left)
	} else {
		b.pushAll(node.Right)
	}
	return node.Val
}

func (b *BSTIterator) PrintStack() {
	ret := make([]int, 0)
	for _, node := range b.Stack {
		ret = append(ret, node.Val)
	}
	fmt.Println(ret)
}

func (b *BSTIterator) pushAll(root *TreeNode) {
	for root != nil {
		b.Stack = append(b.Stack, root)
		if b.Rev {
			// put the maximum first
			root = root.Right
		} else {
			// put the minimum first
			root = root.Left
		}
	}
}
