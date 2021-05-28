package removeelement

func removeElement(nums []int, val int) int {
	tmp := nums[:0]
	for _, v := range nums {
		if v != val {
			tmp = append(tmp, v)
		}
	}
	return len(tmp)
}
