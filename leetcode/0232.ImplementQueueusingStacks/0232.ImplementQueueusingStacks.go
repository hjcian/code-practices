package implementqueueusingstacks

// Runtime: 1 ms, faster than 56.91% of Go online submissions for Implement Queue using Stacks.
// Memory Usage: 1.9 MB, less than 85.95% of Go online submissions for Implement Queue using Stacks.
type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	ret := this.stack[0]
	this.stack = this.stack[1:]
	return ret
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
