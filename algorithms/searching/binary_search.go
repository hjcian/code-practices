package searching

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}
