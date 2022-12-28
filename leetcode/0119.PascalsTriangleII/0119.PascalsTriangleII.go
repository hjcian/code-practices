package pascalstriangleii

func getRow(rowIndex int) []int {
	ret := []int{1}
	for i := 0; i < rowIndex; i++ {
		ret = append(ret, 0)
		for j := len(ret) - 1; j > 0; j-- {
			ret[j] += ret[j-1]
		}
	}
	return ret
}
