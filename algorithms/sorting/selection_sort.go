package sort

func selectionSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := range nums {
		minIdx := findMin(nums, i)
		if nums[minIdx] < nums[i] {
			nums[i], nums[minIdx] = nums[minIdx], nums[i] // swap them
		}
	}
	return nums
}

func findMin(nums []int, start int) int {
	minIdx := start
	for i := start; i < len(nums); i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// My old version
// numsLen := len(nums)
// for i := range nums {
// 	minIdx := i // current index assume is min.
// 	for j := i + 1; j < numsLen; j++ {
// 		if nums[j] < nums[minIdx] {
// 			minIdx = j
// 		}
// 	}
// 	// swap the min. value and current i value
// 	nums[i], nums[minIdx] = nums[minIdx], nums[i]
// }
// return nums
