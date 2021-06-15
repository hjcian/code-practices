package sortcolors

func sortColors(nums []int) {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
	// Memory Usage: 2.1 MB, less than 80.51% of Go online submissions for Sort Colors.
	zeroIdx := -1
	twoIdx := len(nums)
	i := 0
	for i < twoIdx {
		if nums[i] == 0 {
			zeroIdx++
			nums[i], nums[zeroIdx] = nums[zeroIdx], nums[i]
			i++
		} else if nums[i] == 2 {
			twoIdx--
			nums[i], nums[twoIdx] = nums[twoIdx], nums[i]
		} else {
			i++
		}
	}
}
