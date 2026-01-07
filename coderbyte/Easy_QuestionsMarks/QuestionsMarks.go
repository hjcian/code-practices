package easyquestionsmarks

import (
	"strconv"
	"unicode"
)

func QuestionsMarks(str string) string {
	foundDigit := false
	leftDigit := -1
	markCount := 0
	hasPair := false
	for _, c := range str {
		if !foundDigit && unicode.IsDigit(c) {
			foundDigit = true
			markCount = 0
			num, _ := strconv.Atoi(string(c))
			leftDigit = num
			continue
		}
		if foundDigit && c == '?' {
			markCount++
			continue
		}
		if foundDigit && unicode.IsDigit(c) {
			rightDigit, _ := strconv.Atoi(string(c))
			if leftDigit+rightDigit == 10 && markCount == 3 {
				hasPair = true
			}
			if leftDigit+rightDigit == 10 && markCount != 3 {
				// found an invalid pair, so return false
				return "false"
			}
			// reset for next pair
			markCount = 0
			leftDigit = rightDigit
		}
	}
	if hasPair {
		return "true"
	}
	return "false"
}
