package kdiffpairsinanarray

// 解題思路
//
// 將陣列轉成 map 紀錄每個值的出現次數，在第二次 loop 時使用 map 來迭代可避免重複的值出現
// 接下來我們只要檢查（值＋Ｋ）是否有在 map 裡即可
// 為何不用檢查（值－Ｋ）？因為（１，３）、（３，１）僅會計算一次，故我們在迭代的過程中一定會觸碰到一次（１，３）這個 pair
// 然後 k=0 是特別的情境，與 k!=0 的邏輯不同，沿用的話會導致（值＋０）永遠出現在 map 中
// 故單獨拉出來判斷出現次數出現超過１次的值即可

func findPairs(nums []int, k int) int {
	if k < 0 || len(nums) == 0 {
		return 0
	}

	count := 0
	counter := make(map[int]int, len(nums))
	for _, v := range nums {
		counter[v]++
	}

	for v := range counter {
		if k == 0 && counter[v] > 1 {
			count++
		} else if k != 0 && counter[v+k] > 0 {
			count++
		}
	}
	return count
}

// brute force

// func nums2str(a, b int) string {
// 	if a < b {
// 		return fmt.Sprintf("(%v,%v)", a, b)
// 	}
// 	return fmt.Sprintf("(%v,%v)", b, a)
// }

// func abs(c int) int {
// 	if c < 0 {
// 		return -1 * c
// 	}
// 	return c
// }

// func findPairs(nums []int, k int) int {
// 	ret := 0
// 	pairs := make(map[string]struct{}, len(nums))
// 	for i := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if _, ok := pairs[nums2str(nums[i], nums[j])]; !ok && abs(nums[i]- nums[j]) == k {
// 				pairs[nums2str(nums[i], nums[j])] = struct{}{}
// 				ret++
// 			}
// 		}
// 	}
// 	fmt.Println(pairs)
// 	return ret
// }
