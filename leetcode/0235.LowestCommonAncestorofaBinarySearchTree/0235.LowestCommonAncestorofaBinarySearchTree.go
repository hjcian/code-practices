package lowestcommonancestorofabinarysearchtree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPath(n *TreeNode, target int, path *[]int) {
	if n == nil {
		return
	}
	if n.Val == target {
		*path = append(*path, n.Val)
		return
	}
	*path = append(*path, n.Val)
	if target < n.Val {
		findPath(n.Left, target, path)
	} else {
		findPath(n.Right, target, path)
	}
}

type Result struct {
	IsTouchedP bool
	IsTouchedA bool
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := make([]int, 0)
	findPath(root, p.Val, &pathP)
	pathQ := make([]int, 0)
	findPath(root, q.Val, &pathQ)
	fmt.Println(pathP)
	fmt.Println(pathQ)
	ancestor := root.Val
	for i := 0; i < len(pathP) && i < len(pathQ); i++ {
		if pathP[i] == pathQ[i] {
			ancestor = pathP[i]
		}
	}
	return &TreeNode{
		Val: ancestor,
	}
}
