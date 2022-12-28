package firstmissingpositive

func firstMissingPositive(nums []int) int {
	// nums = append(nums, 0)
	// fmt.Println(nums)
	n := len(nums)

	// 3, 4, 1, 1
	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] >= n {
			nums[i] = 0
		}
	}

	// 3, 0, 1, 1 (i=0) ->3, 0, 1, 1+4=5
	// 3, 0, 1, 5 (i=1) ->3+4=7, 0, 1, 5
	// 7, 0, 1, 5 (i=2) ->7, 0+4=4, 1, 5
	// 7, 4, 1, 5 (i=3) 用 5%n 找回 1 -> 7, 4+4=8, 1, 5
	// 7, 8, 1, 5

	// fmt.Println(nums)
	for i := 0; i < n; i++ {
		nums[nums[i]%n] += n
	}
	// fmt.Println(nums)
	for i := 1; i < n; i++ {
		if nums[i] < n {
			return i
		}
	}
	return n
}
