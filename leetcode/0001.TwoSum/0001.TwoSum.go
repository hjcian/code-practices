package twosum

import "fmt"

func twoSum(nums []int, target int) []int {
	diffSet := make(map[int][]int)

	for i, v := range nums {
		indices, ok := diffSet[v]
		if ok {
			diffSet[v] = append(indices, i)
		} else {
			diffSet[v] = []int{i}
		}
	}

	ret := make([]int, 0, len(nums))
	for i, v := range nums {
		indices, ok := diffSet[target-v]
		if ok {

			for _, index := range indices {
				if i != index {
					ret = append(ret, i)
				}
			}
		}
	}

	return ret
}

func twoSum20210528(nums []int, target int) []int {
	// twoSum20210528
	// Runtime: 4 ms, faster than 92.49% of Go online submissions for Two Sum.
	// Memory Usage: 4.6 MB, less than 5.03% of Go online submissions for Two Sum.
	lookupTable := make(map[int]int, len(nums))
	for i := range nums {
		lookupTable[target-nums[i]] = i
	}
	fmt.Println(nums, target)
	fmt.Println(lookupTable)
	res := []int{}
	for i := range nums {
		if pairIdx, ok := lookupTable[nums[i]]; ok && i < pairIdx {
			res = append(res, i, pairIdx)
		}
	}
	fmt.Println(res)
	return res
}
