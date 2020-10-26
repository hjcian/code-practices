package binarysearch

import "fmt"

// Runtime: 292 ms, faster than 6.04% of Go online submissions for Binary Search.
// Memory Usage: 6.7 MB, less than 100.00% of Go online submissions for Binary Search.
func binSearchRecursive(nums []int, target, left, right int) int {
	midIdx := (right + left) / 2
	midVal := nums[midIdx]
	fmt.Println(left, right, midIdx, midVal)
	if midVal == target {
		return midIdx
	}
	if right-left < 2 {
		// 在此情況，若 midVal 沒比到，且可能是此 left=1 right=2 情況，表示沒有下一輪了
		return -1
	}
	if target < midVal {
		return binSearchRecursive(nums, target, left, midIdx)
	}
	return binSearchRecursive(nums, target, midIdx, right)
}

func searchRecursive(nums []int, target int) int {
	fmt.Println(nums)
	return binSearchRecursive(nums, target, 0, len(nums))
}

// Runtime: 36 ms, faster than 65.41% of Go online submissions for Binary Search.
// Memory Usage: 6.5 MB, less than 100.00% of Go online submissions for Binary Search.
func searchWhile(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := (right + left) / 2
	for left <= right {
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (right + left) / 2
	}
	return -1
}

func search(nums []int, target int) int {
	return searchWhile(nums, target)
}
