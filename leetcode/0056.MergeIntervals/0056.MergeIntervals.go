package mergeintervals

import (
	"fmt"
	"sort"
)

// Matrix stroe intervals
type Matrix [][]int

func (m Matrix) Len() int      { return len(m) }
func (m Matrix) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m Matrix) Less(i, j int) bool {
	// i 與 j 若按照順序出現，小於的那邊會擺在前面
	return m[i][0] < m[j][0]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Sort(Matrix(intervals))
	res := [][]int{}
	curr := intervals[0]
	fmt.Println("first: ", curr)
	for i := 1; i < len(intervals); i++ {
		if curr[1] < intervals[i][0] {
			// 此時的尾巴，比下一個頭還要小，表示沒有重疊
			// 將此時的區間記錄下來，然後將下一個指派給此時的區間
			res = append(res, []int{curr[0], curr[1]})
			curr = intervals[i]
		} else if curr[1] < intervals[i][1] {
			// 此時的尾巴，比下一個頭還要大，表示有重疊
			// 再檢查，下一個的尾巴是否比此時的尾巴還要大，若比較大才要修正此時的尾巴
			curr[1] = intervals[i][1]
		}
	}
	res = append(res, curr)

	for i := range res {
		fmt.Println(res[i])
	}
	fmt.Println("============")
	return res
}
