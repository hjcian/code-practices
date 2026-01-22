package validateproperlynestedbrackets

/*
 * Complete the 'areBracketsProperlyMatched' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts STRING code_snippet as parameter.
 */

func isRight(c rune) bool {
	return c == ')' || c == ']' || c == '}'
}

func isLeft(c rune) bool {
	return c == '(' || c == '[' || c == '{'
}

func isPair(left, right rune) bool {
	return (left == '(' && right == ')') ||
		(left == '[' && right == ']') ||
		(left == '{' && right == '}')
}

func areBracketsProperlyMatched(code_snippet string) bool {
	// Write your code here
	stack := make([]rune, 0)
	for _, c := range code_snippet {
		if isRight(c) {
			if len(stack) == 0 {
				return false
			}
			if !isPair(stack[len(stack)-1], c) {
				return false
			}
			stack = stack[:len(stack)-1] // pop the last
		}
		if isLeft(c) {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     code_snippet := readLine(reader)

//     result := areBracketsProperlyMatched(code_snippet)

//     fmt.Printf("%s\n", (map[bool]string{true: "1", false: "0"})[result])
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
