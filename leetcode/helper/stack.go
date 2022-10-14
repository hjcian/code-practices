package helper

type Stack []int

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Top() int {
	if len(*s) == 0 {
		panic("Stack is empty")
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() {
	if len(*s) == 0 {
		panic("Stack is empty")
	}
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
