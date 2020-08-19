package convertipv4

import (
	"math"
	"strconv"
	"strings"
)

func equivalentInterger(ipv4 string) string {
	// Useful tool:
	// 		http://www.aboutmyip.com/AboutMyXApp/IP2Integer.jsp?ipAddress=216.58.200.14
	// golang int/string convert hints:
	// 		https://yourbasic.org/golang/convert-int-to-string/
	//
	values := strings.Split(ipv4, ".")
	// fmt.Println(values)
	var ret float64 = 0
	for i, s := range values {
		if n, err := strconv.Atoi(s); err == nil {
			// [216 58 200 14]
			// [11011000, 00111010, 11001000, 00001110]
			// [ 2^(8*3),  2^(8*2),  2^(8*1),  2^(8*0) ]
			multiplier := math.Pow(float64(2), float64((len(values)-i-1)*8))
			ret += float64(multiplier) * float64(n)
		}
	}
	return strconv.Itoa(int(ret))
}

func convertIPv4(ipv4 string) []string {
	eq := equivalentInterger(ipv4)
	return []string{eq, "not"}
}
