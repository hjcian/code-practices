package problem283

// =================================================================
// ## Success Details
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Move Zeroes.
//
// Memory Usage: 3.8 MB, less than 72.67% of Go online submissions for Move Zeroes.
// =================================================================

// Runtime 4 ms version
// func moveZeroes(nums []int) {
// 	headZero := -1
// 	for idx := 0; idx < len(nums); idx++ {
// 		if headZero == -1 && nums[idx] == 0 {
// 			headZero = idx
// 		} else if headZero != -1 && nums[idx] != 0 {
// 			nums[idx], nums[headZero] = nums[headZero], nums[idx]
// 			headZero++
// 		}
// 	}
// 	return
// }

// Runtime 0 ms version (key is early return)
// func moveZeroes(nums []int) {
// 	if len(nums) < 2 {
// 		return
// 	}
// 	headZero := -1
// 	for idx := 0; idx < len(nums); idx++ {
// 		if headZero == -1 && nums[idx] == 0 {
// 			headZero = idx
// 		} else if headZero != -1 && nums[idx] != 0 {
// 			nums[idx], nums[headZero] = nums[headZero], nums[idx]
// 			headZero++
// 		}
// 	}
// 	return
// }

// refine version
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	headZero := 0
	for idx := 0; idx < len(nums); idx++ {
		// 其實只要檢查目前的元素是否不等於零
		// line 46 可不用沒差，如果計較那一點 swap 的時間
		if nums[idx] != 0 {
			if idx != headZero {
				nums[idx], nums[headZero] = nums[headZero], nums[idx]
			}
			headZero++
		}
	}
	return
}
