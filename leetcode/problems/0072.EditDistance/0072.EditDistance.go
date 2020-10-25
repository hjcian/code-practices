package editdistance

import (
	"fmt"
	"sort"
)

func min(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

// Runtime: 28 ms, faster than 15.45% of Go online submissions for Edit Distance.
// Memory Usage: 6.5 MB, less than 7.32% of Go online submissions for Edit Distance.
func minDistance(word1 string, word2 string) int {
	// fmt.Println(word1, word2)
	// ros horse
	word1Len := len(word1)
	word2Len := len(word2)

	matrix := make([][]int, word1Len+1)
	for i := 0; i < word1Len+1; i++ {
		if i == 0 {
			matrix[i] = make([]int, word2Len+1)
			for j := 0; j < word2Len+1; j++ {
				matrix[i][j] = j
			}
		} else {
			matrix[i] = make([]int, word2Len+1)
			matrix[i][0] = i // 輔助值，用來幫助 routine 中 min() 運算時參考到對的 min. dist.
		}
	}
	// displayMatrix(matrix)
	// [0 1 2 3 4 5]
	// [1 0 0 0 0 0]
	// [2 0 0 0 0 0]
	// [3 0 0 0 0 0]

	// fmt.Println()
	for m := 1; m <= word1Len; m++ {
		for n := 1; n <= word2Len; n++ {
			if word1[m-1] == word2[n-1] {
				matrix[m][n] = matrix[m-1][n-1]
			} else {
				matrix[m][n] = 1 + min(
					matrix[m-1][n],
					matrix[m][n-1],
					matrix[m-1][n-1],
				)
			}
		}
	}
	// displayMatrix(matrix)
	// [0 1 2 3 4 5]
	// [1 1 2 2 3 4]
	// [2 2 1 2 3 4]
	// [3 3 2 2 2 3]

	return matrix[word1Len][word2Len]
}

func displayMatrix(matrix [][]int) {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}
