package sort

// move i, compare the all left j to find i's position
// insertion 的意義為，將 i 插到左邊適當的位置
func insertionSort(nums []int) []int {
	numsLen := len(nums)
	for i := 1; i < numsLen; i++ {
		cursor := i - 1
		hold := nums[i]
		for cursor >= 0 && nums[cursor] > hold {
			nums[cursor+1] = nums[cursor]
			cursor--
		}
		// because nums[cursor] < hold, so put 'hold' in right of cursor
		nums[cursor+1] = hold
	}
	return nums
}

func insertionSort_20210531(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}

func insertionSort_20220417(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			} else {
				break
			}
		}
	}
	return nums
}
