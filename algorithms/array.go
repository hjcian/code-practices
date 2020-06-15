package arrayproblems

import (
	"fmt"
	"sort"
)


// 78. Subsets
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{})
	for _, v := range nums {
		for _, set := range ret {
			// neet to create another clear space for subset appending avoiding "append trap"
			tmpset := make([]int, len(set))
			copy(tmpset, set)
			tmpset = append(tmpset, v)
			// if v == 7 {
			// 	fmt.Printf("		prev set[%v]: %v \t (%v) %p %p \n", i, tmpset, ret[i], tmpset, ret[i])
			// }
			ret = append(ret, tmpset)
		}
	}
	return ret
}

// 581. Shortest Unsorted Continuous Subarray
func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	startIdx := 0
	for i := range nums {
		if nums[i] != sorted[i] {
			startIdx = i
			break
		}
	}

	endIdx := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != sorted[i] {
			endIdx = i
			break
		}
	}
	if endIdx == startIdx {
		// means no sub array
		return 0
	}
	return endIdx - startIdx + 1
}

// 1464. Maximum Product of Two Elements in an Array
func maxProduct(nums []int) int {
	var first int
	var second int
	if nums[0] > nums[1] {
		first = nums[0]
		second = nums[1]
	} else {
		first = nums[1]
		second = nums[0]
	}

	if len(nums) > 2 {
		for _, v := range nums[2:] {
			if v > first {
				second = first
				first = v
			} else if v > second {
				second = v
			}
		}
	}
	return (first - 1) * (second - 1)
}

// 1431. Kids With the Greatest Number of Candies
func kidsWithCandies(candies []int, extraCandies int) []bool {
	ret := make([]bool, len(candies))

	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= max {
			ret[i] = true
		}
	}
	return ret
}
