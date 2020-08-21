package convertipv4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func toHex(n int64) string {
	return strconv.FormatInt(n, 16)
}

func compactGroup(values []string) string {
	str := strings.Join(values[:4], "")
	// remove leading zeros (refer to: https://en.wikipedia.org/wiki/IPv6#Address_representation)
	ret := strings.TrimLeft(str, "0")
	if len(ret) == 0 {
		return "0"
	}
	return ret
}

func ipv4ToHex(nums []int64) string {
	// Hint from: https://www.researchgate.net/figure/IPv4-to-IPv6-Conversion-Method1-In-this-method-firstly-to-convert-the-Decimal-IPv4_fig1_271294793
	// Details: https://zh.wikipedia.org/wiki/IPv6
	hexValues := make([]string, 0, 8)
	for _, n := range nums {
		leftBits := n >> 4
		rightBits := n & 0x0F
		hexValues = append(hexValues, toHex(leftBits), toHex(rightBits))
	}

	return fmt.Sprintf("0:0:0:0:0:ffff:%v:%v", compactGroup(hexValues[:4]), compactGroup(hexValues[4:]))
}

func equivalentInterger(nums []int64) string {
	// Useful tool:
	// 		http://www.aboutmyip.com/AboutMyXApp/IP2Integer.jsp?ipAddress=216.58.200.14
	// golang int/string convert hints:
	// 		https://yourbasic.org/golang/convert-int-to-string/
	//
	var ret float64 = 0
	for i, n := range nums {
		// [216 58 200 14]
		// [11011000, 00111010, 11001000, 00001110]
		// [ 2^(8*3),  2^(8*2),  2^(8*1),  2^(8*0) ]
		multiplier := math.Pow(float64(2), float64((len(nums)-i-1)*8))
		ret += float64(multiplier) * float64(n)
	}
	return strconv.Itoa(int(ret))
}

func convertIPv4(ipv4 string) []string {
	values := strings.Split(ipv4, ".")
	if len(values) != 4 {
		return []string{}
	}
	nums := make([]int64, 0, 4)
	for _, s := range values {
		if n, err := strconv.Atoi(s); err == nil {
			if n > 255 || n < 0 {
				return []string{}
			}
			nums = append(nums, int64(n))
		} else {
			return []string{}
		}
	}
	eqInt := equivalentInterger(nums)
	ipv6 := ipv4ToHex(nums)
	return []string{eqInt, ipv6}
}

func ipv4Toipv6(ipv4 string) string {
	values := strings.Split(ipv4, ".")
	nums := make([]int64, 0, 4)
	for _, s := range values {
		if n, err := strconv.Atoi(s); err == nil {
			nums = append(nums, int64(n))
		}
	}
	return ipv4ToHex(nums)
}
