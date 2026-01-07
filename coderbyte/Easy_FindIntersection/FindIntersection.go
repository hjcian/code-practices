package findintersection

import (
	"strings"
)

func FindIntersection(strArr []string) string {
	first := strings.Split(strArr[0], ", ")
	second := strings.Split(strArr[1], ", ")
	dict := make(map[string]bool)
	for _, word := range first {
		dict[word] = true
	}
	ret := make([]string, 0)
	for _, word := range second {
		if dict[word] {
			ret = append(ret, word)
		}
	}
	// fmt.Println("intersection:", ret)
	if len(ret) > 0 {
		return strings.Join(ret, ",")
	}
	return "false"
}

// do not modify below here, readline is our function
// that properly reads in the input for you
// func main() {

// 	fmt.Println(FindIntersection(readline()))

// }
