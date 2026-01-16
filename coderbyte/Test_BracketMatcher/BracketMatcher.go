// Have the function BracketMatcher(str) take the str parameter being passed and return 1 if the brackets are correctly matched and each one is accounted for. Otherwise return 0. For example: if str is "(hello (world))", then the output should be 1, but if str is "((hello (world))" the the output should be 0 because the brackets do not correctly match up. Only "(" and ")" will be used as brackets. If str contains no brackets return 1.

package testbracketmatcher

func SearchingChallenge(str string) string {
	leftCount := 0
	for _, c := range str {
		if c == ')' {
			if leftCount == 0 {
				// invalid
				return "0"
			}
			leftCount--
		}
		if c == '(' {
			leftCount++
		}
	}
	// code goes here
	if leftCount == 0 {
		return "1"
	}
	return "0"
}

// do not modify below here, readline is our function
// that properly reads in the input for you
// func main() {

// 	fmt.Println(SearchingChallenge(readline()))

// }
