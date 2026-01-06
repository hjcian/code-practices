package sort

import "fmt"

func merge20260106(left, right []int) []int {
	fmt.Println("Left:", left, "  Right:", right)
	totalLen := len(left) + len(right)
	merged := make([]int, 0, totalLen)
	leftIdx, rightIdx := 0, 0
	for i := 0; i < totalLen; i++ {
		if leftIdx == len(left) {
			merged = append(merged, right[rightIdx])
			rightIdx++
			continue
		}
		if rightIdx == len(right) {
			merged = append(merged, left[leftIdx])
			leftIdx++
			continue
		}

		if left[leftIdx] <= right[rightIdx] {
			merged = append(merged, left[leftIdx])
			leftIdx++
		} else {
			merged = append(merged, right[rightIdx])
			rightIdx++
		}
	}
	return merged
}

func mergeSort20260106(nums []int) []int {
	fmt.Println("Given:", nums)
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	leftNums := mergeSort20260106(nums[:mid])
	rightNums := mergeSort20260106(nums[mid:])
	merged := merge20260106(leftNums, rightNums)
	return merged
}

func mergeUtil(left, right []int) []int {
	totalLen := len(left) + len(right)
	// 只有 totalLen 個要分配到新的籃子裡，cap 給 totalLen
	ret := make([]int, 0, totalLen)
	li, ri := 0, 0
	// fmt.Printf("left:%v, right:%v \n", left, right)
	for i := 0; i < totalLen; i++ {
		if li == len(left) {
			// 表示 left 已經分配完，繼續分配 right
			ret = append(ret, right[ri])
			ri++
		} else if ri == len(right) {
			// 表示 right 已經分配完，繼續分配 left
			ret = append(ret, left[li])
			li++
		} else if left[li] < right[ri] {
			// 兩邊都還有，比大小，小的優先
			ret = append(ret, left[li])
			li++
		} else {
			ret = append(ret, right[ri])
			ri++
		}
	}
	// fmt.Printf("ret:%v \n", ret)
	return ret
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	middle := len(nums) / 2
	left := mergeSort(nums[0:middle])
	right := mergeSort(nums[middle:])
	return mergeUtil(left, right)
}
