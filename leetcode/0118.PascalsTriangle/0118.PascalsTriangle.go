package pascalstriangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	ret := make([][]int, numRows)
	ret[0] = []int{1}
	ret[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1 // first
		for j := 1; j < i; j++ {
			// fill the rest of elements
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
		ret[i][len(ret[i])-1] = 1 // last
	}
	return ret
}
