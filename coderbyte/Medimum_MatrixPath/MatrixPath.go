package medimummatrixpath

// 題目：Matrix Path 這題測試你對 DFS 與方向控制的熟練度。

// 說明：給予一個二維陣列（只包含 0 和 1），請判斷是否能從左上角 (0,0) 走到右下角。你只能往「右、下、左、上」移動，且只能走在 1 上面。

func dfs(m, n int, matrix [][]int) bool {
	if m < 0 || n < 0 || m >= len(matrix) || n >= len(matrix[0]) {
		return false
	}
	if matrix[m][n] != 1 {
		return false
	}
	// deal with matrix[m][n] == 1
	if m == len(matrix)-1 && n == len(matrix[0])-1 {
		// reach the destination
		return true
	}
	matrix[m][n] = -1 // mark as visited for avoiding cycle
	result := dfs(m-1, n, matrix) || dfs(m+1, n, matrix) || dfs(m, n-1, matrix) || dfs(m, n+1, matrix)
	if result {
		return true
	}
	matrix[m][n] = 1
	return false
}

func MatrixPath(matrix [][]int) bool {
	return dfs(0, 0, matrix)
}
