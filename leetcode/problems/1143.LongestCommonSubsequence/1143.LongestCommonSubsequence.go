package longestcommonsubsequence

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcsRecursive(a, b string, i, j int) int {
	if len(a) == i || len(b) == j {
		return 0
	}

	var res int
	if a[i] == b[j] {
		res = 1 + lcsRecursive(a, b, i+1, j+1)
	} else {
		res = max(lcsRecursive(a, b, i+1, j), lcsRecursive(a, b, i, j+1))
	}

	return res
}

func longestCommonSubsequenceRecursive(text1 string, text2 string) int {
	return lcsRecursive(text1, text2, 0, 0)
}

// Runtime: 20 ms, faster than 24.11% of Go online submissions for Longest Common Subsequence.
// Memory Usage: 10.4 MB, less than 6.38% of Go online submissions for Longest Common Subsequence.
func lcsCacheMatrix(a, b string, i, j int, matrix [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if matrix[i][j] != -1 {
		return matrix[i][j]
	}

	var res int
	if a[i] == b[j] {
		res = 1 + lcsCacheMatrix(a, b, i-1, j-1, matrix)
	} else {
		res = max(lcsCacheMatrix(a, b, i-1, j, matrix), lcsCacheMatrix(a, b, i, j-1, matrix))
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

func longestCommonSubsequenceCacheMetrix(text1 string, text2 string) int {
	matrix := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		matrix[i] = make([]int, len(text2))
		initValue(matrix[i], -1)
	}

	for i := range text1 {
		for j := range text2 {
			lcsCacheMatrix(text1, text2, i, j, matrix)
		}
	}

	// displayMatrix(matrix)
	return matrix[len(text1)-1][len(text2)-1]
}

// Runtime: 4 ms, faster than 87.94% of Go online submissions for Longest Common Subsequence.
// Memory Usage: 10.5 MB, less than 6.38% of Go online submissions for Longest Common Subsequence.
func longestCommonSubsequence(text1 string, text2 string) int {

	matrix := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		matrix[i] = make([]int, len(text2)+1)
	}
	// displayMatrix(matrix)

	for m := 1; m <= len(text1); m++ {
		for n := 1; n <= len(text2); n++ {
			if text1[m-1] == text2[n-1] {
				matrix[m][n] = 1 + matrix[m-1][n-1]
			} else {
				matrix[m][n] = max(matrix[m-1][n], matrix[m][n-1])
			}
		}
	}

	// displayMatrix(matrix)
	return matrix[len(text1)][len(text2)]
}
