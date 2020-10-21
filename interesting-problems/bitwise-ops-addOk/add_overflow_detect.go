package addoverflowdetect

// Truth table of Signed Bits of Two Integers
//          +   |   -
//      |   000 |   111
//  ----------------------
// +    |   000 |   000    (& and)
//  000 |   000 |   111    (| or )
//      |   000 |   111    (^ xor)
//  ----------------------
// -    |   000 |   111
//  111 |   111 |   111
//      |   111 |   000

// Consider
//         +               |  -
//   +  | (a+b)>>n == 000  |  (always OK)
//   -  | (always OK)      |  (a+b)>>n == 111

func addOK(a, b int8) bool {
	aSignBits := a >> 7
	bSignBits := b >> 7
	compareTarget := aSignBits & bSignBits
	mask := ^(aSignBits ^ bSignBits)
	return ((a+b)>>7)&mask == compareTarget
}

// Given:
//     MinInt = MaxInt + 1

// Truth table of Signed Bits of Two Integers
//          +   |   -
//      |   000 |   111
//  ----------------------
// +    |   000 |   000    (& and)
//  000 |   000 |   111    (| or )
//      |   000 |   111    (^ xor)
//  ----------------------
// -    |   000 |   111
//  111 |   111 |   111
//      |   111 |   000

// Consider:
//                +         |       -
//     +    a + b < MaxInt  |  (always OK)
//     -    (always OK)     |  a + b > MinInt = MaxInt + 1

// const MaxUint8 = ^uint8(0)
// const MinUint8 = 0
// const MaxInt8 = int8(MaxUint8 >> 1)
// const MinInt8 = -MaxInt8 - 1

// func abs8(n int8) int8 {
// 	mask := n >> 7
// 	return (n ^ mask) - mask
// }

// func printBin8(n int8) {
// 	fmt.Printf("%08b | %v \n", uint8(n), n)
// }

// func addOK(a, b int8) bool {
// 	maskA := a >> 7
// 	maskB := b >> 7
// 	mask := maskA & maskB
// 	magic := uint8(maskA^maskB) ^ 0xFF
// 	return (mask|1)*(MaxInt8-mask)-abs8(a) > (abs8(b) & int8(magic))
// }
