package arrayproblems

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

func bubbleSort(nums []int) []int {
	end := len(nums)
	for end > 1 {
		for i := 0; i < end-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		end--
	}
	return nums
}
