package maximumsubarray

// divide and conquer solution

// O(n) solution
func maxSubArray(nums []int) int {
	glbMax := nums[0]
	subMax := nums[0]
	for _, v := range nums[1:] {

		if v+subMax > v {
			subMax += v
		} else {
			subMax = v
		}
		if subMax > glbMax {
			glbMax = subMax
		}
	}
	return glbMax
}
