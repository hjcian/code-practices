package toeplitzmatrix

// 解題思路
// 第一個觀察，常對角矩陣 (Toeplitz matrix)的最右上與最左下不需要關心
// 每一個 [i][j] 其實與 [i+1][j+1] 比較就好（左上與右下），當比較不相等時表示不是常對角矩陣
// 然後 loop 一遍即可（矩陣的 O(N) 就 row 一遍、column 一遍）

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
