package longestcommonsubsequence

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcs(a, b string, i, j int) int {
	if len(a) == 0 || len(b) == 0 || len(a) == i || len(b) == j {
		return 0
	}

	var res int
	if a[i] == b[j] {
		res = 1 + lcs(a, b, i+1, j+1)
	} else {
		res = max(lcs(a, b, i+1, j), lcs(a, b, i, j+1))
	}
	return res
}

func longestCommonSubsequence(text1 string, text2 string) int {
	return lcs(text1, text2, 0, 0)
}
