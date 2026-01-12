package validanagram

func isAnagram(s string, t string) bool {
	bucketS := make(map[rune]int, 0)
	bucketT := make(map[rune]int, 0)

	for _, c := range s {
		bucketS[c]++
	}
	for _, c := range t {
		bucketT[c]++
	}
	if len(bucketS) != len(bucketT) {
		return false
	}

	for char, countS := range bucketS {
		countT, ok := bucketT[char]
		if !ok {
			return false
		}
		if countS != countT {
			return false
		}
	}
	return true
}
