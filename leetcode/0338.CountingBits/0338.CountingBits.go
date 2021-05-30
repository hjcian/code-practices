package countingbits

func countBits(n int) []int {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Counting Bits.
	// Memory Usage: 4.7 MB, less than 32.84% of Go online submissions for Counting Bits.

	countBit := func(n int) int {
		cnt := 0
		for n > 0 {
			cnt += n & 1
			n = n >> 1
		}
		return cnt
	}

	ret := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		ret[i] = countBit(i)
		// log.Println(i, countBit(i))
	}
	return ret
}

// countBitsByKernighansAlgorithm is O(n) solution, only one single pass
//nee to know this table: https://leetcode.com/problems/counting-bits/discuss/79538/Simple-Python-Solution/742084
func countBitsByKernighansAlgorithm(n int) []int {
	// Runtime: 4 ms, faster than 98.51% of Go online submissions for Counting Bits.
	// Memory Usage: 4.7 MB, less than 32.84% of Go online submissions for Counting Bits.
	ret := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		ret[i] = ret[i&(i-1)] + 1
	}
	return ret
}
