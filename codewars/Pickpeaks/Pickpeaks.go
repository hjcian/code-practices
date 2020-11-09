package pickpeaks

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	ret := PosPeaks{[]int{}, []int{}}
	totalNum := len(array)

	for i := 1; i < totalNum-1; {
		left := array[i-1]
		mid := array[i]

		// compare 'left' and 'mid', need a RISE
		if left >= mid {
			i++
			continue
		}

		j := i + 1 // right pointer for going through the plateau
		right := array[j]

		// util find out the drop
		for mid == right && j < totalNum-1 {
			j++
			right = array[j]
		}

		// compare 'mid' and 'right', need a DROP
		if mid > right {
			ret.Peaks = append(ret.Peaks, array[i])
			ret.Pos = append(ret.Pos, i)
		}
		i = j
	}

	return ret
}
