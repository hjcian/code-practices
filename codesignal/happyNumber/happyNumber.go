package happynumber

import (
	"fmt"
	"strconv"
	"strings"
)

func happyNumber(n int) bool {
	type empty struct{}
	memory := make(map[int]empty)
	current := n

	for {
		fmt.Printf("==== Iteration (current:%v)\n", current)
		if current == 1 {
			return true
		} else if _, ok := memory[current]; ok {
			return false
		}
		nStr := strconv.Itoa(current)
		nextValue := 0
		for _, c := range strings.Split(nStr, "") {
			num, _ := strconv.Atoi(c)
			nextValue += num * num
			fmt.Println(num, nextValue)
		}
		memory[current] = empty{}
		current = nextValue
		// break
	}
	// return false
}
