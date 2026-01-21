package checkpalindromebyfilteringnonletters

import (
	"strings"
	"unicode"
)

/*
 * Complete the 'isAlphabeticPalindrome' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts STRING code as parameter.
 */

func isAlphabeticPalindrome(code string) bool {
	// Write your code here
	letters := ""
	for _, c := range code {
		if unicode.IsLetter(c) {
			letters += strings.ToLower(string(c))
		}
	}
	left := 0
	right := len(letters) - 1
	for left <= right {
		if letters[left] != letters[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     code := readLine(reader)

//     result := isAlphabeticPalindrome(code)

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
