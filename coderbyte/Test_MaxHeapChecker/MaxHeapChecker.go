// Have the function MaxHeapChecker(arr) take arr which represents a heap data structure and determine whether or not it is a max heap. A max heap has the property that all nodes in the heap are either greater than or equal to each of its children. For example: if arr is [100,19,36,17] then this is a max heap and your program should return the string max. If the input is not a max heap then your program should return a list of nodes in string format, in the order that they appear in the tree, that currently do not satisfy the max heap property because the child nodes are larger than their parent. For example: if arr is [10,19,52,13,16] then your program should return 19,52.
// Another example: if arr is [10,19,52,104,14] then your program should return 19,52,104

package testmaxheapchecker

import (
	"strconv"
	"strings"
)

func ArrayChallenge(arr []int) string {
	invalidNums := make([]int, 0)
	// code goes here
	// child nodes formula:
	// left: 2i+1.  right: 2i+2. (<len(arr))
	// Max Heap property: each parent node is equal to or greater than its child nodes
	for parentIdx := range arr {
		//fmt.Println(parentIdx)
		leftIdx := 2*parentIdx + 1
		rightIdx := 2*parentIdx + 2
		if leftIdx < len(arr) && arr[parentIdx] < arr[leftIdx] {
			invalidNums = append(invalidNums, arr[leftIdx])
		}
		if rightIdx < len(arr) && arr[parentIdx] < arr[rightIdx] {
			invalidNums = append(invalidNums, arr[rightIdx])
		}
	}
	if len(invalidNums) > 0 {
		// convert []int -> []string -> Join
		numStrs := make([]string, 0)
		for _, num := range invalidNums {
			numStr := strconv.Itoa(num)
			numStrs = append(numStrs, numStr)
		}
		// fmt.Println(invalidNums)
		// fmt.Println(numStrs)
		return strings.Join(numStrs, ",")
	}
	return "max"
}

// do not modify below here, readline is our function
// that properly reads in the input for you
// func main() {

//   fmt.Println(ArrayChallenge(readline()))

// }
