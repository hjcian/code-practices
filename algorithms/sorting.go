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

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left - 1 // 目前可能是左邊或右邊了，所以從左邊的 index-1 開始
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			// 將目前的 nums[j] 丟到最左邊去、原本最左邊的數值就往右丟
			// i 初始為 -1 就是考慮了第一次迭代就遇到比 pivot 小的情況
			i++
			nums[i], nums[j] = nums[j], nums[i]
			// 目前的 i 意義上是方才比 pivot 小的那個值的位置
		}
	}
	// 目前的 nums[i]，是方才比 pivot 小的那個值
	// 而 j 是確實地從左走到右一遍
	// 所以我知道，nums[i+1] 是比 pivot 大的值了，故將 pivot 與 nums[i+1] 互換
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func quickUtil(nums []int, left, right int) {
	if left < right {
		mid := partition(nums, left, right)
		quickUtil(nums, left, mid-1)
		quickUtil(nums, mid+1, right)
	}
}

func quickSort(nums []int) []int {
	quickUtil(nums, 0, len(nums)-1)
	return nums
}
