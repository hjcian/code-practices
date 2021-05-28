package longestsubstringwithoutrepeatingcharacters

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Runtime: 4 ms, faster than 91.03% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 5.7 MB, less than 9.33% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(s string) int {
	usedChar := make(map[string]int, len(s))

	start := 0
	maxLen := 0
	// tmmzuxt
	// 0123456
	// abcabcbb
	// 01234567
	for i := 0; i < len(s); i++ {
		idx, isExists := usedChar[string(s[i])]
		// start <= idx 有兩個目的
		// 1. 如果發現曾經出現的 char 的 index 是在目前的 start 還左邊（< 的用意），表示不影響目前累積的長度
		// 		需要更新 start 位置的情況是發現重複的 char index 比 start 還右邊，則移動 start 到 index + 1 的位置
		// 2. 如果連續幾個字元重複（bbb），表示 start 才剛好是前一個 b 的 index，此時我還是得更新 start 的位置
		if isExists && start <= idx {
			fmt.Printf("[%v] char: %s, prev seem: %v \n", i, string(s[i]), idx)
			start = idx + 1
		} else {
			maxLen = max(maxLen, i-start+1)
		}
		usedChar[string(s[i])] = i
	}

	return maxLen
}
