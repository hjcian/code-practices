package sort

func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for end := len(nums) - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if nums[i] > nums[i+1] {
				swap(nums, i, i+1)
			}
		}
	}
	return nums
}

func bubbleSort_20210531(nums []int) []int {
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

func bubbleSort_20220417(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := len(nums); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
