package stringtointeger_atoi_

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	var sign int32 = 1
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = string(s[1:])
	}
	digits := ""
	for _, c := range s {
		if !unicode.IsDigit(c) {
			break
		}
		digits += string(c)
	}
	digits = strings.TrimLeft(digits, "0")
	fmt.Println(sign, digits)
	if len(digits) == 0 {
		return 0
	}

	i64, err := strconv.ParseInt(digits, 10, 32)
	num := (int32)(i64)
	if err != nil {
		if sign == 1 {
			return int(math.MaxInt32)
		} else {
			return int(math.MinInt32)
		}
	}
	return int(num * sign)
}
