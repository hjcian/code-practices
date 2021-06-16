package climbingstairs

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
