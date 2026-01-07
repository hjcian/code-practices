package easycodelandusernamevalidation

import (
	"fmt"
	"strings"
	"unicode"
)

func CodelandUsernameValidation(str string) string {
	fmt.Println(str)
	if len(str) < 4 || len(str) > 25 {
		return "false"
	}
	if strings.HasSuffix(str, "_") {
		return "false"
	}
	if unicode.IsLetter(rune(str[0])) == false {
		return "false"
	}

	for _, c := range str {
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' {
			continue
		}
		return "false"
	}
	// match, _ := regexp.MatchString("^[A-Za-z][A-Za-z\\d\\_]+", str)
	// if !match {
	// 	return "false"
	// }

	// code goes here
	return "true"
}
