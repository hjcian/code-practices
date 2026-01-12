package validpalindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	cleaned := ""
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			cleaned += strings.ToLower(string(c))
		}
	}
	if len(cleaned) == 0 {
		return true
	}
	left := 0
	right := len(cleaned) - 1
	for left <= right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true
}
