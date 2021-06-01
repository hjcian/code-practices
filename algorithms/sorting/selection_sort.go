package sort

func selectionSort(nums []int) []int {
	numsLen := len(nums)
	for i := range nums {
		minIdx := i // current index assume is min.
		for j := i + 1; j < numsLen; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		// swap the min. value and current i value
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}
