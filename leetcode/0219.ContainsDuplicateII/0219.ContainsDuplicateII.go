package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int)
	for i, v := range nums {
		prevIdx, ok := record[v]
		if ok && i-prevIdx <= k {
			return true
		}
		record[v] = i

	}
	return false
}
