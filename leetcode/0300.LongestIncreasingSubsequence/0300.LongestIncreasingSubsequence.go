package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	for i := range lis {
		lis[i] = 1
	}
	longest := 0
	for i := len(nums) - 1; i >= 0; i-- {
		//fmt.Println(quality[i:])
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
		if lis[i] > longest {
			longest = lis[i]
		}
	}
	return longest
}
