package mediumminwindowsubstring

// Question: Min Window Substring
// Have the function MinWindowSubstring(strArr) take the array of strings stored in strArr, which will contain only two strings, the first parameter being the string N and the second parameter being a string K of some characters, and your goal is to determine the smallest substring of N that contains all the characters in K. For example: if strArr is ["aaabaaddae", "aed"] then the smallest substring of N that contains the characters a, e, and d is "dae" located at the end of the string. So for this example your program should return the string dae.
// Another example: if strArr is ["aabdccdbcacd", "aad"] then the smallest substring of N that contains all of the characters in K is "aabd" which is located at the beginning of the string. Both parameters will be strings ranging in length from 1 to 50 characters and all of K's characters will exist somewhere in the string N. Both strings will only contains lowercase alphabetic characters.

func MinWindowSubstring(strArr []string) string {
	dict := make(map[string]int, 0)
	have := make(map[string]int, 0)

	for _, c := range strArr[1] {
		dict[string(c)]++
		have[string(c)] = 0
	}

	firstMatch := -1
	for i, c := range strArr[0] {
		_, ok := dict[string(c)]
		if !ok {
			continue
		}
		firstMatch = i
		break
	}
	if firstMatch < 0 {
		return ""
	}

	total := len(strArr[1])
	finalMatch := -1
	for j := firstMatch; j < len(strArr[0]); j++ {
		cnt, ok := dict[string(strArr[0][j])]
		if ok && cnt > 0 {
			dict[string(strArr[0][j])]--
			total--
		}
		if ok && cnt == 0 {
			have[string(strArr[0][j])]++
		}
		if total == 0 {
			finalMatch = j
			break
		}
	}
	if finalMatch < 0 {
		return ""
	}

	for i := firstMatch; i < len(strArr[0]); i++ {
		cnt, ok := have[string(strArr[0][i])]
		if ok && cnt == 0 {
			firstMatch = i
			break
		}
		if ok && cnt > 0 {
			have[string(strArr[0][i])]--
		}
	}
	//fmt.Println("Ans:", strArr[0][firstMatch:finalMatch+1])
	return strArr[0][firstMatch : finalMatch+1]
}
