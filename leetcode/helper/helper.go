package helper

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	NULL = -1 << 63
)

func BuildTree(nums []int) (root *TreeNode) {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root = &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	// using BFS to build the tree

	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	// PrintTree(root)
	return root
}

func PrintTree(root *TreeNode) {
	queue := []*TreeNode{root}
	lines := [][]int{}
	for {
		tmpQ := []*TreeNode{}
		line := []int{}
		needNextLevel := false

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			val := NULL
			if node != nil {
				val = node.Val
				needNextLevel = true
			}
			if node == nil {
				node = &TreeNode{}
			}
			tmpQ = append(tmpQ, node.Left, node.Right)
			line = append(line, val)
		}

		lines = append(lines, line)
		if needNextLevel {
			queue = tmpQ
		} else {
			break
		}
	}

	for _, line := range lines {
		print := false
		for _, v := range line {
			if v != NULL {
				print = true
			}
		}
		if print {
			fmt.Println(line)
		}
	}
}
