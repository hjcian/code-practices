package mincostclimbingstairs

// 解題思路
// 在有限範圍內做最佳化選擇，而每次的有限範圍內都是最佳化的話，累加起來的結果也會是整體的最佳化
// 趣味就在於「在有限範圍內」

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func minCostClimbingStairs(cost []int) int {
// 	minCost := make([]int, len(cost), len(cost))
// 	minCost[0], minCost[1] = cost[0], cost[1]
// 	for i := 2; i < len(cost); i++ {
// 		minCost[i] = cost[i] + min(minCost[i-1], minCost[i-2])
// 	}
// 	return min(minCost[len(minCost)-1], minCost[len(minCost)-2])
// }

// 其實只需要常數空間，紀錄前兩個最優解
// [a, b, c, d, e]
//	last1=a
//  last2=b
// 在 c 這步，我只需要知道我是 last1+c 比較優、還是 last2+c 比較優
// 比較優的解存入 last2；原本的 last2 存入 last1

func minCostClimbingStairs(cost []int) int {
	last1, last2 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		last1, last2 = last2, min(cost[i]+last1, cost[i]+last2)
	}
	return min(last1, last2)
}
