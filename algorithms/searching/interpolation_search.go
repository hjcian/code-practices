package searching

import "log"

func interpolationSearch(nums []int, target int) int {
	x1 := 0
	x2 := len(nums)
	y1 := nums[x1]
	y2 := nums[len(nums)-1] // just use largest number as the outer value, should be fine

	for x1 < x2 {
		x := x1 + (x2-x1)*(target-y1)/(y2-y1)
		log.Printf("(%d, %d, %d)", x1, x, x2)
		if x < 0 || x >= len(nums) {
			return -1
		}

		if nums[x] == target {
			return x
		}
		if target > nums[x] {
			x1, y1 = x, nums[x]
		} else {
			x2, y2 = x, nums[x]
		}
	}
	return -1
}
