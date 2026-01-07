package easylongestword

import (
	"strings"
	"unicode"
)

func LongestWord(sen string) string {
	cleaned := ""
	for _, c := range sen {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			cleaned += string(c)
		} else {
			cleaned += " "
		}
	}
	words := strings.Split(cleaned, " ")
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}

	}
	// code goes here
	return longest
}
