package findthesmallestmissingpositiveinteger

// Considered corner cases
// Input: [3, 4, -1, 1]
// Input: [1, 1]
// Input: [1, 2, 3]
// Input: [0, 10, 2, -10, 1]

func findSmallestMissingPositive(orderNumbers []int32) int32 {

	// Write your code here
	if len(orderNumbers) == 0 {
		return 1
	}

	for i := range orderNumbers {
		for {
			current := int64(orderNumbers[i])
			//fmt.Println("idx:", i, "current:", current, orderNumbers)
			if current > 0 && current != int64(i+1) && current <= int64(len(orderNumbers)) && orderNumbers[current-1] != orderNumbers[i] {
				orderNumbers[i], orderNumbers[current-1] = orderNumbers[current-1], orderNumbers[i]
			} else {
				break
			}
		}
	}
	for i := range orderNumbers {
		if int64(orderNumbers[i]) != int64(i+1) {
			return int32(i + 1)
		}
	}
	return int32(len(orderNumbers) + 1) // e.g. [1,2,3,4,5], return 6
}
