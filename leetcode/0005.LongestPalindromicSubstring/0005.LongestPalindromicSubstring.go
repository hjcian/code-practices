package longestpalindromicsubstring

func longestPalindrome(s string) string {
	n := len(s)
	start := 0
	long := 1

	for i := range s {
		// 起點，左右都相等
		l, r := i, i

		// 處理 even case: "baac" or "abbba"
		for r+1 < n && s[l] == s[r+1] {
			r += 1
		}

		// 處理 odd case: 正常來說，選定 i 之後，左右各跑一次，看看是否相等。
		// 若相等，則兩邊 ++
		// 那因為這樣無法在 iteration 中處理 event case，故先在上面處理掉目前的 s[i] 有相鄰的 event case 的情境
		for l-1 >= 0 && r+1 < n && s[l-1] == s[r+1] {
			l -= 1
			r += 1
		}

		if r-l+1 > long {
			start = l
			long = r - l + 1
		}
	}
	return string(s[start : start+long])
}
