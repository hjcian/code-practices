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
	// if end == 1，表示只剩下一個元素可排序 = 不用排序
	for end > 1 {
		for i := 0; i < end-1; i++ {
			// 每次兩兩比較從頭看到尾，可確保最大的元素會被丟到後面
			if nums[i] > nums[i+1] {
				// 發現左邊比右邊大，swap them
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		// 故當我已經知道最大的元素被丟到後面，那麼下一次整理時，就可以不用看到最後一個
		end--
	}
	return nums
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
