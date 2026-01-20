package climbingstairs

import (
	"math"
)

func climbStairs(n int) int {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
	// Memory Usage: 1.9 MB, less than 61.72% of Go online submissions for Climbing Stairs.
	if n <= 2 {
		// base cases
		return n
	}
	first := 1
	second := 2
	possibleWays := 0
	for n > 2 {
		possibleWays = first + second
		first = second
		second = possibleWays
		n--
	}
	return possibleWays
}

func climbStairs20260112(n int) int {
	math.Sqrt(float64(n))
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	nums := make([]int, n)
	nums[0] = 1
	nums[1] = 2
	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n-1]
}

type Record map[int]map[int]interface{}

func (r Record) Lookup(m, n int) bool {
	nPos, ok := r[m]
	if ok {
		_, used := nPos[n]
		if used {
			return true
		}
	}
	return false
}

func (r Record) Add(m, n int) {
	nPos, ok := r[m]
	if !ok {
		nPos = make(map[int]interface{}, 0)
	}
	nPos[n] = struct{}{}
	r[m] = nPos
}

func (r Record) Remove(m, n int) {
	nPos, ok := r[m]
	if ok {
		_, used := nPos[n]
		if used {
			delete(nPos, n)
		}
	}
}

func allLE(upperbound int, nums ...int) bool {
	for _, n := range nums {
		if n >= upperbound {
			return false
		}
	}
	return true
}

func allGE(lowerbound int, nums ...int) bool {
	for _, n := range nums {
		if n <= lowerbound {
			return false
		}
	}
	return true
}
