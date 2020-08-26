package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i := m - 1
	j := n - 1
	k := len(nums1) - 1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	// 思考以下的組合：
	// i: A B C - - -
	// j: D E F
	// 最後，如果 i 或 j 有人已經跑完，此時只要檢查 j 是否還有剩，還有剩也只要直接填進去 nums1 即可
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
	return
}
