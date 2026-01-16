// Have the function LRUCache(strArr) take the array of characters stored in strArr, which will contain characters ranging from A to Z in some arbitrary order, and determine what elements still remain in a virtual cache that can hold up to 5 elements with an LRU cache algorithm implemented. For example: if strArr is ["A", "B", "C", "D", "A", "E", "D", "Z"], then the following steps are taken:
// (1) A does not exist in the cache, so access it and store it in the cache.
// (2) B does not exist in the cache, so access it and store it in the cache as well. So far the cache contains: ["A", "B"].
// (3) Same goes for C, so the cache is now: ["A", "B", "C"].
// (4) Same goes for D, so the cache is now: ["A", "B", "C", "D"].
// (5) Now A is accessed again, but it exists in the cache already so it is brought to the front: ["B", "C", "D", "A"].
// (6) E does not exist in the cache, so access it and store it in the cache: ["B", "C", "D", "A", "E"].
// (7) D is accessed again so it is brought to the front: ["B", "C", "A", "E", "D"].
// (8) Z does not exist in the cache so add it to the front and remove the least recently used element: ["C", "A", "E", "D", "Z"].
// Now the caching steps have been completed and your program should return the order of the cache with the elements joined into a string, separated by a hyphen. Therefore, for the example above your program should return C-A-E-D-Z.

// 1. For input []string {"J", "M", "R", "A", "B", "C", "R", "M", "F", "C"} the output was incorrect. The correct output is B-R-M-F-C

package lrucache

import (
	"fmt"
	"strings"
)

func LRUCache(strArr []string) string {
	dict := make(map[string]int, 0) // char-index
	chars := make([]string, 0)
	for _, str := range strArr {
		fmt.Println("Now:", str)
		idx, ok := dict[str]
		if ok {
			// move the element to the end of chars
			//fmt.Println("Cached:", str, idx, chars)
			left := chars[0:idx]
			right := chars[idx+1:]
			//fmt.Println("left, right:", left, right)
			chars = append([]string{}, left...)
			for _, leftStr := range right {
				dict[leftStr]-- // move left one index
				chars = append(chars, leftStr)
			}
			chars = append(chars, str)
			dict[str] = len(chars) - 1
			// fmt.Println("check char-index", dict)
			// fmt.Println("chars:", chars)
			continue
		}
		if len(chars) < 5 {
			chars = append(chars, str)
			dict[str] = len(chars) - 1
			// fmt.Println("check char-index", dict)
			// fmt.Println("chars:", chars)
			continue
		}
		// len(chars) >= 5, full and incoming char is the new char!
		first := chars[0]
		delete(dict, first)
		for _, str := range chars[1:] {
			dict[str]-- // move left one index
		}
		chars = append(chars[1:], str)
		dict[str] = len(chars) - 1 // missing this line
	}
	//fmt.Println(chars)
	//fmt.Println(dict)
	// code goes here
	return strings.Join(chars, "-")
}

// // do not modify below here, readline is our function
// // that properly reads in the input for you
// func main() {

//   fmt.Println(ArrayChallenge(readline()))

// }
