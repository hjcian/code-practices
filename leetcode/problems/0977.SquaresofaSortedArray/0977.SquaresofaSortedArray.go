package squaresofasortedarray

// O(N log N) solution
// func sortedSquares(A []int) []int {
// 	for i := range A {
// 		A[i] = A[i] * A[i]
// 	}
// 	sort.Ints(A)
// 	return A
// }

// O(N) / O(N) solution
// 解題思路
// 目前最佳解就是創造額外 O(N) 空間儲存答案
// 然後，因為有負數存在，故最大值可能出現在頭，故用頭尾雙指針來走過一遍 slice

func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	for i, j, k := 0, len(A)-1, len(A)-1; i <= j; k-- {
		if A[i]*A[i] > A[j]*A[j] {
			ans[k] = A[i] * A[i]
			i++
		} else {
			ans[k] = A[j] * A[j]
			j--
		}
	}
	return ans
}
