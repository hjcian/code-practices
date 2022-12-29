package maximumproductsubarray

import "fmt"

func max(a int, nums ...int) int {
	ret := a
	for i := range nums {
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}

func min(a int, nums ...int) int {
	ret := a
	for i := range nums {
		if nums[i] < ret {
			ret = nums[i]
		}
	}
	return ret
}

func maxProduct(nums []int) int {
	best := nums[0]
	currMax := 1
	currMin := 1

	for i := range nums {
		nextMax := currMax * nums[i]
		nextMin := currMin * nums[i]
		currMax = max(nextMax, nextMin, nums[i])
		currMin = min(nextMax, nextMin, nums[i])
		best = max(currMax, best)
		fmt.Printf("% 4d % 4d % 4d % 4d \n", nums[i], currMax, currMin, best)
	}
	return best
}
