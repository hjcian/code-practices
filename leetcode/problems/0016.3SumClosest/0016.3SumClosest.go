package threesumclosest

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)
	ret := 0
	closest := math.MaxInt32
	for i := 0; i < n-2; i++ {
		// preserve last two for j and k
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]

			// if hit the target, then just return
			if sum == target {
				return sum
			}

			// check the closest number
			diff := abs(sum - target)
			if diff < closest {
				ret, closest = sum, diff
			}

			// choose one side to approach the closest number in this i iteration
			if sum > target {
				// need to reduce the total, so shrink from larger side
				k--
			} else {
				j++
			}
		}

	}
	return ret
}
