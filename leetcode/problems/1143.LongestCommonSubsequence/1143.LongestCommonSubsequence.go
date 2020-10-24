package longestcommonsubsequence

import "fmt"

// Runtime: 20 ms, faster than 24.11% of Go online submissions for Longest Common Subsequence.
// Memory Usage: 10.4 MB, less than 6.38% of Go online submissions for Longest Common Subsequence.

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcs(a, b string, i, j int, matrix [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if matrix[i][j] != -1 {
		return matrix[i][j]
	}

	var res int
	if a[i] == b[j] {
		res = 1 + lcs(a, b, i-1, j-1, matrix)
	} else {
		res = max(lcs(a, b, i-1, j, matrix), lcs(a, b, i, j-1, matrix))
	}

	matrix[i][j] = res

	return res
}

func initValue(slice []int, value int) {
	for i := range slice {
		slice[i] = value
	}
}

func displayMatrix(matrix [][]int) {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	matrix := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		matrix[i] = make([]int, len(text2))
		initValue(matrix[i], -1)
	}

	for i := range text1 {
		for j := range text2 {
			lcs(text1, text2, i, j, matrix)
		}
	}

	displayMatrix(matrix)
	return matrix[len(text1)-1][len(text2)-1]
}
