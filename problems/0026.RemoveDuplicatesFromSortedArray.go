package leetcode

func removeDuplicates(nums []int) int {
	idx := 0
	for {
		if idx+1 >= len(nums) {
			// prevent index exceed slice length
			break
		}

		if nums[idx] == nums[idx+1] {
			// remove element from slice by concatenate
			nums = append(nums[:idx], nums[idx+1:]...)
		} else {
			idx++
		}
	}
	return len(nums)
}
