package combinationsum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}

	for i := 0; i < len(candidates); i++ {
		if target == candidates[i] {

		}
	}
	return res
}
