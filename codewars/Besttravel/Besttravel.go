package besttravel

import (
	"fmt"
)

func util(t, k int, ls []int) int {
	if k == 1 {
		closedValue := -1
		for _, v := range ls {
			if v > closedValue && v <= t {
				closedValue = v
			}
		}
		return closedValue
	} else if k == 2 {
		closedValue := -1
		for i := range ls {
			for j := i + 1; j < len(ls); j++ {
				sum := ls[i] + ls[j]
				if sum > closedValue && sum <= t {
					// fmt.Println("	", ls[i], ls[j], sum)
					closedValue = sum
				}
			}
		}
		return closedValue
	}
	closedValue := -1
	for i := range ls {
		tmpClosedValue := util(t-ls[i], k-1, ls[i+1:])
		if tmpClosedValue > 0 {
			sum := ls[i] + tmpClosedValue
			if sum > closedValue && sum <= t {
				closedValue = sum
			}
		}
		fmt.Println(ls[i], tmpClosedValue)
	}
	return closedValue
}

func ChooseBestSum(t, k int, ls []int) int {
	// your code
	ans := util(t, k, ls)
	fmt.Println("Answer:", ans, t, k, ls)
	return ans
}
