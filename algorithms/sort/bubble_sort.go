package sort

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
