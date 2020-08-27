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

func minCostClimbingStairs(cost []int) int {
	minCost := make([]int, len(cost), len(cost))
	minCost[0], minCost[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		minCost[i] = cost[i] + min(minCost[i-1], minCost[i-2])
	}
	return min(minCost[len(minCost)-1], minCost[len(minCost)-2])
}
