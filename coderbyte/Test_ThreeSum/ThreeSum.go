// Have the function ThreeSum(arr) take the array of integers stored in arr, and determine if any three distinct numbers (excluding the first element) in the array can sum up to the first element in the array. For example: if arr is [8, 2, 1, 4, 10, 5, -1, -1] then there are actually three sets of triplets that sum to the number 8: [2, 1, 5], [4, 5, -1] and [10, -1, -1]. Your program should return the string true if 3 distinct elements sum to the first element, otherwise your program should return the string false. The input array will always contain at least 4 elements.

package main

import (
	"slices"
)

func ArrayChallenge(arr []int) string {
	target := arr[0]
	nums := arr[1:]
	//fmt.Println(target, nums)
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	uniqNums := make([]int, 0)
	for num := range counts {
		uniqNums = append(uniqNums, num)
	}
	slices.Sort(uniqNums)
	//fmt.Println(uniqNums, counts)
	// (假設 nums 中數字沒有重複的可用)
	// 兩兩確認 (double for loop) remain = target - num[i] - num[j]
	// 如果 remain 不在 counts 裏，表示此兩數無解，繼續下一組
	// 如果 remain 在 counts 裏，且 remain != num[i] 且 remain != num[j]，表示有三個不同數字可用，直接回傳 true
	// (考慮重複數字的情況)
	// 如果 remain 在 counts 裏，且 remain == num[i]，那麼再檢查 num[i] 在 counts 裏的數量是否 >= 2，若是則表示 num[i] 有兩個可用，回傳 true
	// 如果 remain 在 counts 裏，且 remain == num[j]，那麼再檢查 num[j] 在 counts 裏的數量是否 >= 2，若是則表示 num[j] 有兩個可用，回傳 true
	for i := 0; i < len(uniqNums); i++ {
		for j := i + 1; j < len(uniqNums); j++ {
			remain := target - uniqNums[i] - uniqNums[j]
			_, exists := counts[remain]
			if !exists {
				continue
			}
			if remain != uniqNums[i] && remain != uniqNums[j] {
				// has 3 distinct numbers
				return "true"
			}
			if count, _ := counts[uniqNums[i]]; uniqNums[i] == remain && count > 1 {
				return "true"
			}
			if count, _ := counts[uniqNums[j]]; uniqNums[j] == remain && count > 1 {
				return "true"
			}
			//fmt.Println(uniqNums[i], uniqNums[j], remain)
		}
	}
	// code goes here
	//fmt.Println("------------------")
	return "false"
}

// do not modify below here, readline is our function
// that properly reads in the input for you
// func main() {

//   fmt.Println(ArrayChallenge(readline()))

// }
