package searching

import "math"

// Constraint: O(n) time complexity
func secondLargestValue(arr []int) (int, bool) {
	if len(arr) < 2 {
		return 0, false
	} else if len(arr) == 2 {
		if arr[0] == arr[1] {
			return 0, false
		}
		if arr[0] > arr[1] {
			return arr[1], true
		}
		return arr[0], true
	}

	largest := math.MinInt
	secondLargest := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}
	return secondLargest, true
}
