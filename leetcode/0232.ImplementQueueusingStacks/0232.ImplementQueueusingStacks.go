package implementqueueusingstacks

import (
	"codepractices/leetcode/helper"
	"fmt"
)

// Runtime: 1 ms, faster than 55.90% of Go online submissions for Implement Queue using Stacks.
// Memory Usage: 1.9 MB, less than 85.61% of Go online submissions for Implement Queue using Stacks.
//
// 10/14/2022 21:40	Accepted	1 ms	1.9 MB	golang
// 10/14/2022 21:39	Accepted	3 ms	2 MB	golang
// 10/14/2022 21:39	Accepted	2 ms	2.1 MB	golang
type MyQueue struct {
	input  helper.Stack
	output helper.Stack
}

func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}

func (q *MyQueue) printStacks(msg string) {
	fmt.Println("Operation:", msg)
	fmt.Printf("input: %v \n", q.input)
	fmt.Printf("output: %v \n", q.output)
	fmt.Println("================================================================")
}

func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
	q.printStacks("Push")
}

func (q *MyQueue) Pop() int {
	if q.Empty() {
		panic("queue is empty")
	}
	q.move()

	q.printStacks("Pop (before)")
	ret := q.output.Top()
	q.output.Pop()
	q.printStacks("Pop (after)")

	return ret
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		panic("queue is empty")
	}
	q.move()
	q.printStacks("Peek")
	return q.output.Top()
}

func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}

func (q *MyQueue) move() {
	// should only do move when output stack is empty, or the order will be messed up
	if !q.output.Empty() {
		return
	}

	for !q.input.Empty() {
		q.output.Push(q.input.Top())
		q.input.Pop()
	}
}
