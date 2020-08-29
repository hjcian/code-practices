package twosum

func twoSum(nums []int, target int) []int {
	diffSet := make(map[int][]int, 0)

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
