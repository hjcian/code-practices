package transposematrix

// 解題思路
// 由於矩陣可能非方陣，所以還是得創建 O(N) 的空間來儲存回傳值
// 有了Ｂ矩陣之後就做一個 A[i][j] = B[j][i] 的動作

func transpose(A [][]int) [][]int {

	B := make([][]int, len(A[0]))
	for i := range B {
		B[i] = make([]int, len(A))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			B[j][i] = A[i][j]
		}
	}
	return B
}
