package sort

func heapify(nums []int, start, end int) {
	// sort the array in-place by Max Heap
	// index from 0: if we have i-th element
	// i's parent = (i-1)/2
	// i's left child = 2i+1
	// i's right child = 2i+2
	for i := end; i >= start; i-- {
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

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// time complexity: log(n)
func heapify_20220417(nums []int, i, end int) {
	if i >= end {
		return
	}

	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < end && nums[c1] > nums[max] {
		max = c1
	}
	if c2 < end && nums[c2] > nums[max] {
		max = c2
	}
	if max != i {
		swap(nums, i, max)
		heapify_20220417(nums, max, end)
	}
}

// time complexity: n/2 * log(n)
func build_heap_20220417(nums []int) {
	l := (len(nums) - 1) / 2
	for i := l; i >= 0; i-- {
		heapify_20220417(nums, i, len(nums))
	}
}

// time complexity: n/2 * log(n) + n * log(n) = 3n/2 * 2log(n) ~= nlog(n)
func heapSort_20220417(nums []int) []int {
	build_heap_20220417(nums)
	for i := len(nums) - 1; i > 0; i-- {
		swap(nums, 0, i)
		heapify_20220417(nums, 0, i)
	}
	return nums
}
