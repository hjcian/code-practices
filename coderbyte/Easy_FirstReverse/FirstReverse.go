package easyfirstreverse

func FirstReverse(str string) string {
	ret := make([]rune, 0)
	for i := len(str) - 1; i >= 0; i-- {
		ret = append(ret, rune(str[i]))
	}
	// code goes here
	return string(ret)
}

// func FirstReverse(str string) string {

// 	// code goes here
// 	// Note: feel free to modify the return type of this function
// 	ans := ""
// 	for _, c := range str {
// 		ans = string(c) + ans

// 	}
// 	return ans

// }
