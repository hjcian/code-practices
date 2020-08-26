package maximumproductofthreenumbers

import (
	"sort"
)

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	// 排序過之後，比較最後三個、第一與最後兩個、前兩個與最後一個，誰比較大
	// 複雜度為排序時的 O(N Log N)

	sort.Ints(nums)

	last := len(nums) - 1
	product := nums[last] * nums[last-1] * nums[last-2]
	if tmp := nums[last] * nums[last-1] * nums[0]; tmp > product {
		product = tmp
	}
	if tmp := nums[last] * nums[1] * nums[0]; tmp > product {
		product = tmp
	}
	return product
}

// 複雜度 O(N) 的，則善用更多的暫存用變數
// func maximumProduct(nums []int) int {

// 	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
// 	min4, min5 := math.MaxInt32, math.MaxInt32
// 	for _, n := range nums {
// 		if n < min5 {
// 			min4, min5 = min5, n
// 		} else if n < min4 {
// 			min4 = n
// 		}
// 		if n > max1 {
// 			max1, max2, max3 = n, max1, max2
// 		} else if n > max2 {
// 			max2, max3 = n, max2
// 		} else if n > max3 {
// 			max3 = n
// 		}
// 	}
// 	fmt.Println(max1, max2, max3, min4, min5)
// 	topThree := max1 * max2 * max3
// 	topOneBottomTwo := max1 * min4 * min5
// 	if topThree > topOneBottomTwo {
// 		return topThree
// 	}
// 	return topOneBottomTwo
// }
