package mediumtreeconstructor

import (
	"strconv"
	"strings"
)

// Tree Constructor
// Have the function TreeConstructor(strArr) take the array of strings stored in strArr, which will contain pairs of integers in the following format: (i1,i2), where i1 represents a child node in a tree and the second integer i2 signifies that it is the parent of i1. For example: if strArr is ["(1,2)", "(2,4)", "(7,2)"], then this forms the following tree:
// which you can see forms a proper binary tree. Your program should, in this case, return the string true because a valid binary tree can be formed. If a proper binary tree cannot be formed with the integer pairs, then return the string false. All of the integers within the tree will be unique, which means there can only be one node in the tree with the given integer value.

// Examples
// Input: []string {"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"}
// Output: true
// Input: []string {"(1,2)", "(3,2)", "(2,12)", "(5,2)"}
// Output: false

type Node struct {
	Val      int
	Parent   *Node
	Children map[int]*Node
}

func (n *Node) AddParent(p *Node) {
	n.Parent = p
}

func (n *Node) AddChild(c *Node) {
	n.Children[c.Val] = c
}

func MakeNode(val int) *Node {
	return &Node{
		Val:      val,
		Parent:   nil,
		Children: make(map[int]*Node, 0),
	}
}

type Tree struct {
	Nodes map[int]*Node
}

func (t *Tree) Add(childVal, parentVal int) {
	child, ok := t.Nodes[childVal]
	if !ok {
		child = MakeNode(childVal)
		t.Nodes[childVal] = child
	}
	parent, ok := t.Nodes[parentVal]
	if !ok {
		parent = MakeNode(parentVal)
		t.Nodes[parentVal] = parent
	}
	child.AddParent(parent)
	parent.AddChild(child)
}

func parse(str string) (int, int) {
	trimmed := strings.Trim(str, "()")
	nodes := strings.Split(trimmed, ",")
	nodeVal, _ := strconv.Atoi(nodes[0])
	parentVal, _ := strconv.Atoi(nodes[1])
	return nodeVal, parentVal
}

func TreeConstructor(strArr []string) string {
	tree := Tree{
		Nodes: make(map[int]*Node, 0),
	}

	for _, str := range strArr {
		child, parent := parse(str)
		tree.Add(child, parent)
	}
	rootCount := 0
	for _, node := range tree.Nodes {
		//fmt.Printf("node=%d addr=%p parent=%p (%v)\n", val, node, node.Parent, node.Parent)
		if node.Parent == nil {
			rootCount++
		}
		if len(node.Children) > 2 {
			return "false"
		}
	}
	//fmt.Println(rootCount)
	if rootCount != 1 {
		return "false"
	}
	// code goes here
	return "true"
}
