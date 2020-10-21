package addoverflowdetect

import (
	"fmt"
)

const MaxUint8 = ^uint8(0)
const MinUint8 = 0
const MaxInt8 = int8(MaxUint8 >> 1)
const MinInt8 = -MaxInt8 - 1

func abs8(n int8) int8 {
	mask := n >> 7
	return (n ^ mask) - mask
}

func printBin8(n int8) {
	fmt.Printf("%08b | %v \n", uint8(n), n)
}

func addOK(a, b int8) bool {
	maskA := a >> 7
	maskB := b >> 7
	compareTarget := maskA & maskB
	mask := int8(uint8(maskA^maskB) ^ 0xFF)
	// printBin8(int8(magic))
	// fmt.Println("a, b, magic:", a, b, magic)
	return ((a+b)>>7)&mask == compareTarget
}

// func addOK(a, b int8) bool {
// 	maskA := a >> 7
// 	maskB := b >> 7
// 	mask := maskA & maskB

// 	magic := uint8(maskA^maskB) ^ 0xFF

// 	// printBin8(magic)
// 	fmt.Println("a, b, magic:", a, b, magic)
// 	return (mask|1)*(MaxInt8-mask)-abs8(a) > (abs8(b) & int8(magic))
// }
