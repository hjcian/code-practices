package threesum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	counter := make(map[int]int, len(nums))
	for i := range nums {
		counter[nums[i]]++
	}
	fmt.Println(counter)
	// counter 的目的是儲存出現次數及製造 uniqNums ，供實際解題時迭代
	uniqNums := make([]int, 0, len(counter))
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}

	sort.Ints(uniqNums)
	// 排序過即可避免加入重複的答案，及出現不該出現的答案
	// 排序僅花 O(n log n)
	// 此題最佳也必須使用到 O(n^2)，所以就儘管排序吧，反正 bounded 在 n^2
	fmt.Println(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		fmt.Println(counter[uniqNums[i]], uniqNums[i]*3)
		if counter[uniqNums[i]] >= 3 && uniqNums[i]*3 == 0 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			if counter[uniqNums[i]] >= 2 && uniqNums[i]*2+uniqNums[j] == 0 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if counter[uniqNums[j]] >= 2 && uniqNums[j]*2+uniqNums[i] == 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}

			remain := 0 - uniqNums[i] - uniqNums[j]
			if remain > uniqNums[j] && counter[remain] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], remain})
			}
		}
	}
	return res
}

// brute force must be not accepted !
// func threeSum(nums []int) [][]int {
// 	bucket := make(map[[3]int]struct{}, 0)
// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					tmp := []int{nums[i], nums[j], nums[k]}
// 					sort.Ints(tmp)
// 					bucket[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
// 				}
// 			}
// 		}
// 	}
// 	// fill the bucket
// 	fmt.Println(bucket)
// 	ret := make([][]int, 0)
// 	for key := range bucket {
// 		// tmp := make([]int, len(key))
// 		// copy(tmp, key[:])
// 		tmp := key[:]
// 		fmt.Println(&key[0])
// 		fmt.Println(&tmp[0])
// 		fmt.Println(tmp, key)
// 		ret = append(ret, key[:])
// 	}
// 	return ret
// }
