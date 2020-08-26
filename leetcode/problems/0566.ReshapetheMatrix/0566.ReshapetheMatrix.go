package reshapethematrix

// 解題思路
// 首先 early return 輸出陣形與給定陣型數量不一致的結果，專心處理能夠 reshape 的情境；
// 再來，初始化數值皆為 0 的矩陣
// 接著，計算 global index，loop 一遍，然後將舊的位置 assign 到新的位置去即可
// 思考以下的矩陣，有２個 Rows 與３個 Columns：
//    |C_j |0 1 2
// R_i|gidx|
//  0 |     0 1 2
//  1 |     3 4 5
// global index 由 row_i=0 開始數，可以發現規律是：
// gidx  R_i  C_j|  R_i的算法      C_j的算法
//   0    0    1  gidx/len(c)    gidx%len(c)
//   1    0    2
//   2    0    3

func matrixReshape(nums [][]int, r int, c int) [][]int {
	numsR := len(nums)
	numsC := len(nums[0])
	// fmt.Println(numsR, numsC)
	// fmt.Println(r, c)
	if numsR*numsC != r*c {
		return nums
	}
	ret := make([][]int, r, r)
	for i := range ret {
		ret[i] = make([]int, c, c)
	}
	for gidx := 0; gidx < numsR*numsC; gidx++ {
		oldR := gidx / numsC
		oldC := gidx % numsC
		newR := gidx / c
		newC := gidx % c
		// fmt.Println("================================")
		// fmt.Println("        gidx:", gidx, r, c)
		// fmt.Println("old position:", oldR, oldC)
		// fmt.Println("new position:", newR, newC)
		ret[newR][newC] = nums[oldR][oldC]
	}
	return ret
}
