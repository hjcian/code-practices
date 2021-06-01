package sort

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
