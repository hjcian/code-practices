package rotateimage

func rotate(matrix [][]int) {
	// 解題思路
	// clockwise 順時針９０度會等於上下翻轉再沿對角線＼互換

	for i, j := 0, len(matrix)-1; i <= j && i != j; {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 變化題型：Counterclockwise 逆時針
	// 會等於左右翻轉再沿對角線＼互換
}
