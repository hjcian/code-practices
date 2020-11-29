package rgbtohexconversion

import "fmt"

func closestValue(n int) int {
	if n < 0 {
		return 0
	}
	if n > 255 {
		return 255
	}
	return n
}

func intToHex(n int) string {
	return fmt.Sprintf("%02X", n)
}

func rgb(a, b, c int) string {
	a, b, c = closestValue(a), closestValue(b), closestValue(c)
	return intToHex(a) + intToHex(b) + intToHex(c)
}
