package searchinsertposition

func find(nums []int, start, end, target int) int {
	// Runtime: 4 ms, faster than 84.06% of Go online submissions for Search Insert Position.
	// Memory Usage: 3 MB, less than 99.59% of Go online submissions for Search Insert Position.

	if start == end {
		return start
	}

	if start+1 == end {
		if target <= nums[start] {
			return start
		}
		return end
	}

	mid := len(nums[start:end])/2 + start

	if target == nums[mid] || (target < nums[mid] && target > nums[mid-1]) {
		return mid
	}
	if target < nums[mid] {
		return find(nums, start, mid, target)
	}
	return find(nums, mid+1, end, target)
}

func searchInsert(nums []int, target int) int {
	return find(nums, 0, len(nums), target)
}
