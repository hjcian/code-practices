package sort

func heapify(nums []int, left, right int) {
	// sort the array in-place by Max Heap
	// index from 0: if we have i-th element
	// i's parent = (i-1)/2
	// i's left child = 2i+1
	// i's right child = 2i+2
	for i := right; i >= left; i-- {
		for j := i; nums[j] > nums[(j-1)/2]; j = (j - 1) / 2 {
			// swap the i and its parent if i greater than parent
			nums[j], nums[(j-1)/2] = nums[(j-1)/2], nums[j]
		}
	}
}

func heapSort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		// sort the unsorted pile
		heapify(nums, 0, i)
		// swap the maximum value and first value
		nums[0], nums[i] = nums[i], nums[0]
		// next unsorted pile is 0~(i-1)
	}
	return nums
}
