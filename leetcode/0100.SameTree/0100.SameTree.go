package sametree

import (
	"fmt"
	"math"
)

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(t *TreeNode) []int {
	ret := make([]int, 0)
	if t == nil {
		return ret
	}

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, t)
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		ret = append(ret, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		} else {
			ret = append(ret, math.MinInt)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		} else {
			ret = append(ret, math.MaxInt)
		}
	}
	return ret
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pp := serialize(p)
	qq := serialize(q)
	fmt.Println(pp)
	fmt.Println(qq)
	if len(pp) != len(qq) {
		return false
	}
	for i := 0; i < len(pp); i++ {
		if pp[i] != qq[i] {
			return false
		}
	}
	return true
}
