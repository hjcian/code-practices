package validparentheses

// Runtime: 4 ms, faster than 18.58% of Go online submissions for Valid Parentheses.
// Memory Usage: 2 MB, less than 43.44% of Go online submissions for Valid Parentheses.
func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		switch char {
		case ')':
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
